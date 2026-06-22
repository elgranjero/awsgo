package svcruntime

import (
	"bytes"
	"encoding/csv"
	stdjson "encoding/json"
	"fmt"
	"html"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/goccy/go-json"
	"github.com/jmespath/go-jmespath"
	"gopkg.in/yaml.v3"
)

// ServiceDef is the per-service metadata needed to expose AWS-style operation commands
// without linking every generated service into each split service binary.
type ServiceDef struct {
	Operations             []string
	OperationSet           map[string]bool
	OperationInputs        map[string][]string
	OperationInputTypes    map[string]map[string]string
	OperationInputRequired map[string][]string
	Run                    func(args []string) error
}

type state struct {
	outputFormat string
	awsProfile   string
	awsRegion    string
	queryExpr    string
	verbose      bool
	inputJSON    string
	inputFile    string
	outputSet    bool
}

func ExecuteService(service string, svc ServiceDef, args []string) error {
	st := &state{}
	if len(args) == 0 || isRootHelp(args[0]) {
		printServiceHelp(service, svc)
		return nil
	}
	if args[0] == "help" {
		if len(args) > 1 && svc.OperationSet[args[1]] {
			printOperationHelp(service, svc, args[1])
			return nil
		}
		printServiceHelp(service, svc)
		return nil
	}

	operation := args[0]
	passthrough := args[1:]
	if strings.HasPrefix(operation, "--") {
		candidate := strings.TrimPrefix(operation, "--")
		if svc.OperationSet[candidate] {
			operation = candidate
		} else {
			return fmt.Errorf("unknown operation %q for service %q", operation, service)
		}
	} else if !svc.OperationSet[operation] {
		return fmt.Errorf("unknown operation %q for service %q", operation, service)
	}

	passthrough = st.consumeGlobalFlags(passthrough)
	fieldOverrides, passthrough, err := consumeOperationFlags(svc, operation, passthrough)
	if err != nil {
		return err
	}
	if hasHelpFlag(passthrough) {
		printOperationHelp(service, svc, operation)
		return nil
	}

	effectiveOutput := st.resolveOutputFormat()
	forwarded := make([]string, 0, len(passthrough)+5)
	forwarded = append(forwarded, "--"+operation)
	forwarded = append(forwarded, "--output", "json")
	if st.awsProfile != "" && !hasLongFlag(passthrough, "profile") {
		forwarded = append(forwarded, "--profile", st.awsProfile)
	}
	if st.awsRegion != "" && !hasLongFlag(passthrough, "region") {
		forwarded = append(forwarded, "--region", st.awsRegion)
	}
	forwarded = append(forwarded, passthrough...)
	st.debugf("service=%s operation=%s forwarded_args=%q", service, operation, forwarded)

	restoreEnv := st.setInputEnv(fieldOverrides, svc.OperationInputTypes[operation])
	defer restoreEnv()
	if strings.TrimSpace(st.inputJSON) != "" || strings.TrimSpace(st.inputFile) != "" || len(fieldOverrides) > 0 {
		st.debugf("service=%s operation=%s input_override=true", service, operation)
	}

	raw, stderr, err := captureOutput(func() error {
		return svc.Run(forwarded)
	})
	if len(stderr) > 0 {
		_, _ = os.Stderr.Write(stderr)
	}
	if err != nil {
		st.debugf("service=%s operation=%s run_error=%v", service, operation, err)
		return err
	}
	if bytes.Contains(stderr, []byte("level=error")) {
		st.debugf("service=%s operation=%s detected_error_log=true", service, operation)
		return fmt.Errorf("service reported an error")
	}
	if len(bytes.TrimSpace(raw)) == 0 {
		st.debugf("service=%s operation=%s empty_stdout=true", service, operation)
		return fmt.Errorf("operation produced no output")
	}

	normalizedRaw := normalizeAWSCLIOutputRaw(raw)
	queried, err := applyQuery(normalizedRaw, st.queryExpr)
	if err != nil {
		return err
	}
	formatted, err := formatOutput(queried, effectiveOutput)
	if err != nil {
		return err
	}
	if len(formatted) > 0 {
		_, _ = os.Stdout.Write(formatted)
		if formatted[len(formatted)-1] != '\n' {
			_, _ = os.Stdout.Write([]byte("\n"))
		}
	}
	return nil
}

func isRootHelp(arg string) bool {
	return arg == "--help" || arg == "-h"
}

func (st *state) debugf(format string, args ...any) {
	if !st.verbose {
		return
	}
	_, _ = fmt.Fprintf(os.Stderr, "verbose: "+format+"\n", args...)
}

func hasLongFlag(args []string, name string) bool {
	prefix := "--" + name
	for _, arg := range args {
		if arg == prefix || strings.HasPrefix(arg, prefix+"=") {
			return true
		}
	}
	return false
}

func hasHelpFlag(args []string) bool {
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			return true
		}
	}
	return false
}

func (st *state) resolveOutputFormat() string {
	if st.outputSet {
		if v := strings.TrimSpace(st.outputFormat); v != "" {
			return v
		}
	}
	if v := strings.TrimSpace(os.Getenv("AWS_DEFAULT_OUTPUT")); v != "" {
		return v
	}
	if v := sharedConfigProfileOutput(st.awsProfile); v != "" {
		return v
	}
	return "json"
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

func (st *state) consumeGlobalFlags(args []string) []string {
	kept := make([]string, 0, len(args))
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == "--verbose":
			st.verbose = true
		case strings.HasPrefix(arg, "--verbose="):
			v := strings.TrimPrefix(arg, "--verbose=")
			st.verbose = v == "1" || strings.EqualFold(v, "true")
		case arg == "--profile":
			if i+1 < len(args) {
				st.awsProfile = args[i+1]
				i++
			}
		case strings.HasPrefix(arg, "--profile="):
			st.awsProfile = strings.TrimPrefix(arg, "--profile=")
		case arg == "--region":
			if i+1 < len(args) {
				st.awsRegion = args[i+1]
				i++
			}
		case strings.HasPrefix(arg, "--region="):
			st.awsRegion = strings.TrimPrefix(arg, "--region=")
		case arg == "--output" || arg == "-o":
			if i+1 < len(args) {
				st.outputFormat = args[i+1]
				st.outputSet = true
				i++
			}
		case strings.HasPrefix(arg, "--output="):
			st.outputFormat = strings.TrimPrefix(arg, "--output=")
			st.outputSet = true
		case arg == "--query":
			if i+1 < len(args) {
				st.queryExpr = args[i+1]
				i++
			}
		case strings.HasPrefix(arg, "--query="):
			st.queryExpr = strings.TrimPrefix(arg, "--query=")
		case arg == "--input-json" || arg == "--cli-input-json":
			if i+1 < len(args) {
				st.inputJSON = args[i+1]
				i++
			}
		case strings.HasPrefix(arg, "--input-json="):
			st.inputJSON = strings.TrimPrefix(arg, "--input-json=")
		case strings.HasPrefix(arg, "--cli-input-json="):
			st.inputJSON = strings.TrimPrefix(arg, "--cli-input-json=")
		case arg == "--input-file":
			if i+1 < len(args) {
				st.inputFile = args[i+1]
				i++
			}
		case strings.HasPrefix(arg, "--input-file="):
			st.inputFile = strings.TrimPrefix(arg, "--input-file=")
		default:
			kept = append(kept, arg)
		}
	}
	return kept
}

func consumeOperationFlags(svc ServiceDef, operation string, args []string) (map[string]any, []string, error) {
	fieldTypes := svc.OperationInputTypes[operation]
	if len(fieldTypes) == 0 {
		return nil, args, nil
	}
	flagToField := map[string]string{}
	for field := range fieldTypes {
		flagName := toKebab(field)
		if isReservedOperationFlag(flagName) {
			continue
		}
		flagToField[flagName] = field
	}

	kept := make([]string, 0, len(args))
	overrides := map[string]any{}
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if !strings.HasPrefix(arg, "--") || arg == "--" {
			kept = append(kept, arg)
			continue
		}
		nameVal := strings.TrimPrefix(arg, "--")
		name := nameVal
		value := ""
		hasValue := false
		if idx := strings.Index(nameVal, "="); idx >= 0 {
			name = nameVal[:idx]
			value = nameVal[idx+1:]
			hasValue = true
		}
		field, ok := flagToField[name]
		if !ok {
			kept = append(kept, arg)
			continue
		}
		if !hasValue {
			if i+1 >= len(args) {
				return nil, nil, fmt.Errorf("missing value for --%s", name)
			}
			value = args[i+1]
			i++
		}
		typed, err := coerceValue(fieldTypes[field], value)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid value for --%s: %w", name, err)
		}
		if strings.HasPrefix(fieldTypes[field], "[]") {
			prev, exists := overrides[field]
			if !exists {
				overrides[field] = typed
				continue
			}
			prevArr, ok := prev.([]any)
			if !ok {
				return nil, nil, fmt.Errorf("internal type error for field %s", field)
			}
			nextArr, ok := typed.([]any)
			if !ok {
				return nil, nil, fmt.Errorf("internal type error for field %s", field)
			}
			overrides[field] = append(prevArr, nextArr...)
			continue
		}
		overrides[field] = typed
	}
	if len(overrides) == 0 {
		return nil, kept, nil
	}
	return overrides, kept, nil
}

func isReservedOperationFlag(flagName string) bool {
	switch flagName {
	case "help", "h",
		"profile", "region", "output", "o", "query", "verbose",
		"input-json", "cli-input-json", "input-file":
		return true
	default:
		return false
	}
}

func toKebab(s string) string {
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
		if p != "" {
			if p == "qindex" {
				out = append(out, "q", "index")
				continue
			}
			out = append(out, p)
		}
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

func parseShorthandMap(v string) map[string]any {
	m, _ := parseShorthandMapAny(v)
	return m
}

func parseShorthandMapAny(v string) (map[string]any, error) {
	raw := strings.TrimSpace(v)
	if raw == "" {
		return map[string]any{}, nil
	}
	if strings.HasPrefix(raw, "{") && strings.HasSuffix(raw, "}") {
		raw = strings.TrimSpace(raw[1 : len(raw)-1])
	}
	items, err := splitTopLevel(raw, ',')
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

func printServiceHelp(service string, svc ServiceDef) {
	fmt.Println("Dynamic AWS-style CLI over generated SDK commands")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Printf("  awsgo %s <operation> [flags]\n", service)
	fmt.Println()
	fmt.Println("Available Operations:")
	ops := append([]string(nil), svc.Operations...)
	sort.Strings(ops)
	for _, op := range ops {
		fmt.Printf("  %s\n", op)
	}
}

func printOperationHelp(service string, svc ServiceDef, operation string) {
	_, _ = fmt.Fprintf(os.Stdout, "Usage:\n  aws %s %s [flags]\n\n", service, operation)
	_, _ = fmt.Fprintln(os.Stdout, "Global Flags:")
	_, _ = fmt.Fprintln(os.Stdout, "      --profile string   AWS shared config profile")
	_, _ = fmt.Fprintln(os.Stdout, "      --region string    AWS region")
	_, _ = fmt.Fprintln(os.Stdout, "      --output string    Output format: json|yaml|text|table|csv|markdown|html")
	_, _ = fmt.Fprintln(os.Stdout, "      --query string     Query JSON output (JMESPath)")
	_, _ = fmt.Fprintln(os.Stdout, "      --input-json str   Inline JSON object merged into SDK input")
	_, _ = fmt.Fprintln(os.Stdout, "      --cli-input-json   Alias of --input-json")
	_, _ = fmt.Fprintln(os.Stdout, "      --input-file path  JSON file merged into SDK input")
	_, _ = fmt.Fprintln(os.Stdout, "      --verbose          Print forwarded arguments and diagnostics")
	_, _ = fmt.Fprintln(os.Stdout, "  -h, --help             help for this operation")
	fieldHelp := knownInputFlagHelp(svc, operation)
	if len(fieldHelp) > 0 {
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = fmt.Fprintln(os.Stdout, "Operation Flags:")
		for _, line := range fieldHelp {
			_, _ = fmt.Fprintln(os.Stdout, line)
		}
	}
	_, _ = fmt.Fprintln(os.Stdout)
	_, _ = fmt.Fprintln(os.Stdout, "Notes:")
	_, _ = fmt.Fprintln(os.Stdout, "  Operation input flags are generated from service code and may be incomplete.")
	_, _ = fmt.Fprintln(os.Stdout, "  Internal selector flags (for example, --describe-images) are not user-facing.")
	fields := knownInputFields(svc, operation)
	types := knownInputTypes(svc, operation)
	if len(fields) > 0 {
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = fmt.Fprintln(os.Stdout, "Known Input Fields:")
		for _, f := range fields {
			_, _ = fmt.Fprintf(os.Stdout, "  - %s\n", f)
		}
		_, _ = fmt.Fprintln(os.Stdout)
		_, _ = fmt.Fprintln(os.Stdout, "Example:")
		_, _ = fmt.Fprintf(os.Stdout, "  aws %s %s --input-json '%s'\n", service, operation, inputJSONTemplate(fields, types))
	}
}

func knownInputFlagHelp(svc ServiceDef, operation string) []string {
	fields := knownInputFields(svc, operation)
	if len(fields) == 0 {
		return nil
	}
	types := knownInputTypes(svc, operation)
	required := knownRequiredInputFieldSet(svc, operation)
	type row struct {
		flag string
		typ  string
		req  string
	}
	rows := make([]row, 0, len(fields))
	maxFlag := 0
	maxType := 0
	for _, field := range fields {
		flagName := toKebab(field)
		if isReservedOperationFlag(flagName) {
			continue
		}
		typ := "string"
		if types != nil {
			if t := strings.TrimSpace(types[field]); t != "" {
				typ = helpType(t)
			}
		}
		reqLabel := "Optional"
		if required[field] {
			reqLabel = "Required"
		}
		r := row{
			flag: "--" + flagName,
			typ:  typ,
			req:  reqLabel,
		}
		if len(r.flag) > maxFlag {
			maxFlag = len(r.flag)
		}
		if len(r.typ) > maxType {
			maxType = len(r.typ)
		}
		rows = append(rows, r)
	}
	lines := make([]string, 0, len(rows))
	for _, r := range rows {
		lines = append(lines, fmt.Sprintf("      %-*s  %-*s  %s", maxFlag, r.flag, maxType, r.typ, r.req))
	}
	return lines
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

func (st *state) setInputEnv(overrides map[string]any, fieldTypes map[string]string) func() {
	prev, hadPrev := os.LookupEnv("AWSGO_INPUT_JSON")
	prevPaginate, hadPrevPaginate := os.LookupEnv("AWSGO_DISABLE_PAGINATOR")
	raw := strings.TrimSpace(st.inputJSON)
	if st.inputFile != "" {
		b, err := os.ReadFile(st.inputFile)
		if err != nil {
			st.debugf("input-file read error: %v", err)
		} else {
			raw = strings.TrimSpace(string(b))
		}
	}
	raw = normalizeInputJSON(raw)
	base := map[string]any{}
	if strings.TrimSpace(raw) != "" {
		_ = stdjson.Unmarshal([]byte(raw), &base)
	}
	for field, typ := range fieldTypes {
		v, ok := base[field]
		if !ok {
			continue
		}
		cv, err := coerceJSONValue(typ, v)
		if err != nil {
			continue
		}
		base[field] = cv
	}
	if len(overrides) > 0 {
		for k, v := range overrides {
			base[k] = v
		}
	}
	if len(base) > 0 {
		b, _ := stdjson.Marshal(base)
		raw = string(b)
	}
	if shouldDisablePaginator(base) {
		_ = os.Setenv("AWSGO_DISABLE_PAGINATOR", "1")
	} else {
		_ = os.Unsetenv("AWSGO_DISABLE_PAGINATOR")
	}
	if raw == "" {
		_ = os.Unsetenv("AWSGO_INPUT_JSON")
	} else {
		_ = os.Setenv("AWSGO_INPUT_JSON", raw)
	}

	return func() {
		if hadPrev {
			_ = os.Setenv("AWSGO_INPUT_JSON", prev)
		} else {
			_ = os.Unsetenv("AWSGO_INPUT_JSON")
		}
		if hadPrevPaginate {
			_ = os.Setenv("AWSGO_DISABLE_PAGINATOR", prevPaginate)
		} else {
			_ = os.Unsetenv("AWSGO_DISABLE_PAGINATOR")
		}
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

func knownInputFields(svc ServiceDef, operation string) []string {
	if svc.OperationInputs == nil {
		return nil
	}
	fields := append([]string(nil), svc.OperationInputs[operation]...)
	sort.Strings(fields)
	return fields
}

func knownInputTypes(svc ServiceDef, operation string) map[string]string {
	if svc.OperationInputTypes == nil {
		return nil
	}
	return svc.OperationInputTypes[operation]
}

func knownRequiredInputFieldSet(svc ServiceDef, operation string) map[string]bool {
	if svc.OperationInputRequired == nil {
		return map[string]bool{}
	}
	fields := svc.OperationInputRequired[operation]
	out := make(map[string]bool, len(fields))
	for _, f := range fields {
		out[f] = true
	}
	return out
}

func inputJSONTemplate(fields []string, types map[string]string) string {
	m := map[string]any{}
	limit := len(fields)
	if limit > 8 {
		limit = 8
	}
	for i := 0; i < limit; i++ {
		field := fields[i]
		typ := ""
		if types != nil {
			typ = types[field]
		}
		m[field] = templateValueForType(typ)
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

func applyQuery(raw []byte, query string) ([]byte, error) {
	q := strings.TrimSpace(query)
	if q == "" {
		return raw, nil
	}

	var v any
	if err := json.Unmarshal(bytes.TrimSpace(raw), &v); err != nil {
		return nil, fmt.Errorf("--query requires JSON output: %w", err)
	}
	if rootSlice, ok := v.([]any); ok && len(rootSlice) == 1 {
		v = rootSlice[0]
	}

	result, err := jmespath.Search(q, v)
	if err != nil {
		return nil, fmt.Errorf("invalid --query: %w", err)
	}
	if result == nil && isProjectionQuery(q) {
		result = []any{}
	}
	b, err := stdjson.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func normalizeAWSCLIOutputRaw(raw []byte) []byte {
	trimmed := bytes.TrimSpace(raw)
	if len(trimmed) == 0 {
		return raw
	}
	var v any
	if err := json.Unmarshal(trimmed, &v); err != nil {
		return raw
	}
	v = normalizeAWSCLIOutputDocument(v)
	b, err := stdjson.MarshalIndent(v, "", "  ")
	if err != nil {
		return raw
	}
	return b
}

func isProjectionQuery(q string) bool {
	return strings.Contains(q, "[*]") || strings.Contains(q, "[]")
}

func captureOutput(run func() error) ([]byte, []byte, error) {
	orig := os.Stdout
	origErr := os.Stderr
	stdoutR, stdoutW, err := os.Pipe()
	if err != nil {
		return nil, nil, err
	}
	stderrR, stderrW, err := os.Pipe()
	if err != nil {
		return nil, nil, err
	}

	os.Stdout = stdoutW
	os.Stderr = stderrW
	outCh := make(chan []byte, 1)
	errOutCh := make(chan []byte, 1)
	errCh := make(chan error, 1)

	go func() {
		var buf bytes.Buffer
		if _, copyErr := io.Copy(&buf, stdoutR); copyErr != nil {
			errCh <- copyErr
			return
		}
		outCh <- buf.Bytes()
	}()
	go func() {
		var buf bytes.Buffer
		if _, copyErr := io.Copy(&buf, stderrR); copyErr != nil {
			errCh <- copyErr
			return
		}
		errOutCh <- buf.Bytes()
	}()

	runErr := run()
	_ = stdoutW.Close()
	_ = stderrW.Close()
	os.Stdout = orig
	os.Stderr = origErr

	select {
	case copyErr := <-errCh:
		if runErr != nil {
			return nil, nil, runErr
		}
		return nil, nil, copyErr
	case out := <-outCh:
		errOut := <-errOutCh
		return out, errOut, runErr
	}
}

func formatOutput(raw []byte, format string) ([]byte, error) {
	outFmt := strings.ToLower(strings.TrimSpace(format))
	if outFmt == "yml" {
		outFmt = "yaml"
	}
	trimmed := bytes.TrimSpace(raw)
	if len(trimmed) == 0 {
		return raw, nil
	}
	var structured any
	if err := json.Unmarshal(trimmed, &structured); err != nil {
		return raw, nil
	}
	structured = normalizeAWSCLIOutputDocument(structured)
	headers, rows := buildOutputRows(structured)
	out, err := writeOutput(headers, rows, structured, outFmt)
	if err != nil {
		return nil, err
	}
	if outFmt == "" {
		return raw, nil
	}
	return out, nil
}

func normalizeAWSCLIOutput(v any) any {
	switch x := v.(type) {
	case []any:
		out := make([]any, 0, len(x))
		for _, item := range x {
			out = append(out, normalizeAWSCLIOutput(item))
		}
		return out
	case map[string]any:
		return normalizeAWSCLIMap(x)
	default:
		return normalizeAWSCLIScalar(x)
	}
}

func normalizeAWSCLIOutputDocument(v any) any {
	if merged, ok := mergeRawPaginatorPages(v); ok {
		return normalizeAWSCLIOutput(merged)
	}
	return normalizeAWSCLIOutput(v)
}

func normalizeAWSCLIMap(m map[string]any) map[string]any {
	out := make(map[string]any, len(m))
	truncated, _ := m["IsTruncated"].(bool)
	for key, val := range m {
		if key == "ResultMetadata" || key == "IsTruncated" {
			continue
		}
		if val == nil {
			continue
		}
		if shouldOmitEmptyNestedList(key, val, m) {
			continue
		}
		if key == "Marker" {
			if truncated {
				out["NextToken"] = normalizeAWSCLIOutput(val)
			}
			continue
		}
		normalized := normalizeAWSCLIOutput(val)
		if s, ok := normalized.(string); ok {
			normalized = normalizeAWSCLIStringField(key, s)
		}
		out[key] = normalized
	}
	return out
}

func shouldOmitEmptyNestedList(key string, val any, parent map[string]any) bool {
	if key != "Tags" || len(parent) <= 1 {
		return false
	}
	list, ok := val.([]any)
	return ok && len(list) == 0
}

func normalizeAWSCLIScalar(v any) any {
	if s, ok := v.(string); ok {
		return normalizeAWSTimestamp(s)
	}
	return v
}

func normalizeAWSCLIStringField(key, value string) any {
	if decoded, ok := decodeAWSJSONDocument(value); ok && looksLikePolicyDocumentField(key) {
		return decoded
	}
	return normalizeAWSTimestamp(value)
}

func looksLikePolicyDocumentField(key string) bool {
	return key == "Document" || strings.HasSuffix(key, "PolicyDocument")
}

func decodeAWSJSONDocument(value string) (any, bool) {
	s := strings.TrimSpace(value)
	if s == "" {
		return nil, false
	}
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
	return normalizeAWSCLIOutput(out), true
}

func normalizeAWSTimestamp(value string) any {
	if !looksLikeTimestamp(value) {
		return value
	}
	t, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		return value
	}
	return t.Format("2006-01-02T15:04:05-07:00")
}

func looksLikeTimestamp(value string) bool {
	if len(value) < len("2006-01-02T15:04:05Z") {
		return false
	}
	return value[4] == '-' && value[7] == '-' && value[10] == 'T' && value[13] == ':' && value[16] == ':'
}

func mergeRawPaginatorPages(v any) (any, bool) {
	items, ok := v.([]any)
	if !ok || len(items) == 0 {
		return v, false
	}
	pages := make([]map[string]any, 0, len(items))
	for _, item := range items {
		page, ok := item.(map[string]any)
		if !ok || !looksLikeRawPaginatorPage(page) {
			return v, false
		}
		pages = append(pages, page)
	}
	if len(pages) == 1 {
		return pages[0], true
	}

	merged := map[string]any{}
	for _, page := range pages {
		for key, val := range page {
			if existing, ok := merged[key]; ok {
				existingList, existingOK := existing.([]any)
				newList, newOK := val.([]any)
				if existingOK && newOK {
					merged[key] = append(existingList, newList...)
					continue
				}
			}
			if key == "NextToken" {
				merged[key] = val
				continue
			}
			if _, ok := merged[key]; !ok {
				merged[key] = val
			}
		}
	}
	return merged, true
}

func looksLikeRawPaginatorPage(page map[string]any) bool {
	if _, ok := page["ResultMetadata"]; ok {
		return true
	}
	if _, ok := page["IsTruncated"]; ok {
		return true
	}
	if _, ok := page["Marker"]; ok {
		return true
	}
	if _, ok := page["NextToken"]; ok {
		return true
	}
	return false
}

func writeOutput(headers []string, rows [][]string, structured any, format string) ([]byte, error) {
	var b bytes.Buffer
	structured = normalizeNilSlices(structured)
	switch strings.ToLower(format) {
	case "md", "markdown":
		writeMarkdown(&b, headers, rows)
	case "csv":
		writeCSV(&b, headers, rows)
	case "table":
		writeTable(&b, headers, rows)
	case "html":
		writeHTML(&b, headers, rows)
	case "text", "yaml":
		data, err := yaml.Marshal(structured)
		if err != nil {
			return nil, err
		}
		b.Write(data)
	case "", "json":
		data, err := stdjson.MarshalIndent(structured, "", "  ")
		if err != nil {
			return nil, err
		}
		b.Write(data)
	default:
		return nil, fmt.Errorf("unsupported --output %q (valid: json|yaml|text|table|csv|markdown|html)", format)
	}
	if b.Len() > 0 && b.Bytes()[b.Len()-1] != '\n' {
		b.WriteByte('\n')
	}
	return b.Bytes(), nil
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
		case nil, bool, string, float64, float32, int, int32, int64, uint, uint32, uint64:
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
