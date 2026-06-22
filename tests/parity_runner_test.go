package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

type cmdResult struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

func awsgoBinPath(t *testing.T) string {
	t.Helper()
	if v := strings.TrimSpace(os.Getenv("AWSGO_BIN")); v != "" {
		return v
	}
	return filepath.Join(repoRoot(t), "bin", "awsgo")
}

func awsBinPath() string {
	if v := strings.TrimSpace(os.Getenv("AWS_BIN")); v != "" {
		return v
	}
	return "aws"
}

func runCLI(t *testing.T, bin string, args []string, extraEnv map[string]string) cmdResult {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Env = mergedEnv(extraEnv)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	exitCode := 0
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			exitCode = ee.ExitCode()
		} else {
			t.Fatalf("failed to run %s %v: %v", bin, args, err)
		}
	}
	return cmdResult{
		Stdout:   strings.TrimSpace(stdout.String()),
		Stderr:   strings.TrimSpace(stderr.String()),
		ExitCode: exitCode,
	}
}

func mergedEnv(extra map[string]string) []string {
	if len(extra) == 0 {
		return os.Environ()
	}
	env := map[string]string{}
	for _, item := range os.Environ() {
		k, v, ok := strings.Cut(item, "=")
		if ok {
			env[k] = v
		}
	}
	for k, v := range extra {
		env[k] = v
	}
	out := make([]string, 0, len(env))
	for k, v := range env {
		out = append(out, k+"="+v)
	}
	return out
}

func requireBinary(t *testing.T, bin string) {
	t.Helper()
	if _, err := exec.LookPath(bin); err != nil {
		t.Skipf("binary %q not found in PATH", bin)
	}
}

func canonicalJSON(t *testing.T, raw string) string {
	t.Helper()
	var v any
	if err := json.Unmarshal([]byte(strings.TrimSpace(raw)), &v); err != nil {
		t.Fatalf("invalid JSON: %v\nraw=%s", err, raw)
	}
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("failed to marshal JSON: %v", err)
	}
	return string(b)
}

func sameErrorClass(a, b cmdResult) bool {
	if a.ExitCode == 0 || b.ExitCode == 0 {
		return false
	}
	x := strings.ToLower(a.Stderr)
	y := strings.ToLower(b.Stderr)
	tokens := []string{
		"accessdenied", "unauthorized", "not authorized", "forbidden",
		"validation", "invalid", "missing required", "throttl",
		"optinrequired", "expiredtoken", "signature", "region",
	}
	for _, tok := range tokens {
		if strings.Contains(x, tok) && strings.Contains(y, tok) {
			return true
		}
	}
	return false
}

func mustContainAll(t *testing.T, haystack string, needles []string) {
	t.Helper()
	for _, n := range needles {
		if !strings.Contains(haystack, n) {
			t.Fatalf("expected output to contain %q\n--- output ---\n%s", n, haystack)
		}
	}
}

func mustNotContainAny(t *testing.T, haystack string, needles []string) {
	t.Helper()
	for _, n := range needles {
		if strings.Contains(haystack, n) {
			t.Fatalf("expected output not to contain %q\n--- output ---\n%s", n, haystack)
		}
	}
}

func commandString(bin string, args []string) string {
	return fmt.Sprintf("%s %s", bin, strings.Join(args, " "))
}
