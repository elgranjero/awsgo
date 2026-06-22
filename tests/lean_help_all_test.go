package tests

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestLeanHelpAllGeneratedOperations(t *testing.T) {
	if os.Getenv("AWSGO_LEAN_HELP_ALL") != "1" {
		t.Skip("set AWSGO_LEAN_HELP_ALL=1 to run every generated operation help path")
	}
	bin := awsgoBinPath(t)
	requireBinary(t, bin)

	maxOps := 0
	if raw := strings.TrimSpace(os.Getenv("AWSGO_LEAN_HELP_MAX_OPS")); raw != "" {
		n, err := strconv.Atoi(raw)
		if err != nil {
			t.Fatalf("invalid AWSGO_LEAN_HELP_MAX_OPS=%q: %v", raw, err)
		}
		maxOps = n
	}

	checked := 0
	services := loadManifestV2(t)
	for _, svc := range services {
		serviceHelp := runHelpCLI(t, bin, svc.Name, "--help")
		if serviceHelp.ExitCode != 0 {
			t.Fatalf("service help failed for %s: %s", svc.Name, serviceHelp.Stderr)
		}
		if !strings.Contains(serviceHelp.Stdout, "Available Operations:") {
			t.Fatalf("service help for %s missing operations\n%s", svc.Name, serviceHelp.Stdout)
		}

		for _, op := range svc.Operations {
			if maxOps > 0 && checked >= maxOps {
				t.Logf("stopped after AWSGO_LEAN_HELP_MAX_OPS=%d", maxOps)
				return
			}
			result := runHelpCLI(t, bin, svc.Name, op, "--help")
			if result.ExitCode != 0 {
				t.Fatalf("operation help failed for %s %s: %s", svc.Name, op, result.Stderr)
			}
			if strings.Contains(strings.ToLower(result.Stdout), "unknown operation") {
				t.Fatalf("operation help reported unknown operation for %s %s\n%s", svc.Name, op, result.Stdout)
			}
			mustContainAll(t, result.Stdout, []string{
				"Usage:",
				"Global Flags:",
				"--profile string",
				"--output string",
				"--input-json",
				"--cli-input-json",
			})
			checked++
		}
	}
	t.Logf("checked %d generated operation help paths across %d services", checked, len(services))
}

func runHelpCLI(t *testing.T, bin string, args ...string) cmdResult {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Env = mergedEnv(map[string]string{
		"AWSGO_DISABLE_LEAN": "",
	})
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
