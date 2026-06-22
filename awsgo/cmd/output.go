package cmd

import (
	"bytes"
	"encoding/csv"
	stdjson "encoding/json"
	"fmt"
	"html"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"text/tabwriter"
	"time"

	"gopkg.in/yaml.v3"
)

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

func normalizeAWSCLIOutputRaw(raw []byte) []byte {
	trimmed := bytes.TrimSpace(raw)
	if len(trimmed) == 0 {
		return raw
	}
	var v any
	if err := stdjson.Unmarshal(trimmed, &v); err != nil {
		return raw
	}
	v = normalizeAWSCLIOutputDocument(v)
	b, err := stdjson.MarshalIndent(v, "", "  ")
	if err != nil {
		return raw
	}
	return b
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
	sortStrings(headers)
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
	sortStrings(keys)
	return keys
}

func sortStrings(v []string) {
	sort.Strings(v)
}
