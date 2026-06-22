package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type benchCase struct {
	Name      string
	Service   string
	Operation string
	AWSArgs   []string
	AWSGOArgs []string
}

type task struct {
	Suite     string
	Case      string
	Service   string
	Operation string
	Repeat    int
	Order     int
	CLI       string
	Cmd       []string
	Env       map[string]string
}

type row struct {
	task
	Status      string
	Exit        int
	Real        float64
	User        float64
	Sys         float64
	StdoutBytes int
	StderrBytes int
}

var leanCases = []benchCase{
	{
		Name:      "sts_get_caller_identity_help",
		Service:   "sts",
		Operation: "get-caller-identity",
		AWSArgs:   []string{"sts", "get-caller-identity", "help"},
		AWSGOArgs: []string{"sts", "get-caller-identity", "--help"},
	},
	{
		Name:      "ec2_describe_instances_help",
		Service:   "ec2",
		Operation: "describe-instances",
		AWSArgs:   []string{"ec2", "describe-instances", "help"},
		AWSGOArgs: []string{"ec2", "describe-instances", "--help"},
	},
	{
		Name:      "ec2_describe_volumes_help",
		Service:   "ec2",
		Operation: "describe-volumes",
		AWSArgs:   []string{"ec2", "describe-volumes", "help"},
		AWSGOArgs: []string{"ec2", "describe-volumes", "--help"},
	},
	{
		Name:      "s3_list_buckets_help",
		Service:   "s3",
		Operation: "list-buckets",
		AWSArgs:   []string{"s3api", "list-buckets", "help"},
		AWSGOArgs: []string{"s3", "list-buckets", "--help"},
	},
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}

	suite := flag.String("suite", envDefault("BENCH_STARTUP_SUITE", "lean"), "benchmark suite")
	repeat := flag.Int("repeat", envInt("BENCH_STARTUP_REPEAT", 10), "runs per CLI/case")
	workers := flag.Int("workers", envInt("BENCH_STARTUP_WORKERS", 1), "parallel workers")
	restMS := flag.Int("rest-ms", envInt("BENCH_STARTUP_REST_MS", 100), "rest interval after each command in milliseconds")
	timeoutSeconds := flag.Float64("timeout", envFloat("BENCH_STARTUP_TIMEOUT", 15), "per-command timeout in seconds")
	compare := flag.String("compare", envDefault("BENCH_STARTUP_COMPARE", "aws"), "comparison target: aws or full")
	awsBin := flag.String("aws-bin", envDefault("AWS_BIN", "aws"), "AWS CLI binary")
	awsgoBin := flag.String("awsgo-bin", envDefault("AWSGO_BIN", filepath.Join(root, "bin", "awsgo")), "awsgo binary")
	outPath := flag.String("out", envDefault("BENCH_STARTUP_OUT", ""), "TSV output path")
	dryRun := flag.Bool("dry-run", false, "print commands without running them")
	flag.Parse()

	if *repeat < 1 || *workers < 1 || *restMS < 0 || *timeoutSeconds <= 0 {
		return errors.New("repeat/workers must be positive, rest-ms must be non-negative, and timeout must be positive")
	}

	cases, err := casesForSuite(*suite)
	if err != nil {
		return err
	}
	tasks, err := buildTasks(*suite, *compare, cases, *repeat, *awsBin, *awsgoBin)
	if err != nil {
		return err
	}
	out := *outPath
	if strings.TrimSpace(out) == "" {
		out = filepath.Join(root, "benchmarks", fmt.Sprintf("startup-%s-%s-w%d-%s.tsv", *suite, *compare, *workers, time.Now().UTC().Format("20060102-150405")))
	}

	fmt.Printf("suite=%s compare=%s cases=%d repeat=%d workers=%d rest_ms=%d timeout=%.1f\n", *suite, *compare, len(cases), *repeat, *workers, *restMS, *timeoutSeconds)
	fmt.Printf("commands=%d output=%s\n", len(tasks), out)

	if *dryRun {
		for _, task := range tasks {
			fmt.Printf("%s: %s\n", task.CLI, shellJoinWithEnv(task.Env, task.Cmd))
		}
		return nil
	}

	rows := runTasks(tasks, envForRun(), *workers, time.Duration(*timeoutSeconds*float64(time.Second)), time.Duration(*restMS)*time.Millisecond)
	sortRows(rows)
	if err := writeRows(out, rows); err != nil {
		return err
	}
	summarize(rows)
	return nil
}

func casesForSuite(suite string) ([]benchCase, error) {
	switch suite {
	case "lean":
		return leanCases, nil
	default:
		return nil, fmt.Errorf("unsupported suite %q", suite)
	}
}

func buildTasks(suite, compare string, cases []benchCase, repeat int, awsBin, awsgoBin string) ([]task, error) {
	tasks := make([]task, 0, len(cases)*repeat*2)
	for r := 1; r <= repeat; r++ {
		order, err := compareOrder(compare, r)
		if err != nil {
			return nil, err
		}
		for _, c := range cases {
			for i, cli := range order {
				t := task{
					Suite:     suite,
					Case:      c.Name,
					Service:   c.Service,
					Operation: c.Operation,
					Repeat:    r,
					Order:     i + 1,
					CLI:       cli,
				}
				switch cli {
				case "aws":
					t.Cmd = append([]string{awsBin}, c.AWSArgs...)
				case "awsgo", "lean":
					t.Cmd = append([]string{awsgoBin}, c.AWSGOArgs...)
				case "full":
					t.Cmd = append([]string{awsgoBin}, c.AWSGOArgs...)
					t.Env = map[string]string{"AWSGO_DISABLE_LEAN": "1"}
				default:
					return nil, fmt.Errorf("unsupported cli %q", cli)
				}
				tasks = append(tasks, t)
			}
		}
	}
	return tasks, nil
}

func compareOrder(compare string, repeat int) ([]string, error) {
	var order []string
	switch compare {
	case "aws":
		order = []string{"aws", "awsgo"}
	case "full":
		order = []string{"full", "lean"}
	default:
		return nil, fmt.Errorf("unsupported compare %q", compare)
	}
	if repeat%2 == 0 {
		order[0], order[1] = order[1], order[0]
	}
	return order, nil
}

func runTasks(tasks []task, env []string, workers int, timeout, rest time.Duration) []row {
	jobs := make(chan task)
	results := make(chan row)
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range jobs {
				results <- runOne(task, env, timeout, rest)
			}
		}()
	}

	go func() {
		for _, task := range tasks {
			jobs <- task
		}
		close(jobs)
		wg.Wait()
		close(results)
	}()

	rows := make([]row, 0, len(tasks))
	for result := range results {
		rows = append(rows, result)
		if len(rows)%25 == 0 || len(rows) == len(tasks) {
			fmt.Printf("completed %d/%d\n", len(rows), len(tasks))
		}
	}
	return rows
}

func runOne(t task, env []string, timeout, rest time.Duration) row {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, t.Cmd[0], t.Cmd[1:]...)
	cmd.Env = envWithOverrides(env, t.Env)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	status := "ok"
	exitCode := 0
	err := cmd.Run()
	realSeconds := time.Since(start).Seconds()
	userSeconds := 0.0
	sysSeconds := 0.0

	if cmd.ProcessState != nil {
		userSeconds = cmd.ProcessState.UserTime().Seconds()
		sysSeconds = cmd.ProcessState.SystemTime().Seconds()
	}

	if ctx.Err() == context.DeadlineExceeded {
		status = "timeout"
		exitCode = 124
	} else if err != nil {
		status = "exit"
		exitCode = exitCodeFor(err)
	}

	if rest > 0 {
		time.Sleep(rest)
	}

	return row{
		task:        t,
		Status:      status,
		Exit:        exitCode,
		Real:        realSeconds,
		User:        userSeconds,
		Sys:         sysSeconds,
		StdoutBytes: stdout.Len(),
		StderrBytes: stderr.Len(),
	}
}

func exitCodeFor(err error) int {
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		return exitErr.ExitCode()
	}
	if errors.Is(err, exec.ErrNotFound) {
		return 127
	}
	return 1
}

func writeRows(path string, rows []row) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Comma = '\t'
	defer w.Flush()

	header := []string{"suite", "case", "service", "operation", "repeat", "order", "cli", "status", "exit", "real", "user", "sys", "stdout_bytes", "stderr_bytes", "cmd"}
	if err := w.Write(header); err != nil {
		return err
	}

	for _, r := range rows {
		record := []string{
			r.Suite,
			r.Case,
			r.Service,
			r.Operation,
			strconv.Itoa(r.Repeat),
			strconv.Itoa(r.Order),
			r.CLI,
			r.Status,
			strconv.Itoa(r.Exit),
			formatFloat(r.Real),
			formatFloat(r.User),
			formatFloat(r.Sys),
			strconv.Itoa(r.StdoutBytes),
			strconv.Itoa(r.StderrBytes),
			shellJoinWithEnv(r.Env, r.Cmd),
		}
		if err := w.Write(record); err != nil {
			return err
		}
	}
	return w.Error()
}

func summarize(rows []row) {
	fmt.Println()
	fmt.Println("summary (successful runs only):")
	fmt.Printf("%-8s %9s %9s %9s %9s %8s\n", "cli", "ok/total", "avg_real", "med_real", "avg_user", "avg_sys")

	averages := map[string]float64{}
	medians := map[string]float64{}
	for _, cli := range cliOrder(rows) {
		subset := filterCLI(rows, cli)
		ok := filterOK(subset)
		avgReal := average(values(ok, func(r row) float64 { return r.Real }))
		medReal := median(values(ok, func(r row) float64 { return r.Real }))
		avgUser := average(values(ok, func(r row) float64 { return r.User }))
		avgSys := average(values(ok, func(r row) float64 { return r.Sys }))
		fmt.Printf("%-8s %3d/%-5d %9.3f %9.3f %9.3f %8.3f\n", cli, len(ok), len(subset), avgReal, medReal, avgUser, avgSys)
		averages[cli] = avgReal
		medians[cli] = medReal
	}
	if averages["aws"] > 0 && averages["awsgo"] > 0 {
		fmt.Printf("\nwall-clock speedup: %.2fx\n", averages["aws"]/averages["awsgo"])
	}
	if averages["full"] > 0 && averages["lean"] > 0 {
		fmt.Printf("\nlean-vs-full avg speedup: %.2fx\n", averages["full"]/averages["lean"])
		if medians["full"] > 0 && medians["lean"] > 0 {
			fmt.Printf("lean-vs-full median speedup: %.2fx\n", medians["full"]/medians["lean"])
		}
		if leanFullSameOrBetter(rows, medians["full"], medians["lean"]) {
			fmt.Println("lean-vs-full verdict: same-or-better")
		} else {
			fmt.Println("lean-vs-full verdict: needs-review")
		}
	}
}

func cliOrder(rows []row) []string {
	seen := map[string]bool{}
	for _, r := range rows {
		seen[r.CLI] = true
	}
	preferred := []string{"aws", "awsgo", "full", "lean"}
	out := []string{}
	for _, cli := range preferred {
		if seen[cli] {
			out = append(out, cli)
			delete(seen, cli)
		}
	}
	rest := make([]string, 0, len(seen))
	for cli := range seen {
		rest = append(rest, cli)
	}
	sort.Strings(rest)
	return append(out, rest...)
}

func leanFullSameOrBetter(rows []row, fullMedian, leanMedian float64) bool {
	if leanMedian <= 0 || fullMedian <= 0 || leanMedian > fullMedian {
		return false
	}
	for _, r := range rows {
		if (r.CLI == "lean" || r.CLI == "full") && (r.Status != "ok" || r.Exit != 0) {
			return false
		}
	}
	return true
}

func filterCLI(rows []row, cli string) []row {
	out := make([]row, 0, len(rows))
	for _, r := range rows {
		if r.CLI == cli {
			out = append(out, r)
		}
	}
	return out
}

func filterOK(rows []row) []row {
	out := make([]row, 0, len(rows))
	for _, r := range rows {
		if r.Status == "ok" && r.Exit == 0 {
			out = append(out, r)
		}
	}
	return out
}

func values(rows []row, fn func(row) float64) []float64 {
	out := make([]float64, 0, len(rows))
	for _, r := range rows {
		out = append(out, fn(r))
	}
	return out
}

func average(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	total := 0.0
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func median(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	sort.Float64s(xs)
	mid := len(xs) / 2
	if len(xs)%2 == 1 {
		return xs[mid]
	}
	return (xs[mid-1] + xs[mid]) / 2
}

func sortRows(rows []row) {
	sort.Slice(rows, func(i, j int) bool {
		a, b := rows[i], rows[j]
		switch {
		case a.Case != b.Case:
			return a.Case < b.Case
		case a.Repeat != b.Repeat:
			return a.Repeat < b.Repeat
		case a.Order != b.Order:
			return a.Order < b.Order
		default:
			return a.CLI < b.CLI
		}
	})
}

func envForRun() []string {
	env := map[string]string{}
	for _, item := range os.Environ() {
		key, value, ok := strings.Cut(item, "=")
		if ok {
			env[key] = value
		}
	}
	env["AWS_PAGER"] = ""
	env["AWS_CLI_AUTO_PROMPT"] = "off"
	env["AWS_EC2_METADATA_DISABLED"] = "true"
	env["PAGER"] = "cat"
	env["CLICOLOR"] = "0"
	env["AWS_DEFAULT_OUTPUT"] = firstNonEmpty(env["AWS_DEFAULT_OUTPUT"], "json")
	delete(env, "AWSGO_DISABLE_LEAN")

	keys := make([]string, 0, len(env))
	for key := range env {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	out := make([]string, 0, len(keys))
	for _, key := range keys {
		out = append(out, key+"="+env[key])
	}
	return out
}

func envWithOverrides(base []string, overrides map[string]string) []string {
	if len(overrides) == 0 {
		return base
	}
	env := map[string]string{}
	for _, item := range base {
		key, value, ok := strings.Cut(item, "=")
		if ok {
			env[key] = value
		}
	}
	for key, value := range overrides {
		env[key] = value
	}
	keys := make([]string, 0, len(env))
	for key := range env {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	out := make([]string, 0, len(keys))
	for _, key := range keys {
		out = append(out, key+"="+env[key])
	}
	return out
}

func envDefault(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return fallback
}

func envInt(key string, fallback int) int {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	n, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return n
}

func envFloat(key string, fallback float64) float64 {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	n, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fallback
	}
	return n
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func formatFloat(v float64) string {
	return strconv.FormatFloat(v, 'f', 6, 64)
}

func shellJoin(args []string) string {
	parts := make([]string, 0, len(args))
	for _, arg := range args {
		parts = append(parts, shellQuote(arg))
	}
	return strings.Join(parts, " ")
}

func shellJoinWithEnv(env map[string]string, args []string) string {
	if len(env) == 0 {
		return shellJoin(args)
	}
	keys := make([]string, 0, len(env))
	for key := range env {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys)+len(args))
	for _, key := range keys {
		parts = append(parts, key+"="+shellQuote(env[key]))
	}
	for _, arg := range args {
		parts = append(parts, shellQuote(arg))
	}
	return strings.Join(parts, " ")
}

func shellQuote(s string) string {
	if s == "" {
		return "''"
	}
	if strings.IndexFunc(s, func(r rune) bool {
		return !(r >= 'A' && r <= 'Z') &&
			!(r >= 'a' && r <= 'z') &&
			!(r >= '0' && r <= '9') &&
			!strings.ContainsRune("@%_+=:,./-", r)
	}) == -1 {
		return s
	}
	return "'" + strings.ReplaceAll(s, "'", "'\\''") + "'"
}
