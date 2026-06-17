package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type serviceMeta struct {
	Name                   string
	CmdVar                 string
	Operations             []string
	OperationInputs        map[string][]string
	OperationInputTypes    map[string]map[string]string
	OperationInputRequired map[string][]string
}

type sdkOpMeta struct {
	Name         string
	InputFields  []string
	InputTypes   map[string]string
	RequiredOnly []string
}

var (
	cmdVarPattern   = regexp.MustCompile(`var\s+([A-Za-z0-9_]+)\s*=\s*&cobra\.Command\s*{`)
	opFlagRegex     = regexp.MustCompile(`\.Flags\(\)\.BoolVarP\(&[A-Za-z0-9_]+,\s*"([^"]+)"`)
	funcRegex       = regexp.MustCompile(`^func\s+[A-Za-z0-9_]+_([A-Za-z0-9]+)\s*\(`)
	inputFieldRegex = regexp.MustCompile(`^([A-Za-z0-9_]+)\s*:\s*([^,]+),`)
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	root, err := findRepoRoot(wd)
	if err != nil {
		panic(err)
	}

	modulePath, err := readModulePath(filepath.Join(root, "go.mod"))
	if err != nil {
		panic(err)
	}

	services, err := discoverServices(root)
	if err != nil {
		panic(err)
	}

	for _, svc := range services {
		if err := writeAdapter(filepath.Join(root, "generated", svc.Name, "cmd"), svc.CmdVar); err != nil {
			panic(err)
		}
	}

	out := filepath.Join(root, "awsgo", "cmd", "registry_gen.go")
	if err := writeRegistry(out, modulePath, services); err != nil {
		panic(err)
	}
	if err := writeManifest(filepath.Join(root, "awsgo", "cmd", "manifest_gen.json"), services); err != nil {
		panic(err)
	}
	if err := writeSplitEntrypoints(root, modulePath, services); err != nil {
		panic(err)
	}

	fmt.Printf("generated awsgo registry for %d services\n", len(services))
}

func discoverServices(root string) ([]serviceMeta, error) {
	serviceRoot := filepath.Join(root, "generated")
	entries, err := os.ReadDir(serviceRoot)
	if err != nil {
		return nil, err
	}

	services := make([]serviceMeta, 0, 512)
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		if shouldSkipDir(name) {
			continue
		}

		cmdFile := filepath.Join(root, "generated", name, "cmd", name+".go")
		if _, err := os.Stat(cmdFile); err != nil {
			continue
		}

		content, err := os.ReadFile(cmdFile)
		if err != nil {
			continue
		}

		cmdVarMatch := cmdVarPattern.FindSubmatch(content)
		if len(cmdVarMatch) < 2 {
			continue
		}
		cmdVar := string(cmdVarMatch[1])

		opMatches := opFlagRegex.FindAllSubmatch(content, -1)
		opSet := map[string]struct{}{}
		for _, m := range opMatches {
			if len(m) >= 2 {
				opSet[string(m[1])] = struct{}{}
			}
		}
		rawOps := make([]string, 0, len(opSet))
		for op := range opSet {
			rawOps = append(rawOps, op)
		}
		sort.Strings(rawOps)

		opInputs := map[string][]string{}
		opInputTypes := map[string]map[string]string{}
		opInputRequired := map[string][]string{}
		ops := []string{}

		sdkOps := loadSDKOperations(sdkServiceRoot(root), name)
		if len(sdkOps) > 0 {
			rawByKey := map[string]string{}
			for _, raw := range rawOps {
				rawByKey[canonicalOpKey(raw)] = raw
			}
			for _, meta := range sdkOps {
				_, ok := rawByKey[canonicalOpKey(meta.Name)]
				if !ok {
					continue
				}
				ops = append(ops, meta.Name)
				opInputs[meta.Name] = meta.InputFields
				opInputTypes[meta.Name] = meta.InputTypes
				opInputRequired[meta.Name] = meta.RequiredOnly
			}
			sort.Strings(ops)
		}

		if len(ops) == 0 {
			ops = append(ops, rawOps...)
			sort.Strings(ops)

			rawInputs := extractOperationInputs(string(content), rawOps)
			for raw, fields := range rawInputs {
				opInputs[raw] = fields
			}
			rawInputTypes := extractOperationInputTypes(string(content), rawOps)
			for raw, fields := range rawInputTypes {
				opInputTypes[raw] = fields
			}
			rawRequired := extractOperationRequiredInputs(string(content), rawOps)
			for raw, fields := range rawRequired {
				opInputRequired[raw] = fields
			}
		}

		services = append(services, serviceMeta{
			Name:                   name,
			CmdVar:                 cmdVar,
			Operations:             ops,
			OperationInputs:        opInputs,
			OperationInputTypes:    opInputTypes,
			OperationInputRequired: opInputRequired,
		})
	}

	sort.Slice(services, func(i, j int) bool { return services[i].Name < services[j].Name })
	return services, nil
}

func sdkServiceRoot(repoRoot string) string {
	if v := strings.TrimSpace(os.Getenv("AWS_SDK_V2_SERVICE_DIR")); v != "" {
		return v
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, "git", "aws-sdk-go-v2", "service")
}

func loadSDKOperations(serviceRoot, service string) []sdkOpMeta {
	if serviceRoot == "" {
		return nil
	}
	pattern := filepath.Join(serviceRoot, service, "api_op_*.go")
	files, err := filepath.Glob(pattern)
	if err != nil || len(files) == 0 {
		return nil
	}
	sort.Strings(files)

	fset := token.NewFileSet()
	out := make([]sdkOpMeta, 0, len(files))
	for _, p := range files {
		base := filepath.Base(p)
		if !strings.HasPrefix(base, "api_op_") || !strings.HasSuffix(base, ".go") {
			continue
		}
		opPascal := strings.TrimSuffix(strings.TrimPrefix(base, "api_op_"), ".go")
		inputType := opPascal + "Input"

		parsed, err := parser.ParseFile(fset, p, nil, parser.ParseComments)
		if err != nil {
			continue
		}

		fields := []string{}
		required := []string{}
		fieldTypes := map[string]string{}
		for _, decl := range parsed.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if !ok || gen.Tok != token.TYPE {
				continue
			}
			for _, spec := range gen.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok || ts.Name == nil || ts.Name.Name != inputType {
					continue
				}
				st, ok := ts.Type.(*ast.StructType)
				if !ok || st.Fields == nil {
					continue
				}
				for _, f := range st.Fields.List {
					if len(f.Names) == 0 {
						continue
					}
					isRequired := commentContainsRequired(f.Doc) || commentContainsRequired(f.Comment)
					typ := exprString(fset, f.Type)
					for _, n := range f.Names {
						if n == nil {
							continue
						}
						fields = append(fields, n.Name)
						fieldTypes[n.Name] = typ
						if isRequired {
							required = append(required, n.Name)
						}
					}
				}
			}
		}
		sort.Strings(fields)
		sort.Strings(required)
		out = append(out, sdkOpMeta{
			Name:         kebabFromPascal(opPascal),
			InputFields:  fields,
			InputTypes:   fieldTypes,
			RequiredOnly: required,
		})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

func exprString(fset *token.FileSet, e ast.Expr) string {
	var b bytes.Buffer
	if err := format.Node(&b, fset, e); err != nil {
		return ""
	}
	return strings.TrimSpace(b.String())
}

func commentContainsRequired(cg *ast.CommentGroup) bool {
	if cg == nil {
		return false
	}
	return strings.Contains(cg.Text(), "This member is required.")
}

func kebabFromPascal(s string) string {
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

func canonicalOpKey(s string) string {
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			if r >= 'A' && r <= 'Z' {
				r = r - 'A' + 'a'
			}
			b.WriteRune(r)
		}
	}
	return b.String()
}

func shouldSkipDir(name string) bool {
	if strings.HasPrefix(name, ".") {
		return true
	}
	if unsupportedService(name) {
		return true
	}
	switch name {
	case "aws", "awsgo", "internal":
		return true
	default:
		return false
	}
}

func unsupportedService(name string) bool {
	switch name {
	case "alexaforbusiness", "bedrockdataautomationruntime", "qapps", "securityir":
		return true
	default:
		return false
	}
}

func writeAdapter(cmdDir, cmdVar string) error {
	path := filepath.Join(cmdDir, "zz_execute.go")
	content := fmt.Sprintf(`package cmd

func Execute(args []string) error {
	if p := %s.Parent(); p != nil {
		p.SetArgs(append([]string{%s.Name()}, args...))
		return p.Execute()
	}
	%s.SetArgs(args)
	return %s.Execute()
}
`, cmdVar, cmdVar, cmdVar, cmdVar)
	return os.WriteFile(path, []byte(content), 0o644)
}

func writeRegistry(path, modulePath string, services []serviceMeta) error {
	var b bytes.Buffer
	b.WriteString("// Code generated by awsgo/gen; DO NOT EDIT.\n")
	b.WriteString("package cmd\n\n")
	b.WriteString("import (\n")
	for _, svc := range services {
		fmt.Fprintf(&b, "\t%s \"%s/generated/%s/cmd\"\n", importAlias(svc.Name), modulePath, svc.Name)
	}
	b.WriteString(")\n\n")

	b.WriteString("type serviceDef struct {\n")
	b.WriteString("\tOperations []string\n")
	b.WriteString("\tOperationSet map[string]bool\n")
	b.WriteString("\tOperationInputs map[string][]string\n")
	b.WriteString("\tOperationInputTypes map[string]map[string]string\n")
	b.WriteString("\tOperationInputRequired map[string][]string\n")
	b.WriteString("\tRun        func(args []string) error\n")
	b.WriteString("}\n\n")

	b.WriteString("var serviceRegistry = map[string]serviceDef{\n")
	for _, svc := range services {
		alias := importAlias(svc.Name)
		fmt.Fprintf(&b, "\t\"%s\": {\n", svc.Name)
		b.WriteString("\t\tOperations: []string{")
		for i, op := range svc.Operations {
			if i > 0 {
				b.WriteString(", ")
			}
			fmt.Fprintf(&b, "\"%s\"", op)
		}
		b.WriteString("},\n")
		b.WriteString("\t\tOperationSet: map[string]bool{")
		for i, op := range svc.Operations {
			if i > 0 {
				b.WriteString(" ")
			}
			fmt.Fprintf(&b, "\"%s\": true,", op)
		}
		b.WriteString("},\n")
		b.WriteString("\t\tOperationInputs: map[string][]string{\n")
		opNames := make([]string, 0, len(svc.OperationInputs))
		for op := range svc.OperationInputs {
			opNames = append(opNames, op)
		}
		sort.Strings(opNames)
		for _, op := range opNames {
			fields := append([]string(nil), svc.OperationInputs[op]...)
			sort.Strings(fields)
			fmt.Fprintf(&b, "\t\t\t%q: {", op)
			for i, f := range fields {
				if i > 0 {
					b.WriteString(", ")
				}
				fmt.Fprintf(&b, "%q", f)
			}
			b.WriteString("},\n")
		}
		b.WriteString("\t\t},\n")
		b.WriteString("\t\tOperationInputTypes: map[string]map[string]string{\n")
		for _, op := range opNames {
			typeMap := svc.OperationInputTypes[op]
			typeFields := make([]string, 0, len(typeMap))
			for f := range typeMap {
				typeFields = append(typeFields, f)
			}
			sort.Strings(typeFields)
			fmt.Fprintf(&b, "\t\t\t%q: {", op)
			for i, f := range typeFields {
				if i > 0 {
					b.WriteString(", ")
				}
				fmt.Fprintf(&b, "%q: %q", f, typeMap[f])
			}
			b.WriteString("},\n")
		}
		b.WriteString("\t\t},\n")
		b.WriteString("\t\tOperationInputRequired: map[string][]string{\n")
		for _, op := range opNames {
			requiredFields := append([]string(nil), svc.OperationInputRequired[op]...)
			sort.Strings(requiredFields)
			fmt.Fprintf(&b, "\t\t\t%q: {", op)
			for i, f := range requiredFields {
				if i > 0 {
					b.WriteString(", ")
				}
				fmt.Fprintf(&b, "%q", f)
			}
			b.WriteString("},\n")
		}
		b.WriteString("\t\t},\n")
		fmt.Fprintf(&b, "\t\tRun: %s.Execute,\n", alias)
		b.WriteString("\t},\n")
	}
	b.WriteString("}\n")

	return os.WriteFile(path, b.Bytes(), 0o644)
}

func extractOperationInputs(content string, operations []string) map[string][]string {
	funcToOp := make(map[string]string, len(operations))
	for _, op := range operations {
		funcToOp[pascalFromOperation(op)] = op
	}

	out := map[string][]string{}
	lines := strings.Split(content, "\n")
	currentFunc := ""
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if m := funcRegex.FindStringSubmatch(line); len(m) == 2 {
			currentFunc = m[1]
		}
		if currentFunc == "" {
			continue
		}
		opName, ok := funcToOp[currentFunc]
		if !ok {
			continue
		}
		if !strings.Contains(line, "input := &") || !strings.Contains(line, "Input{") {
			continue
		}
		fields := map[string]struct{}{}
		for j := i + 1; j < len(lines); j++ {
			l := strings.TrimSpace(lines[j])
			if strings.HasPrefix(l, "}") {
				break
			}
			if strings.HasPrefix(l, "//") {
				l = strings.TrimSpace(strings.TrimPrefix(l, "//"))
			}
			if l == "" {
				continue
			}
			m := inputFieldRegex.FindStringSubmatch(l)
			if len(m) >= 2 {
				fields[m[1]] = struct{}{}
			}
		}
		if len(fields) == 0 {
			continue
		}
		items := make([]string, 0, len(fields))
		for f := range fields {
			items = append(items, f)
		}
		sort.Strings(items)
		out[opName] = items
	}
	return out
}

func extractOperationInputTypes(content string, operations []string) map[string]map[string]string {
	funcToOp := make(map[string]string, len(operations))
	for _, op := range operations {
		funcToOp[pascalFromOperation(op)] = op
	}

	out := map[string]map[string]string{}
	lines := strings.Split(content, "\n")
	currentFunc := ""
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if m := funcRegex.FindStringSubmatch(line); len(m) == 2 {
			currentFunc = m[1]
		}
		if currentFunc == "" {
			continue
		}
		opName, ok := funcToOp[currentFunc]
		if !ok {
			continue
		}
		if !strings.Contains(line, "input := &") || !strings.Contains(line, "Input{") {
			continue
		}
		fields := map[string]string{}
		for j := i + 1; j < len(lines); j++ {
			l := strings.TrimSpace(lines[j])
			if strings.HasPrefix(l, "}") {
				break
			}
			lineForType := l
			if strings.HasPrefix(l, "//") {
				lineForType = strings.TrimSpace(strings.TrimPrefix(l, "//"))
			}
			if lineForType == "" {
				continue
			}
			m := inputFieldRegex.FindStringSubmatch(lineForType)
			if len(m) != 3 {
				continue
			}
			field := m[1]
			valueExpr := strings.TrimSpace(m[2])
			fields[field] = inferType(valueExpr)
		}
		if len(fields) > 0 {
			out[opName] = fields
		}
	}
	return out
}

func extractOperationRequiredInputs(content string, operations []string) map[string][]string {
	funcToOp := make(map[string]string, len(operations))
	for _, op := range operations {
		funcToOp[pascalFromOperation(op)] = op
	}

	out := map[string][]string{}
	lines := strings.Split(content, "\n")
	currentFunc := ""
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if m := funcRegex.FindStringSubmatch(line); len(m) == 2 {
			currentFunc = m[1]
		}
		if currentFunc == "" {
			continue
		}
		opName, ok := funcToOp[currentFunc]
		if !ok {
			continue
		}
		if !strings.Contains(line, "input := &") || !strings.Contains(line, "Input{") {
			continue
		}
		required := map[string]struct{}{}
		for j := i + 1; j < len(lines); j++ {
			l := strings.TrimSpace(lines[j])
			if strings.HasPrefix(l, "}") {
				break
			}
			if strings.HasPrefix(l, "//") {
				l = strings.TrimSpace(strings.TrimPrefix(l, "//"))
			}
			if l == "" || !strings.Contains(l, "Required") {
				continue
			}
			m := inputFieldRegex.FindStringSubmatch(l)
			if len(m) >= 2 {
				required[m[1]] = struct{}{}
			}
		}
		items := make([]string, 0, len(required))
		for f := range required {
			items = append(items, f)
		}
		sort.Strings(items)
		out[opName] = items
	}
	return out
}

func inferType(valueExpr string) string {
	switch {
	case strings.HasPrefix(valueExpr, "aws.String("):
		return "string"
	case strings.HasPrefix(valueExpr, "aws.Int32("):
		return "int32"
	case strings.HasPrefix(valueExpr, "aws.Int64("):
		return "int64"
	case strings.HasPrefix(valueExpr, "aws.Bool("):
		return "bool"
	case strings.HasPrefix(valueExpr, "[]"):
		return valueExpr
	case strings.HasPrefix(valueExpr, "*"):
		return valueExpr
	case strings.Contains(valueExpr, "time.Time"):
		return "*time.Time"
	default:
		return valueExpr
	}
}

func pascalFromOperation(op string) string {
	parts := strings.Split(op, "-")
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		if len(p) == 1 {
			b.WriteString(strings.ToUpper(p))
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]))
		b.WriteString(p[1:])
	}
	return b.String()
}

func importAlias(name string) string {
	return "svc_" + identName(name)
}

func identName(s string) string {
	var b strings.Builder
	for i, r := range s {
		isAlpha := (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
		isDigit := r >= '0' && r <= '9'
		if isAlpha || isDigit {
			if i == 0 && isDigit {
				b.WriteRune('_')
			}
			b.WriteRune(r)
		} else {
			b.WriteRune('_')
		}
	}
	return b.String()
}

func readModulePath(goModPath string) (string, error) {
	b, err := os.ReadFile(goModPath)
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(string(b), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}
	return "", fmt.Errorf("module path not found in %s", goModPath)
}

func findRepoRoot(start string) (string, error) {
	curr := start
	for {
		if _, err := os.Stat(filepath.Join(curr, "go.mod")); err == nil {
			return curr, nil
		}
		next := filepath.Dir(curr)
		if next == curr {
			break
		}
		curr = next
	}
	return "", fmt.Errorf("could not locate repo root from %s", start)
}

func serviceCmdCompiles(root, service string) bool {
	cmd := exec.Command("go", "build", "./generated/"+service+"/cmd")
	cmd.Dir = root
	cmd.Env = os.Environ()
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func writeManifest(path string, services []serviceMeta) error {
	b, err := json.MarshalIndent(services, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0o644)
}

func writeSplitEntrypoints(root, modulePath string, services []serviceMeta) error {
	dispatcherDir := filepath.Join(root, "awsgo", "dispatcher")
	if err := os.MkdirAll(dispatcherDir, 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(dispatcherDir, "main.go"), renderDispatcherMain(services), 0o644); err != nil {
		return err
	}

	for _, svc := range services {
		serviceDir := filepath.Join(root, "awsgo", "services", svc.Name)
		if err := os.MkdirAll(serviceDir, 0o755); err != nil {
			return err
		}
		src := renderServiceMain(modulePath, svc.Name)
		if err := os.WriteFile(filepath.Join(serviceDir, "main.go"), src, 0o644); err != nil {
			return err
		}
	}
	return nil
}

func renderServiceMain(modulePath, service string) []byte {
	var b bytes.Buffer
	b.WriteString("package main\n\n")
	b.WriteString("import (\n")
	b.WriteString("\t\"fmt\"\n")
	b.WriteString("\t\"os\"\n\n")
	fmt.Fprintf(&b, "\tservicecmd %q\n", modulePath+"/generated/"+service+"/cmd")
	b.WriteString(")\n\n")
	b.WriteString("func main() {\n")
	b.WriteString("\tif err := servicecmd.Execute(os.Args[1:]); err != nil {\n")
	b.WriteString("\t\tfmt.Fprintln(os.Stderr, err)\n")
	b.WriteString("\t\tos.Exit(1)\n")
	b.WriteString("\t}\n")
	b.WriteString("}\n")
	return gofmtBytes(b.Bytes())
}

func renderDispatcherMain(services []serviceMeta) []byte {
	var b bytes.Buffer
	b.WriteString("package main\n\n")
	b.WriteString("import (\n")
	b.WriteString("\t\"fmt\"\n")
	b.WriteString("\t\"os\"\n")
	b.WriteString("\t\"os/exec\"\n")
	b.WriteString("\t\"path/filepath\"\n")
	b.WriteString("\t\"runtime\"\n")
	b.WriteString(")\n\n")

	b.WriteString("var serviceNames = []string{\n")
	for _, svc := range services {
		fmt.Fprintf(&b, "\t%q,\n", svc.Name)
	}
	b.WriteString("}\n\n")

	b.WriteString("var serviceOps = map[string]map[string]bool{\n")
	for _, svc := range services {
		fmt.Fprintf(&b, "\t%q: {", svc.Name)
		for _, op := range svc.Operations {
			fmt.Fprintf(&b, "%q: true, ", op)
		}
		b.WriteString("},\n")
	}
	b.WriteString("}\n\n")

	b.WriteString(`func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		printRootHelp()
		return
	}

	service := args[0]
	ops, ok := serviceOps[service]
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown service %q\n", service)
		os.Exit(2)
	}

	forwarded := args[1:]
	if len(forwarded) > 0 {
		if _, ok := ops[forwarded[0]]; ok {
			forwarded = append([]string{"--" + forwarded[0]}, forwarded[1:]...)
		}
	}

	bin, err := findServiceBinary(service)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(127)
	}

	cmd := exec.Command(bin, forwarded...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func printRootHelp() {
	fmt.Println("Dynamic AWS-style CLI over generated SDK commands")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  awsgo <service> <operation> [flags]")
	fmt.Println()
	fmt.Println("Available Services:")
	for _, name := range serviceNames {
		fmt.Printf("  %s\n", name)
	}
}

func findServiceBinary(service string) (string, error) {
	name := "awsgo-" + service
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	if dir := os.Getenv("AWSGO_SERVICE_DIR"); dir != "" {
		if p := filepath.Join(dir, name); fileExists(p) {
			return p, nil
		}
	}
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	dir := filepath.Dir(exe)
	for _, p := range []string{
		filepath.Join(dir, "awsgo-services", name),
		filepath.Join(dir, name),
	} {
		if fileExists(p) {
			return p, nil
		}
	}
	return "", fmt.Errorf("service binary for %q not found; install %s under awsgo-services/ or set AWSGO_SERVICE_DIR", service, name)
}

func fileExists(path string) bool {
	st, err := os.Stat(path)
	return err == nil && !st.IsDir()
}
`)
	return gofmtBytes(b.Bytes())
}

func gofmtBytes(src []byte) []byte {
	out, err := format.Source(src)
	if err != nil {
		return src
	}
	return out
}
