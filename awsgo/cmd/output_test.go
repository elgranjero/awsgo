package cmd

import (
	"strings"
	"testing"
)

func TestFormatOutputTextUsesStructuredYAML(t *testing.T) {
	raw := []byte(`{"Volumes":[{"VolumeId":"vol-1"}]}`)
	out, err := formatOutput(raw, "text")
	if err != nil {
		t.Fatalf("formatOutput returned error: %v", err)
	}
	got := string(out)
	if !strings.Contains(got, "Volumes:") || !strings.Contains(got, "VolumeId: vol-1") {
		t.Fatalf("expected YAML-like text output, got %q", got)
	}
}

func TestFormatOutputMarkdown(t *testing.T) {
	raw := []byte(`[{"A":"x","B":"y"}]`)
	out, err := formatOutput(raw, "markdown")
	if err != nil {
		t.Fatalf("formatOutput returned error: %v", err)
	}
	got := string(out)
	if !strings.Contains(got, "| A | B |") {
		t.Fatalf("expected markdown table header, got %q", got)
	}
	if !strings.Contains(got, "| x | y |") {
		t.Fatalf("expected markdown table row, got %q", got)
	}
}

func TestFormatOutputHTML(t *testing.T) {
	raw := []byte(`[{"A":"x"}]`)
	out, err := formatOutput(raw, "html")
	if err != nil {
		t.Fatalf("formatOutput returned error: %v", err)
	}
	got := string(out)
	if !strings.Contains(got, "<html") || !strings.Contains(got, "<table>") {
		t.Fatalf("expected html table output, got %q", got)
	}
}

func TestWriteOutputNormalizesNilSlices(t *testing.T) {
	type response struct {
		Datapoints []string
		Label      string
	}
	out, err := writeOutput(nil, nil, response{Label: "VolumeWriteIOPs"}, "json")
	if err != nil {
		t.Fatalf("writeOutput returned error: %v", err)
	}
	got := string(out)
	if !strings.Contains(got, `"Datapoints": []`) {
		t.Fatalf("expected nil slice to render as empty array, got %q", got)
	}
}

func TestFormatOutputNormalizesIAMListRolesLikeAWSCLI(t *testing.T) {
	raw := []byte(`{
		"IsTruncated": true,
		"Marker": "TOKEN",
		"ResultMetadata": {},
		"Roles": [
			{
				"Arn": "arn:aws:iam::123456789012:role/Administrator",
				"AssumeRolePolicyDocument": "%7B%22Version%22%3A%222012-10-17%22%2C%22Statement%22%3A%5B%7B%22Effect%22%3A%22Allow%22%2C%22Action%22%3A%22sts%3AAssumeRole%22%7D%5D%7D",
				"CreateDate": "2017-03-28T20:33:52Z",
				"Description": null,
				"Path": "/",
				"RoleName": "Administrator",
				"Tags": []
			}
		]
	}`)
	out, err := formatOutput(raw, "json")
	if err != nil {
		t.Fatalf("formatOutput returned error: %v", err)
	}
	got := string(out)
	for _, unwanted := range []string{
		"ResultMetadata",
		"IsTruncated",
		`"Marker"`,
		`"Description"`,
		`"Tags"`,
		`%7B%22Version%22`,
	} {
		if strings.Contains(got, unwanted) {
			t.Fatalf("output still contains %q:\n%s", unwanted, got)
		}
	}
	for _, want := range []string{
		`"NextToken": "TOKEN"`,
		`"AssumeRolePolicyDocument": {`,
		`"Version": "2012-10-17"`,
		`"CreateDate": "2017-03-28T20:33:52+00:00"`,
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("output missing %q:\n%s", want, got)
		}
	}
}

func TestFormatOutputMergesPaginatorPages(t *testing.T) {
	raw := []byte(`[
		{"ResultMetadata":{},"Roles":[{"RoleName":"one"}]},
		{"ResultMetadata":{},"Roles":[{"RoleName":"two"}]}
	]`)
	out, err := formatOutput(raw, "json")
	if err != nil {
		t.Fatalf("formatOutput returned error: %v", err)
	}
	got := string(out)
	if !strings.Contains(got, `"RoleName": "one"`) || !strings.Contains(got, `"RoleName": "two"`) {
		t.Fatalf("expected merged paginator page roles, got:\n%s", got)
	}
	if strings.Contains(got, "ResultMetadata") {
		t.Fatalf("expected ResultMetadata to be removed, got:\n%s", got)
	}
}
