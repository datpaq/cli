# Datpaq CLI — Release Process

How to cut a tagged release of the CLI and publish it to Homebrew once
`github.com/datpaq/cli` is public and `github.com/datpaq/homebrew-tap`
exists.

This is operational documentation for your future self. Nothing in here
needs to happen for everyday development — `go install` and `go build`
work right now without any of it.

---

## Prerequisites (one-time setup)

### 1. Create the GitHub repos

Two repos under the `datpaq` org:

| Repo | Purpose | Visibility |
| --- | --- | --- |
| `github.com/datpaq/cli` | This CLI (the source you're reading) | Public |
| `github.com/datpaq/homebrew-tap` | Homebrew formula host | Public |

The tap repo can start empty. `goreleaser` will commit `Formula/datpaq.rb`
into it on every release.

### 2. Push the CLI

```bash
cd ~/Projects/Datpaq/CLI/datpaq
git remote add origin git@github.com:datpaq/cli.git
git push -u origin main
```

### 3. Install goreleaser

```bash
brew install goreleaser
goreleaser --version   # should be >= 2.0
```

### 4. Create a GitHub Personal Access Token (PAT) for the tap

`goreleaser` needs write access to `homebrew-tap` to commit the formula.
Make a fine-grained PAT scoped to that one repo:

1. GitHub → **Settings** → **Developer settings** → **Personal access
   tokens** → **Fine-grained tokens** → **Generate new token**
2. Name: `datpaq-goreleaser-homebrew-tap`
3. Repository access: **Only select repositories** → `datpaq/homebrew-tap`
4. Permissions: **Contents: Read and write**
5. Generate, copy the token, store it somewhere safe (a password manager).
6. Export it whenever you run `goreleaser release`:

```bash
export GITHUB_TOKEN=ghp_xxxxxxxxxxxxxxxxxxxx
```

For day-to-day use, put it in `~/.config/datpaq/release.env` (gitignored
on your machine) and `source` it before releasing.

---

## Cutting a release

### 1. Make sure `main` is clean and pushed

```bash
cd ~/Projects/Datpaq/CLI/datpaq
git status                       # working tree clean
git pull --rebase origin main    # in sync with remote
go build ./... && go vet ./...   # sanity check
```

### 2. Tag the version (semver)

```bash
git tag -a v0.1.0 -m "v0.1.0: first release"
git push origin v0.1.0
```

Conventions:

- `v0.x.y` — pre-1.0, breaking changes allowed between minors.
- `v1.0.0` — first stable. After this, follow strict semver.
- Patch releases (`v0.1.1`, `v1.0.1`) for bug fixes only — no API changes.

### 3. Dry-run the release

This builds for every target in `.goreleaser.yaml` and renders the formula
without publishing anything. Catches misconfiguration before it's
permanent.

```bash
source ~/.config/datpaq/release.env   # exports GITHUB_TOKEN
goreleaser release --snapshot --clean --skip=publish
```

Inspect:
- `dist/` for the built binaries (darwin/linux/windows × amd64/arm64)
- `dist/datpaq.rb` for the rendered Homebrew formula

If anything looks wrong, fix it on `main`, delete the tag locally and on
GitHub, and re-tag:

```bash
git tag -d v0.1.0
git push origin :refs/tags/v0.1.0
# fix things, commit, push
git tag -a v0.1.0 -m "v0.1.0: first release"
git push origin v0.1.0
```

### 4. Real release

```bash
source ~/.config/datpaq/release.env
goreleaser release --clean
```

This will:

1. Build `datpaq` and `datpaq-mcp` for darwin/linux/windows × amd64/arm64
   (12 binaries total).
2. Package each as a `.tar.gz` (Windows: `.zip`) named per
   `.goreleaser.yaml`'s template.
3. Compute SHA-256 checksums into `checksums.txt`.
4. Create a GitHub Release on `github.com/datpaq/cli` tagged `v0.1.0`
   with all the archives attached.
5. Commit a fresh `Formula/datpaq.rb` to
   `github.com/datpaq/homebrew-tap`, pointing at the new archives.

Total runtime: ~2-3 minutes.

### 5. Verify

```bash
brew tap datpaq/tap                # one-time per user
brew install datpaq                # pulls the latest formula
datpaq version                     # should print v0.1.0
```

After tapping once, users get future releases via:

```bash
brew upgrade datpaq
```

And the non-Homebrew install path keeps working in parallel:

```bash
go install github.com/datpaq/cli/cmd/datpaq@latest
```

---

## What `.goreleaser.yaml` already has wired

Lives at the repo root, set up during CLI generation. The current
configuration:

- Builds two binaries: `datpaq` (the CLI) and `datpaq-mcp` (the MCP
  server bundle, useful for Claude Desktop and other MCP clients).
- 6 platforms each: darwin/linux/windows × amd64/arm64.
- `CGO_ENABLED=0` so binaries are statically linked and need no system
  libraries on the target.
- `ldflags: -s -w` strips debug info to shrink binaries.
- Embeds the build version via `-X github.com/datpaq/cli/internal/cli.version={{ .Version }}`
  so `datpaq version` reports the tag.
- Brew formula goes to `datpaq/homebrew-tap`, installs both binaries.

Nothing in `.goreleaser.yaml` should need changing for routine releases.
If you ever want to:

- **Drop the Homebrew tap** — delete the `brews:` block at the bottom.
- **Add Windows MSI / Linux .deb / RPM packaging** — add `nfpms:` and
  `chocolateys:` blocks (see goreleaser docs).
- **Add Docker images** — add a `dockers:` block.

---

## Rollback

If a release is broken and you need to yank it:

1. Delete the GitHub Release (web UI) and the tag:

   ```bash
   git push origin :refs/tags/v0.1.0
   git tag -d v0.1.0
   ```

2. Revert the formula commit on `homebrew-tap`:

   ```bash
   cd /path/to/homebrew-tap
   git revert HEAD
   git push
   ```

3. Users on the broken version can downgrade with:

   ```bash
   brew uninstall datpaq
   brew install datpaq@<previous-version>   # if you tag prior versions explicitly
   ```

   …or just wait for the next good release and `brew upgrade datpaq`.

Better: catch issues with `--snapshot --skip=publish` before tagging.

---

## CI automation (optional, later)

Once releases happen often enough to be tedious, move them into GitHub
Actions. Skeleton:

```yaml
# .github/workflows/release.yml
name: release
on:
  push:
    tags: ["v*"]
jobs:
  release:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - uses: actions/setup-go@v5
        with: { go-version: stable }
      - uses: goreleaser/goreleaser-action@v6
        with: { args: release --clean }
        env:
          GITHUB_TOKEN: ${{ secrets.HOMEBREW_TAP_TOKEN }}
```

Add `HOMEBREW_TAP_TOKEN` as a repo secret (same fine-grained PAT from
step 4 above). Then a release is just `git tag … && git push --tags`.

---

## Quick reference

```bash
# Cut a release
git tag -a vX.Y.Z -m "vX.Y.Z: description"
git push origin vX.Y.Z
source ~/.config/datpaq/release.env
goreleaser release --clean

# Dry-run first if unsure
goreleaser release --snapshot --clean --skip=publish

# Install via Homebrew (end users)
brew tap datpaq/tap
brew install datpaq

# Install via Go (end users)
go install github.com/datpaq/cli/cmd/datpaq@latest
```
