package cmd

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestCoerceValueInt(t *testing.T) {
	v, err := coerceValue("*int32", "1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	n, ok := v.(int64)
	if !ok || n != 1 {
		t.Fatalf("expected int64(1), got %#v", v)
	}
}

func TestCoerceValueFloat(t *testing.T) {
	v, err := coerceValue("*float64", "1.25")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	f, ok := v.(float64)
	if !ok || f != 1.25 {
		t.Fatalf("expected float64(1.25), got %#v", v)
	}
}

func TestCoerceValueBool(t *testing.T) {
	v, err := coerceValue("*bool", "true")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	b, ok := v.(bool)
	if !ok || !b {
		t.Fatalf("expected bool(true), got %#v", v)
	}
}

func TestSetInputEnvCoercesInputJSONStrings(t *testing.T) {
	restore := setInputEnv(`{"MaxItems":"1","DryRun":"true"}`, "", nil, map[string]string{
		"MaxItems": "*int32",
		"DryRun":   "*bool",
	})
	defer restore()

	raw := os.Getenv("AWSGO_INPUT_JSON")
	var m map[string]any
	if err := json.Unmarshal([]byte(raw), &m); err != nil {
		t.Fatalf("invalid env json: %v", err)
	}
	if _, ok := m["MaxItems"].(float64); !ok {
		t.Fatalf("MaxItems should be numeric, got %#v", m["MaxItems"])
	}
	if _, ok := m["DryRun"].(bool); !ok {
		t.Fatalf("DryRun should be bool, got %#v", m["DryRun"])
	}
}

func TestSetInputEnvMergesTypedOverrides(t *testing.T) {
	restore := setInputEnv(`{"Marker":"abc"}`, "", map[string]any{"MaxItems": int64(1)}, map[string]string{
		"MaxItems": "*int32",
		"Marker":   "*string",
	})
	defer restore()

	raw := os.Getenv("AWSGO_INPUT_JSON")
	var m map[string]any
	if err := json.Unmarshal([]byte(raw), &m); err != nil {
		t.Fatalf("invalid env json: %v", err)
	}
	if _, ok := m["MaxItems"].(float64); !ok {
		t.Fatalf("MaxItems should be numeric, got %#v", m["MaxItems"])
	}
	if marker, ok := m["Marker"].(string); !ok || marker != "abc" {
		t.Fatalf("Marker mismatch: %#v", m["Marker"])
	}
}

func TestParseShorthandMapAnyBasic(t *testing.T) {
	got, err := parseShorthandMapAny("Name=DBInstanceIdentifier,Value=sample-db")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := map[string]any{
		"Name":  "DBInstanceIdentifier",
		"Value": "sample-db",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("mismatch\n got=%#v\nwant=%#v", got, want)
	}
}

func TestParseShorthandMapAnyNested(t *testing.T) {
	got, err := parseShorthandMapAny("MetricName=VolumeWriteIOPs,Dimensions=[{Name=DBInstanceIdentifier,Value=sample-db}],Period=3600")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got["MetricName"] != "VolumeWriteIOPs" {
		t.Fatalf("MetricName mismatch: %#v", got["MetricName"])
	}
	if got["Period"] != int64(3600) {
		t.Fatalf("Period mismatch: %#v", got["Period"])
	}
	dims, ok := got["Dimensions"].([]any)
	if !ok || len(dims) != 1 {
		t.Fatalf("Dimensions mismatch: %#v", got["Dimensions"])
	}
	d0, ok := dims[0].(map[string]any)
	if !ok {
		t.Fatalf("Dimension[0] type mismatch: %#v", dims[0])
	}
	if d0["Name"] != "DBInstanceIdentifier" || d0["Value"] != "sample-db" {
		t.Fatalf("Dimension[0] values mismatch: %#v", d0)
	}
}

func TestCoerceValueListStructFromShorthand(t *testing.T) {
	v, err := coerceValue("[]types.Dimension", "Name=DBInstanceIdentifier,Value=sample-db")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	arr, ok := v.([]any)
	if !ok || len(arr) != 1 {
		t.Fatalf("expected single element list, got %#v", v)
	}
	m, ok := arr[0].(map[string]any)
	if !ok {
		t.Fatalf("expected map element, got %#v", arr[0])
	}
	if m["Name"] != "DBInstanceIdentifier" || m["Value"] != "sample-db" {
		t.Fatalf("unexpected map contents: %#v", m)
	}
}

func TestCoerceValueListScalarCommaSeparated(t *testing.T) {
	v, err := coerceValue("[]string", "Average,Sum")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	arr, ok := v.([]any)
	if !ok {
		t.Fatalf("expected []any, got %#v", v)
	}
	if len(arr) != 2 || arr[0] != "Average" || arr[1] != "Sum" {
		t.Fatalf("unexpected values: %#v", arr)
	}
}

func TestCoerceValueListEnumString(t *testing.T) {
	v, err := coerceValue("[]types.Statistic", "Average")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	arr, ok := v.([]any)
	if !ok {
		t.Fatalf("expected []any, got %#v", v)
	}
	if len(arr) != 1 || arr[0] != "Average" {
		t.Fatalf("enum values should stay strings, got %#v", arr)
	}
}

func TestHelpTypeEnumString(t *testing.T) {
	if got := helpType("types.StandardUnit"); got != "string" {
		t.Fatalf("enum help type mismatch: %q", got)
	}
	if got := helpType("*types.DefaultAction"); got != "object" {
		t.Fatalf("object help type mismatch: %q", got)
	}
}
