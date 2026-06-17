#!/usr/bin/env bash
set -euo pipefail

usage() {
	cat <<'USAGE'
Usage:
  scripts/private-first-push.sh [--repo NAME] [--remote URL] [--message MSG] [--tag TAG] [--skip-tests] [--no-push]

Purpose:
  Prepare and push this project to an existing private GitHub repo using your
  default Git user.name/user.email and GitHub SSH account.

Notes:
  - The GitHub repo must already exist and should be private.
  - Do not initialize the GitHub repo with README/license/gitignore.
  - If --remote is omitted, the script uses git@github.com:<ssh-user>/<repo>.git.
USAGE
}

repo_name=""
remote_url=""
commit_message="Initial awsgo parity prototype"
tag_name="pre-split-default"
skip_tests=0
no_push=0

while [[ $# -gt 0 ]]; do
	case "$1" in
		--repo)
			repo_name="${2:-}"
			shift 2
			;;
		--remote)
			remote_url="${2:-}"
			shift 2
			;;
		--message)
			commit_message="${2:-}"
			shift 2
			;;
		--tag)
			tag_name="${2:-}"
			shift 2
			;;
		--skip-tests)
			skip_tests=1
			shift
			;;
		--no-push)
			no_push=1
			shift
			;;
		-h|--help)
			usage
			exit 0
			;;
		*)
			if [[ -z "$repo_name" ]]; then
				repo_name="$1"
				shift
			else
				echo "unexpected argument: $1" >&2
				usage >&2
				exit 2
			fi
			;;
	esac
done

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
root="$(cd "$script_dir/.." && pwd)"
cd "$root"

echo "repo root: $root"

nested_git="$(find . -mindepth 2 -type d -name .git -print)"
if [[ -n "$nested_git" ]]; then
	echo "nested Git repositories found; remove these before pushing:" >&2
	echo "$nested_git" >&2
	exit 1
fi

if ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
	git init
fi
git branch -M main

global_name="$(git config --global --get user.name || true)"
global_email="$(git config --global --get user.email || true)"
if [[ -z "$global_name" || -z "$global_email" ]]; then
	echo "global Git identity is not configured." >&2
	echo "Set it first:" >&2
	echo '  git config --global user.name "Your Name"' >&2
	echo '  git config --global user.email "you@example.com"' >&2
	exit 1
fi

git config user.name "$global_name"
git config user.email "$global_email"
echo "commit author: $(git config --get user.name) <$(git config --get user.email)>"

tracked_artifacts="$(git ls-files bin .gocache .gomodcache 2>/dev/null || true)"
if [[ -n "$tracked_artifacts" ]]; then
	echo "build/cache artifacts are tracked; remove them from the index before pushing:" >&2
	echo "$tracked_artifacts" >&2
	exit 1
fi

if [[ "$skip_tests" -eq 0 ]]; then
	if [[ -f Makefile ]]; then
		make test-safe
	else
		echo "Makefile not found; skipping test-safe."
	fi
fi

git add .

if git diff --cached --quiet && git rev-parse --verify HEAD >/dev/null 2>&1; then
	echo "no staged changes to commit."
else
	git commit -m "$commit_message"
fi

head_commit="$(git rev-parse HEAD)"
if git rev-parse -q --verify "refs/tags/$tag_name" >/dev/null; then
	tag_commit="$(git rev-list -n 1 "$tag_name")"
	if [[ "$tag_commit" != "$head_commit" ]]; then
		echo "moving local tag $tag_name to HEAD"
		git tag -f "$tag_name"
	fi
else
	git tag "$tag_name"
fi

if [[ -z "$remote_url" ]]; then
	if git remote get-url origin >/dev/null 2>&1; then
		remote_url="$(git remote get-url origin)"
	else
		if [[ -z "$repo_name" ]]; then
			read -r -p "Private GitHub repo name: " repo_name
		fi
		if [[ -z "$repo_name" ]]; then
			echo "repo name is required." >&2
			exit 1
		fi

		ssh_output="$(ssh -T git@github.com 2>&1 || true)"
		github_user="$(printf '%s\n' "$ssh_output" | sed -n 's/^Hi \([^!]*\)!.*/\1/p' | head -n 1)"
		if [[ -z "$github_user" ]]; then
			echo "$ssh_output" >&2
			read -r -p "GitHub username: " github_user
		fi
		if [[ -z "$github_user" ]]; then
			echo "GitHub username is required." >&2
			exit 1
		fi
		remote_url="git@github.com:${github_user}/${repo_name}.git"
	fi
fi

if git remote get-url origin >/dev/null 2>&1; then
	git remote set-url origin "$remote_url"
else
	git remote add origin "$remote_url"
fi
echo "origin: $(git remote get-url origin)"

if [[ "$no_push" -eq 1 ]]; then
	echo "--no-push set; stopping before network push."
	exit 0
fi

git push -u origin main
git push origin "$tag_name"
