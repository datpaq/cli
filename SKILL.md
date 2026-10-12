---
name: pp-datpaq
description: "Printing Press CLI for Datpaq. Comprehensive multi-service API platform providing aviation, geolocation, financial, utilities, image processing,..."
author: "datpaq"
license: "Apache-2.0"
argument-hint: "<command> [args] | install cli|mcp"
allowed-tools: "Read Bash"
metadata:
  openclaw:
    requires:
      bins:
        - datpaq
---

# Datpaq — Printing Press CLI

## Prerequisites: Install the CLI

This skill drives the `datpaq` binary. **You must verify the CLI is installed before invoking any command from this skill.** If it is missing, install it first:

1. Install via the Printing Press installer:
   ```bash
   npx -y @mvanhorn/printing-press install datpaq --cli-only
   ```
2. Verify: `datpaq --version`
3. Ensure `$GOPATH/bin` (or `$HOME/go/bin`) is on `$PATH`.

If the `npx` install fails before this CLI has a public-library category, install Node or use the category-specific Go fallback after publish.

If `--version` reports "command not found" after install, the install step did not put the binary on `$PATH`. Do not proceed with skill commands until verification succeeds.

Comprehensive multi-service API platform providing aviation, geolocation, financial, utilities, image processing, and data enrichment capabilities. All endpoints require authentication via API key passed as an x-api-key header or api_key query parameter.

## Command Reference

63 active APIs. Run `datpaq api` for the live list and `datpaq <api> <command> --help` for flags.

**aircraft** — FAA aircraft registry lookups by tail number, ICAO hex, or type code

- `datpaq aircraft batch-lookup` — Batch aircraft lookups by tail number
- `datpaq aircraft lookup-by-icao` — Look up aircraft by ICAO hex code
- `datpaq aircraft lookup-by-tail` — Returns aircraft ownership, manufacturer, model, type, and year from the FAA registry using an N-number (tail number).
- `datpaq aircraft lookup-by-type` — Look up aircraft type specifications

**breach-check** — Email breach exposure, password k-anonymity checks, and breach catalog

- `datpaq breach-check breaches` — List known breaches
- `datpaq breach-check check` — Check one email for known breaches
- `datpaq breach-check check-batch` — Check up to 25 emails for known breaches
- `datpaq breach-check password` — Check a SHA-1 hash prefix against breached passwords

**calendar** — Calendar month, day, range, and holiday lookups

- `datpaq calendar day` — Get calendar metadata for a single date
- `datpaq calendar holidays` — List holidays for a year (optional month filter)
- `datpaq calendar month` — Get a calendar month grid
- `datpaq calendar range` — Get calendar days across a date range

**color-convert** — Color conversion across formats with contrast and gradient helpers

- `datpaq color-convert convert` — Convert a color and optionally check contrast

**convert-time** — Converts a datetime value from a source timezone to a target timezone. Accepts ISO 8601 strings, Unix epoch...

- `datpaq convert-time` — Converts a datetime value from a source timezone to a target timezone. Accepts ISO 8601 strings, Unix epoch...

**country-codes** — ISO country code lookup and metadata

- `datpaq country-codes country-by-iso` — Get a country by ISO 2 or ISO 3 code
- `datpaq country-codes country-export` — Export country data in bulk
- `datpaq country-codes country-list` — List all countries with ISO codes and metadata
- `datpaq country-codes country-search` — Search countries by full or partial name

**current-time** — Returns the current date/time with UTC offset and DST status for a given timezone.

- `datpaq current-time` — Returns the current date/time with UTC offset and DST status for a given timezone.

**data-format-convert** — Convert data between JSON, CSV, XML, YAML, and spreadsheet formats

- `datpaq data-format-convert convert` — Convert data from one format to another

**define** — Returns definitions, part of speech, pronunciation, and etymology for the given word.

- `datpaq define` — Returns definitions, part of speech, pronunciation, and etymology for the given word.

**deslop** — Detect and rewrite AI writing patterns in text

- `datpaq deslop analyze` — Analyze text for AI writing patterns
- `datpaq deslop analyze-get` — Analyze short text via GET
- `datpaq deslop rewrite` — Rewrite text to remove AI writing patterns
- `datpaq deslop rules` — List detection rules

**diffy** — Web page snapshots and change diffs against a previous snapshot

- `datpaq diffy diff` — Diff a page against a previous snapshot state
- `datpaq diffy snapshot` — Capture a page snapshot and its diffable state

**dns** — Supports record types A, AAAA, MX, CNAME, TXT, NS, SOA, and ALL.

- `datpaq dns` — Supports record types A, AAAA, MX, CNAME, TXT, NS, SOA, and ALL.

**document-convert** — Convert office documents, PDF, and text formats

- `datpaq document-convert convert` — Convert one document from a local file or URL
- `datpaq document-convert formats` — List supported conversion pairs

**document-generation** — Generate business documents (invoices, letters, forms) as PDF, DOCX, or HTML

- `datpaq document-generation batch` — Generate up to 5 documents as base64
- `datpaq document-generation documents` — List available document types
- `datpaq document-generation example` — Get example data for a document type
- `datpaq document-generation generate` — Generate a PDF or DOCX document
- `datpaq document-generation json-schema` — Get a document type's JSON Schema
- `datpaq document-generation openapi` — Get the service OpenAPI document
- `datpaq document-generation render` — Render a document with data as HTML
- `datpaq document-generation render-blank` — Render a blank document as HTML
- `datpaq document-generation schema` — Get a document type's field schema

**domain-lookup** — Domain availability and registration information

- `datpaq domain-lookup get` — Check domain availability and registration info
- `datpaq domain-lookup post` — Check domain availability (POST)

**economic-indicators** — Macroeconomic indicator catalog, series history, and interest rates

- `datpaq economic-indicators indicators` — List available indicators
- `datpaq economic-indicators interest-rates` — Get key interest rates
- `datpaq economic-indicators latest` — Get the latest value for an indicator
- `datpaq economic-indicators series` — Get history for an indicator

**encoding-convert** — Encode, decode, hash, JWT-decode, normalize, and data-URI conversions

- `datpaq encoding-convert catalog` — List supported operations, encodings, and algorithms
- `datpaq encoding-convert convert` — Run an encoding operation

**ev-charger** — Electric vehicle charging station discovery and status

- `datpaq ev-charger ev-networks` — List all supported charging networks
- `datpaq ev-charger ev-station-by-id` — Get specific charging station by ID
- `datpaq ev-charger ev-station-status` — Get real-time station status
- `datpaq ev-charger ev-stations-by-address` — Search EV charging stations by address
- `datpaq ev-charger ev-stations-by-coords` — Find EV charging stations near coordinates

**gdp** — US GDP levels, history, growth, and dataset comparison

- `datpaq gdp catalog` — List datasets, sources, and periods
- `datpaq gdp compare` — Compare GDP datasets
- `datpaq gdp growth` — Get GDP growth rates
- `datpaq gdp history` — Get GDP history
- `datpaq gdp latest` — Get the latest GDP observation

**geocoding** — Forward, reverse, and batch geocoding

- `datpaq geocoding batch` — Batch forward and/or reverse geocoding
- `datpaq geocoding forward` — Geocode an address or place name to coordinates
- `datpaq geocoding reverse` — Reverse-geocode coordinates to an address

**helicopter** — Helicopter registry lookups by registration, ICAO, or serial number

- `datpaq helicopter lookup-by-icao` — Look up helicopter by ICAO hex code
- `datpaq helicopter lookup-by-registration` — Look up helicopter by registration number
- `datpaq helicopter lookup-by-serial` — Look up helicopter by manufacturer serial number
- `datpaq helicopter manufacturers` — List helicopter manufacturers

**iban** — IBAN structure and mod-97 checksum validation with country metadata

- `datpaq iban countries` — List supported IBAN countries
- `datpaq iban lookup` — Validate one IBAN from the URL path
- `datpaq iban validate` — Validate one IBAN
- `datpaq iban validate-batch` — Validate up to 100 IBANs

**image-processing** — Image resize, compress, convert, crop, and pipeline operations

- `datpaq image-processing image-compress` — Compress an image
- `datpaq image-processing image-convert` — Convert image to a different format
- `datpaq image-processing image-crop` — Crop an image to a region
- `datpaq image-processing image-metadata` — Extract metadata from an image
- `datpaq image-processing image-pipeline` — Executes an ordered array of image operations in a single request. Recommended over individual endpoints when...
- `datpaq image-processing image-resize` — Resize an image to specified dimensions
- `datpaq image-processing info` — Returns supported image operations, formats, and service metadata.

**income-tax** — US federal income tax estimates and brackets

- `datpaq income-tax brackets` — Get federal tax brackets
- `datpaq income-tax catalog` — List supported years, filing statuses, and sources
- `datpaq income-tax estimate` — Estimate federal and state income tax

**interest-rate** — Current benchmark interest rates and the Treasury yield curve

- `datpaq interest-rate current` — Get current benchmark rates by category
- `datpaq interest-rate yield-curve` — Get the ordered Treasury yield curve

**ip-geolocation** — Geographic location data for IPv4 and IPv6 addresses

- `datpaq ip-geolocation batch` — Batch IP geolocation lookup
- `datpaq ip-geolocation get` — Returns country, region, city, coordinates, ISP, ASN, and timezone for any IPv4 or IPv6 address.
- `datpaq ip-geolocation post` — Get geographic location for an IP address (POST)

**ip-intelligence** — Threat intelligence and risk scoring for IP addresses

- `datpaq ip-intelligence batch` — Batch IP threat intelligence lookup
- `datpaq ip-intelligence batch-info` — Returns batch processing limits, example request shape, and response schema details.
- `datpaq ip-intelligence get` — Returns trust score, risk classification, VPN/proxy/Tor detection, datacenter detection, and threat categories. Omit...

**libor** — SOFR-first benchmark rates with archival LIBOR history and tenor spreads

- `datpaq libor catalog` — List supported benchmarks, tenors, and coverage
- `datpaq libor history` — Get benchmark rate history
- `datpaq libor latest` — Get the latest benchmark rate
- `datpaq libor spread` — Get the spread between two tenors

**logo** — Discover a company's logos, icons, and brand palette by domain

- `datpaq logo discover` — Discover logos and brand colors for a domain

**lorem-ipsum** — Placeholder text generation as JSON, text, or HTML

- `datpaq lorem-ipsum generate` — Generate lorem ipsum words, sentences, or paragraphs

**mac-address** — Lookup vendor by full MAC address, 6-character OUI prefix, or organization name. Provide exactly one of the three...

- `datpaq mac-address` — Lookup vendor by full MAC address, 6-character OUI prefix, or organization name. Provide exactly one of the three...

**markdown-convert** — Markdown to HTML, HTML to Markdown, plaintext, and document outlines

- `datpaq markdown-convert batch` — Process up to 25 documents in one request
- `datpaq markdown-convert capabilities` — List supported targets, limits, and Markdown flavor
- `datpaq markdown-convert outline` — Get document structure and statistics
- `datpaq markdown-convert to-html` — Render Markdown to sanitized HTML
- `datpaq markdown-convert to-markdown` — Convert HTML to Markdown
- `datpaq markdown-convert to-text` — Render Markdown or HTML to plaintext

**market-cap** — Company market capitalization lookups, history, trends, and rankings

- `datpaq market-cap catalog` — List supported symbols, sectors, and periods
- `datpaq market-cap compare` — Compare market caps across symbols
- `datpaq market-cap history` — Get market cap history for a symbol
- `datpaq market-cap lookup` — Get market cap for a symbol
- `datpaq market-cap top` — List the largest companies by market cap
- `datpaq market-cap trend` — Get market cap trend statistics for a symbol

**mortgage-calculator** — Mortgage payment, refinance break-even, and current mortgage rates

- `datpaq mortgage-calculator current-rates` — Get weekly average mortgage rates from FRED
- `datpaq mortgage-calculator payment` — Calculate monthly payment, total interest, and optional amortization
- `datpaq mortgage-calculator refinance` — Simulate a refinance and its break-even point

**mx-lookup** — Mail exchange record lookup with optional SMTP and SPF validation

- `datpaq mx-lookup batch` — Batch MX record lookup for multiple domains
- `datpaq mx-lookup get` — Returns mail exchange records with optional IP resolution, SMTP testing, and SPF extraction.
- `datpaq mx-lookup post` — Look up MX records for a domain (POST)

**ocr** — Extract text from PDFs and images, with optional OCR

- `datpaq ocr parse` — Extract text from a local file or URL

**pdf-generation** — Generate PDFs from a URL or HTML

- `datpaq pdf-generation from-html` — Render HTML content to PDF
- `datpaq pdf-generation from-url` — Render a public URL to PDF

**pdf-toolkit** — Merge, split, extract, rotate, watermark, compress, encrypt, and inspect PDFs

- `datpaq pdf-toolkit catalog` — List supported operations and limits
- `datpaq pdf-toolkit run` — Run a PDF operation on local files

**ping-human** — Escalate a decision to a human via Slack or email and collect the reply

- `datpaq ping-human cancel` — Cancel an open escalation
- `datpaq ping-human create` — Create an escalation and notify humans
- `datpaq ping-human get` — Get an escalation's status and reply

**profanity** — Content moderation and profanity detection

- `datpaq profanity analyze` — Comprehensive profanity analysis with severity and word positions
- `datpaq profanity detect` — Detect whether text contains profanity
- `datpaq profanity filter` — Replaces profane words with asterisks. Returns cleaned text.

**public-holidays** — Global public holiday data by country and year

- `datpaq public-holidays countries` — List all supported countries for holiday data
- `datpaq public-holidays` — Accepts ISO-2, ISO-3, or full country names. Returns holidays with ISO, US, and EU date formats. Filter by year...

**qr-code** — Generate and decode QR codes

- `datpaq qr-code decode` — Decode QR code(s) from a base64 image
- `datpaq qr-code generate` — Generate a QR code image from text

**sample-data** — Mock data generation for testing and development (Faker.js)

- `datpaq sample-data batch-get` — Generates mock data using Faker.js. Supports 20+ schema types including user, person, address, company, product, and...
- `datpaq sample-data batch-post` — Generate multiple types of sample data (POST)
- `datpaq sample-data generate-get` — GET-compatible sample data generation endpoint for scripts and browser testing.
- `datpaq generate` — Generate realistic mock data using a JSON request body. Recommended for full control over type, count, fields, and...
- `datpaq generate-batch` — Submit multiple sample-data generation requests in one call.
- `datpaq schemas` — Returns all available sample-data schema names with default field lists.

**sec-filings** — SEC EDGAR company profiles, filings, and XBRL facts

- `datpaq sec-filings catalog` — List supported forms, taxonomies, and limits
- `datpaq sec-filings company` — Get a company profile by CIK or ticker
- `datpaq sec-filings facts` — Get XBRL concept facts for a company
- `datpaq sec-filings filings` — List recent filings for a company
- `datpaq sec-filings forms` — List a company's filings of one form type

**spell-check** — Multi-language spell checking with correction suggestions

- `datpaq spell-check batch` — Spell check multiple texts
- `datpaq spell-check languages` — List supported languages for spell checking
- `datpaq spell-check` — Detects spelling errors with ranked correction suggestions. Supports English, Spanish, French, German, and Italian....

**structured-extraction** — Extract structured fields from a web page or HTML using CSS selectors

- `datpaq structured-extraction extract` — Extract fields from a URL or HTML

**tech-stack-detection** — Detect the technologies a website uses

- `datpaq tech-stack-detection detect` — Detect technologies for a domain

**text-language** — Text language detection supporting 176+ languages

- `datpaq text-language detect` — Identifies the language of text with a confidence score. Supports 176+ languages with ISO 639-3 output. Requires at...
- `datpaq text-language detect-batch` — Batch language detection for multiple texts
- `datpaq text-language supported` — Returns all 176+ supported languages with ISO 639-1 and ISO 639-3 codes.

**thesaurus** — Returns synonyms, antonyms, or both for a given word. Filter results by part of speech and limit the number of results.

- `datpaq thesaurus` — Returns synonyms, antonyms, or both for a given word. Filter results by part of speech and limit the number of results.

**translate** — Text translation between supported languages

- `datpaq translate languages` — List supported languages
- `datpaq translate text` — Translate one text or a batch of texts

**unemployment** — US unemployment rates, labor force, claims, demographics, and states

- `datpaq unemployment claims` — Get initial and continued jobless claims
- `datpaq unemployment demographic` — Get unemployment by demographic group
- `datpaq unemployment labor-force` — Get participation and employment-population ratios
- `datpaq unemployment long-term` — Get long-term unemployment
- `datpaq unemployment measures` — Get U-1 through U-6 measures
- `datpaq unemployment national` — Get the national unemployment rate
- `datpaq unemployment series` — List the unemployment series catalog
- `datpaq unemployment state` — Get unemployment for one state or DC
- `datpaq unemployment states` — Get latest unemployment for all states and DC

**unit-conversion** — Unit conversion across 100+ measurement types

- `datpaq unit-conversion unit-categories` — List available measurement categories
- `datpaq unit-conversion unit-convert-get` — Supports 100+ unit types across length, weight, temperature, volume, area, speed, digital storage, and time....
- `datpaq unit-conversion unit-convert-post` — Convert a value between units (POST)
- `datpaq unit-conversion units-by-category` — List units available in a category

**us-states** — U.S. state metadata by abbreviation or FIPS code

- `datpaq us-states by-abbr` — Get U.S. state by 2-letter abbreviation
- `datpaq us-states by-fips` — Get U.S. state by FIPS code
- `datpaq us-states list` — List or search U.S. states

**user-avatar** — Dynamic user avatar generation in PNG, WebP, or SVG

- `datpaq user-avatar generate` — Creates avatar images using Canvas (PNG/WebP) or native SVG. Supports initials-based avatars, custom colors,...
- `datpaq user-avatar upload-icon` — Upload a custom icon for avatar generation

**validate-ip** — IP address validation and classification (IPv4, IPv6, CIDR)

- `datpaq validate-ip get` — Validates IPv4, IPv6, or CIDR notation and returns classification flags: private, reserved, loopback, multicast,...
- `datpaq validate-ip post` — Validate and classify IP addresses (POST)

**vat-validation-and-rates** — VAT number validation, VAT rates by country, and VAT price calculation

- `datpaq vat-validation-and-rates calculate` — Calculate VAT for an amount
- `datpaq vat-validation-and-rates calculate-batch` — Calculate VAT for multiple amounts
- `datpaq vat-validation-and-rates calculate-batch-get` — Calculate VAT for multiple amounts via GET
- `datpaq vat-validation-and-rates catalog` — List supported countries, categories, and limits
- `datpaq vat-validation-and-rates rate-countries` — List countries with VAT rate data
- `datpaq vat-validation-and-rates rates` — Get VAT rates for a country
- `datpaq vat-validation-and-rates validate` — Validate a VAT number
- `datpaq vat-validation-and-rates validate-batch` — Validate up to 50 VAT numbers
- `datpaq vat-validation-and-rates validate-batch-get` — Validate up to 50 comma-separated VAT numbers via GET
- `datpaq vat-validation-and-rates validate-number` — Validate a VAT number from the URL path

**vin-lookup** — Vehicle identification number decoding via NHTSA

- `datpaq vin-lookup vin-decode` — Decodes 17-character Vehicle Identification Numbers using the official NHTSA vPIC database. Validates ISO 3779 check...
- `datpaq vin-lookup vin-makes` — List all vehicle makes in the NHTSA database
- `datpaq vin-lookup vin-sample-get` — Decode a single VIN via GET
- `datpaq vin-lookup vin-suggest` — Get VIN prefix suggestions for autocomplete

**weather** — Current conditions and forecasts by latitude/longitude

- `datpaq weather current` — Get current weather conditions for a location
- `datpaq weather forecast` — Get weather forecast for a location

**web-scraping** — SSRF-safe public web page content extraction

- `datpaq web-scraping` — SSRF-safe public web page content extraction

**web-screenshot** — Renders a web page via Puppeteer headless browser and returns a screenshot. Supports viewport simulation, format...

- `datpaq web-screenshot` — Renders a web page via Puppeteer headless browser and returns a screenshot. Supports viewport simulation, format...

**web-search** — Web search for a single query or a batch of queries

- `datpaq web-search batch` — Batch web search for up to 10 queries
- `datpaq web-search search` — Search the web for a single query

**whois** — Returns registrar, registration date, expiration date, name servers, and raw WHOIS data. Supports single domain or...

- `datpaq whois` — Returns registrar, registration date, expiration date, name servers, and raw WHOIS data. Supports single domain or...

**working-days** — Business day calculations with regional holiday support

- `datpaq working-days batch` — Batch working day calculations
- `datpaq working-days calculate-get` — Calculate working days (GET)
- `datpaq working-days calculate-post` — For range calculations: provide start_date and end_date to get the number of working days between them. For offset...


### Finding the right command

When you know what you want to do but not which command does it, ask the CLI directly:

```bash
datpaq which "<capability in your own words>"
```

`which` resolves a natural-language capability query to the best matching command from this CLI's curated feature index. Exit code `0` means at least one match; exit code `2` means no confident match — fall back to `--help` or use a narrower query.

## Auth Setup
Run `datpaq auth setup` to print the URL and steps for getting a key (add `--launch` to open the URL). Then set:

```bash
export DATPAQ_API_KEY_HEADER="<your-key>"
```

Or persist it in `~/.config/datpaq-proapi-pp-cli/config.toml`.

Run `datpaq doctor` to verify setup.

## Agent Mode

Add `--agent` to any command. Expands to: `--json --compact --no-input --no-color --yes`.

- **Pipeable** — JSON on stdout, errors on stderr
- **Filterable** — `--select` keeps a subset of fields. Dotted paths descend into nested structures; arrays traverse element-wise. Critical for keeping context small on verbose APIs:

  ```bash
  datpaq domain-lookup get --domain example-value --agent --select id,name,status
  ```
- **Previewable** — `--dry-run` shows the request without sending
- **Offline-friendly** — sync/search commands can use the local SQLite store when available
- **Non-interactive** — never prompts, every input is a flag
- **Explicit retries** — use `--idempotent` only when an already-existing create should count as success

### Response envelope

Commands that read from the local store or the API wrap output in a provenance envelope:

```json
{
  "meta": {"source": "live" | "local", "synced_at": "...", "reason": "..."},
  "results": <data>
}
```

Parse `.results` for data and `.meta.source` to know whether it's live or local. A human-readable `N results (live)` summary is printed to stderr only when stdout is a terminal AND no machine-format flag (`--json`, `--csv`, `--compact`, `--quiet`, `--plain`, `--select`) is set — piped/agent consumers and explicit-format runs get pure JSON on stdout.

## Agent Feedback

When you (or the agent) notice something off about this CLI, record it:

```
datpaq feedback "the --since flag is inclusive but docs say exclusive"
datpaq feedback --stdin < notes.txt
datpaq feedback list --json --limit 10
```

Entries are stored locally at `~/.datpaq/feedback.jsonl`. They are never POSTed unless `DATPAQ_FEEDBACK_ENDPOINT` is set AND either `--send` is passed or `DATPAQ_FEEDBACK_AUTO_SEND=true`. Default behavior is local-only.

Write what *surprised* you, not a bug report. Short, specific, one line: that is the part that compounds.

## Output Delivery

Every command accepts `--deliver <sink>`. The output goes to the named sink in addition to (or instead of) stdout, so agents can route command results without hand-piping. Three sinks are supported:

| Sink | Effect |
|------|--------|
| `stdout` | Default; write to stdout only |
| `file:<path>` | Atomically write output to `<path>` (tmp + rename) |
| `webhook:<url>` | POST the output body to the URL (`application/json` or `application/x-ndjson` when `--compact`) |

Unknown schemes are refused with a structured error naming the supported set. Webhook failures return non-zero and log the URL + HTTP status on stderr.

## Named Profiles

A profile is a saved set of flag values, reused across invocations. Use it when a scheduled agent calls the same command every run with the same configuration - HeyGen's "Beacon" pattern.

```
datpaq profile save briefing --json
datpaq --profile briefing domain-lookup get --domain example-value
datpaq profile list --json
datpaq profile show briefing
datpaq profile delete briefing --yes
```

Explicit flags always win over profile values; profile values win over defaults. `agent-context` lists all available profiles under `available_profiles` so introspecting agents discover them at runtime.

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 2 | Usage error (wrong arguments) |
| 3 | Resource not found |
| 4 | Authentication required |
| 5 | API error (upstream issue) |
| 7 | Rate limited (wait and retry) |
| 10 | Config error |

## Argument Parsing

Parse `$ARGUMENTS`:

1. **Empty, `help`, or `--help`** → show `datpaq --help` output
2. **Starts with `install`** → ends with `mcp` → MCP installation; otherwise → see Prerequisites above
3. **Anything else** → Direct Use (execute as CLI command with `--agent`)

## MCP Server Installation

Install the MCP binary from this CLI's published public-library entry or pre-built release, then register it:

```bash
claude mcp add datpaq-mcp -- datpaq-mcp
```

Verify: `claude mcp list`

## Direct Use

1. Check if installed: `which datpaq`
   If not found, offer to install (see Prerequisites at the top of this skill).
2. Match the user query to the best command from the Unique Capabilities and Command Reference above.
3. Execute with the `--agent` flag:
   ```bash
   datpaq <command> [subcommand] [args] --agent
   ```
4. If ambiguous, drill into subcommand help: `datpaq <command> --help`.
