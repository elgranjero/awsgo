package cmd

import (
	"encoding/json"
	"testing"
)

func TestApplyQueryJMESPathFunctions(t *testing.T) {
	raw := []byte(`{"Label":"VolumeWriteIOPs","Datapoints":[{"Average":2},{"Average":1}]}`)
	out, err := applyQuery(raw, "sort_by(Datapoints,&Average)[*].Average")
	if err != nil {
		t.Fatalf("applyQuery returned error: %v", err)
	}

	var got []int
	if err := json.Unmarshal(out, &got); err != nil {
		t.Fatalf("failed to decode query output: %v", err)
	}
	if len(got) != 2 || got[0] != 1 || got[1] != 2 {
		t.Fatalf("unexpected query output: %#v", got)
	}
}

func TestApplyQueryJMESPathObject(t *testing.T) {
	raw := []byte(`{"Label":"VolumeWriteIOPs","Datapoints":[]}`)
	out, err := applyQuery(raw, "{Label: Label, Count: length(Datapoints)}")
	if err != nil {
		t.Fatalf("applyQuery returned error: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(out, &got); err != nil {
		t.Fatalf("failed to decode query output: %v", err)
	}
	if got["Label"] != "VolumeWriteIOPs" {
		t.Fatalf("unexpected label: %#v", got["Label"])
	}
	if count, ok := got["Count"].(float64); !ok || count != 0 {
		t.Fatalf("unexpected count: %#v", got["Count"])
	}
}

func TestApplyQueryProjectionNullReturnsEmptyArray(t *testing.T) {
	raw := []byte(`{"Label":"VolumeWriteIOPs","Datapoints":null}`)
	out, err := applyQuery(raw, "Datapoints[*].{Average:Average,Timestamp:Timestamp,Unit:Unit}")
	if err != nil {
		t.Fatalf("applyQuery returned error: %v", err)
	}
	if string(out) != "[]" && string(out) != "[]\n" {
		t.Fatalf("expected empty array, got %q", string(out))
	}
}
