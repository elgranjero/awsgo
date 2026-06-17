package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"
)

type config struct {
	sdkRoot     string
	outRoot     string
	servicesRaw string
	write       bool
	jobs        int
}

type serviceMeta struct {
	Name          string
	Operations    []opMeta
	RequiredVars  []stringVarMeta
	VarForFlag    map[string]string
	VarTypeByFlag map[string]string
	Region        regionBehavior
}

type regionBehavior struct {
	VarName   string
	Shorthand string
	Default   string
}

type opMeta struct {
	Service       string
	NamePascal    string
	NameKebab     string
	NameWords     string
	FuncName      string
	BoolVar       string
	DocLines      []string
	InputFields   []fieldMeta
	OutputFields  []fieldMeta
	HasPaginator  bool
	SourcePath    string
	GeneratedFile string
}

type fieldMeta struct {
	Name     string
	Type     string
	Required bool
}

type stringVarMeta struct {
	Name     string
	Flag     string
	Desc     string
	Type     string
	Required bool
}

func main() {
	cfg := parseFlags()
	if cfg.sdkRoot == "" {
		cfg.sdkRoot = filepath.Join(userHomeDir(), "git", "aws-sdk-go-v2", "service")
	}
	if cfg.outRoot == "" {
		cfg.outRoot = filepath.Join(userHomeDir(), "aws", "service", "generated")
	}

	services, err := discoverServices(cfg.sdkRoot, cfg.servicesRaw)
	if err != nil {
		fatal(err)
	}
	if len(services) == 0 {
		fatal(fmt.Errorf("no services selected under %s", cfg.sdkRoot))
	}
	if cfg.write && strings.TrimSpace(cfg.servicesRaw) == "" {
		if err := pruneStaleServiceDirs(cfg.outRoot, services); err != nil {
			fatal(err)
		}
	}

	if cfg.jobs < 1 {
		cfg.jobs = 1
	}
	if cfg.jobs > len(services) {
		cfg.jobs = len(services)
	}
	type result struct {
		service string
		skipped bool
		err     error
	}
	workCh := make(chan string)
	resCh := make(chan result, len(services))
	var wg sync.WaitGroup
	for i := 0; i < cfg.jobs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for svc := range workCh {
				err := generateService(cfg, svc)
				if err == errNoOperations {
					resCh <- result{service: svc, skipped: true}
					continue
				}
				resCh <- result{service: svc, err: err}
			}
		}()
	}
	for _, svc := range services {
		workCh <- svc
	}
	close(workCh)
	wg.Wait()
	close(resCh)

	generated := 0
	for r := range resCh {
		if r.err != nil {
			fatal(fmt.Errorf("%s: %w", r.service, r.err))
		}
		if r.skipped {
			continue
		}
		generated++
	}

	mode := "dry-run"
	if cfg.write {
		mode = "write"
	}
	fmt.Printf("servicegen complete: %d service(s) (%s)\n", generated, mode)
}

func parseFlags() config {
	cfg := config{}
	var outAlias string
	flag.StringVar(&cfg.sdkRoot, "sdk-root", "", "Input AWS SDK service root (default: ~/git/aws-sdk-go-v2/service)")
	flag.StringVar(&cfg.outRoot, "out-root", "", "Output root for generated files (default: ~/aws/service/generated)")
	flag.StringVar(&outAlias, "out", "", "Output root alias (same as --out-root)")
	flag.StringVar(&cfg.servicesRaw, "services", "", "Comma-separated services to generate (default: auto-discover from sdk-root)")
	flag.BoolVar(&cfg.write, "write", true, "Write generated files")
	flag.IntVar(&cfg.jobs, "jobs", runtime.NumCPU(), "Number of concurrent services to generate")
	flag.Parse()
	if cfg.outRoot == "" && outAlias != "" {
		cfg.outRoot = outAlias
	}
	return cfg
}

var (
	errNoOperations      = fmt.Errorf("no operations")
	regionFlagLineRegexp = regexp.MustCompile(`Flags\(\)\.StringVarP\(&([A-Za-z0-9_]+), "region", "([^"]*)", "([^"]*)", "Set AWS Region"\)`)
)

func generateService(cfg config, svc string) error {
	meta, err := loadServiceMeta(cfg.sdkRoot, cfg.outRoot, svc)
	if err != nil {
		return err
	}
	if len(meta.Operations) == 0 {
		if cfg.write {
			_ = os.RemoveAll(filepath.Join(cfg.outRoot, svc))
		}
		return errNoOperations
	}

	serviceOutDir := filepath.Join(cfg.outRoot, svc)
	if cfg.write {
		if err := os.MkdirAll(filepath.Join(serviceOutDir, "cmd"), 0o755); err != nil {
			return err
		}
		if err := pruneStaleServiceFiles(serviceOutDir, meta); err != nil {
			return err
		}
	}

	for _, op := range meta.Operations {
		opPath := filepath.Join(serviceOutDir, op.NamePascal+".go")
		src := renderOperationReference(meta.Name, op)
		if cfg.write {
			if err := os.WriteFile(opPath, src, 0o644); err != nil {
				return err
			}
		}
	}

	if cfg.write {
		cmdPath := filepath.Join(serviceOutDir, "cmd", svc+".go")
		execPath := filepath.Join(serviceOutDir, "cmd", "zz_execute.go")
		sharedPath := filepath.Join(serviceOutDir, "cmd", "zz_shared.go")
		if err := os.WriteFile(cmdPath, renderServiceCmd(meta), 0o644); err != nil {
			return err
		}
		if err := os.WriteFile(execPath, renderExecute(meta.Name), 0o644); err != nil {
			return err
		}
		if err := os.WriteFile(sharedPath, renderShared(), 0o644); err != nil {
			return err
		}
	}
	return nil
}

func pruneStaleServiceFiles(serviceOutDir string, meta serviceMeta) error {
	keepRoot := map[string]struct{}{}
	for _, op := range meta.Operations {
		keepRoot[op.NamePascal+".go"] = struct{}{}
	}
	rootEntries, err := os.ReadDir(serviceOutDir)
	if err != nil {
		return err
	}
	for _, e := range rootEntries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !strings.HasSuffix(name, ".go") {
			continue
		}
		if _, ok := keepRoot[name]; ok {
			continue
		}
		if err := os.Remove(filepath.Join(serviceOutDir, name)); err != nil {
			return err
		}
	}

	keepCmd := map[string]struct{}{
		meta.Name + ".go": {},
		"zz_execute.go":   {},
		"zz_shared.go":    {},
	}
	cmdDir := filepath.Join(serviceOutDir, "cmd")
	cmdEntries, err := os.ReadDir(cmdDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	for _, e := range cmdEntries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !strings.HasSuffix(name, ".go") {
			continue
		}
		if _, ok := keepCmd[name]; ok {
			continue
		}
		if err := os.Remove(filepath.Join(cmdDir, name)); err != nil {
			return err
		}
	}
	return nil
}

func discoverServices(sdkRoot, raw string) ([]string, error) {
	if strings.TrimSpace(raw) != "" {
		parts := strings.Split(raw, ",")
		out := make([]string, 0, len(parts))
		for _, p := range parts {
			s := strings.TrimSpace(p)
			if s != "" {
				out = append(out, s)
			}
		}
		sort.Strings(out)
		return uniq(out), nil
	}

	entries, err := os.ReadDir(sdkRoot)
	if err != nil {
		return nil, err
	}
	out := []string{}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		name := e.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		if _, err := os.Stat(filepath.Join(sdkRoot, name, "api_client.go")); err != nil {
			continue
		}
		out = append(out, name)
	}
	sort.Strings(out)
	return out, nil
}

func loadServiceMeta(sdkRoot, outRoot, service string) (serviceMeta, error) {
	svcDir := filepath.Join(sdkRoot, service)
	pattern := filepath.Join(svcDir, "api_op_*.go")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return serviceMeta{}, err
	}
	sort.Strings(files)

	ops := make([]opMeta, 0, len(files))
	requiredByFlag := map[string]stringVarMeta{}
	hasInputRegion := false
	for _, path := range files {
		op, err := parseOperationFile(service, path)
		if err != nil {
			return serviceMeta{}, err
		}
		if op.NamePascal == "" {
			continue
		}
		ops = append(ops, op)
	}
	sort.Slice(ops, func(i, j int) bool { return ops[i].NameKebab < ops[j].NameKebab })
	opFlags := map[string]bool{}
	for _, op := range ops {
		opFlags[op.NameKebab] = true
	}
	for _, op := range ops {
		for _, f := range op.InputFields {
			if f.Type == "*string" && f.Name == "Region" {
				hasInputRegion = true
			}
			flagName := kebabFromPascal(f.Name)
			if isReservedServiceFlag(flagName) || opFlags[flagName] {
				continue
			}
			varType := "string"
			if f.Type == "[]string" {
				varType = "[]string"
			}
			v := stringVarMeta{
				Name:     "_" + service + f.Name,
				Flag:     flagName,
				Desc:     wordsFromPascal(f.Name),
				Type:     varType,
				Required: f.Required,
			}
			existing, ok := requiredByFlag[v.Flag]
			if !ok {
				requiredByFlag[v.Flag] = v
				continue
			}
			if !existing.Required && v.Required {
				requiredByFlag[v.Flag] = v
				continue
			}
			if existing.Required == v.Required && existing.Type != "[]string" && v.Type == "[]string" {
				requiredByFlag[v.Flag] = v
				continue
			}
			if chooseCanonicalVar(existing.Name, v.Name) == v.Name {
				requiredByFlag[v.Flag] = v
			}
		}
	}

	varList := make([]stringVarMeta, 0, len(requiredByFlag))
	varForFlag := make(map[string]string, len(requiredByFlag))
	varTypeByFlag := make(map[string]string, len(requiredByFlag))
	for _, v := range requiredByFlag {
		varList = append(varList, v)
		varForFlag[v.Flag] = v.Name
		varTypeByFlag[v.Flag] = v.Type
	}
	sort.Slice(varList, func(i, j int) bool { return varList[i].Flag < varList[j].Flag })

	rb := detectRegionBehavior(outRoot, service, hasInputRegion)
	if rb.VarName != "_awsRegion" {
		varForFlag["region"] = rb.VarName
		varTypeByFlag["region"] = "*string"
		filtered := make([]stringVarMeta, 0, len(varList))
		for _, v := range varList {
			if v.Name == rb.VarName || v.Flag == "region" {
				continue
			}
			filtered = append(filtered, v)
		}
		varList = filtered
	}

	return serviceMeta{
		Name:          service,
		Operations:    ops,
		RequiredVars:  varList,
		VarForFlag:    varForFlag,
		VarTypeByFlag: varTypeByFlag,
		Region:        rb,
	}, nil
}

func detectRegionBehavior(outRoot, service string, hasInputRegion bool) regionBehavior {
	out := regionBehavior{VarName: "_awsRegion", Shorthand: "", Default: ""}
	cmdPath := filepath.Join(outRoot, service, "cmd", service+".go")
	if b, err := os.ReadFile(cmdPath); err == nil {
		if m := regionFlagLineRegexp.FindSubmatch(b); len(m) == 4 {
			out.VarName = string(m[1])
			out.Shorthand = string(m[2])
			out.Default = string(m[3])
			if out.VarName == "" {
				out.VarName = "_awsRegion"
			}
			return out
		}
	}
	if hasInputRegion {
		out.VarName = "_" + service + "Region"
	}
	return out
}

func parseOperationFile(service, path string) (opMeta, error) {
	base := filepath.Base(path)
	if !strings.HasPrefix(base, "api_op_") || !strings.HasSuffix(base, ".go") {
		return opMeta{}, nil
	}
	opPascal := strings.TrimSuffix(strings.TrimPrefix(base, "api_op_"), ".go")
	if strings.HasSuffix(opPascal, "_test") {
		return opMeta{}, nil
	}
	inputType := opPascal + "Input"
	outputType := opPascal + "Output"

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return opMeta{}, err
	}

	inputFields := []fieldMeta{}
	outputFields := []fieldMeta{}
	hasPaginator := false
	docLines := []string{}

	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Name == nil {
				continue
			}
			if d.Name.Name == "New"+opPascal+"Paginator" {
				hasPaginator = true
			}
			if d.Name.Name == opPascal && d.Doc != nil {
				docLines = cleanDocLines(d.Doc)
			}
		case *ast.GenDecl:
			if d.Tok != token.TYPE {
				continue
			}
			for _, spec := range d.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok || ts.Name == nil {
					continue
				}
				st, ok := ts.Type.(*ast.StructType)
				if !ok || st.Fields == nil {
					continue
				}
				if ts.Name.Name == inputType {
					inputFields = parseStructFields(fset, st)
				}
				if ts.Name.Name == outputType {
					outputFields = parseStructFields(fset, st)
				}
			}
		}
	}

	return opMeta{
		Service:      service,
		NamePascal:   opPascal,
		NameKebab:    kebabFromPascal(opPascal),
		NameWords:    wordsFromPascal(opPascal),
		FuncName:     service + "_" + opPascal,
		BoolVar:      "_" + service + opPascal,
		DocLines:     docLines,
		InputFields:  inputFields,
		OutputFields: outputFields,
		HasPaginator: hasPaginator,
		SourcePath:   path,
	}, nil
}

func parseStructFields(fset *token.FileSet, st *ast.StructType) []fieldMeta {
	fields := []fieldMeta{}
	for _, f := range st.Fields.List {
		if len(f.Names) == 0 {
			continue
		}
		typ := strings.TrimSpace(exprString(fset, f.Type))
		required := commentHasRequired(f.Doc) || commentHasRequired(f.Comment)
		for _, n := range f.Names {
			if n == nil || !n.IsExported() {
				continue
			}
			if n.Name == "noSmithyDocumentSerde" {
				continue
			}
			fields = append(fields, fieldMeta{Name: n.Name, Type: typ, Required: required})
		}
	}
	return fields
}

func exprString(fset *token.FileSet, e ast.Expr) string {
	var b bytes.Buffer
	if err := format.Node(&b, fset, e); err != nil {
		return ""
	}
	return b.String()
}

func commentHasRequired(cg *ast.CommentGroup) bool {
	if cg == nil {
		return false
	}
	for _, c := range cg.List {
		t := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
		if strings.Contains(t, "This member is required") {
			return true
		}
	}
	return false
}

func cleanDocLines(cg *ast.CommentGroup) []string {
	if cg == nil {
		return nil
	}
	out := []string{}
	for _, c := range cg.List {
		t := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
		if t == "" {
			out = append(out, "")
			continue
		}
		t = strings.ReplaceAll(t, "@", "(at)")
		out = append(out, t)
	}
	for len(out) > 0 && strings.TrimSpace(out[len(out)-1]) == "" {
		out = out[:len(out)-1]
	}
	return out
}

func renderOperation(op opMeta, varForFlag map[string]string, varTypeByFlag map[string]string) []byte {
	var b bytes.Buffer
	if len(op.DocLines) > 0 {
		for _, line := range op.DocLines {
			if line == "" {
				fmt.Fprintf(&b, "//\n")
			} else {
				fmt.Fprintf(&b, "// %s\n", line)
			}
		}
	}
	fmt.Fprintf(&b, "func %s(cfg aws.Config, client *%s.Client) {\n", op.FuncName, op.Service)
	requiredCount := 0
	for _, f := range op.InputFields {
		if f.Required {
			requiredCount++
		}
	}
	if requiredCount == 0 {
		fmt.Fprintf(&b, "\tinput := &%s.%sInput{}\n\n", op.Service, op.NamePascal)
	} else {
		fmt.Fprintf(&b, "\tinput := &%s.%sInput{\n", op.Service, op.NamePascal)
		for _, f := range op.InputFields {
			if !f.Required {
				continue
			}
			fmt.Fprintf(&b, "\t\t// %s: %s, // Required\n", f.Name, f.Type)
		}
		fmt.Fprintf(&b, "\t}\n\n")
	}
	for _, f := range op.InputFields {
		flag := kebabFromPascal(f.Name)
		varName, ok := varForFlag[flag]
		if !ok {
			continue
		}
		varType := varTypeByFlag[flag]
		if f.Type == "*string" {
			if varType == "[]string" {
				fmt.Fprintf(&b, "\tif len(%s) > 0 {\n", varName)
				fmt.Fprintf(&b, "\t\tinput.%s = aws.String(%s[0])\n", f.Name, varName)
				fmt.Fprintf(&b, "\t}\n")
				continue
			}
			fmt.Fprintf(&b, "\tif len(%s) > 0 {\n", varName)
			fmt.Fprintf(&b, "\t\tinput.%s = aws.String(%s)\n", f.Name, varName)
			fmt.Fprintf(&b, "\t}\n")
			continue
		}
		if f.Type == "[]string" {
			if varType == "[]string" {
				fmt.Fprintf(&b, "\tif len(%s) > 0 {\n", varName)
				fmt.Fprintf(&b, "\t\tinput.%s = append([]string(nil), %s...)\n", f.Name, varName)
				fmt.Fprintf(&b, "\t}\n")
				continue
			}
			fmt.Fprintf(&b, "\tif len(%s) > 0 {\n", varName)
			fmt.Fprintf(&b, "\t\tinput.%s = []string{%s}\n", f.Name, varName)
			fmt.Fprintf(&b, "\t}\n")
			continue
		}
		rawExpr := varName
		if varType == "[]string" {
			fmt.Fprintf(&b, "\tif len(%s) > 0 {\n", varName)
			rawExpr = varName + "[0]"
		} else {
			fmt.Fprintf(&b, "\tif len(%s) > 0 {\n", varName)
		}
		fmt.Fprintf(&b, "\tif err := assignInputField(input, %q, %s); err != nil {\n", f.Name, rawExpr)
		fmt.Fprintf(&b, "\t\tlog.Errorf(\"invalid --%s: %%s\", err.Error())\n", flag)
		fmt.Fprintf(&b, "\t\treturn\n")
		fmt.Fprintf(&b, "\t}\n")
		fmt.Fprintf(&b, "\t}\n")
	}
	fmt.Fprintf(&b, "\n")

	if op.HasPaginator {
		fmt.Fprintf(&b, "\tif disablePaginator() {\n")
		fmt.Fprintf(&b, "\t\tif resp, err := client.%s(context.TODO(), input); err != nil {\n", op.NamePascal)
		fmt.Fprintf(&b, "\t\t\tlog.Errorf(\"%%s\", err.Error())\n")
		fmt.Fprintf(&b, "\t\t\treturn\n")
		fmt.Fprintf(&b, "\t\t} else {\n")
		fmt.Fprintf(&b, "\t\t\twriteOutput(nil, nil, resp, _awsOutput)\n")
		fmt.Fprintf(&b, "\t\t}\n")
		fmt.Fprintf(&b, "\t\treturn\n")
		fmt.Fprintf(&b, "\t}\n\n")

		fmt.Fprintf(&b, "\tvar results []*%s.%sOutput\n", op.Service, op.NamePascal)
		fmt.Fprintf(&b, "\tp := %s.New%sPaginator(client, input)\n", op.Service, op.NamePascal)
		fmt.Fprintf(&b, "\tfor p.HasMorePages() {\n")
		fmt.Fprintf(&b, "\t\tif resp, err := p.NextPage(context.TODO()); err != nil {\n")
		fmt.Fprintf(&b, "\t\t\tlog.Errorf(\"%%s\", err.Error())\n")
		fmt.Fprintf(&b, "\t\t\treturn\n")
		fmt.Fprintf(&b, "\t\t} else {\n")
		fmt.Fprintf(&b, "\t\t\tresults = append(results, resp)\n")
		fmt.Fprintf(&b, "\t\t}\n")
		fmt.Fprintf(&b, "\t}\n")
		fmt.Fprintf(&b, "\twriteOutput(nil, nil, results, _awsOutput)\n")
	} else {
		fmt.Fprintf(&b, "\tif resp, err := client.%s(context.TODO(), input); err != nil {\n", op.NamePascal)
		fmt.Fprintf(&b, "\t\tlog.Errorf(\"%%s\", err.Error())\n")
		fmt.Fprintf(&b, "\t\treturn\n")
		fmt.Fprintf(&b, "\t} else {\n")
		fmt.Fprintf(&b, "\t\twriteOutput(nil, nil, resp, _awsOutput)\n")
		fmt.Fprintf(&b, "\t}\n")
	}
	fmt.Fprintf(&b, "}\n")
	return gofmtBytes(b.Bytes())
}

func renderOperationReference(service string, op opMeta) []byte {
	var b bytes.Buffer
	fmt.Fprintf(&b, "package %s\n\n", service)
	fmt.Fprintf(&b, "// %s is generated as a reference stub.\n", op.NamePascal)
	fmt.Fprintf(&b, "// Executable command wiring lives under cmd/%s.go.\n", service)
	if len(op.DocLines) > 0 {
		fmt.Fprintf(&b, "//\n")
		for _, line := range op.DocLines {
			if line == "" {
				fmt.Fprintf(&b, "//\n")
			} else {
				fmt.Fprintf(&b, "// %s\n", line)
			}
		}
	}
	return gofmtBytes(b.Bytes())
}

func renderServiceCmd(svc serviceMeta) []byte {
	var b bytes.Buffer
	regionVar := svc.Region.VarName
	regionFlagShort := svc.Region.Shorthand
	regionDefault := svc.Region.Default
	fmt.Fprintf(&b, "package cmd\n\n")
	fmt.Fprintf(&b, "import (\n")
	fmt.Fprintf(&b, "\t\"context\"\n")
	fmt.Fprintf(&b, "\n")
	fmt.Fprintf(&b, "\t\"github.com/aws/aws-sdk-go-v2/aws\"\n")
	fmt.Fprintf(&b, "\t\"github.com/aws/aws-sdk-go-v2/service/%s\"\n", svc.Name)
	fmt.Fprintf(&b, "\tlog \"github.com/sirupsen/logrus\"\n")
	fmt.Fprintf(&b, "\t\"github.com/spf13/cobra\"\n")
	fmt.Fprintf(&b, ")\n\n")

	fmt.Fprintf(&b, "// %sCmd represents the %s command\n", svc.Name, svc.Name)
	fmt.Fprintf(&b, "var _%sCmd = &cobra.Command{\n", svc.Name)
	fmt.Fprintf(&b, "\tUse:   %q,\n", svc.Name)
	fmt.Fprintf(&b, "\tShort: %q,\n", "AWS "+svc.Name+" CLI")
	fmt.Fprintf(&b, "\tRun: func(cmd *cobra.Command, args []string) {\n")
	fmt.Fprintf(&b, "\t\t_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed(\"output\"))\n")
	fmt.Fprintf(&b, "\t\tcfg, err := LoadAWSConfigWithMiddleware(_awsProfile)\n")
	fmt.Fprintf(&b, "\t\tif err != nil {\n")
	fmt.Fprintf(&b, "\t\t\tlog.Errorf(\"Failed to load configuration: %%s\", err.Error())\n")
	fmt.Fprintf(&b, "\t\t\treturn\n")
	fmt.Fprintf(&b, "\t\t}\n")
	fmt.Fprintf(&b, "\t\tif len(%s) > 0 {\n", regionVar)
	fmt.Fprintf(&b, "\t\t\tcfg.Region = %s\n", regionVar)
	fmt.Fprintf(&b, "\t\t}\n")
	fmt.Fprintf(&b, "\t\tclient := %s.NewFromConfig(cfg)\n", svc.Name)
	for _, op := range svc.Operations {
		fmt.Fprintf(&b, "\t\tif %s {\n", op.BoolVar)
		fmt.Fprintf(&b, "\t\t\t%s(cfg, client)\n", op.FuncName)
		fmt.Fprintf(&b, "\t\t\treturn\n")
		fmt.Fprintf(&b, "\t\t}\n")
	}
	fmt.Fprintf(&b, "\n\t},\n")
	fmt.Fprintf(&b, "}\n\n")

	fmt.Fprintf(&b, "var (\n")
	for _, op := range svc.Operations {
		fmt.Fprintf(&b, "\t%s bool\n", op.BoolVar)
	}
	if regionVar != "_awsRegion" {
		fmt.Fprintf(&b, "\n\t%s string\n", regionVar)
	}
	if len(svc.RequiredVars) > 0 {
		fmt.Fprintf(&b, "\n")
		for _, v := range svc.RequiredVars {
			if v.Type == "[]string" {
				fmt.Fprintf(&b, "\t%s []string\n", v.Name)
			} else {
				fmt.Fprintf(&b, "\t%s string\n", v.Name)
			}
		}
	}
	fmt.Fprintf(&b, ")\n\n")

	for _, op := range svc.Operations {
		fmt.Fprintf(&b, "%s\n\n", strings.TrimSpace(string(renderOperation(op, svc.VarForFlag, svc.VarTypeByFlag))))
	}

	fmt.Fprintf(&b, "func init() {\n")
	fmt.Fprintf(&b, "\t_rootCmd.AddCommand(_%sCmd)\n", svc.Name)
	fmt.Fprintf(&b, "\t_%sCmd.Flags().SortFlags = false\n\n", svc.Name)
	fmt.Fprintf(&b, "\t_%sCmd.Flags().StringVarP(&_awsProfile, \"profile\", \"\", \"\", \"AWS shared config profile\")\n", svc.Name)
	fmt.Fprintf(&b, "\t_%sCmd.Flags().StringVarP(&%s, \"region\", %q, %q, \"Set AWS Region\")\n\n", svc.Name, regionVar, regionFlagShort, regionDefault)
	fmt.Fprintf(&b, "\t_%sCmd.Flags().StringVarP(&_awsOutput, \"output\", \"o\", \"\", \"Output format: json|yaml|text|table|csv|markdown|html\")\n\n", svc.Name)
	for _, v := range svc.RequiredVars {
		if v.Type == "[]string" {
			fmt.Fprintf(&b, "\t_%sCmd.Flags().StringSliceVarP(&%s, %q, \"\", nil, %q)\n", svc.Name, v.Name, v.Flag, v.Desc)
		} else {
			fmt.Fprintf(&b, "\t_%sCmd.Flags().StringVarP(&%s, %q, \"\", \"\", %q)\n", svc.Name, v.Name, v.Flag, v.Desc)
		}
	}
	if len(svc.RequiredVars) > 0 {
		fmt.Fprintf(&b, "\n")
	}
	for _, op := range svc.Operations {
		fmt.Fprintf(&b, "\t_%sCmd.Flags().BoolVarP(&%s, %q, \"\", false, %q)\n", svc.Name, op.BoolVar, op.NameKebab, op.NameWords)
	}
	fmt.Fprintf(&b, "\n}\n")

	return gofmtBytes(b.Bytes())
}

func renderExecute(service string) []byte {
	src := fmt.Sprintf(`package cmd

func Execute(args []string) error {
	if p := _%sCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_%sCmd.Name()}, args...))
		return p.Execute()
	}
	_%sCmd.SetArgs(args)
	return _%sCmd.Execute()
}
`, service, service, service, service)
	return gofmtBytes([]byte(src))
}

func renderShared() []byte {
	src := `package cmd

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html"
	"os"
	"path/filepath"
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
	safeStructured := normalizeStructured(structured)
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
		data, err := yaml.Marshal(safeStructured)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		b.Write(data)
	case "", "json":
		data, err := json.MarshalIndent(safeStructured, "", "  ")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		b.Write(data)
	default:
		data, err := json.MarshalIndent(safeStructured, "", "  ")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		b.Write(data)
	}
	if b.Len() > 0 && b.Bytes()[b.Len()-1] != '\n' {
		b.WriteByte('\n')
	}
	_, _ = os.Stdout.Write(b.Bytes())
}

func normalizeStructured(v any) any {
	data, err := json.Marshal(v)
	if err != nil {
		return v
	}
	var out any
	if err := json.Unmarshal(data, &out); err != nil {
		return v
	}
	return out
}

func resolveAWSOutput(profile string, explicit bool) string {
	if explicit {
		if v := strings.TrimSpace(_awsOutput); v != "" {
			return v
		}
	}
	if v := strings.TrimSpace(os.Getenv("AWS_DEFAULT_OUTPUT")); v != "" {
		return v
	}
	if v := sharedConfigProfileOutput(profile); v != "" {
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
`
	return gofmtBytes([]byte(src))
}

func kebabFromPascal(s string) string {
	if s == "" {
		return ""
	}
	return strings.Join(splitIdentifierWords(s), "-")
}

func wordsFromPascal(s string) string {
	if s == "" {
		return ""
	}
	parts := splitIdentifierWords(s)
	for i := range parts {
		u := strings.ToUpper(parts[i])
		if isAcronym(u) {
			parts[i] = u
		} else {
			parts[i] = strings.ToUpper(parts[i][:1]) + strings.ToLower(parts[i][1:])
		}
	}
	return strings.Join(parts, " ")
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

func chooseCanonicalVar(a, b string) string {
	score := func(v string) int {
		s := 0
		for _, r := range v {
			if r >= 'A' && r <= 'Z' {
				s++
			}
		}
		return s
	}
	if score(a) > score(b) {
		return a
	}
	if score(b) > score(a) {
		return b
	}
	if a < b {
		return a
	}
	return b
}

func isAcronym(s string) bool {
	switch s {
	case "ACL", "AES", "AMI", "API", "ARN", "ASCII", "AWS", "AZ", "CIDR", "CPU", "CRC32", "CRC32C", "CRC64NVME", "CSV", "DB", "DDOS", "DNS", "EC2", "ECR", "EFS", "EIP", "GPU", "HTML", "HTTP", "HTTPS", "IAM", "ID", "IOPS", "IP", "IPV4", "IPV6", "JSON", "JWK", "JWS", "KMS", "MD5", "MFA", "NAT", "NVME", "OIDC", "OS", "RAM", "RDS", "RSA", "S3", "SES", "SHA1", "SHA256", "SNS", "SQS", "SSE", "SSH", "SSL", "TCP", "TLS", "TTL", "UDP", "UID", "URI", "URL", "UUID", "VPC", "VPN", "VTL", "WAF", "XML", "YAML":
		return true
	default:
		return false
	}
}

func isReservedServiceFlag(flag string) bool {
	switch flag {
	case "help", "h", "profile", "region", "output":
		return true
	default:
		return false
	}
}

func gofmtBytes(src []byte) []byte {
	out, err := format.Source(src)
	if err != nil {
		return src
	}
	return out
}

func uniq(in []string) []string {
	seen := map[string]struct{}{}
	out := make([]string, 0, len(in))
	for _, s := range in {
		if _, ok := seen[s]; ok {
			continue
		}
		seen[s] = struct{}{}
		out = append(out, s)
	}
	return out
}

func pruneStaleServiceDirs(outRoot string, services []string) error {
	entries, err := os.ReadDir(outRoot)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	keep := map[string]struct{}{}
	for _, s := range services {
		keep[s] = struct{}{}
	}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		name := e.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		if _, ok := keep[name]; ok {
			continue
		}
		if err := os.RemoveAll(filepath.Join(outRoot, name)); err != nil {
			return err
		}
	}
	return nil
}

func userHomeDir() string {
	h, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return h
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "servicegen error: %v\n", err)
	os.Exit(1)
}
