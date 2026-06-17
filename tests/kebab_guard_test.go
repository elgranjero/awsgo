package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGeneratedFlagsNoBrokenKebabFragments(t *testing.T) {
	root := repoRoot(t)
	pattern := filepath.Join(root, "generated", "*", "cmd", "*.go")
	files, err := filepath.Glob(pattern)
	if err != nil {
		t.Fatalf("glob failed: %v", err)
	}
	if len(files) == 0 {
		t.Fatalf("no generated cmd files found")
	}

	badFragments := []string{
		"--s-h-a",
		"--c-r-c",
		"--m-d5",
		"--m-f-a",
		"--d-do-s",
		"--n-v-m-e",
		"--o-auth",
		"--qn-a",
		"--b-box",
		"--a-zs",
		"--t-bs",
		"--g-bs",
		"--ur-ls",
		"--ar-ns",
		"--quick-sight-",
	}

	for _, p := range files {
		b, err := os.ReadFile(p)
		if err != nil {
			t.Fatalf("read %s: %v", p, err)
		}
		s := string(b)
		for _, bad := range badFragments {
			if strings.Contains(s, bad) {
				t.Fatalf("found broken flag fragment %q in %s", bad, p)
			}
		}
	}
}
