package converter

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/toon-format/toon-go"
)

const (
	sampleJSONFile = "../../examples/sample.json"
	sampleTOONFile = "../../examples/sample.toon"
)

func TestJSONToTOON(t *testing.T) {
	jsonData, err := os.ReadFile(sampleJSONFile)
	if err != nil {
		t.Fatalf("failed to read sample.json: %v", err)
	}

	expectedTOON, err := os.ReadFile(sampleTOONFile)
	if err != nil {
		t.Fatalf("failed to read sample.toon: %v", err)
	}

	got, err := JSONToTOON(string(jsonData))
	if err != nil {
		t.Fatalf("JSONToTOON failed: %v", err)
	}

	// Semantic check: decode both TOON outputs into map[string]any
	var wantMap, gotMap map[string]any
	if err := toon.Unmarshal(expectedTOON, &wantMap); err != nil {
		t.Fatalf("failed to decode expected TOON: %v", err)
	}
	if err := toon.Unmarshal([]byte(got), &gotMap); err != nil {
		t.Fatalf("failed to decode actual TOON: %v", err)
	}
	if !deepEqual(wantMap, gotMap) {
		t.Errorf("JSONToTOON semantic mismatch")
	}

	// log exact string outputs
	t.Logf("Expected TOON:\n%s\n", string(expectedTOON))
	t.Logf("Generated TOON:\n%s\n", got)
}

func TestTOONToJSON(t *testing.T) {
	toonData, err := os.ReadFile(sampleTOONFile)
	if err != nil {
		t.Fatalf("failed to read sample.toon: %v", err)
	}

	originalJSON, err := os.ReadFile(sampleJSONFile)
	if err != nil {
		t.Fatalf("failed to read sample.json: %v", err)
	}

	got, err := TOONToJSON(string(toonData))
	if err != nil {
		t.Fatalf("TOONToJSON failed: %v", err)
	}

	// Semantic check
	if !jsonEqual(originalJSON, []byte(got)) {
		t.Errorf("TOONToJSON semantic mismatch")
	}

	// pretty-print and log outputs for inspection
	var origPretty, gotPretty bytes.Buffer
	if err := json.Indent(&origPretty, originalJSON, "", "  "); err != nil {
		t.Fatalf("failed to indent original JSON: %v", err)
	}
	if err := json.Indent(&gotPretty, []byte(got), "", "  "); err != nil {
		t.Fatalf("failed to indent generated JSON: %v", err)
	}

	t.Logf("Original JSON:\n%s\n", origPretty.String())
	t.Logf("Generated JSON:\n%s\n", gotPretty.String())
}

func TestRoundTrip(t *testing.T) {
	jsonData, err := os.ReadFile(sampleJSONFile)
	if err != nil {
		t.Fatalf("failed to read sample.json: %v", err)
	}

	toonStr, err := JSONToTOON(string(jsonData))
	if err != nil {
		t.Fatalf("JSONToTOON failed: %v", err)
	}

	backJSON, err := TOONToJSON(toonStr)
	if err != nil {
		t.Fatalf("TOONToJSON failed: %v", err)
	}

	if !jsonEqual(jsonData, []byte(backJSON)) {
		t.Error("round-trip JSON → TOON → JSON failed: result not equal to original")
	}
}

// compares two JSON blobs semantically
func jsonEqual(a, b []byte) bool {
	var ja, jb any
	if err := json.Unmarshal(a, &ja); err != nil {
		return false
	}
	if err := json.Unmarshal(b, &jb); err != nil {
		return false
	}
	return deepEqual(ja, jb)
}

// performs semantic comparison of JSON structures
func deepEqual(a, b any) bool {
	switch ta := a.(type) {
	case map[string]any:
		tb, ok := b.(map[string]any)
		if !ok || len(ta) != len(tb) {
			return false
		}
		for k, v := range ta {
			if !deepEqual(v, tb[k]) {
				return false
			}
		}
		return true
	case []any:
		tb, ok := b.([]any)
		if !ok || len(ta) != len(tb) {
			return false
		}
		for i := range ta {
			if !deepEqual(ta[i], tb[i]) {
				return false
			}
		}
		return true
	default:
		return a == b
	}
}
