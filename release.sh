#!/usr/bin/env bash
set -euo pipefail

CURRENT=$(cat VERSION | tr -d '[:space:]')
echo "Current version: $CURRENT"

# Parse semver (strip leading 'v')
BASE="${CURRENT#v}"
IFS='.' read -r MAJOR MINOR PATCH <<< "$BASE"

# Default bump type
BUMP="${1:-patch}"

case "$BUMP" in
  major) MAJOR=$((MAJOR + 1)); MINOR=0; PATCH=0 ;;
  minor) MINOR=$((MINOR + 1)); PATCH=0 ;;
  patch) PATCH=$((PATCH + 1)) ;;
  *)
    echo "Usage: $0 [major|minor|patch]  (default: patch)"
    exit 1
    ;;
esac

NEW="v${MAJOR}.${MINOR}.${PATCH}"
echo "New version:     $NEW"

# Confirm
read -rp "Commit, tag, and push $NEW? [y/N] " CONFIRM
[[ "$CONFIRM" =~ ^[Yy]$ ]] || { echo "Aborted."; exit 0; }

# Write new version
echo "$NEW" > VERSION

# Stage all changes (including VERSION)
git add -A

# Commit
git commit -m "chore: release $NEW"

# Tag
git tag "$NEW"

# Push commit + tag
git push origin HEAD
git push origin "$NEW"

echo "Released $NEW"
