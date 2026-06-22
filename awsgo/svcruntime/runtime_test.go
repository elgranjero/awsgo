package svcruntime

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestExecuteServiceHandlesCLIInputJSONAndQuery(t *testing.T) {
	var forwarded []string
	var input map[string]any
	svc := ServiceDef{
		Operations:   []string{"get-metric-statistics"},
		OperationSet: map[string]bool{"get-metric-statistics": true},
		OperationInputs: map[string][]string{
			"get-metric-statistics": {"Dimensions", "MetricName", "Period", "StartTime", "Statistics"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-metric-statistics": {
				"Dimensions": "[]types.Dimension",
				"MetricName": "*string",
				"Period":     "*int32",
				"StartTime":  "*time.Time",
				"Statistics": "[]types.Statistic",
			},
		},
		OperationInputRequired: map[string][]string{
			"get-metric-statistics": {"MetricName", "Period", "StartTime", "Statistics"},
		},
		Run: func(args []string) error {
			forwarded = append([]string(nil), args...)
			if err := json.Unmarshal([]byte(os.Getenv("AWSGO_INPUT_JSON")), &input); err != nil {
				t.Fatalf("unmarshal AWSGO_INPUT_JSON: %v", err)
			}
			_, _ = os.Stdout.Write([]byte(`{"Datapoints":[],"Label":"VolumeWriteIOPs"}`))
			return nil
		},
	}

	out := captureStdout(t, func() {
		err := ExecuteService("cloudwatch", svc, []string{
			"get-metric-statistics",
			"--profile", "prod",
			"--cli-input-json", `{"MetricName":"VolumeWriteIOPs","Dimensions":[{"Name":"DBInstanceIdentifier","Value":"example-db"}],"Period":"3600","StartTime":"2026-06-17T00:00:00Z","Statistics":["Average"]}`,
			"--query", "Datapoints[*]",
		})
		if err != nil {
			t.Fatalf("ExecuteService returned error: %v", err)
		}
	})

	wantForwarded := []string{"--get-metric-statistics", "--output", "json", "--profile", "prod"}
	if !reflect.DeepEqual(forwarded, wantForwarded) {
		t.Fatalf("forwarded args mismatch:\n got: %#v\nwant: %#v", forwarded, wantForwarded)
	}
	if input["Period"] != float64(3600) {
		t.Fatalf("Period should be numeric after coercion, got %#v", input["Period"])
	}
	statistics, ok := input["Statistics"].([]any)
	if !ok || len(statistics) != 1 || statistics[0] != "Average" {
		t.Fatalf("Statistics should preserve enum string, got %#v", input["Statistics"])
	}
	dimensions, ok := input["Dimensions"].([]any)
	if !ok || len(dimensions) != 1 {
		t.Fatalf("Dimensions should preserve object list, got %#v", input["Dimensions"])
	}
	if dim, ok := dimensions[0].(map[string]any); !ok || dim["Name"] != "DBInstanceIdentifier" || dim["Value"] != "example-db" {
		t.Fatalf("Dimensions element mismatch, got %#v", dimensions[0])
	}
	if got := strings.TrimSpace(out); got != "[]" {
		t.Fatalf("query output mismatch: %q", got)
	}
}

func TestExecuteServiceOperationHelpIncludesRuntimeFlags(t *testing.T) {
	svc := ServiceDef{
		Operations:   []string{"list-roles"},
		OperationSet: map[string]bool{"list-roles": true},
		OperationInputs: map[string][]string{
			"list-roles": {"Marker", "MaxItems", "PathPrefix"},
		},
		OperationInputTypes: map[string]map[string]string{
			"list-roles": {
				"Marker":     "*string",
				"MaxItems":   "*int32",
				"PathPrefix": "*string",
			},
		},
	}

	out := captureStdout(t, func() {
		err := ExecuteService("iam", svc, []string{"list-roles", "--help"})
		if err != nil {
			t.Fatalf("ExecuteService returned error: %v", err)
		}
	})

	for _, want := range []string{
		"--cli-input-json",
		"--query string",
		"--max-items",
		"Known Input Fields:",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("help output missing %q:\n%s", want, out)
		}
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

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	orig := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w
	done := make(chan string, 1)
	go func() {
		var b bytes.Buffer
		_, _ = io.Copy(&b, r)
		done <- b.String()
	}()
	fn()
	_ = w.Close()
	os.Stdout = orig
	return <-done
}
