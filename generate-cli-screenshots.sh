#!/usr/bin/env bash
# Regenerate the CLI's banner and social-share images from the vhs tape at
# docs/cli-terminal.tape.
#
# Outputs:
#   docs/cli-terminal.webp  — animated banner used on the landing page README
#   docs/cli-terminal.png   — static final frame (JSON payload), social card
#   docs/cli-terminal.gif   — vhs intermediate; safe to delete or commit
#
# Requirements (one-time):
#   brew install vhs webp
#
# Why this script exists instead of `vhs Output …webp` directly:
# Homebrew's stock ffmpeg (8.x and later) ships without libwebp, so the
# webp encoder isn't available to vhs. We work around it by having vhs
# emit a GIF and converting it with gif2webp (part of libwebp).
#
# Run from anywhere — the script cd's to the repo root itself.

set -euo pipefail

# ── Resolve the repo root regardless of where this script lives ──────────
# Walks symlinks one hop at a time (BSD/macOS `readlink` lacks `-f`, so we
# can't shell out to `readlink -f` or `realpath` portably). After the walk,
# we cd to the script's true directory — works whether the script is
# invoked by absolute path, relative path, on $PATH, or via symlink.
SCRIPT_PATH="${BASH_SOURCE[0]}"
while [ -L "${SCRIPT_PATH}" ]; do
  LINK_TARGET="$(readlink "${SCRIPT_PATH}")"
  case "${LINK_TARGET}" in
    /*) SCRIPT_PATH="${LINK_TARGET}" ;;
    *)  SCRIPT_PATH="$(cd "$(dirname "${SCRIPT_PATH}")" && pwd)/${LINK_TARGET}" ;;
  esac
done
SCRIPT_DIR="$(cd "$(dirname "${SCRIPT_PATH}")" && pwd)"
cd "${SCRIPT_DIR}"

# Verify we landed in the CLI repo. Two markers must both exist; either
# alone could match an unrelated Go project or doc folder.
if [ ! -f "go.mod" ] || [ ! -f "docs/cli-terminal.tape" ]; then
  echo "error: this script must live in the Datpaq CLI repo root" >&2
  echo "  expected go.mod and docs/cli-terminal.tape under: ${SCRIPT_DIR}" >&2
  exit 1
fi

for tool in go vhs gif2webp; do
  if ! command -v "${tool}" >/dev/null 2>&1; then
    echo "error: ${tool} is not on PATH" >&2
    echo "  hint: brew install vhs webp" >&2
    exit 1
  fi
done

# Build into a temp dir so we don't clobber whatever's on the user's PATH.
# The tape itself doesn't reference the binary path — it relies on `datpaq`
# being resolvable from $PATH inside the vhs sub-shell, so we prepend the
# temp dir before invoking vhs.
BIN_DIR="$(mktemp -d)"
trap 'rm -rf "${BIN_DIR}"' EXIT

echo "Building datpaq binary..."
go build -o "${BIN_DIR}/datpaq" ./cmd/datpaq

echo "Recording with vhs (this takes ~30s for the current tape)..."
PATH="${BIN_DIR}:${PATH}" vhs docs/cli-terminal.tape

echo "Converting GIF -> WebP..."
gif2webp -lossy -q 80 docs/cli-terminal.gif -o docs/cli-terminal.webp >/dev/null

# Keep the GIF intermediate by default — it's small (~few hundred KB) and
# useful as a fallback for any reader that doesn't render animated WebP.
# Uncomment to discard it after conversion:
# rm docs/cli-terminal.gif

echo
echo "Done. Updated:"
echo "  docs/cli-terminal.webp  (landing page banner — animated)"
echo "  docs/cli-terminal.png   (social share card — final frame)"
echo "  docs/cli-terminal.gif   (intermediate; safe to delete)"
