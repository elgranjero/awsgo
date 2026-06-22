package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveOutputFormatFromSharedConfig(t *testing.T) {
	configPath := writeTestAWSConfig(t, `
[default]
region = us-east-1
output = text

[profile prod]
region = us-west-2
output = yaml
`)
	t.Setenv("AWS_CONFIG_FILE", configPath)
	t.Setenv("AWS_PROFILE", "")
	t.Setenv("AWS_DEFAULT_OUTPUT", "")
	outputFormat = ""

	if got := resolveOutputFormat("default", false); got != "text" {
		t.Fatalf("default profile output = %q, want text", got)
	}
	if got := resolveOutputFormat("prod", false); got != "yaml" {
		t.Fatalf("prod profile output = %q, want yaml", got)
	}
}

func TestResolveOutputFormatPrecedence(t *testing.T) {
	configPath := writeTestAWSConfig(t, `
[default]
output = text
`)
	t.Setenv("AWS_CONFIG_FILE", configPath)
	t.Setenv("AWS_PROFILE", "")
	t.Setenv("AWS_DEFAULT_OUTPUT", "table")

	outputFormat = ""
	if got := resolveOutputFormat("default", false); got != "table" {
		t.Fatalf("env output = %q, want table", got)
	}

	outputFormat = "json"
	if got := resolveOutputFormat("default", true); got != "json" {
		t.Fatalf("explicit output = %q, want json", got)
	}
}

func writeTestAWSConfig(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "config")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}
