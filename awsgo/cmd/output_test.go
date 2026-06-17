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
