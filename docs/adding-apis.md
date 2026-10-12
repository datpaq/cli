# Adding ProApi services to the CLI

New services are declared once in a table. The same definition builds the
CLI command, its typed MCP tool, `datpaq which` entries, `datpaq sample`
snippets, and (via a test helper) the `spec.yaml` paths. `spec.yaml` is what
staging.datpaq.com and datpaq.com serve as `/openapi.yaml`, so this is also how
the public OpenAPI gets updated.

## Source of truth

- **Which APIs are live:** `https://staging.datpaq.com/api-list` (the website catalog).
  Production follows staging at release.
- **Routes and params:** the ProApi service code, `ProApi/v1/<slug>/index.js`
  (Joi schemas and `app.get/post` registrations). The Markdown docs in
  `ProApi/docs/api-docs/main-docs/<slug>.md` are useful but drift from the
  code; when they disagree, follow the code.

## Steps

1. **Write `internal/cli/<slug>.go`.** Copy a similar existing file
   (e.g. `libor.go` for GET query APIs, `iban.go` for path and body params,
   `ocr.go` for file uploads):

   ```go
   func init() {
   	registerDeclaredAPI(APIGroup{
   		Name:  "<slug>",               // must equal the ProApi/website slug
   		Short: "One-line description",
   		Endpoints: []APIEndpoint{
   			{Name: "lookup", Short: "...", Method: "GET", Path: "/<slug>/{id}",
   				Params: []APIParam{
   					{Flag: "id", In: "path", Usage: "...", Example: "abc"},
   					{Flag: "start-date", Wire: "startDate", Usage: "...", Example: "2026-01-01"},
   					{Flag: "limit", Type: "int", Usage: "...", Example: "10", Required: true},
   				}},
   		},
   	})
   }
   ```

   - `Flag` is kebab-case; set `Wire` when the API name differs (camelCase or snake_case).
   - `In`: `query` (default), `body`, or `path` (path params become positionals and `{Wire}` placeholders).
   - `Type`: `string` (default), `int`, `float`, `bool`, `array` (JSON or comma-separated),
     `object` (JSON), `file` (local path sent as base64), `files` (comma-separated paths).
   - `Example` is required: it drives `--help`, the request tests, and the spec.
   - Mark `Required` only for unconditionally required params (Joi `.required()`).
   - When one route has both GET and POST, give POST the plain name and GET a `-get` suffix.
   - Skip routes that are aliases, human-facing, or hard-coded 503 stubs, and say so in a file comment.

2. **Run the finisher** from the repo root:

   ```bash
   scripts/api-sync/add-api.sh <patch-id> "<summary>" "<reason>" <slug> [<slug>...]
   ```

   It adds the slugs to `active-apis.json` and `active_apis_test.go`, emits and splices
   the `spec.yaml` paths, appends a `.printing-press-patches.json` entry, and runs
   `go vet` and `go test ./...`.

3. **Review and commit.** Keep each PR under ~400 changed lines excluding `spec.yaml`
   (about 5–8 APIs). Group related APIs per PR.

4. **Mirror the active list** to the MCP repo if needed: `scripts/sync-active-apis.sh`.

5. **Refresh SKILL.md's command reference** from the live command tree
   (`datpaq api --agent` and `datpaq api <slug> --agent`).

## What the tests enforce

`go test ./internal/cli ./internal/mcp` fails if any declared endpoint:

- builds the wrong method, path, or wire parameter names (`TestDeclaredEndpointsBuildRequests`);
- does not enforce its required flags (`TestDeclaredRequiredFlagsAreEnforced`);
- is missing from `active-apis.json` (`TestDeclaredAPIGroupsAreActive`);
- is missing an operation or parameter in `spec.yaml` (`TestDeclaredEndpointsDocumentedInSpec`);
- lacks a typed MCP tool with every parameter (`TestDeclaredAPIToolsRegisteredWithParams`).

## Deactivating an API

Remove the slug from `active-apis.json`, move it to the inactive list in
`active_apis_test.go`, and delete its paths and tag from `spec.yaml`. Keep the
command registered so existing scripts keep working.

## Live smoke tests

Staging rejects production credentials. Log in against staging before
smoke-testing; `DATPAQ_BASE_URL=https://staging.datpaq.com/api/v1` overrides
the base URL for a single command.
