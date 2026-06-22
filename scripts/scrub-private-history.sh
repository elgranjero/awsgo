#!/usr/bin/env bash
set -euo pipefail

usage() {
	cat <<'USAGE'
Usage:
  scripts/scrub-private-history.sh [--name NAME] [--email EMAIL] [--checkpoint MESSAGE] [--check-only]

Purpose:
  Rewrite local Git history before public release so private local handles,
  organization strings, and local absolute paths are not present in reachable
  commits, paths, tags, reflogs, or backup refs.

What it does:
  - Optionally commits current work with a sanitized local Git identity.
  - Removes generated/accessanalyzer/tidy from every reachable commit.
  - Rewrites commit author/committer identity to the sanitized values.
  - Deletes filter-branch backup refs and expires reflogs.
  - Runs Git GC with prune=now.
  - Verifies reachable history no longer contains the blocked strings.

Notes:
  - This is local-only. If the old history was pushed, force-push afterward.
  - Run from the repo root or any subdirectory.
USAGE
}

safe_name="El Granjero"
safe_email="elgranjero@example.invalid"
checkpoint_message=""
check_only=0

while [[ $# -gt 0 ]]; do
	case "$1" in
		--name)
			safe_name="${2:-}"
			shift 2
			;;
		--email)
			safe_email="${2:-}"
			shift 2
			;;
		--checkpoint)
			checkpoint_message="${2:-}"
			shift 2
			;;
		--check-only)
			check_only=1
			shift
			;;
		-h|--help)
			usage
			exit 0
			;;
		*)
			echo "unknown argument: $1" >&2
			usage >&2
			exit 2
			;;
	esac
done

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
root="$(cd "$script_dir/.." && pwd)"
cd "$root"

bad_user="cro""bles"
bad_person="chris"".""ro""bles"
bad_org="pp""fa"
bad_patterns=(
	"$bad_user"
	"$bad_person"
	"$bad_org-prod"
	"@$bad_org.org"
	"$bad_org.org"
	"/Users/$bad_user"
)

print_matches() {
	local found=1 pattern
	for pattern in "${bad_patterns[@]}"; do
		if grep -R -I -n -i --exclude-dir=.git --exclude-dir=.gocache --exclude-dir=.gomodcache --exclude-dir=bin --exclude-dir=benchmarks -e "$pattern" .; then
			found=0
		fi
	done
	return "$found"
}

history_matches() {
	local found=1 pattern commits
	commits="$(git rev-list --all 2>/dev/null || true)"
	for pattern in "${bad_patterns[@]}"; do
		if git log --all --format='%H%x09%an <%ae>%x09%cn <%ce>%x09%s' | grep -i -e "$pattern"; then
			found=0
		fi
		if [[ -n "$commits" ]] && git grep -I -n -i -e "$pattern" $commits -- 2>/dev/null; then
			found=0
		fi
		if git rev-list --objects --all | grep -i -e "$pattern"; then
			found=0
		fi
	done
	return "$found"
}

if [[ "$check_only" -eq 1 ]]; then
	status=0
	if print_matches; then
		status=1
	fi
	if history_matches; then
		status=1
	fi
	exit "$status"
fi

if print_matches; then
	echo "blocked strings are still present in the current working tree; fix those before rewriting history." >&2
	exit 1
fi

git config user.name "$safe_name"
git config user.email "$safe_email"

if [[ -n "$(git status --porcelain)" ]]; then
	if [[ -z "$checkpoint_message" ]]; then
		echo "working tree has changes. Either commit/stash them first, or pass --checkpoint MESSAGE." >&2
		exit 1
	fi
	git add -A
	GIT_AUTHOR_NAME="$safe_name" \
	GIT_AUTHOR_EMAIL="$safe_email" \
	GIT_COMMITTER_NAME="$safe_name" \
	GIT_COMMITTER_EMAIL="$safe_email" \
		git commit -m "$checkpoint_message"
fi

export SAFE_NAME="$safe_name"
export SAFE_EMAIL="$safe_email"
export BAD_USER="$bad_user"
export BAD_PERSON="$bad_person"
export BAD_ORG="$bad_org"

privacy_guard_template="$(mktemp "${TMPDIR:-/tmp}/awsgo-privacy-guard.XXXXXX")"
cat >"$privacy_guard_template" <<'GO'
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
GO
export PRIVACY_GUARD_TEMPLATE="$privacy_guard_template"
trap 'rm -f "$privacy_guard_template"' EXIT

FILTER_BRANCH_SQUELCH_WARNING=1 git filter-branch --force \
	--env-filter '
GIT_AUTHOR_NAME="$SAFE_NAME"
GIT_AUTHOR_EMAIL="$SAFE_EMAIL"
GIT_COMMITTER_NAME="$SAFE_NAME"
GIT_COMMITTER_EMAIL="$SAFE_EMAIL"
export GIT_AUTHOR_NAME GIT_AUTHOR_EMAIL GIT_COMMITTER_NAME GIT_COMMITTER_EMAIL
' \
	--tree-filter '
rm -f generated/accessanalyzer/tidy
if [ -f tests/privacy_guard_test.go ]; then
	cp "$PRIVACY_GUARD_TEMPLATE" tests/privacy_guard_test.go
fi
' \
	--tag-name-filter cat -- --all

while IFS= read -r ref; do
	git update-ref -d "$ref"
done < <(git for-each-ref --format='%(refname)' refs/original/)

git reflog expire --expire=now --expire-unreachable=now --all
git gc --prune=now --aggressive

if history_matches; then
	echo "history scrub verification failed; blocked strings remain in reachable history." >&2
	exit 1
fi

if grep -R -I -n -i -e "$bad_user" -e "$bad_person" -e "$bad_org-prod" -e "@$bad_org.org" -e "$bad_org.org" .git 2>/dev/null; then
	echo "blocked strings remain somewhere under .git after GC." >&2
	exit 1
fi

cat <<DONE
History scrub complete.

If this repo was already pushed, update the private remote with:
  git push --force-with-lease origin main
  git push --force origin --tags

Then ask GitHub support to purge cached unreachable objects before making it public,
or create a brand-new public repo from this scrubbed checkout.
DONE
