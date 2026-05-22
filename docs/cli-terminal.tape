# Hero banner capture for the Datpaq CLI.
#
# Narrative — what a new developer sees end-to-end:
#   1. `datpaq`                  — DATPAQ banner + the splash menu (the cold
#                                  open: this is what they see on first run)
#   2. `datpaq exec list`        — the curated catalog with the excluded footer
#   3. `datpaq exec whois.lookup` — interactive prompt UX (header + the
#                                   modern `?` prompt) and the pretty-JSON
#                                   dry-run envelope
#
# Outputs (both written on every run):
#   docs/cli-terminal.gif  — intermediate; converted to .webp by the
#                            generate-cli-screenshots.sh helper at the repo root.
#   docs/cli-terminal.png  — static frame of the final scene (the JSON
#                            payload). Suitable for social share cards.
#
# Why GIF instead of native vhs WebP output: Homebrew's stock ffmpeg
# (since 8.x) is built without libwebp, so vhs's `Output …webp` fails
# with "Encoder not found". GIF works on stock ffmpeg, and `gif2webp`
# (from `brew install webp`) converts it after.
#
# Use the helper at the repo root: ./generate-cli-screenshots.sh
# Or run directly:
#   vhs docs/cli-terminal.tape
#   gif2webp -lossy -q 80 docs/cli-terminal.gif -o docs/cli-terminal.webp

Output docs/cli-terminal.gif
Output docs/cli-terminal.png

# ── Look & feel ───────────────────────────────────────────────────────────
# 1200×800 leaves headroom for the DATPAQ ASCII banner + splash menu + a
# scene of output without the banner scrolling off too quickly. The README's
# <img width="720"> downsamples; rendering large keeps text crisp on retina.
Set Width 1200
Set Height 800
Set FontFamily "JetBrains Mono"
Set FontSize 18
Set Padding 24
Set Theme "Catppuccin Mocha"
Set TypingSpeed 55ms
Set Shell "zsh"

# ── Scene 0: cold open — the DATPAQ welcome banner ────────────────────────
# Hidden clear so the recording starts on a fresh prompt rather than
# whatever the user had in their shell history. The visible recording
# begins with the developer typing `datpaq`.
Hide
Type "clear"
Enter
Sleep 200ms
Show
Sleep 500ms

Type "datpaq"
Sleep 400ms
Enter

# Hold on the banner + splash menu long enough that a viewer can read
# all six pointers (including the new `datpaq exec` entry).
Sleep 3.5s

# ── Scene 1: the curated catalog ──────────────────────────────────────────
# Clear hidden so the transition feels like the developer just hit Cmd+K
# rather than scrolling through a typed `clear`. Keeps the demo seamless.
Hide
Type "clear"
Enter
Sleep 200ms
Show

Type "datpaq exec list"
Sleep 400ms
Enter
Sleep 3.5s

# ── Scene 2: interactive exec with the polished prompt ────────────────────
# --human-friendly enables the green/bold `?` accent on the prompt.
# --dry-run lets the recording run without a real API key and shows the
# pretty-JSON envelope so the final scene is the payload shape.
Hide
Type "clear"
Enter
Sleep 200ms
Show

Type "datpaq --human-friendly --dry-run exec whois.lookup"
Sleep 400ms
Enter

# Header banner + first prompt land here. Give viewers time to read them.
Sleep 1.5s

Type "example.com"
Sleep 400ms
Enter

# Hold the final frame so a static thumbnail of the .png lands on the
# payload — that's the moment that pays off the whole demo.
Sleep 3.5s
