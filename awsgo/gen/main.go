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
	OperationMethods       map[string]string
	OperationHasPaginator  map[string]bool
	OperationInputs        map[string][]string
	OperationInputTypes    map[string]map[string]string
	OperationInputRequired map[string][]string
}

type sdkOpMeta struct {
	Name         string
	GoName       string
	InputFields  []string
	InputTypes   map[string]string
	RequiredOnly []string
	HasPaginator bool
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
		opMethods := map[string]string{}
		opHasPaginator := map[string]bool{}
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
				opMethods[meta.Name] = meta.GoName
				opHasPaginator[meta.Name] = meta.HasPaginator
				opInputs[meta.Name] = meta.InputFields
				opInputTypes[meta.Name] = meta.InputTypes
				opInputRequired[meta.Name] = meta.RequiredOnly
			}
			sort.Strings(ops)
		}

		if len(ops) == 0 {
			ops = append(ops, rawOps...)
			sort.Strings(ops)
			for _, op := range ops {
				opMethods[op] = pascalFromOperation(op)
			}

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
			OperationMethods:       opMethods,
			OperationHasPaginator:  opHasPaginator,
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
		hasPaginator := false
		for _, decl := range parsed.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Name == nil {
				continue
			}
			if fn.Name.Name == "New"+opPascal+"Paginator" {
				hasPaginator = true
				break
			}
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
			GoName:       opPascal,
			InputFields:  fields,
			InputTypes:   fieldTypes,
			RequiredOnly: required,
			HasPaginator: hasPaginator,
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
		src := renderServiceMain(modulePath, svc)
		if err := os.WriteFile(filepath.Join(serviceDir, "main.go"), src, 0o644); err != nil {
			return err
		}
	}

	leanDispatcherDir := filepath.Join(root, "awsgo", "lean-dispatcher")
	if err := os.MkdirAll(leanDispatcherDir, 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(leanDispatcherDir, "main.go"), renderLeanDispatcherMain(services), 0o644); err != nil {
		return err
	}
	for _, svc := range services {
		serviceDir := filepath.Join(root, "awsgo", "lean-services", svc.Name)
		if err := os.MkdirAll(serviceDir, 0o755); err != nil {
			return err
		}
		src := renderLeanServiceMain(modulePath, svc)
		if err := os.WriteFile(filepath.Join(serviceDir, "main.go"), src, 0o644); err != nil {
			return err
		}
	}
	return nil
}

func renderServiceMain(modulePath string, svc serviceMeta) []byte {
	var b bytes.Buffer
	b.WriteString("package main\n\n")
	b.WriteString("import (\n")
	b.WriteString("\t\"fmt\"\n")
	b.WriteString("\t\"os\"\n\n")
	fmt.Fprintf(&b, "\truntime %q\n", modulePath+"/awsgo/svcruntime")
	fmt.Fprintf(&b, "\tservicecmd %q\n", modulePath+"/generated/"+svc.Name+"/cmd")
	b.WriteString(")\n\n")
	b.WriteString("func main() {\n")
	b.WriteString("\tsvc := runtime.ServiceDef{\n")
	b.WriteString("\t\tOperations: []string{")
	for i, op := range svc.Operations {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "%q", op)
	}
	b.WriteString("},\n")
	b.WriteString("\t\tOperationSet: map[string]bool{")
	for i, op := range svc.Operations {
		if i > 0 {
			b.WriteString(" ")
		}
		fmt.Fprintf(&b, "%q: true,", op)
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
	b.WriteString("\t\tRun: servicecmd.Execute,\n")
	b.WriteString("\t}\n")
	fmt.Fprintf(&b, "\tif err := runtime.ExecuteService(%q, svc, os.Args[1:]); err != nil {\n", svc.Name)
	b.WriteString("\t\tfmt.Fprintln(os.Stderr, err)\n")
	b.WriteString("\t\tos.Exit(1)\n")
	b.WriteString("\t}\n")
	b.WriteString("}\n")
	return gofmtBytes(b.Bytes())
}

func renderLeanServiceMain(modulePath string, svc serviceMeta) []byte {
	var b bytes.Buffer
	b.WriteString("package main\n\n")
	b.WriteString("import (\n")
	b.WriteString("\t\"context\"\n")
	b.WriteString("\t\"fmt\"\n")
	b.WriteString("\t\"os\"\n\n")
	fmt.Fprintf(&b, "\t%q\n", modulePath+"/awsgo/leanruntime")
	b.WriteString("\t\"github.com/aws/aws-sdk-go-v2/aws\"\n")
	fmt.Fprintf(&b, "\tsvc %q\n", "github.com/aws/aws-sdk-go-v2/service/"+svc.Name)
	b.WriteString(")\n\n")

	opNames := append([]string(nil), svc.Operations...)
	sort.Strings(opNames)
	for _, op := range opNames {
		fieldVar := leanFieldsIdent(op)
		fmt.Fprintf(&b, "var %s = []leanruntime.Field{\n", fieldVar)
		fields := append([]string(nil), svc.OperationInputs[op]...)
		sort.Strings(fields)
		required := map[string]bool{}
		for _, f := range svc.OperationInputRequired[op] {
			required[f] = true
		}
		typeMap := svc.OperationInputTypes[op]
		for _, field := range fields {
			typ := "string"
			if typeMap != nil {
				if v := strings.TrimSpace(typeMap[field]); v != "" {
					typ = v
				}
			}
			fmt.Fprintf(&b, "\t{Name: %q, Flag: %q, Type: %q, Required: %t},\n", field, kebabFromPascal(field), typ, required[field])
		}
		b.WriteString("}\n\n")
	}

	b.WriteString("func main() {\n")
	b.WriteString("\tops := map[string]leanruntime.Operation{\n")
	for _, op := range opNames {
		method := svc.OperationMethods[op]
		if method == "" {
			method = pascalFromOperation(op)
		}
		fieldVar := leanFieldsIdent(op)
		fmt.Fprintf(&b, "\t\t%q: {\n", op)
		fmt.Fprintf(&b, "\t\t\tName: %q,\n", op)
		fmt.Fprintf(&b, "\t\t\tFields: %s,\n", fieldVar)
		b.WriteString("\t\t\tRun: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {\n")
		fmt.Fprintf(&b, "\t\t\t\tinput := &svc.%sInput{}\n", method)
		if svc.OperationHasPaginator[op] {
			fmt.Fprintf(&b, "\t\t\t\tdisablePaginator, err := leanruntime.ApplyInput(input, %s, values)\n", fieldVar)
			b.WriteString("\t\t\t\tif err != nil {\n")
			b.WriteString("\t\t\t\t\treturn nil, err\n")
			b.WriteString("\t\t\t\t}\n")
			b.WriteString("\t\t\t\tclient := svc.NewFromConfig(cfg)\n")
			b.WriteString("\t\t\t\tif disablePaginator || leanruntime.PaginatorDisabled() {\n")
			fmt.Fprintf(&b, "\t\t\t\t\treturn client.%s(ctx, input)\n", method)
			b.WriteString("\t\t\t\t}\n")
			fmt.Fprintf(&b, "\t\t\t\tvar results []*svc.%sOutput\n", method)
			fmt.Fprintf(&b, "\t\t\t\tp := svc.New%sPaginator(client, input)\n", method)
			b.WriteString("\t\t\t\tfor p.HasMorePages() {\n")
			b.WriteString("\t\t\t\t\tresp, err := p.NextPage(ctx)\n")
			b.WriteString("\t\t\t\t\tif err != nil {\n")
			b.WriteString("\t\t\t\t\t\treturn nil, err\n")
			b.WriteString("\t\t\t\t\t}\n")
			b.WriteString("\t\t\t\t\tresults = append(results, resp)\n")
			b.WriteString("\t\t\t\t}\n")
			b.WriteString("\t\t\t\treturn results, nil\n")
		} else {
			fmt.Fprintf(&b, "\t\t\t\tif _, err := leanruntime.ApplyInput(input, %s, values); err != nil {\n", fieldVar)
			b.WriteString("\t\t\t\t\treturn nil, err\n")
			b.WriteString("\t\t\t\t}\n")
			b.WriteString("\t\t\t\tclient := svc.NewFromConfig(cfg)\n")
			fmt.Fprintf(&b, "\t\t\t\treturn client.%s(ctx, input)\n", method)
		}
		b.WriteString("\t\t\t},\n")
		b.WriteString("\t\t},\n")
	}
	b.WriteString("\t}\n")
	fmt.Fprintf(&b, "\tif err := leanruntime.Execute(%q, ops, os.Args[1:]); err != nil {\n", svc.Name)
	b.WriteString("\t\tfmt.Fprintln(os.Stderr, err)\n")
	b.WriteString("\t\tos.Exit(1)\n")
	b.WriteString("\t}\n")
	b.WriteString("}\n")
	return gofmtBytes(b.Bytes())
}

func leanFieldsIdent(op string) string {
	return "fields_" + identName(strings.ReplaceAll(op, "-", "_"))
}

func renderLeanDispatcherMain(services []serviceMeta) []byte {
	var b bytes.Buffer
	b.WriteString("package main\n\n")
	b.WriteString("import (\n")
	b.WriteString("\t\"fmt\"\n")
	b.WriteString("\t\"os\"\n")
	b.WriteString("\t\"os/exec\"\n")
	b.WriteString("\t\"path/filepath\"\n")
	b.WriteString("\t\"runtime\"\n")
	b.WriteString("\t\"sort\"\n")
	b.WriteString("\t\"strings\"\n")
	b.WriteString(")\n\n")
	b.WriteString("var services = map[string]bool{\n")
	for _, svc := range services {
		fmt.Fprintf(&b, "\t%q: true,\n", svc.Name)
	}
	b.WriteString("}\n\n")
	b.WriteString(`func main() {
	if len(os.Args) < 2 || os.Args[1] == "--help" || os.Args[1] == "-h" || os.Args[1] == "help" {
		printHelp()
		return
	}
	service := os.Args[1]
	if !services[service] {
		fmt.Fprintf(os.Stderr, "unknown lean service %q\n", service)
		os.Exit(1)
	}
	bin, err := findServiceBinary(service)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	cmd := exec.Command(bin, os.Args[1:]...)
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

func printHelp() {
	fmt.Println("Lean AWS-style CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  awsgo-lean <service> <operation> [flags]")
	fmt.Println()
	fmt.Println("Available Services:")
	names := make([]string, 0, len(services))
	for name := range services {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("  %s\n", name)
	}
}

func findServiceBinary(service string) (string, error) {
	name := "awsgo-lean-" + service
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	if dir := strings.TrimSpace(os.Getenv("AWSGO_LEAN_SERVICE_DIR")); dir != "" {
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
		filepath.Join(dir, "awsgo-lean-services", name),
		filepath.Join(dir, name),
	} {
		if fileExists(p) {
			return p, nil
		}
	}
	return "", fmt.Errorf("lean service binary for %q not found", service)
}

func fileExists(path string) bool {
	st, err := os.Stat(path)
	return err == nil && !st.IsDir()
}
`)
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
	b.WriteString("\t\"sort\"\n")
	b.WriteString("\t\"strings\"\n")
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

	b.WriteString("var leanServiceOps = map[string]map[string]bool{\n")
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
	if !serviceEnabled(service) {
		fmt.Fprintf(os.Stderr, "service %q is not part of this split build\n", service)
		os.Exit(127)
	}

	if shouldUseLean(args) {
		bin, err := findLeanServiceBinary(service)
		if err == nil {
			runCommand(bin, args)
			return
		}
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

	runCommand(bin, forwarded)
}

func runCommand(bin string, args []string) {
	cmd := exec.Command(bin, args...)
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

func shouldUseLean(args []string) bool {
	if leanDisabled() {
		return false
	}
	service, operation, ok := leanCandidateOperation(args)
	if !ok {
		return false
	}
	ops, ok := leanServiceOps[service]
	return ok && ops[operation]
}

func leanDisabled() bool {
	switch strings.ToLower(strings.TrimSpace(os.Getenv("AWSGO_DISABLE_LEAN"))) {
	case "1", "true", "yes", "on":
		return true
	default:
		return false
	}
}

func leanCandidateOperation(args []string) (string, string, bool) {
	if len(args) < 2 {
		return "", "", false
	}
	service := args[0]
	tail := args[1:]
	for i := 0; i < len(tail); i++ {
		arg := tail[i]
		if arg == "help" {
			if i+1 < len(tail) {
				return service, tail[i+1], true
			}
			return "", "", false
		}
		if strings.HasPrefix(arg, "--") {
			name, _, hasValue := strings.Cut(strings.TrimPrefix(arg, "--"), "=")
			if name == "help" {
				return "", "", false
			}
			if !hasValue && flagConsumesNextArg(name) && i+1 < len(tail) {
				i++
			}
			continue
		}
		if strings.HasPrefix(arg, "-") {
			if arg == "-h" {
				return "", "", false
			}
			if arg == "-o" && i+1 < len(tail) {
				i++
			}
			continue
		}
		return service, arg, true
	}
	return "", "", false
}

func flagConsumesNextArg(name string) bool {
	switch name {
	case "profile", "region", "output", "query", "input-json", "cli-input-json", "input-file":
		return true
	default:
		return false
	}
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

func printRootHelp() {
	fmt.Println("Dynamic AWS-style CLI over generated SDK commands")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  awsgo <service> <operation> [flags]")
	fmt.Println()
	fmt.Println("Available Services:")
	names := enabledServiceNames()
	if len(names) == 0 {
		fmt.Println("  (none found; run make build-split)")
		return
	}
	for _, name := range names {
		fmt.Printf("  %s\n", name)
	}
}

func serviceEnabled(service string) bool {
	if names, ok := builtManifestServices(); ok {
		for _, name := range names {
			if name == service {
				return true
			}
		}
		return false
	}
	if _, err := findServiceBinary(service); err == nil {
		return true
	}
	_, err := findLeanServiceBinary(service)
	return err == nil
}

func enabledServiceNames() []string {
	if names, ok := builtManifestServices(); ok {
		return names
	}
	names := make([]string, 0, len(serviceNames))
	for _, name := range serviceNames {
		if _, err := findServiceBinary(name); err == nil {
			names = append(names, name)
			continue
		}
		if _, err := findLeanServiceBinary(name); err == nil {
			names = append(names, name)
		}
	}
	return names
}

func builtManifestServices() ([]string, bool) {
	paths := append(serviceManifestPaths(), leanServiceManifestPaths()...)
	out := make([]string, 0)
	seen := map[string]bool{}
	found := false
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		found = true
		for _, raw := range strings.Split(string(data), "\n") {
			name := strings.TrimSpace(raw)
			if name == "" || strings.HasPrefix(name, "#") {
				continue
			}
			if _, ok := serviceOps[name]; !ok || seen[name] {
				continue
			}
			seen[name] = true
			out = append(out, name)
		}
	}
	if !found {
		return nil, false
	}
	sort.Strings(out)
	return out, true
}

func serviceManifestPaths() []string {
	paths := []string{}
	if dir := os.Getenv("AWSGO_SERVICE_DIR"); dir != "" {
		paths = append(paths, filepath.Join(dir, "manifest.txt"))
	}
	exe, err := os.Executable()
	if err != nil {
		return paths
	}
	dir := filepath.Dir(exe)
	return append(paths,
		filepath.Join(dir, "awsgo-services", "manifest.txt"),
		filepath.Join(dir, "manifest.txt"),
	)
}

func leanServiceManifestPaths() []string {
	paths := []string{}
	if dir := os.Getenv("AWSGO_LEAN_SERVICE_DIR"); dir != "" {
		paths = append(paths, filepath.Join(dir, "manifest.txt"))
	}
	exe, err := os.Executable()
	if err != nil {
		return paths
	}
	dir := filepath.Dir(exe)
	return append(paths,
		filepath.Join(dir, "awsgo-lean-services", "manifest.txt"),
	)
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

func findLeanServiceBinary(service string) (string, error) {
	name := "awsgo-lean-" + service
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	if dir := strings.TrimSpace(os.Getenv("AWSGO_LEAN_SERVICE_DIR")); dir != "" {
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
		filepath.Join(dir, "awsgo-lean-services", name),
		filepath.Join(dir, name),
	} {
		if fileExists(p) {
			return p, nil
		}
	}
	return "", fmt.Errorf("lean service binary for %q not found", service)
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
