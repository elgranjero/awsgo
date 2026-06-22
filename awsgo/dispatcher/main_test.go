package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestEnabledServiceNamesUsesSplitManifest(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("AWSGO_SERVICE_DIR", dir)
	writeManifest(t, dir, "sts\niam\nmissing\nsts\n")

	got := enabledServiceNames()
	want := []string{"iam", "sts"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("enabled services = %#v, want %#v", got, want)
	}
}

func TestServiceEnabledUsesSplitManifest(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("AWSGO_SERVICE_DIR", dir)
	writeManifest(t, dir, "ec2\n")

	if !serviceEnabled("ec2") {
		t.Fatal("ec2 should be enabled by manifest")
	}
	if serviceEnabled("s3") {
		t.Fatal("s3 should not be enabled when absent from manifest")
	}
}

func TestLeanCandidateOperation(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		service   string
		operation string
		ok        bool
	}{
		{
			name:      "aws style service operation",
			args:      []string{"sts", "get-caller-identity", "--profile", "prod"},
			service:   "sts",
			operation: "get-caller-identity",
			ok:        true,
		},
		{
			name:      "global flag before operation",
			args:      []string{"ec2", "--profile", "prod", "describe-instances"},
			service:   "ec2",
			operation: "describe-instances",
			ok:        true,
		},
		{
			name:      "help operation form",
			args:      []string{"s3", "help", "list-buckets"},
			service:   "s3",
			operation: "list-buckets",
			ok:        true,
		},
		{
			name: "service help is not an operation",
			args: []string{"s3", "--help"},
			ok:   false,
		},
		{
			name: "legacy selector is full-runtime only",
			args: []string{"ec2", "--describe-instances"},
			ok:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, operation, ok := leanCandidateOperation(tt.args)
			if ok != tt.ok || service != tt.service || operation != tt.operation {
				t.Fatalf("lean candidate = (%q, %q, %v), want (%q, %q, %v)", service, operation, ok, tt.service, tt.operation, tt.ok)
			}
		})
	}
}

func TestShouldUseLean(t *testing.T) {
	t.Setenv("AWSGO_DISABLE_LEAN", "")
	if !shouldUseLean([]string{"sts", "get-caller-identity"}) {
		t.Fatal("supported lean operation should use lean")
	}
	if shouldUseLean([]string{"iam", "definitely-not-an-operation"}) {
		t.Fatal("unknown operation should not use lean")
	}
}

func TestShouldUseLeanHonorsDisableEnv(t *testing.T) {
	t.Setenv("AWSGO_DISABLE_LEAN", "1")
	if shouldUseLean([]string{"sts", "get-caller-identity"}) {
		t.Fatal("AWSGO_DISABLE_LEAN=1 should force full runtime")
	}
}

func TestShouldUseLeanSupportsNonJSONOutputFlag(t *testing.T) {
	t.Setenv("AWSGO_DISABLE_LEAN", "")
	if !shouldUseLean([]string{"ec2", "describe-instances", "--output", "text"}) {
		t.Fatal("--output text should use lean now that lean supports formatted output")
	}
	if !shouldUseLean([]string{"ec2", "describe-instances", "--output", "json"}) {
		t.Fatal("--output json should remain eligible for lean runtime")
	}
}

func TestShouldUseLeanSupportsSharedConfigOutput(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, "config")
	if err := os.WriteFile(configPath, []byte("[profile prod]\noutput = text\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("AWS_CONFIG_FILE", configPath)
	t.Setenv("AWS_PROFILE", "")
	t.Setenv("AWS_DEFAULT_OUTPUT", "")

	if !shouldUseLean([]string{"sts", "--profile", "prod", "get-caller-identity"}) {
		t.Fatal("profile output=text should use lean now that lean supports formatted output")
	}
	if !shouldUseLean([]string{"sts", "--profile", "prod", "get-caller-identity", "--output", "json"}) {
		t.Fatal("explicit --output json should override shared config output")
	}
}

func TestFindLeanServiceBinaryUsesEnvDir(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("AWSGO_LEAN_SERVICE_DIR", dir)
	path := filepath.Join(dir, "awsgo-lean-sts")
	if err := os.WriteFile(path, []byte("#!/bin/sh\n"), 0o700); err != nil {
		t.Fatal(err)
	}

	got, err := findLeanServiceBinary("sts")
	if err != nil {
		t.Fatal(err)
	}
	if got != path {
		t.Fatalf("lean binary = %q, want %q", got, path)
	}
}

func writeManifest(t *testing.T, dir, content string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, "manifest.txt"), []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
}
