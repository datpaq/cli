# Datpaq CLI

Comprehensive multi-service API platform providing aviation, geolocation, financial, utilities, image processing, and data enrichment capabilities. All endpoints require authentication via API key passed as an x-api-key header or api_key query parameter.

Learn more at [Datpaq](https://datpaq.com).

## Install

The recommended path installs both the `datpaq` binary and the `pp-datpaq` agent skill (Claude Code, Codex, Cursor, Gemini CLI, GitHub Copilot, and other agents supported by the upstream [`skills`](https://github.com/vercel-labs/skills) CLI) in one shot:

```bash
npx -y @mvanhorn/printing-press install datpaq
```

For CLI only (no skill):

```bash
npx -y @mvanhorn/printing-press install datpaq --cli-only
```

For skill only — installs the skill into the same agents as the default command above, but skips the CLI binary (use this to update or reinstall just the skill):

```bash
npx -y @mvanhorn/printing-press install datpaq --skill-only
```

To constrain the skill install to one or more specific agents (repeatable — agent names match the [`skills`](https://github.com/vercel-labs/skills) CLI):

```bash
npx -y @mvanhorn/printing-press install datpaq --agent claude-code
npx -y @mvanhorn/printing-press install datpaq --agent claude-code --agent codex
```

### Without Node

The generated install path is category-agnostic until this CLI is published. If `npx` is not available before publish, install Node or use the category-specific Go fallback from the public-library entry after publish.

### Pre-built binary

Download a pre-built binary for your platform from the [latest release](https://github.com/mvanhorn/printing-press-library/releases/tag/datpaq-current). On macOS, clear the Gatekeeper quarantine: `xattr -d com.apple.quarantine <binary>`. On Unix, mark it executable: `chmod +x <binary>`.

<!-- pp-hermes-install-anchor -->
## Install for Hermes

From the Hermes CLI:

```bash
hermes skills install mvanhorn/printing-press-library/cli-skills/pp-datpaq --force
```

Inside a Hermes chat session:

```bash
/skills install mvanhorn/printing-press-library/cli-skills/pp-datpaq --force
```

## Install for OpenClaw

Tell your OpenClaw agent (copy this):

```
Install the pp-datpaq skill from https://github.com/mvanhorn/printing-press-library/tree/main/cli-skills/pp-datpaq. The skill defines how its required CLI can be installed.
```

## Use with Claude Desktop

This CLI ships an [MCPB](https://github.com/modelcontextprotocol/mcpb) bundle — Claude Desktop's standard format for one-click MCP extension installs (no JSON config required).

To install:

1. Download the `.mcpb` for your platform from the [latest release](https://github.com/mvanhorn/printing-press-library/releases/tag/datpaq-current).
2. Double-click the `.mcpb` file. Claude Desktop opens and walks you through the install.
3. Fill in `DATPAQ_API_KEY_HEADER` when Claude Desktop prompts you.

Requires Claude Desktop 1.0.0 or later. Pre-built bundles ship for macOS Apple Silicon (`darwin-arm64`) and Windows (`amd64`, `arm64`); for other platforms, use the manual config below.

<details>
<summary>Manual JSON config (advanced)</summary>

If you can't use the MCPB bundle (older Claude Desktop, unsupported platform), install the MCP binary and configure it manually.


Install the MCP binary from this CLI's published public-library entry or pre-built release.

Add to your Claude Desktop config (`~/Library/Application Support/Claude/claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "datpaq": {
      "command": "datpaq-mcp",
      "env": {
        "DATPAQ_API_KEY_HEADER": "<your-key>"
      }
    }
  }
}
```

</details>

## Quick Start

### 1. Install

See [Install](#install) above.

### 2. Set Up Credentials

Get your API key from your API provider's developer portal. The key typically looks like a long alphanumeric string.

```bash
export DATPAQ_API_KEY_HEADER="<paste-your-key>"
```

You can also persist this in your config file at `~/.config/datpaq-proapi-pp-cli/config.toml`.

### 3. Verify Setup

```bash
datpaq doctor
```

This checks your configuration and credentials.

### 4. Try Your First Command

```bash
datpaq domain-lookup get --domain example-value
```

## Usage

Run `datpaq --help` for the full command reference and flag list.

## Commands

### aircraft

FAA aircraft registry lookups by tail number, ICAO hex, or type code

- **`datpaq aircraft batch-lookup`** - Batch aircraft lookups by tail number
- **`datpaq aircraft lookup-by-icao`** - Look up aircraft by ICAO hex code
- **`datpaq aircraft lookup-by-tail`** - Returns aircraft ownership, manufacturer, model, type, and year from the FAA registry using an N-number (tail number).
- **`datpaq aircraft lookup-by-type`** - Look up aircraft type specifications

### company-enrichment

Company data enrichment by name or domain

- **`datpaq company-enrichment`** - Returns structured company data including industry, size, social profiles, and contact information. Provide company_name, domain, or both.

### convert-time

Timezone-aware datetime conversion

- **`datpaq convert-time`** - Converts a datetime value from a source timezone to a target timezone. Accepts ISO 8601 strings, Unix epoch timestamps, or omit sourceTime to convert the current time.

### country-codes

ISO country code lookup and metadata

- **`datpaq country-codes country-by-iso`** - Get a country by ISO 2 or ISO 3 code
- **`datpaq country-codes country-export`** - Export country data in bulk
- **`datpaq country-codes country-list`** - List all countries with ISO codes and metadata
- **`datpaq country-codes country-search`** - Search countries by full or partial name

### current-time

Current time by IANA timezone identifier

- **`datpaq current-time`** - Returns the current date/time with UTC offset and DST status for a given timezone.

### define

Manage define

- **`datpaq define`** - Returns definitions, part of speech, pronunciation, and etymology for the given word.

### dns

Manage dns

- **`datpaq dns`** - Supports record types A, AAAA, MX, CNAME, TXT, NS, SOA, and ALL.

### domain-lookup

Domain availability and registration information

- **`datpaq domain-lookup get`** - Check domain availability and registration info
- **`datpaq domain-lookup post`** - Check domain availability (POST)

### email-validation

Email address format and deliverability validation

- **`datpaq email-validation email-validate-batch`** - Validate multiple email addresses
- **`datpaq email-validation email-validate-single`** - Checks email format, domain MX records, and optional deliverability. Detects disposable/temporary address providers.

### ev-charger

Electric vehicle charging station discovery and status

- **`datpaq ev-charger ev-networks`** - List all supported charging networks
- **`datpaq ev-charger ev-station-by-id`** - Get specific charging station by ID
- **`datpaq ev-charger ev-station-status`** - Get real-time station status
- **`datpaq ev-charger ev-stations-by-address`** - Search EV charging stations by address
- **`datpaq ev-charger ev-stations-by-coords`** - Find EV charging stations near coordinates

### exchange-rates-and-currency

Manage exchange rates and currency

- **`datpaq exchange-rates-and-currency exchange-rate-bulk`** - Batch currency conversions
- **`datpaq exchange-rates-and-currency exchange-rate-get`** - Returns the current exchange rate between two currencies. Provide an amount to perform conversion. Supports fiat (ISO 4217) and crypto (BTC, ETH).

### generate

Manage generate

- **`datpaq generate`** - Generate realistic mock data using a JSON request body. Recommended for full control over type, count, fields, and format.

### generate-batch

Manage generate batch

- **`datpaq generate-batch`** - Submit multiple sample-data generation requests in one call.

### helicopter

Helicopter registry lookups by registration, ICAO, or serial number

- **`datpaq helicopter lookup-by-icao`** - Look up helicopter by ICAO hex code
- **`datpaq helicopter lookup-by-registration`** - Look up helicopter by registration number
- **`datpaq helicopter lookup-by-serial`** - Look up helicopter by manufacturer serial number
- **`datpaq helicopter manufacturers`** - List helicopter manufacturers

### image-processing

Image resize, compress, convert, crop, and pipeline operations

- **`datpaq image-processing image-compress`** - Compress an image
- **`datpaq image-processing image-convert`** - Convert image to a different format
- **`datpaq image-processing image-crop`** - Crop an image to a region
- **`datpaq image-processing image-metadata`** - Extract metadata from an image
- **`datpaq image-processing image-pipeline`** - Executes an ordered array of image operations in a single request. Recommended over individual endpoints when multiple operations are needed. Supports resize, compress, convert, crop, and metadata.
- **`datpaq image-processing image-resize`** - Resize an image to specified dimensions
- **`datpaq image-processing info`** - Returns supported image operations, formats, and service metadata.

### ip-geolocation

Geographic location data for IPv4 and IPv6 addresses

- **`datpaq ip-geolocation batch`** - Batch IP geolocation lookup
- **`datpaq ip-geolocation get`** - Returns country, region, city, coordinates, ISP, ASN, and timezone for any IPv4 or IPv6 address.
- **`datpaq ip-geolocation post`** - Get geographic location for an IP address (POST)

### ip-intelligence

Threat intelligence and risk scoring for IP addresses

- **`datpaq ip-intelligence batch`** - Batch IP threat intelligence lookup
- **`datpaq ip-intelligence batch-info`** - Returns batch processing limits, example request shape, and response schema details.
- **`datpaq ip-intelligence get`** - Returns trust score, risk classification, VPN/proxy/Tor detection, datacenter detection, and threat categories. Omit the ip parameter to analyze the caller's IP.

### mac-address

MAC address vendor and manufacturer resolution

- **`datpaq mac-address`** - Lookup vendor by full MAC address, 6-character OUI prefix, or organization name. Provide exactly one of the three parameters.

### mx-lookup

Mail exchange record lookup with optional SMTP and SPF validation

- **`datpaq mx-lookup batch`** - Batch MX record lookup for multiple domains
- **`datpaq mx-lookup get`** - Returns mail exchange records with optional IP resolution, SMTP testing, and SPF extraction.
- **`datpaq mx-lookup post`** - Look up MX records for a domain (POST)

### precious-metals

Precious metal and cryptocurrency pricing data

- **`datpaq precious-metals assets`** - Returns metadata for all supported assets — gold (XAU), silver (XAG), palladium (XPD), copper, BTC, ETH, etc.
- **`datpaq precious-metals history`** - Get historical price data for an asset
- **`datpaq precious-metals prices`** - Get current prices for precious metals and crypto

### profanity

Content moderation and profanity detection

- **`datpaq profanity analyze`** - Comprehensive profanity analysis with severity and word positions
- **`datpaq profanity detect`** - Detect whether text contains profanity
- **`datpaq profanity filter`** - Replaces profane words with asterisks. Returns cleaned text.

### public-holidays

Global public holiday data by country and year

- **`datpaq public-holidays countries`** - List all supported countries for holiday data
- **`datpaq public-holidays public-holidays`** - Accepts ISO-2, ISO-3, or full country names. Returns holidays with ISO, US, and EU date formats. Filter by year and/or month.

### sample-data

Mock data generation for testing and development (Faker.js)

- **`datpaq sample-data batch-get`** - Generates mock data using Faker.js. Supports 20+ schema types including user, person, address, company, product, and more. Use the requests parameter as a JSON-encoded array.
- **`datpaq sample-data batch-post`** - Generate multiple types of sample data (POST)
- **`datpaq sample-data generate-get`** - GET-compatible sample data generation endpoint for scripts and browser testing.

### schemas

Manage schemas

- **`datpaq schemas`** - Returns all available sample-data schema names with default field lists.

### secure-relay

Secure HMAC-verified HTTP request forwarding

- **`datpaq secure-relay custom`** - Legacy relay endpoint with extended configuration options for custom header injection and response transformation.
- **`datpaq secure-relay relay`** - Forwards HTTP requests to HTTPS targets with no server-side storage. Requests must include HMAC-SHA256 signature headers for verification. Supports replay attack protection via timestamp validation.

### spell-check

Multi-language spell checking with correction suggestions

- **`datpaq spell-check batch`** - Spell check multiple texts
- **`datpaq spell-check languages`** - List supported languages for spell checking
- **`datpaq spell-check spell-check`** - Detects spelling errors with ranked correction suggestions. Supports English, Spanish, French, German, and Italian. Also catches grammar patterns such as repeated words and spacing errors.

### states

U.S. state metadata by abbreviation or FIPS code

- **`datpaq states by-abbr`** - Get U.S. state by 2-letter abbreviation
- **`datpaq states by-fips`** - Get U.S. state by FIPS code
- **`datpaq states list`** - List or search U.S. states

### text-language

Text language detection supporting 176+ languages

- **`datpaq text-language detect`** - Identifies the language of text with a confidence score. Supports 176+ languages with ISO 639-3 output. Requires at least 10 characters for reliable detection.
- **`datpaq text-language detect-batch`** - Batch language detection for multiple texts
- **`datpaq text-language supported`** - Returns all 176+ supported languages with ISO 639-1 and ISO 639-3 codes.

### thesaurus

Synonyms, antonyms, and related word lookup

- **`datpaq thesaurus`** - Returns synonyms, antonyms, or both for a given word. Filter results by part of speech and limit the number of results.

### unit-conversion

Unit conversion across 100+ measurement types

- **`datpaq unit-conversion unit-categories`** - List available measurement categories
- **`datpaq unit-conversion unit-convert-get`** - Supports 100+ unit types across length, weight, temperature, volume, area, speed, digital storage, and time. Converts value from source unit to target unit.
- **`datpaq unit-conversion unit-convert-post`** - Convert a value between units (POST)
- **`datpaq unit-conversion units-by-category`** - List units available in a category

### user-avatar

Dynamic user avatar generation in PNG, WebP, or SVG

- **`datpaq user-avatar generate`** - Creates avatar images using Canvas (PNG/WebP) or native SVG. Supports initials-based avatars, custom colors, external image integration, and uploaded icons.
- **`datpaq user-avatar upload-icon`** - Upload a custom icon for avatar generation

### validate-ip

IP address validation and classification (IPv4, IPv6, CIDR)

- **`datpaq validate-ip get`** - Validates IPv4, IPv6, or CIDR notation and returns classification flags: private, reserved, loopback, multicast, link-local, bogon, and public. Supports comma-separated batch input (max 100).
- **`datpaq validate-ip post`** - Validate and classify IP addresses (POST)

### vin-lookup

Vehicle identification number decoding via NHTSA

- **`datpaq vin-lookup vin-decode`** - Decodes 17-character Vehicle Identification Numbers using the official NHTSA vPIC database. Validates ISO 3779 check digit. Supports single VIN or batch array (max 50).
- **`datpaq vin-lookup vin-makes`** - List all vehicle makes in the NHTSA database
- **`datpaq vin-lookup vin-sample-get`** - Decode a single VIN via GET
- **`datpaq vin-lookup vin-suggest`** - Get VIN prefix suggestions for autocomplete

### web-screenshot

Web page screenshot capture via headless browser

- **`datpaq web-screenshot`** - Renders a web page via Puppeteer headless browser and returns a screenshot. Supports viewport simulation, format selection, quality control, and crop.

### whois

Domain WHOIS registration data lookup

- **`datpaq whois`** - Returns registrar, registration date, expiration date, name servers, and raw WHOIS data. Supports single domain or batch array (max 50).

### working-days

Business day calculations with regional holiday support

- **`datpaq working-days batch`** - Batch working day calculations
- **`datpaq working-days calculate-get`** - Calculate working days (GET)
- **`datpaq working-days calculate-post`** - For range calculations: provide start_date and end_date to get the number of working days between them. For offset calculations: provide start_date and days_offset to find the resulting business date. Supports US, DE, and IN regional holiday calendars.


## Output Formats

```bash
# Human-readable table (default in terminal, JSON when piped)
datpaq domain-lookup get --domain example-value

# JSON for scripting and agents
datpaq domain-lookup get --domain example-value --json

# Filter to specific fields
datpaq domain-lookup get --domain example-value --json --select id,name,status

# Dry run — show the request without sending
datpaq domain-lookup get --domain example-value --dry-run

# Agent mode — JSON + compact + no prompts in one flag
datpaq domain-lookup get --domain example-value --agent
```

## Agent Usage

This CLI is designed for AI agent consumption:

- **Non-interactive** - never prompts, every input is a flag
- **Pipeable** - `--json` output to stdout, errors to stderr
- **Filterable** - `--select id,name` returns only fields you need
- **Previewable** - `--dry-run` shows the request without sending
- **Explicit retries** - add `--idempotent` to create retries when a no-op success is acceptable
- **Confirmable** - `--yes` for explicit confirmation of destructive actions
- **Piped input** - write commands can accept structured input when their help lists `--stdin`
- **Offline-friendly** - sync/search commands can use the local SQLite store when available
- **Agent-safe by default** - no colors or formatting unless `--human-friendly` is set

Exit codes: `0` success, `2` usage error, `3` not found, `4` auth error, `5` API error, `7` rate limited, `10` config error.

## Health Check

```bash
datpaq doctor
```

Verifies configuration, credentials, and connectivity to the API.

## Configuration

Config file: `~/.config/datpaq-proapi-pp-cli/config.toml`

Static request headers can be configured under `headers`; per-command header overrides take precedence.

Environment variables:

| Name | Kind | Required | Description |
| --- | --- | --- | --- |
| `DATPAQ_API_KEY_HEADER` | per_call | Yes | Set to your API credential. |

## Troubleshooting
**Authentication errors (exit code 4)**
- Run `datpaq doctor` to check credentials
- Verify the environment variable is set: `echo $DATPAQ_API_KEY_HEADER`
**Not found errors (exit code 3)**
- Check the resource ID is correct
- Run the `list` command to see available items

---

Generated by [CLI Printing Press](https://github.com/mvanhorn/cli-printing-press)
