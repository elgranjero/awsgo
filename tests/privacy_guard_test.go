package tests

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestNoPersonalIdentifiersLeakInRepoFiles(t *testing.T) {
	root := repoRoot(t)
	patterns := privacyLeakPatterns()
	files := trackedRepoFiles(t, root)
	if len(files) == 0 {
		t.Fatalf("no files discovered for privacy guard")
	}

	for _, p := range files {
		if strings.HasSuffix(filepath.ToSlash(p), "/tests/privacy_guard_test.go") {
			continue
		}
		b, err := os.ReadFile(p)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			t.Fatalf("read %s: %v", p, err)
		}
		s := strings.ToLower(string(b))
		for _, pat := range patterns {
			if strings.Contains(s, strings.ToLower(pat)) {
				t.Fatalf("privacy leak pattern %q found in %s", pat, p)
			}
		}
	}
}

func privacyLeakPatterns() []string {
	user := "cro" + "bles"
	org := "pp" + "fa"
	return []string{
		user,
		"chris" + "." + "ro" + "bles",
		org + "-prod",
		"@" + org + ".org",
		org + ".org",
		filepath.ToSlash(filepath.Join("/Users", user)),
	}
}

func trackedRepoFiles(t *testing.T, root string) []string {
	t.Helper()
	cmd := exec.Command("git", "-C", root, "ls-files", "-z")
	out, err := cmd.Output()
	if err == nil {
		files := []string{}
		for _, rel := range strings.Split(string(out), "\x00") {
			rel = strings.TrimSpace(rel)
			if rel == "" {
				continue
			}
			files = append(files, filepath.Join(root, rel))
		}
		return files
	}

	files := []string{}
	err = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			switch d.Name() {
			case ".git", ".gocache", ".gomodcache", "bin":
				return filepath.SkipDir
			default:
				return nil
			}
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		t.Fatalf("discover files for privacy guard: %v", err)
	}
	return files
}
