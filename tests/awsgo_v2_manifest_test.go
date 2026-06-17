package tests

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

type serviceManifestV2 struct {
	Name       string   `json:"Name"`
	CmdVar     string   `json:"CmdVar"`
	Operations []string `json:"Operations"`
}

var opFlagRegexV2 = regexp.MustCompile(`\.Flags\(\)\.BoolVarP\(&[A-Za-z0-9_]+,\s*"([^"]+)"`)

func TestAWSGoV2ManifestHasServices(t *testing.T) {
	services := loadManifestV2(t)
	if len(services) == 0 {
		t.Fatalf("awsgo manifest has no services")
	}
}

func TestAWSGoV2ManifestOperationsMatchServiceFlags(t *testing.T) {
	root := repoRoot(t)
	services := loadManifestV2(t)

	for _, svc := range services {
		svcCmdPath := filepath.Join(root, "generated", svc.Name, "cmd", svc.Name+".go")
		b, err := os.ReadFile(svcCmdPath)
		if err != nil {
			t.Fatalf("failed reading %s: %v", svcCmdPath, err)
		}

		matches := opFlagRegexV2.FindAllSubmatch(b, -1)
		flagSet := make(map[string]bool, len(matches))
		for _, m := range matches {
			if len(m) > 1 {
				flagSet[canonicalOpKeyV2(string(m[1]))] = true
			}
		}

		for _, op := range svc.Operations {
			if !flagSet[canonicalOpKeyV2(op)] {
				t.Fatalf("v2 service %q operation %q missing --%s in %s", svc.Name, op, op, svcCmdPath)
			}
		}
	}
}

func canonicalOpKeyV2(op string) string {
	b := make([]byte, 0, len(op))
	for i := 0; i < len(op); i++ {
		c := op[i]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			if c >= 'A' && c <= 'Z' {
				c = c - 'A' + 'a'
			}
			b = append(b, c)
		}
	}
	return string(b)
}

func loadManifestV2(t *testing.T) []serviceManifestV2 {
	t.Helper()
	root := repoRoot(t)
	path := filepath.Join(root, "awsgo", "cmd", "manifest_gen.json")
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read manifest %s: %v", path, err)
	}

	var services []serviceManifestV2
	if err := json.Unmarshal(b, &services); err != nil {
		t.Fatalf("failed to parse manifest %s: %v", path, err)
	}
	return services
}
