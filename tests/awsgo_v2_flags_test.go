package tests

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

var anyFlagRegexV2 = regexp.MustCompile(`\.Flags\(\)\.[A-Za-z]+VarP?\(&[A-Za-z0-9_]+,\s*"([^"]+)"`)

func TestAWSGoV2NoDuplicateServiceFlags(t *testing.T) {
	root := repoRoot(t)
	services := loadManifestV2(t)

	for _, svc := range services {
		path := filepath.Join(root, "generated", svc.Name, "cmd", svc.Name+".go")
		b, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed reading %s: %v", path, err)
		}

		seen := map[string]bool{}
		matches := anyFlagRegexV2.FindAllSubmatch(b, -1)
		for _, m := range matches {
			if len(m) < 2 {
				continue
			}
			flag := string(m[1])
			if seen[flag] {
				t.Fatalf("duplicate flag %q detected in %s", flag, path)
			}
			seen[flag] = true
		}
	}
}
