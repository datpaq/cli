<div align="center">

<img src="docs/cli-terminal.webp" alt="Datpaq CLI" width="720">

# Datpaq CLI

**The official command-line interface for the [Datpaq API](https://datpaq.com).**

[![Go Reference](https://pkg.go.dev/badge/github.com/datpaq/cli.svg)](https://pkg.go.dev/github.com/datpaq/cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/datpaq/cli)](https://goreportcard.com/report/github.com/datpaq/cli)
[![Latest Release](https://img.shields.io/github/v/release/datpaq/cli?color=6b21a8)](https://github.com/datpaq/cli/releases/latest)
[![License](https://img.shields.io/badge/license-Apache_2.0-6b21a8.svg)](LICENSE)

[Install](#install) · [Updating](#updating) · [Quickstart](#quickstart) · [Documentation](#documentation) · [Dashboard ↗](https://datpaq.com)

</div>

---

## Install

```bash
go install github.com/datpaq/cli/cmd/datpaq@latest
```

Requires Go 1.26.3+. Once installed, `datpaq` is on your `PATH` (assuming `$HOME/go/bin` is exported).

**Windows (Node.js / npm):**

```powershell
npm install -g datpaq
```

Requires Node.js 18+. Once installed, `datpaq` is available in your terminal. Verify with `datpaq --version`.

Homebrew install coming soon:

```bash
brew install datpaq/tap/datpaq          # not yet available
```

## Updating

**From a cloned repo** — pull and reinstall from the repo root:

```bash
git pull
go install ./cmd/datpaq
```

Verify with `which datpaq` and `datpaq --version`. The binary lands in `$HOME/go/bin` (see [Install](#install) for `PATH` setup).

**From a release** (no local clone):

```bash
go install github.com/datpaq/cli/cmd/datpaq@latest
```

## Quickstart

```bash
datpaq auth login                            # sign in via browser (Clerk SSO)
datpaq api                                   # list available APIs
datpaq ip-geolocation --ip 8.8.8.8           # call any endpoint
datpaq sample ip-geolocation --lang py       # generate copy-paste code
```

That's the loop: discover → call → integrate.

## What's inside

- **88 endpoints across 13 APIs**, one consistent interface — geolocation, image processing, sample data, WHOIS, working days, and more.
- **`datpaq api`** — browse every endpoint by interface, with descriptions pulled from the live OpenAPI spec.
- **`datpaq sample`** — emit ready-to-paste snippets in `curl`, `JavaScript`, `Python`, or `Go`. Uses `$DATPAQ_API_KEY` as the credential placeholder so snippets are safe to share.
- **Browser-based SSO** via `datpaq auth login` — no copy-pasting tokens from the dashboard.
- **Agent-friendly mode** — append `--agent` to any command for JSON output, non-interactive defaults, and stable exit codes. Drop it into a Claude Code workflow or a CI job.
- **Bundled MCP server** (`datpaq-mcp`) — every command auto-mirrored as an MCP tool for Claude Desktop and other MCP clients. For a remote/hosted MCP server, see [`github.com/datpaq/mcp`](https://github.com/datpaq/mcp).
- **Generated, not hand-stitched** — built with [`cli-printing-press`](https://github.com/mvanhorn/cli-printing-press) from the live OpenAPI spec. New endpoints land in the CLI when the spec updates.

## Authentication

Three paths, listed in order of how a real developer typically uses them:

```bash
# 1. Browser-based SSO (recommended for human use)
datpaq auth login

# 2. Paste an API key from https://datpaq.com (CI, headless, or SSH sessions)
datpaq auth set-token YOUR_KEY

# 3. Environment variable (CI, secret-managers, ephemeral runs)
export DATPAQ_API_KEY=YOUR_KEY
```

Credentials persist to `~/.config/datpaq/config.toml` (mode `0600`, never committed). The `DATPAQ_API_KEY` env var always wins when set. Run `datpaq doctor` to verify auth and connectivity end-to-end.

## A few things you can do

**Look up an aircraft tail number:**

```bash
$ datpaq aircraft lookup-by-tail --tail N12345
```

**Validate a batch of email addresses from a file:**

```bash
$ datpaq email-validation validate-batch --emails-file ./list.txt --json | jq '.[] | select(.deliverable == false)'
```

**Resize an image inline:**

```bash
$ datpaq image-processing image-resize --url https://example.com/photo.jpg --width 800 > photo-800.jpg
```

**Generate Python code for the IP intelligence endpoint:**

```bash
$ datpaq sample ip-intelligence get --lang py
```

**Use in a Claude Code agent or CI:**

```bash
$ datpaq dns example.com --agent | jq '.records'
```

## Managing active APIs

The CLI's discovery surface (`datpaq api`, `datpaq sample`, splash) only shows endpoints listed in [`internal/cli/active-apis.json`](internal/cli/active-apis.json). When a new API ships on [datpaq.com](https://datpaq.com), add its slug there and rebuild.

The same file powers the hosted MCP server at [mcp.datpaq.com](https://mcp.datpaq.com/) and lives at [`github.com/datpaq/mcp`](https://github.com/datpaq/mcp/blob/main/internal/cli/active-apis.json) — keep both in sync. If the MCP repo is checked out as a sibling directory, mirror the file with:

```bash
./scripts/sync-active-apis.sh
```

## Documentation

- [docs/cli-auth-spec.md](docs/cli-auth-spec.md) — the website-side OAuth contract for `datpaq auth login`
- [docs/release-process.md](docs/release-process.md) — how to cut a tagged release and publish to Homebrew
- [API reference](https://datpaq.com/docs) — full endpoint documentation on the website
- [Dashboard ↗](https://datpaq.com) — manage your API keys and usage

## Configuration

| Variable | Purpose |
| --- | --- |
| `DATPAQ_API_KEY` | API credential — overrides any value in the config file |
| `DATPAQ_BASE_URL` | Override the API base URL (default: `https://datpaq.com/api/v1`) — handy for pointing at staging |
| `DATPAQ_CONFIG` | Override the config file path (default: `~/.config/datpaq/config.toml`) |

## Development

```bash
git clone https://github.com/datpaq/cli && cd cli
go build ./cmd/datpaq
./datpaq doctor
```

The CLI is generated from the OpenAPI spec at `spec.yaml`. To regenerate after the spec changes:

```bash
printing-press generate --spec spec.yaml --name datpaq --force
```

`--force` uses AST-based merge to preserve hand-edits (the OAuth `auth login` command, the gradient banner, the splash menu, the `sample` command, the active-APIs filter).

## License

Apache 2.0 — see [LICENSE](LICENSE).

Generated by [CLI Printing Press](https://github.com/mvanhorn/cli-printing-press).
