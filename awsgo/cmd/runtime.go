package cmd

import (
	"bytes"
	stdjson "encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/goccy/go-json"
	"github.com/jmespath/go-jmespath"
)

func runServiceOperation(service, operation string, passthrough []string) error {
	svc, ok := serviceRegistry[service]
	if !ok {
		return fmt.Errorf("unknown service %q", service)
	}
	if !svc.OperationSet[operation] {
		return fmt.Errorf("unknown operation %q for service %q", operation, service)
	}
	passthrough = consumeGlobalFlags(passthrough)
	fieldOverrides, passthrough, err := consumeOperationFlags(svc, operation, passthrough)
	if err != nil {
		return err
	}
	if hasHelpFlag(passthrough) {
		printOperationHelp(service, operation)
		return nil
	}

	args := make([]string, 0, len(passthrough)+5)
	args = append(args, "--"+operation)
	if awsProfile != "" && !hasLongFlag(passthrough, "profile") {
		args = append(args, "--profile", awsProfile)
	}
	if awsRegion != "" && !hasLongFlag(passthrough, "region") {
		args = append(args, "--region", awsRegion)
	}
	args = append(args, passthrough...)
	debugf("service=%s operation=%s forwarded_args=%q", service, operation, args)

	restoreEnv := setInputEnv(inputJSON, inputFile, fieldOverrides, svc.OperationInputTypes[operation])
	defer restoreEnv()
	if strings.TrimSpace(inputJSON) != "" || strings.TrimSpace(inputFile) != "" || len(fieldOverrides) > 0 {
		debugf("service=%s operation=%s input_override=true", service, operation)
	}

	raw, stderr, err := captureOutput(func() error {
		return svc.Run(args)
	})
	if len(stderr) > 0 {
		_, _ = os.Stderr.Write(stderr)
	}
	if err != nil {
		debugf("service=%s operation=%s run_error=%v", service, operation, err)
		return err
	}
	if bytes.Contains(stderr, []byte("level=error")) {
		debugf("service=%s operation=%s detected_error_log=true", service, operation)
		return fmt.Errorf("service reported an error")
	}
	if len(bytes.TrimSpace(raw)) == 0 {
		debugf("service=%s operation=%s empty_stdout=true", service, operation)
		return fmt.Errorf("operation produced no output")
	}
	queried, err := applyQuery(raw, queryExpr)
	if err != nil {
		return err
	}

	formatted, err := formatOutput(queried, outputFormat)
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

func debugf(format string, args ...any) {
	if !verbose {
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

func consumeGlobalFlags(args []string) []string {
	kept := make([]string, 0, len(args))
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == "--verbose":
			verbose = true
		case strings.HasPrefix(arg, "--verbose="):
			v := strings.TrimPrefix(arg, "--verbose=")
			verbose = v == "1" || strings.EqualFold(v, "true")
		case arg == "--profile":
			if i+1 < len(args) {
				awsProfile = args[i+1]
				i++
			}
		case strings.HasPrefix(arg, "--profile="):
			awsProfile = strings.TrimPrefix(arg, "--profile=")
		case arg == "--region":
			if i+1 < len(args) {
				awsRegion = args[i+1]
				i++
			}
		case strings.HasPrefix(arg, "--region="):
			awsRegion = strings.TrimPrefix(arg, "--region=")
		case arg == "--output":
			if i+1 < len(args) {
				outputFormat = args[i+1]
				i++
			}
		case strings.HasPrefix(arg, "--output="):
			outputFormat = strings.TrimPrefix(arg, "--output=")
		case arg == "--query":
			if i+1 < len(args) {
				queryExpr = args[i+1]
				i++
			}
		case strings.HasPrefix(arg, "--query="):
			queryExpr = strings.TrimPrefix(arg, "--query=")
		case arg == "--input-json":
			if i+1 < len(args) {
				inputJSON = args[i+1]
				i++
			}
		case strings.HasPrefix(arg, "--input-json="):
			inputJSON = strings.TrimPrefix(arg, "--input-json=")
		case arg == "--cli-input-json":
			if i+1 < len(args) {
				inputJSON = args[i+1]
				i++
			}
		case strings.HasPrefix(arg, "--cli-input-json="):
			inputJSON = strings.TrimPrefix(arg, "--cli-input-json=")
		case arg == "--input-file":
			if i+1 < len(args) {
				inputFile = args[i+1]
				i++
			}
		case strings.HasPrefix(arg, "--input-file="):
			inputFile = strings.TrimPrefix(arg, "--input-file=")
		default:
			kept = append(kept, arg)
		}
	}
	return kept
}

func consumeOperationFlags(svc serviceDef, operation string, args []string) (map[string]any, []string, error) {
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
		"profile", "region", "output", "query", "verbose",
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
		if cur == '_' || cur == '-' {
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

func isUpper(r rune) bool { return r >= 'A' && r <= 'Z' }
func isLower(r rune) bool { return r >= 'a' && r <= 'z' }
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
	t := normalizeType(typ)
	if strings.HasPrefix(t, "[]") {
		elem := strings.TrimPrefix(t, "[]")
		trimmed := strings.TrimSpace(raw)
		if strings.HasPrefix(strings.TrimSpace(raw), "[") {
			var arr []any
			if err := stdjson.Unmarshal([]byte(trimmed), &arr); err == nil {
				return coerceJSONArray(elem, arr)
			}
			if arr, err := parseShorthandList(trimmed); err == nil {
				return coerceJSONArray(elem, arr)
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
	return coerceScalar(t, raw)
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
	case strings.HasPrefix(t, "types."), strings.HasPrefix(t, "map["):
		if strings.HasPrefix(strings.TrimSpace(raw), "{") {
			var obj map[string]any
			if err := stdjson.Unmarshal([]byte(raw), &obj); err == nil {
				return obj, nil
			}
		}
		return parseShorthandMap(raw), nil
	default:
		return raw, nil
	}
}

func printOperationHelp(service, operation string) {
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
	fieldHelp := knownInputFlagHelp(service, operation)
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
	fields := knownInputFields(service, operation)
	types := knownInputTypes(service, operation)
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

func knownInputFlagHelp(service, operation string) []string {
	fields := knownInputFields(service, operation)
	if len(fields) == 0 {
		return nil
	}
	types := knownInputTypes(service, operation)
	required := knownRequiredInputFieldSet(service, operation)
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
	case strings.HasPrefix(t, "types."), strings.HasPrefix(t, "map["):
		return "object"
	default:
		return "string"
	}
}

func setInputEnv(inlineJSON, filePath string, overrides map[string]any, fieldTypes map[string]string) func() {
	prev, hadPrev := os.LookupEnv("AWSGO_INPUT_JSON")
	prevPaginate, hadPrevPaginate := os.LookupEnv("AWSGO_DISABLE_PAGINATOR")
	raw := strings.TrimSpace(inlineJSON)
	if filePath != "" {
		b, err := os.ReadFile(filePath)
		if err != nil {
			debugf("input-file read error: %v", err)
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
	// Common shell pattern: escaped quotes inside heredoc JSON (\"key\").
	unescapedQuotes := strings.ReplaceAll(trimmed, `\"`, `"`)
	if stdjson.Valid([]byte(unescapedQuotes)) {
		return unescapedQuotes
	}
	// Handle a quoted JSON string payload.
	if u, err := strconv.Unquote(trimmed); err == nil && stdjson.Valid([]byte(u)) {
		return u
	}
	return trimmed
}

func knownInputFields(service, operation string) []string {
	svc, ok := serviceRegistry[service]
	if !ok || svc.OperationInputs == nil {
		return nil
	}
	fields := append([]string(nil), svc.OperationInputs[operation]...)
	sort.Strings(fields)
	return fields
}

func knownInputTypes(service, operation string) map[string]string {
	svc, ok := serviceRegistry[service]
	if !ok || svc.OperationInputTypes == nil {
		return nil
	}
	return svc.OperationInputTypes[operation]
}

func knownRequiredInputFieldSet(service, operation string) map[string]bool {
	svc, ok := serviceRegistry[service]
	if !ok || svc.OperationInputRequired == nil {
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
		// preserve source output when response isn't valid JSON
		return raw, nil
	}
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
