#!/usr/bin/env bash
# Finish adding table-declared APIs after writing internal/cli/<slug>.go files.
# Activates the slugs, emits + splices their spec.yaml paths, records a
# .printing-press-patches.json entry, and runs the test suite.
#
# Usage: scripts/api-sync/add-api.sh <patch-id> "<summary>" "<reason>" <slug>...
# See docs/adding-apis.md.

set -euo pipefail
cd "$(git rev-parse --show-toplevel)"
here=scripts/api-sync

[[ $# -ge 4 ]] || { sed -n '2,7p' "$0"; exit 2; }
pid=$1 summary=$2 reason=$3
shift 3

for s in "$@"; do
  [[ -f internal/cli/$s.go ]] || { echo "missing internal/cli/$s.go" >&2; exit 1; }
done

python3 $here/catalog.py activate "$@"

frag=$(mktemp)
trap 'rm -f "$frag"' EXIT
DATPAQ_SPEC_FRAGMENT_OUT="$frag" go test ./internal/cli -run TestEmitDeclaredSpecFragment -count=1 >/dev/null
python3 $here/catalog.py splice-spec "$frag"

files=(internal/cli/active-apis.json internal/cli/active_apis_test.go spec.yaml)
for s in "$@"; do files+=("internal/cli/$s.go"); done
python3 $here/catalog.py manifest "$pid" "$summary" "$reason" \
  "go test ./... passes: table-driven request, required-flag, MCP-param, and spec.yaml parity tests cover every new endpoint." \
  "${files[@]}"

unformatted=$(gofmt -l "${files[@]:3}")
[[ -z $unformatted ]] || { echo "run gofmt -w on: $unformatted" >&2; exit 1; }
go vet ./internal/cli ./internal/mcp
go test ./...
echo
echo "Done. Review 'git diff', then commit. To mirror the active list to the MCP repo: scripts/sync-active-apis.sh"
