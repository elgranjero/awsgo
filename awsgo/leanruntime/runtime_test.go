package leanruntime

import (
	"context"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func TestParseGlobalAndOperationFlags(t *testing.T) {
	ops := map[string]Operation{
		"describe-instances": {
			Fields: []Field{{Name: "MaxResults", Flag: "max-results", Type: "int"}},
			Run:    func(context.Context, aws.Config, Values) (any, error) { return nil, nil },
		},
	}
	parsed, err := Parse("ec2", ops, []string{"ec2", "describe-instances", "--profile", "prod", "--region=us-east-1", "--max-results", "20", "--query", "Reservations[*]"})
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}
	if parsed.Operation != "describe-instances" || parsed.Profile != "prod" || parsed.Region != "us-east-1" || parsed.Query != "Reservations[*]" {
		t.Fatalf("unexpected parsed globals: %#v", parsed)
	}
	if got := StringValue(parsed.Values, "MaxResults"); got != "20" {
		t.Fatalf("MaxResults mismatch: %q", got)
	}
}

func TestFormatNormalizesSDKStyleOutput(t *testing.T) {
	out, err := Format(map[string]any{
		"ResultMetadata": map[string]any{},
		"IsTruncated":    true,
		"Marker":         "TOKEN",
		"Roles": []any{
			map[string]any{
				"RoleName":                 "Administrator",
				"CreateDate":               "2017-03-28T20:33:52Z",
				"AssumeRolePolicyDocument": "%7B%22Version%22%3A%222012-10-17%22%7D",
				"Description":              nil,
			},
		},
	}, "json", "")
	if err != nil {
		t.Fatalf("Format returned error: %v", err)
	}
	got := string(out)
	for _, unwanted := range []string{"ResultMetadata", "IsTruncated", `"Marker"`, `%7B%22Version%22`, `"Description"`} {
		if strings.Contains(got, unwanted) {
			t.Fatalf("output still contains %q:\n%s", unwanted, got)
		}
	}
	for _, want := range []string{`"NextToken": "TOKEN"`, `"AssumeRolePolicyDocument": {`, `"CreateDate": "2017-03-28T20:33:52+00:00"`} {
		if !strings.Contains(got, want) {
			t.Fatalf("output missing %q:\n%s", want, got)
		}
	}
}

func TestFormatAppliesJMESPathQuery(t *testing.T) {
	out, err := Format(map[string]any{"Reservations": []any{map[string]any{"ReservationId": "r-1"}}}, "json", "Reservations[*].ReservationId")
	if err != nil {
		t.Fatalf("Format returned error: %v", err)
	}
	if got := strings.TrimSpace(string(out)); got != `[
  "r-1"
]` {
		t.Fatalf("query output mismatch: %s", got)
	}
}
