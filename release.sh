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

# Update go.mod module path for major versions >= 2 (Go semver rule)
if [[ "$BUMP" == "major" && "$MAJOR" -ge 2 ]]; then
  OLD_MAJOR=$((MAJOR - 1))
  MODULE=$(grep '^module ' go.mod | awk '{print $2}')

  if [[ "$OLD_MAJOR" -ge 2 ]]; then
    # e.g. v2 -> v3: replace /v2 with /v3
    NEW_MODULE="${MODULE%/v${OLD_MAJOR}}/v${MAJOR}"
  else
    # e.g. v1 -> v2: append /v2
    NEW_MODULE="${MODULE}/v${MAJOR}"
  fi

  sed -i '' "s|^module .*|module ${NEW_MODULE}|" go.mod
  echo "Updated go.mod: ${MODULE} → ${NEW_MODULE}"
  echo ""
  echo "  NOTE: Update import paths in all consumer services:"
  echo "    find . -name '*.go' -exec sed -i '' 's|${MODULE}/|${NEW_MODULE}/|g' {} +"
  echo ""
fi

# Confirm
read -rp "Commit, tag, and push $NEW? [y/N] " CONFIRM
[[ "$CONFIRM" =~ ^[Yy]$ ]] || { echo "Aborted."; exit 0; }

# Write new version
echo "$NEW" > VERSION

# Stage all changes (including VERSION and go.mod)
git add -A

# Commit
git commit -m "chore: release $NEW"

# Tag
git tag "$NEW"

# Push commit + tag
git push origin HEAD
git push origin "$NEW"

echo "Released $NEW"
