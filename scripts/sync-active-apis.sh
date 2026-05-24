#!/usr/bin/env bash
# Mirror internal/cli/active-apis.json to the datpaq/mcp repo.
#
# Usage: ./scripts/sync-active-apis.sh [path-to-mcp-repo]
# Default target: ../MCP (sibling layout)
#
# Workflow: edit the file in this repo, commit, then run this script.
# It copies the file to the MCP repo working tree but does NOT commit
# there — review and commit in the MCP repo manually.

set -euo pipefail

ROOT="$(git rev-parse --show-toplevel)"
SOURCE="$ROOT/internal/cli/active-apis.json"
TARGET_ROOT="${1:-$ROOT/../MCP}"
TARGET="$TARGET_ROOT/internal/cli/active-apis.json"

if [[ ! -f "$SOURCE" ]]; then
  echo "error: source not found: $SOURCE" >&2
  exit 1
fi

if [[ ! -d "$TARGET_ROOT" ]]; then
  echo "error: target repo not found: $TARGET_ROOT" >&2
  echo "       pass the path as the first argument" >&2
  exit 1
fi

if [[ ! -f "$TARGET" ]]; then
  echo "error: target file does not exist: $TARGET" >&2
  exit 1
fi

if cmp -s "$SOURCE" "$TARGET"; then
  echo "already in sync."
  exit 0
fi

cp "$SOURCE" "$TARGET"
echo "synced: $SOURCE"
echo "    -> $TARGET"
echo ""
echo "Next: cd $TARGET_ROOT && git add internal/cli/active-apis.json && git commit"
