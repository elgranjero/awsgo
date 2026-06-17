package cmd

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html"
	"os"
	"reflect"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/smithy-go/middleware"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	_awsProfile string
	_awsRegion  string
	_awsOutput  string
	_rootCmd    = &cobra.Command{Use: "service"}
)

func writeOutput(headers []string, rows [][]string, structured any, format string) {
	var b bytes.Buffer
	switch strings.ToLower(strings.TrimSpace(format)) {
	case "md", "markdown":
		writeMarkdown(&b, headers, rows)
	case "csv":
		writeCSV(&b, headers, rows)
	case "table":
		writeTable(&b, headers, rows)
	case "html":
		writeHTML(&b, headers, rows)
	case "text", "yaml":
		data, _ := yaml.Marshal(structured)
		b.Write(data)
	case "", "json":
		data, _ := json.MarshalIndent(structured, "", "  ")
		b.Write(data)
	default:
		data, _ := json.MarshalIndent(structured, "", "  ")
		b.Write(data)
	}
	if b.Len() > 0 && b.Bytes()[b.Len()-1] != '\n' {
		b.WriteByte('\n')
	}
	_, _ = os.Stdout.Write(b.Bytes())
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

func LoadAWSConfigWithMiddleware(profile string) (aws.Config, error) {
	opts := make([]func(*config.LoadOptions) error, 0, 1)
	if profile != "" {
		opts = append(opts, config.WithSharedConfigProfile(profile))
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(), opts...)
	if err != nil {
		return aws.Config{}, err
	}
	raw := strings.TrimSpace(os.Getenv("AWSGO_INPUT_JSON"))
	if raw != "" {
		cfg.APIOptions = append(cfg.APIOptions, func(stack *middleware.Stack) error {
			return stack.Initialize.Add(inputOverrideMiddleware{rawJSON: raw}, middleware.Before)
		})
	}
	return cfg, nil
}

type inputOverrideMiddleware struct {
	rawJSON string
}

func disablePaginator() bool {
	v := strings.TrimSpace(os.Getenv("AWSGO_DISABLE_PAGINATOR"))
	return v == "1" || strings.EqualFold(v, "true")
}

func (m inputOverrideMiddleware) ID() string {
	return "awsgo-input-json-override"
}

func (m inputOverrideMiddleware) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (middleware.InitializeOutput, middleware.Metadata, error) {
	if in.Parameters != nil && m.rawJSON != "" {
		if err := json.Unmarshal([]byte(m.rawJSON), in.Parameters); err != nil {
			return middleware.InitializeOutput{}, middleware.Metadata{}, err
		}
	}
	return next.HandleInitialize(ctx, in)
}

func assignInputField(input any, fieldName, raw string) error {
	v := reflect.ValueOf(input)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("input must be pointer")
	}
	elem := v.Elem()
	if elem.Kind() != reflect.Struct {
		return fmt.Errorf("input must point to struct")
	}
	f := elem.FieldByName(fieldName)
	if !f.IsValid() || !f.CanSet() {
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return setFromString(f, raw)
}

func setFromString(dst reflect.Value, raw string) error {
	if dst.Kind() == reflect.Ptr {
		if dst.IsNil() {
			dst.Set(reflect.New(dst.Type().Elem()))
		}
		return setFromString(dst.Elem(), raw)
	}
	switch dst.Kind() {
	case reflect.String:
		dst.SetString(raw)
		return nil
	case reflect.Bool:
		v, err := strconv.ParseBool(raw)
		if err != nil {
			return err
		}
		dst.SetBool(v)
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return err
		}
		dst.SetInt(v)
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, err := strconv.ParseUint(raw, 10, 64)
		if err != nil {
			return err
		}
		dst.SetUint(v)
		return nil
	case reflect.Float32, reflect.Float64:
		v, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return err
		}
		dst.SetFloat(v)
		return nil
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array, reflect.Interface:
		return json.Unmarshal([]byte(raw), dst.Addr().Interface())
	default:
		return json.Unmarshal([]byte(raw), dst.Addr().Interface())
	}
}
