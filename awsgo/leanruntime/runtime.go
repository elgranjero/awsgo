package leanruntime

import (
	"bytes"
	"context"
	"encoding/csv"
	stdjson "encoding/json"
	"fmt"
	"html"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/goccy/go-json"
	"github.com/jmespath/go-jmespath"
	"gopkg.in/yaml.v3"
)

type Operation struct {
	Name   string
	Fields []Field
	Run    func(context.Context, aws.Config, Values) (any, error)
}

type Field struct {
	Name     string
	Flag     string
	Type     string
	Required bool
}

type Values map[string][]string

const (
	specialInputJSON = "__awsgo_input_json"
	specialInputFile = "__awsgo_input_file"
)

type Parsed struct {
	Operation string
	Profile   string
	Region    string
	Output    string
	OutputSet bool
	Query     string
	InputJSON string
	InputFile string
	Verbose   bool
	Help      bool
	Values    Values
}

func Execute(service string, operations map[string]Operation, args []string) error {
	parsed, err := Parse(service, operations, args)
	if err != nil {
		return err
	}
	if parsed.Help || parsed.Operation == "" {
		PrintHelp(service, operations, parsed.Operation)
		return nil
	}
	op, ok := operations[parsed.Operation]
	if !ok {
		return fmt.Errorf("unknown operation %q for service %q", parsed.Operation, service)
	}
	cfg, err := LoadConfig(parsed.Profile, parsed.Region)
	if err != nil {
		return err
	}
	if parsed.Verbose {
		_, _ = fmt.Fprintf(os.Stderr, "verbose: lean service=%s operation=%s fields=%v\n", service, parsed.Operation, parsed.Values)
	}
	resp, err := op.Run(context.Background(), cfg, parsed.Values)
	if err != nil {
		return err
	}
	out, err := Format(resp, parsed.ResolveOutput(), parsed.Query)
	if err != nil {
		return err
	}
	_, _ = os.Stdout.Write(out)
	if len(out) > 0 && out[len(out)-1] != '\n' {
		_, _ = os.Stdout.Write([]byte("\n"))
	}
	return nil
}

func Parse(service string, operations map[string]Operation, args []string) (Parsed, error) {
	p := Parsed{Output: "json", Values: Values{}}
	if len(args) > 0 && args[0] == service {
		args = args[1:]
	}
	if len(args) > 0 && args[0] == "help" {
		p.Help = true
		if len(args) > 1 {
			p.Operation = args[1]
		}
		return p, nil
	}

	flagToField := map[string]string{}
	for _, op := range operations {
		for _, field := range op.Fields {
			flag := field.Flag
			if flag == "" {
				flag = ToKebab(field.Name)
			}
			flagToField[flag] = field.Name
		}
	}

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "--help" || arg == "-h" {
			p.Help = true
			continue
		}
		if !strings.HasPrefix(arg, "-") {
			if p.Operation == "" {
				p.Operation = arg
			} else {
				return p, fmt.Errorf("unexpected argument %q", arg)
			}
			continue
		}
		name, val, hasValue := splitFlag(arg)
		switch name {
		case "profile":
			v, next, err := flagValue(name, val, hasValue, args, i)
			if err != nil {
				return p, err
			}
			p.Profile = v
			i = next
		case "region":
			v, next, err := flagValue(name, val, hasValue, args, i)
			if err != nil {
				return p, err
			}
			p.Region = v
			i = next
		case "output", "o":
			v, next, err := flagValue(name, val, hasValue, args, i)
			if err != nil {
				return p, err
			}
			p.Output = v
			p.OutputSet = true
			i = next
		case "query":
			v, next, err := flagValue(name, val, hasValue, args, i)
			if err != nil {
				return p, err
			}
			p.Query = v
			i = next
		case "input-json", "cli-input-json":
			v, next, err := flagValue(name, val, hasValue, args, i)
			if err != nil {
				return p, err
			}
			p.InputJSON = v
			i = next
		case "input-file":
			v, next, err := flagValue(name, val, hasValue, args, i)
			if err != nil {
				return p, err
			}
			p.InputFile = v
			i = next
		case "verbose":
			p.Verbose = true
		default:
			field, ok := flagToField[name]
			if !ok {
				return p, fmt.Errorf("unknown flag: --%s", name)
			}
			if hasValue {
				p.Values[field] = append(p.Values[field], val)
				continue
			}
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				p.Values[field] = append(p.Values[field], args[i+1])
				i++
			} else {
				p.Values[field] = append(p.Values[field], "true")
			}
		}
	}
	if p.Profile == "" {
		p.Profile = strings.TrimSpace(os.Getenv("AWS_PROFILE"))
	}
	if p.InputJSON != "" {
		p.Values[specialInputJSON] = []string{p.InputJSON}
	}
	if p.InputFile != "" {
		p.Values[specialInputFile] = []string{p.InputFile}
	}
	return p, nil
}

func (p Parsed) ResolveOutput() string {
	if p.OutputSet {
		if v := strings.TrimSpace(p.Output); v != "" {
			return v
		}
	}
	if v := strings.TrimSpace(os.Getenv("AWS_DEFAULT_OUTPUT")); v != "" {
		return v
	}
	if v := sharedConfigProfileOutput(p.Profile); v != "" {
		return v
	}
	return "json"
}

func splitFlag(arg string) (string, string, bool) {
	arg = strings.TrimLeft(arg, "-")
	name, val, ok := strings.Cut(arg, "=")
	return name, val, ok
}

func flagValue(name, val string, hasValue bool, args []string, i int) (string, int, error) {
	if hasValue {
		return val, i, nil
	}
	if i+1 >= len(args) {
		return "", i, fmt.Errorf("missing --%s value", name)
	}
	return args[i+1], i + 1, nil
}

func LoadConfig(profile, region string) (aws.Config, error) {
	var opts []func(*config.LoadOptions) error
	if profile != "" {
		opts = append(opts, config.WithSharedConfigProfile(profile))
	}
	if region != "" {
		opts = append(opts, config.WithRegion(region))
	}
	return config.LoadDefaultConfig(context.Background(), opts...)
}

func sharedConfigProfileOutput(profile string) string {
	profile = strings.TrimSpace(profile)
	if profile == "" {
		profile = strings.TrimSpace(os.Getenv("AWS_PROFILE"))
	}
	if profile == "" {
		profile = "default"
	}

	sectionName := "default"
	if profile != "default" {
		sectionName = "profile " + profile
	}

	configPath := strings.TrimSpace(os.Getenv("AWS_CONFIG_FILE"))
	if configPath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return ""
		}
		configPath = filepath.Join(home, ".aws", "config")
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return ""
	}

	currentSection := ""
	for _, rawLine := range strings.Split(string(data), "\n") {
		line := strings.TrimSpace(rawLine)
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}
		if strings.HasPrefix(line, "[") {
			if end := strings.Index(line, "]"); end >= 0 {
				currentSection = strings.TrimSpace(line[1:end])
			}
			continue
		}
		if currentSection != sectionName {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok || strings.TrimSpace(key) != "output" {
			continue
		}
		return strings.Trim(strings.TrimSpace(value), "\"'")
	}
	return ""
}

func Format(v any, output, query string) ([]byte, error) {
	structured, err := Normalize(v)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(query) != "" {
		structured, err = jmespath.Search(query, structured)
		if err != nil {
			return nil, fmt.Errorf("invalid --query: %w", err)
		}
		if structured == nil && (strings.Contains(query, "[*]") || strings.Contains(query, "[]")) {
			structured = []any{}
		}
	}
	outFmt := strings.ToLower(strings.TrimSpace(output))
	if outFmt == "yml" {
		outFmt = "yaml"
	}
	headers, rows := buildOutputRows(structured)
	switch outFmt {
	case "", "json":
		b, err := stdjson.MarshalIndent(structured, "", "  ")
		if err != nil {
			return nil, err
		}
		return append(b, '\n'), nil
	case "text", "yaml":
		b, err := yaml.Marshal(normalizeNilSlices(structured))
		if err != nil {
			return nil, err
		}
		return ensureTrailingNewline(b), nil
	case "table":
		var b bytes.Buffer
		writeTable(&b, headers, rows)
		return ensureTrailingNewline(b.Bytes()), nil
	case "csv":
		var b bytes.Buffer
		writeCSV(&b, headers, rows)
		return ensureTrailingNewline(b.Bytes()), nil
	case "md", "markdown":
		var b bytes.Buffer
		writeMarkdown(&b, headers, rows)
		return ensureTrailingNewline(b.Bytes()), nil
	case "html":
		var b bytes.Buffer
		writeHTML(&b, headers, rows)
		return ensureTrailingNewline(b.Bytes()), nil
	default:
		return nil, fmt.Errorf("unsupported --output %q (valid: json|yaml|text|table|csv|markdown|html)", output)
	}
}

func ensureTrailingNewline(b []byte) []byte {
	if len(b) > 0 && b[len(b)-1] != '\n' {
		return append(b, '\n')
	}
	return b
}

func Normalize(v any) (any, error) {
	b, err := json.Marshal(normalizeNilSlices(v))
	if err != nil {
		return nil, err
	}
	var out any
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return normalizeValue(out), nil
}

func normalizeValue(v any) any {
	switch x := v.(type) {
	case []any:
		out := make([]any, 0, len(x))
		for _, item := range x {
			out = append(out, normalizeValue(item))
		}
		return out
	case map[string]any:
		out := map[string]any{}
		truncated, _ := x["IsTruncated"].(bool)
		for k, val := range x {
			if k == "ResultMetadata" || k == "IsTruncated" || val == nil {
				continue
			}
			if k == "Marker" {
				if truncated {
					out["NextToken"] = normalizeValue(val)
				}
				continue
			}
			n := normalizeValue(val)
			if s, ok := n.(string); ok {
				n = normalizeStringField(k, s)
			}
			out[k] = n
		}
		return out
	case string:
		return normalizeTimestamp(x)
	default:
		return v
	}
}

func normalizeStringField(key, value string) any {
	if key == "Document" || strings.HasSuffix(key, "PolicyDocument") {
		if decoded, ok := decodeJSONDocument(value); ok {
			return decoded
		}
	}
	return normalizeTimestamp(value)
}

func decodeJSONDocument(value string) (any, bool) {
	s := strings.TrimSpace(value)
	if strings.Contains(s, "%") {
		if decoded, err := url.QueryUnescape(s); err == nil {
			s = strings.TrimSpace(decoded)
		}
	}
	if !strings.HasPrefix(s, "{") && !strings.HasPrefix(s, "[") {
		return nil, false
	}
	var out any
	if err := stdjson.Unmarshal([]byte(s), &out); err != nil {
		return nil, false
	}
	return normalizeValue(out), true
}

func normalizeTimestamp(value string) any {
	if len(value) < len("2006-01-02T15:04:05Z") || value[4] != '-' || value[7] != '-' || value[10] != 'T' {
		return value
	}
	t, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		return value
	}
	return t.Format("2006-01-02T15:04:05-07:00")
}

func PrintHelp(service string, operations map[string]Operation, operation string) {
	if operation == "" {
		fmt.Println("Dynamic AWS-style CLI over generated SDK commands")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Printf("  awsgo %s <operation> [flags]\n", service)
		fmt.Println()
		fmt.Println("Available Operations:")
		names := make([]string, 0, len(operations))
		for name := range operations {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			fmt.Printf("  %s\n", name)
		}
		return
	}
	op, ok := operations[operation]
	if !ok {
		fmt.Printf("unknown operation %q for service %q\n", operation, service)
		return
	}
	fmt.Printf("Usage:\n  awsgo %s %s [flags]\n\n", service, operation)
	fmt.Println("Global Flags:")
	fmt.Println("      --profile string   AWS shared config profile")
	fmt.Println("      --region string    AWS region")
	fmt.Println("      --output string    Output format: json|yaml|text|table|csv|markdown|html")
	fmt.Println("      --query string     Query JSON output (JMESPath)")
	fmt.Println("      --input-json str   Inline JSON object merged into SDK input")
	fmt.Println("      --cli-input-json   Alias of --input-json")
	fmt.Println("      --input-file path  JSON file merged into SDK input")
	fmt.Println("      --verbose          Print forwarded arguments and diagnostics")
	fmt.Println("  -h, --help             help for this operation")
	if len(op.Fields) > 0 {
		fmt.Println("\nOperation Flags:")
		for _, f := range op.Fields {
			req := "Optional"
			if f.Required {
				req = "Required"
			}
			fmt.Printf("      --%-32s %-8s %s\n", fieldFlagName(f), helpType(f.Type), req)
		}
	}
	fmt.Println()
	fmt.Println("Notes:")
	fmt.Println("  Operation input flags are generated from AWS SDK input structs.")
	fmt.Println("  Internal selector flags (for example, --describe-images) are not user-facing.")
	if len(op.Fields) > 0 {
		fmt.Println()
		fmt.Println("Known Input Fields:")
		for _, f := range op.Fields {
			fmt.Printf("  - %s\n", f.Name)
		}
		fmt.Println()
		fmt.Println("Example:")
		fmt.Printf("  awsgo %s %s --input-json '%s'\n", service, operation, inputJSONTemplate(op.Fields))
	}
}

func ToKebab(s string) string {
	if s == "" {
		return ""
	}
	return strings.Join(splitIdentifierWords(s), "-")
}

func splitIdentifierWords(s string) []string {
	if s == "" {
		return nil
	}
	s = strings.ReplaceAll(s, "QnA", "QNA")
	s = strings.ReplaceAll(s, "QuickSight", "Quicksight")
	s = strings.ReplaceAll(s, "CRC32C", "CRC32-C")
	s = strings.ReplaceAll(s, "CRC64NVMe", "CRC64-NVME")
	s = strings.ReplaceAll(s, "CRC64NVME", "CRC64-NVME")
	s = strings.ReplaceAll(s, "DDoS", "DDOS")
	s = strings.ReplaceAll(s, "DoS", "DOS")
	s = strings.ReplaceAll(s, "NVMe", "NVME")

	runes := []rune(s)
	parts := make([]string, 0, 8)
	start := 0
	for i := 1; i < len(runes); i++ {
		prev := runes[i-1]
		cur := runes[i]
		next := rune(0)
		if i+1 < len(runes) {
			next = runes[i+1]
		}
		if cur == '_' || cur == '-' || cur == ' ' {
			if start < i {
				parts = append(parts, strings.ToLower(string(runes[start:i])))
			}
			start = i + 1
			continue
		}
		boundary := false
		if isLower(prev) && isUpper(cur) {
			boundary = true
		}
		if isUpper(prev) && isUpper(cur) && isLower(next) {
			if i-start >= 2 && lowerRunLenFrom(runes, i+1) > 1 {
				boundary = true
			}
		}
		if isDigit(prev) && isUpper(cur) && isLower(next) {
			boundary = true
		}
		if boundary {
			parts = append(parts, strings.ToLower(string(runes[start:i])))
			start = i
		}
	}
	if start < len(runes) {
		parts = append(parts, strings.ToLower(string(runes[start:])))
	}
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p == "" {
			continue
		}
		if p == "qindex" {
			out = append(out, "q", "index")
			continue
		}
		out = append(out, p)
	}
	return out
}

func isLower(r rune) bool { return r >= 'a' && r <= 'z' }
func isUpper(r rune) bool { return r >= 'A' && r <= 'Z' }
func isDigit(r rune) bool { return r >= '0' && r <= '9' }
func lowerRunLenFrom(runes []rune, i int) int {
	n := 0
	for ; i < len(runes); i++ {
		if !isLower(runes[i]) {
			break
		}
		n++
	}
	return n
}

func helpType(typ string) string {
	t := normalizeType(typ)
	switch {
	case strings.HasPrefix(t, "[]"):
		return "list"
	case strings.Contains(t, "int32"), strings.Contains(t, "int64"), t == "int":
		return "int"
	case strings.Contains(t, "float32"), strings.Contains(t, "float64"):
		return "float"
	case strings.Contains(t, "bool"):
		return "bool"
	case strings.Contains(t, "time.Time"):
		return "timestamp"
	case strings.HasPrefix(t, "map["):
		return "object"
	case strings.HasPrefix(strings.TrimSpace(typ), "*types."):
		return "object"
	case strings.HasPrefix(t, "types."):
		return "string"
	default:
		return "string"
	}
}

func inputJSONTemplate(fields []Field) string {
	m := map[string]any{}
	limit := len(fields)
	if limit > 8 {
		limit = 8
	}
	for i := 0; i < limit; i++ {
		m[fields[i].Name] = templateValueForType(fields[i].Type)
	}
	b, err := stdjson.Marshal(m)
	if err != nil {
		return "{}"
	}
	return string(b)
}

func templateValueForType(typ string) any {
	t := strings.TrimSpace(typ)
	switch {
	case strings.Contains(t, "int32"), strings.Contains(t, "int64"), t == "int":
		return 0
	case strings.Contains(t, "float32"), strings.Contains(t, "float64"):
		return 0
	case strings.Contains(t, "bool"):
		return false
	case strings.Contains(t, "time.Time"):
		return "2026-01-01T00:00:00Z"
	case strings.HasPrefix(t, "[]"):
		return []any{}
	default:
		return ""
	}
}

func StringValue(values Values, name string) string {
	v := values[name]
	if len(v) == 0 {
		return ""
	}
	return v[len(v)-1]
}

func Int32Value(values Values, name string) (*int32, error) {
	s := StringValue(values, name)
	if strings.TrimSpace(s) == "" {
		return nil, nil
	}
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid --%s: %w", ToKebab(name), err)
	}
	vv := int32(v)
	return &vv, nil
}

func BoolValue(values Values, name string) (*bool, error) {
	s := StringValue(values, name)
	if strings.TrimSpace(s) == "" {
		return nil, nil
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return nil, fmt.Errorf("invalid --%s: %w", ToKebab(name), err)
	}
	return &v, nil
}

func ApplyInput(input any, fields []Field, values Values) (bool, error) {
	base, err := ReadInputJSON(StringValue(values, specialInputJSON), StringValue(values, specialInputFile))
	if err != nil {
		return false, err
	}

	fieldTypes := map[string]string{}
	for _, field := range fields {
		fieldTypes[field.Name] = field.Type
	}
	for name, typ := range fieldTypes {
		v, ok := base[name]
		if !ok {
			continue
		}
		cv, err := coerceJSONValue(typ, v)
		if err != nil {
			return false, fmt.Errorf("invalid input-json field %s: %w", name, err)
		}
		base[name] = cv
	}

	for _, field := range fields {
		rawValues := values[field.Name]
		if len(rawValues) == 0 {
			continue
		}
		v, err := coerceFlagValues(field.Type, rawValues)
		if err != nil {
			return false, fmt.Errorf("invalid --%s: %w", fieldFlagName(field), err)
		}
		base[field.Name] = v
	}

	if len(base) == 0 {
		return false, nil
	}
	b, err := stdjson.Marshal(base)
	if err != nil {
		return false, err
	}
	dec := stdjson.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	if err := dec.Decode(input); err != nil {
		return false, err
	}
	return shouldDisablePaginator(base), nil
}

func fieldFlagName(field Field) string {
	if strings.TrimSpace(field.Flag) != "" {
		return field.Flag
	}
	return ToKebab(field.Name)
}

func PaginatorDisabled() bool {
	switch strings.ToLower(strings.TrimSpace(os.Getenv("AWSGO_DISABLE_PAGINATOR"))) {
	case "1", "true", "yes", "on":
		return true
	default:
		return false
	}
}

func shouldDisablePaginator(input map[string]any) bool {
	if len(input) == 0 {
		return false
	}
	for _, k := range []string{"MaxItems", "MaxResults", "Limit", "PageSize"} {
		v, ok := input[k]
		if !ok || v == nil {
			continue
		}
		switch n := v.(type) {
		case stdjson.Number:
			if parsed, err := n.Float64(); err == nil && parsed > 0 {
				return true
			}
		case float64:
			if n > 0 {
				return true
			}
		case int64:
			if n > 0 {
				return true
			}
		case int:
			if n > 0 {
				return true
			}
		case string:
			if strings.TrimSpace(n) != "" {
				return true
			}
		default:
			return true
		}
	}
	return false
}

func coerceFlagValues(typ string, rawValues []string) (any, error) {
	if len(rawValues) == 0 {
		return nil, nil
	}
	if strings.HasPrefix(normalizeType(typ), "[]") {
		out := []any{}
		for _, raw := range rawValues {
			v, err := coerceValue(typ, raw)
			if err != nil {
				return nil, err
			}
			if arr, ok := v.([]any); ok {
				out = append(out, arr...)
			} else {
				out = append(out, v)
			}
		}
		return out, nil
	}
	return coerceValue(typ, rawValues[len(rawValues)-1])
}

func coerceValue(typ, raw string) (any, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return coerceScalar(typ, raw)
	}
	if strings.HasPrefix(trimmed, "[") {
		var arr []any
		if err := stdjson.Unmarshal([]byte(trimmed), &arr); err == nil {
			return coerceJSONArray(strings.TrimPrefix(normalizeType(typ), "[]"), arr)
		}
	}
	if strings.HasPrefix(trimmed, "{") {
		var obj map[string]any
		if err := stdjson.Unmarshal([]byte(trimmed), &obj); err == nil {
			return obj, nil
		}
	}
	if strings.HasPrefix(normalizeType(typ), "[]") {
		elem := strings.TrimPrefix(normalizeType(typ), "[]")
		if strings.HasPrefix(trimmed, "[") {
			items, err := parseShorthandList(trimmed)
			if err == nil {
				return coerceJSONArray(elem, items)
			}
		}
		if findTopLevel(trimmed, ',') >= 0 && findTopLevel(trimmed, '=') < 0 {
			parts, err := splitTopLevel(trimmed, ',')
			if err == nil {
				arr := make([]any, 0, len(parts))
				for _, p := range parts {
					arr = append(arr, strings.TrimSpace(p))
				}
				return coerceJSONArray(elem, arr)
			}
		}
		v, err := coerceScalar(elem, raw)
		if err != nil {
			return nil, err
		}
		return []any{v}, nil
	}
	return coerceScalar(typ, raw)
}

func parseShorthandMap(raw string) map[string]any {
	m, _ := parseShorthandMapAny(raw)
	return m
}

func parseShorthandMapAny(raw string) (map[string]any, error) {
	s := strings.TrimSpace(raw)
	if s == "" {
		return map[string]any{}, nil
	}
	if strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}") {
		s = strings.TrimSpace(s[1 : len(s)-1])
	}
	items, err := splitTopLevel(s, ',')
	if err != nil {
		return nil, err
	}
	out := map[string]any{}
	for _, item := range items {
		p := strings.TrimSpace(item)
		if p == "" {
			continue
		}
		idx := findTopLevel(p, '=')
		if idx < 0 {
			return nil, fmt.Errorf("invalid shorthand map item %q", p)
		}
		k := strings.TrimSpace(p[:idx])
		rhs := strings.TrimSpace(p[idx+1:])
		if k == "" {
			return nil, fmt.Errorf("empty shorthand key in %q", p)
		}
		val, err := parseShorthandValue(rhs)
		if err != nil {
			return nil, err
		}
		out[k] = val
	}
	return out, nil
}

func parseShorthandValue(raw string) (any, error) {
	s := strings.TrimSpace(raw)
	if s == "" {
		return "", nil
	}
	if len(s) >= 2 && ((s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'')) {
		if s[0] == '\'' {
			return s[1 : len(s)-1], nil
		}
		u, err := strconv.Unquote(s)
		if err != nil {
			return s[1 : len(s)-1], nil
		}
		return u, nil
	}
	if strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}") {
		return parseShorthandMapAny(s)
	}
	if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
		return parseShorthandList(s)
	}
	if findTopLevel(s, '=') >= 0 {
		return parseShorthandMapAny(s)
	}
	if strings.EqualFold(s, "true") {
		return true, nil
	}
	if strings.EqualFold(s, "false") {
		return false, nil
	}
	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return i, nil
	}
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f, nil
	}
	return s, nil
}

func parseShorthandList(raw string) ([]any, error) {
	s := strings.TrimSpace(raw)
	if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
		s = strings.TrimSpace(s[1 : len(s)-1])
	}
	if s == "" {
		return []any{}, nil
	}
	items, err := splitTopLevel(s, ',')
	if err != nil {
		return nil, err
	}
	out := make([]any, 0, len(items))
	for _, item := range items {
		v, err := parseShorthandValue(item)
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, nil
}

func splitTopLevel(s string, sep rune) ([]string, error) {
	parts := []string{}
	var b strings.Builder
	depthBrace := 0
	depthBracket := 0
	quote := rune(0)
	escape := false
	for _, r := range s {
		if escape {
			b.WriteRune(r)
			escape = false
			continue
		}
		if quote != 0 {
			if r == '\\' {
				escape = true
				b.WriteRune(r)
				continue
			}
			if r == quote {
				quote = 0
			}
			b.WriteRune(r)
			continue
		}
		switch r {
		case '\'', '"':
			quote = r
			b.WriteRune(r)
		case '{':
			depthBrace++
			b.WriteRune(r)
		case '}':
			depthBrace--
			if depthBrace < 0 {
				return nil, fmt.Errorf("unbalanced '}' in %q", s)
			}
			b.WriteRune(r)
		case '[':
			depthBracket++
			b.WriteRune(r)
		case ']':
			depthBracket--
			if depthBracket < 0 {
				return nil, fmt.Errorf("unbalanced ']' in %q", s)
			}
			b.WriteRune(r)
		default:
			if r == sep && depthBrace == 0 && depthBracket == 0 {
				parts = append(parts, b.String())
				b.Reset()
				continue
			}
			b.WriteRune(r)
		}
	}
	if quote != 0 || depthBrace != 0 || depthBracket != 0 {
		return nil, fmt.Errorf("unterminated shorthand value %q", s)
	}
	parts = append(parts, b.String())
	return parts, nil
}

func findTopLevel(s string, target rune) int {
	depthBrace := 0
	depthBracket := 0
	quote := rune(0)
	escape := false
	for i, r := range s {
		if escape {
			escape = false
			continue
		}
		if quote != 0 {
			if r == '\\' {
				escape = true
				continue
			}
			if r == quote {
				quote = 0
			}
			continue
		}
		switch r {
		case '\'', '"':
			quote = r
		case '{':
			depthBrace++
		case '}':
			depthBrace--
		case '[':
			depthBracket++
		case ']':
			depthBracket--
		default:
			if r == target && depthBrace == 0 && depthBracket == 0 {
				return i
			}
		}
	}
	return -1
}

func normalizeType(t string) string {
	s := strings.TrimSpace(t)
	for strings.HasPrefix(s, "*") {
		s = strings.TrimPrefix(s, "*")
	}
	return s
}

func coerceJSONArray(elemType string, arr []any) ([]any, error) {
	out := make([]any, 0, len(arr))
	for _, item := range arr {
		v, err := coerceJSONValue(elemType, item)
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, nil
}

func coerceJSONValue(typ string, in any) (any, error) {
	t := normalizeType(typ)
	if strings.HasPrefix(t, "[]") {
		elem := strings.TrimPrefix(t, "[]")
		switch v := in.(type) {
		case []any:
			return coerceJSONArray(elem, v)
		case string:
			sv, err := coerceScalar(elem, v)
			if err != nil {
				return nil, err
			}
			return []any{sv}, nil
		default:
			return []any{v}, nil
		}
	}
	switch v := in.(type) {
	case string:
		return coerceScalar(t, v)
	default:
		return in, nil
	}
}

func coerceScalar(typ, raw string) (any, error) {
	t := normalizeType(typ)
	switch {
	case strings.Contains(t, "time.Time"):
		tt, err := time.Parse(time.RFC3339, raw)
		if err != nil {
			return nil, err
		}
		return tt.Format(time.RFC3339), nil
	case strings.Contains(t, "int32"), strings.Contains(t, "int64"), t == "int":
		n, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return nil, err
		}
		return n, nil
	case strings.Contains(t, "float32"), strings.Contains(t, "float64"):
		f, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return nil, err
		}
		return f, nil
	case strings.Contains(t, "bool"):
		v, err := strconv.ParseBool(raw)
		if err != nil {
			return nil, err
		}
		return v, nil
	case strings.HasPrefix(t, "map["):
		if strings.HasPrefix(strings.TrimSpace(raw), "{") {
			var obj map[string]any
			if err := stdjson.Unmarshal([]byte(raw), &obj); err == nil {
				return obj, nil
			}
		}
		return parseShorthandMap(raw), nil
	case strings.HasPrefix(t, "types."):
		trimmed := strings.TrimSpace(raw)
		if strings.HasPrefix(trimmed, "{") {
			var obj map[string]any
			if err := stdjson.Unmarshal([]byte(trimmed), &obj); err == nil {
				return obj, nil
			}
			return parseShorthandMap(raw), nil
		}
		if findTopLevel(trimmed, '=') >= 0 {
			return parseShorthandMap(raw), nil
		}
		return raw, nil
	default:
		return raw, nil
	}
}

func ReadInputJSON(inlineJSON, filePath string) (map[string]any, error) {
	raw := normalizeInputJSON(inlineJSON)
	if filePath != "" {
		b, err := os.ReadFile(filepath.Clean(filePath))
		if err != nil {
			return nil, err
		}
		raw = normalizeInputJSON(string(b))
	}
	if raw == "" {
		return map[string]any{}, nil
	}
	var out map[string]any
	dec := stdjson.NewDecoder(bytes.NewBufferString(raw))
	dec.UseNumber()
	if err := dec.Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

func normalizeInputJSON(raw string) string {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return ""
	}
	if stdjson.Valid([]byte(trimmed)) {
		return trimmed
	}
	unescapedQuotes := strings.ReplaceAll(trimmed, `\"`, `"`)
	if stdjson.Valid([]byte(unescapedQuotes)) {
		return unescapedQuotes
	}
	if u, err := strconv.Unquote(trimmed); err == nil && stdjson.Valid([]byte(u)) {
		return u
	}
	return trimmed
}

func normalizeNilSlices(v any) any {
	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return v
	}
	normalized := normalizeNilSlicesValue(rv)
	if !normalized.IsValid() {
		return v
	}
	return normalized.Interface()
}

func normalizeNilSlicesValue(v reflect.Value) reflect.Value {
	if !v.IsValid() {
		return v
	}
	if isJSONMarshaler(v) {
		return v
	}
	switch v.Kind() {
	case reflect.Interface:
		if v.IsNil() {
			return v
		}
		return normalizeNilSlicesValue(v.Elem())
	case reflect.Ptr:
		if v.IsNil() {
			return v
		}
		elem := normalizeNilSlicesValue(v.Elem())
		out := reflect.New(v.Type().Elem())
		setNormalizedValue(out.Elem(), elem)
		return out
	case reflect.Struct:
		out := reflect.New(v.Type()).Elem()
		for i := 0; i < v.NumField(); i++ {
			dst := out.Field(i)
			if !dst.CanSet() {
				continue
			}
			src := v.Field(i)
			if !src.CanInterface() {
				continue
			}
			setNormalizedValue(dst, normalizeNilSlicesValue(src))
		}
		return out
	case reflect.Slice:
		if v.IsNil() {
			return reflect.MakeSlice(v.Type(), 0, 0)
		}
		out := reflect.MakeSlice(v.Type(), v.Len(), v.Len())
		for i := 0; i < v.Len(); i++ {
			setNormalizedValue(out.Index(i), normalizeNilSlicesValue(v.Index(i)))
		}
		return out
	case reflect.Array:
		out := reflect.New(v.Type()).Elem()
		for i := 0; i < v.Len(); i++ {
			setNormalizedValue(out.Index(i), normalizeNilSlicesValue(v.Index(i)))
		}
		return out
	case reflect.Map:
		if v.IsNil() {
			return v
		}
		out := reflect.MakeMapWithSize(v.Type(), v.Len())
		iter := v.MapRange()
		for iter.Next() {
			val := normalizeNilSlicesValue(iter.Value())
			if val.IsValid() && val.Type().AssignableTo(v.Type().Elem()) {
				out.SetMapIndex(iter.Key(), val)
			} else if val.IsValid() && val.Type().ConvertibleTo(v.Type().Elem()) {
				out.SetMapIndex(iter.Key(), val.Convert(v.Type().Elem()))
			} else {
				out.SetMapIndex(iter.Key(), iter.Value())
			}
		}
		return out
	default:
		return v
	}
}

func isJSONMarshaler(v reflect.Value) bool {
	marshalerType := reflect.TypeOf((*stdjson.Marshaler)(nil)).Elem()
	if v.CanInterface() && v.Type().Implements(marshalerType) {
		return true
	}
	if v.CanAddr() && v.Addr().CanInterface() && v.Addr().Type().Implements(marshalerType) {
		return true
	}
	return false
}

func setNormalizedValue(dst, src reflect.Value) {
	if !src.IsValid() || !dst.CanSet() {
		return
	}
	if src.Type().AssignableTo(dst.Type()) {
		dst.Set(src)
		return
	}
	if src.Type().ConvertibleTo(dst.Type()) {
		dst.Set(src.Convert(dst.Type()))
		return
	}
	if dst.Kind() == reflect.Interface && src.Type().Implements(dst.Type()) {
		dst.Set(src)
	}
}

func writeCSV(out *bytes.Buffer, headers []string, rows [][]string) {
	w := csv.NewWriter(out)
	if len(headers) > 0 {
		_ = w.Write(headers)
	}
	_ = w.WriteAll(rows)
	w.Flush()
}

func writeTable(out *bytes.Buffer, headers []string, rows [][]string) {
	tw := tabwriter.NewWriter(out, 0, 4, 2, ' ', 0)
	if len(headers) > 0 {
		_, _ = fmt.Fprintln(tw, strings.Join(headers, "\t"))
	}
	for _, row := range rows {
		_, _ = fmt.Fprintln(tw, strings.Join(row, "\t"))
	}
	_ = tw.Flush()
}

func writeMarkdown(out *bytes.Buffer, headers []string, rows [][]string) {
	if len(headers) == 0 {
		return
	}
	escape := func(s string) string {
		s = strings.ReplaceAll(s, "\\", "\\\\")
		s = strings.ReplaceAll(s, "|", "\\|")
		s = strings.ReplaceAll(s, "\n", " ")
		s = strings.ReplaceAll(s, "\r", "")
		return s
	}
	escapedHeaders := make([]string, len(headers))
	for i, h := range headers {
		escapedHeaders[i] = escape(h)
	}
	_, _ = fmt.Fprintf(out, "| %s |\n", strings.Join(escapedHeaders, " | "))
	sep := make([]string, len(headers))
	for i := range headers {
		sep[i] = "---"
	}
	_, _ = fmt.Fprintf(out, "| %s |\n", strings.Join(sep, " | "))
	for _, row := range rows {
		escaped := make([]string, len(headers))
		for i := range headers {
			val := ""
			if i < len(row) {
				val = row[i]
			}
			escaped[i] = escape(val)
		}
		_, _ = fmt.Fprintf(out, "| %s |\n", strings.Join(escaped, " | "))
	}
}

func writeHTML(out *bytes.Buffer, headers []string, rows [][]string) {
	var b strings.Builder
	b.WriteString("<!doctype html>\n<html lang=\"en\">\n<head>\n")
	b.WriteString("<meta charset=\"utf-8\"/>\n")
	b.WriteString("<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"/>\n")
	b.WriteString("<title>Output</title>\n")
	b.WriteString("<style>\n")
	b.WriteString("  body { font-family: system-ui, -apple-system, Segoe UI, Roboto, sans-serif; }\n")
	b.WriteString("  table { border-collapse: collapse; width: 100%; }\n")
	b.WriteString("  th, td { padding: 8px 12px; border: 1px solid #ddd; }\n")
	b.WriteString("  thead th { background-color: #222; color: #fff; text-align: left; }\n")
	b.WriteString("  tbody tr:nth-child(even) { background-color: #f6f6f6; }\n")
	b.WriteString("</style>\n")
	b.WriteString("</head>\n<body>\n<table>\n")
	if len(headers) > 0 {
		b.WriteString("<thead><tr>")
		for _, h := range headers {
			b.WriteString("<th>")
			b.WriteString(html.EscapeString(h))
			b.WriteString("</th>")
		}
		b.WriteString("</tr></thead>\n")
	}
	b.WriteString("<tbody>\n")
	for _, row := range rows {
		b.WriteString("<tr>")
		for i := range headers {
			val := ""
			if i < len(row) {
				val = row[i]
			}
			b.WriteString("<td>")
			b.WriteString(html.EscapeString(val))
			b.WriteString("</td>")
		}
		b.WriteString("</tr>\n")
	}
	b.WriteString("</tbody>\n</table>\n</body>\n</html>\n")
	out.WriteString(b.String())
}

func buildOutputRows(v any) ([]string, [][]string) {
	switch x := v.(type) {
	case []any:
		if headers, rows, ok := rowsFromListOfMaps(x); ok {
			return headers, rows
		}
		return []string{"Value"}, scalarRows(x)
	case map[string]any:
		if allScalars(x) {
			keys := sortedMapKeys(x)
			rows := make([][]string, 0, len(keys))
			for _, k := range keys {
				rows = append(rows, []string{k, fmt.Sprint(x[k])})
			}
			return []string{"Key", "Value"}, rows
		}
		var rows [][]string
		flattenRows("", x, &rows)
		return []string{"Path", "Value"}, rows
	default:
		return []string{"Value"}, [][]string{{fmt.Sprint(x)}}
	}
}

func rowsFromListOfMaps(arr []any) ([]string, [][]string, bool) {
	if len(arr) == 0 {
		return []string{"Value"}, nil, true
	}
	keySet := map[string]struct{}{}
	objs := make([]map[string]any, 0, len(arr))
	for _, item := range arr {
		obj, ok := item.(map[string]any)
		if !ok || !allScalars(obj) {
			return nil, nil, false
		}
		objs = append(objs, obj)
		for k := range obj {
			keySet[k] = struct{}{}
		}
	}
	headers := make([]string, 0, len(keySet))
	for k := range keySet {
		headers = append(headers, k)
	}
	sort.Strings(headers)
	rows := make([][]string, 0, len(objs))
	for _, obj := range objs {
		row := make([]string, 0, len(headers))
		for _, h := range headers {
			row = append(row, fmt.Sprint(obj[h]))
		}
		rows = append(rows, row)
	}
	return headers, rows, true
}

func scalarRows(arr []any) [][]string {
	rows := make([][]string, 0, len(arr))
	for _, item := range arr {
		rows = append(rows, []string{scalarString(item)})
	}
	return rows
}

func scalarString(v any) string {
	switch vv := v.(type) {
	case nil:
		return ""
	case map[string]any, []any:
		b, _ := stdjson.Marshal(vv)
		return string(b)
	default:
		return fmt.Sprint(v)
	}
}

func allScalars(m map[string]any) bool {
	for _, v := range m {
		switch v.(type) {
		case nil, bool, string, float64, float32, int, int32, int64, uint, uint32, uint64, stdjson.Number:
		default:
			return false
		}
	}
	return true
}

func flattenRows(path string, v any, rows *[][]string) {
	switch x := v.(type) {
	case map[string]any:
		keys := sortedMapKeys(x)
		for _, k := range keys {
			next := k
			if path != "" {
				next = path + "." + k
			}
			flattenRows(next, x[k], rows)
		}
	case []any:
		for i, item := range x {
			next := fmt.Sprintf("%s[%d]", path, i)
			if path == "" {
				next = fmt.Sprintf("[%d]", i)
			}
			flattenRows(next, item, rows)
		}
	default:
		*rows = append(*rows, []string{path, fmt.Sprint(x)})
	}
}

func sortedMapKeys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
