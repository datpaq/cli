// Copyright 2026 datpaq. Licensed under Apache-2.0. See LICENSE.
// PATCH: secret staging/production environment toggle. Intentionally absent
// from --help, agent-context, MCP, README, SKILL.md, and all user-facing docs.
// Do not advertise or document the "/gummie" invocation.

package cli

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/datpaq/cli/internal/config"
)

// API base URLs for the secret environment toggle.
// Staging matches OpenAPI servers; production uses www.datpaq.com as the
// canonical host (legacy datpaq.com without www is still detected as production).
const (
	gummieProductionBaseURL = "https://www.datpaq.com/api/v1"
	gummieStagingBaseURL    = "https://staging.datpaq.com/api/v1"
	// Legacy default from config.Load / OpenAPI production server.
	gummieLegacyProductionBaseURL = "https://datpaq.com/api/v1"
)

// tryGummie handles the undocumented "/gummie" argv token before Cobra runs.
// Returns handled=true when the token was present as the first positional arg
// so Execute can skip the normal command tree (keeping this feature invisible
// to help, completions, agent-context, and MCP registration).
func tryGummie(args []string, out io.Writer) (handled bool, err error) {
	ok, configPath := parseGummieArgs(args)
	if !ok {
		return false, nil
	}
	if out == nil {
		out = os.Stdout
	}

	// Env override wins over file for every other command; refusing here avoids
	// printing a success message while the effective URL is unchanged.
	if os.Getenv("DATPAQ_BASE_URL") != "" {
		return true, fmt.Errorf("DATPAQ_BASE_URL is set; unset it before toggling environment")
	}

	// Load disk-only so we never persist env-merged credentials (DATPAQ_API_KEY).
	cfg, err := config.LoadDisk(configPath)
	if err != nil {
		return true, configErr(err)
	}

	nextURL, label, err := gummieToggleTarget(cfg.BaseURL)
	if err != nil {
		return true, err
	}
	if err := cfg.SaveBaseURL(nextURL); err != nil {
		return true, configErr(err)
	}
	// Include the URL so engineers/agents can confirm the live API target.
	fmt.Fprintf(out, "Environment is set to: %s (%s)\n", label, nextURL)
	return true, nil
}

// parseGummieArgs reports whether argv's first positional argument is "/gummie".
// Supports an optional --config / --config= path ahead of the token.
func parseGummieArgs(args []string) (ok bool, configPath string) {
	i := 1 // skip argv[0]
	for i < len(args) {
		a := args[i]
		switch {
		case a == "--config":
			if i+1 >= len(args) {
				return false, ""
			}
			configPath = args[i+1]
			i += 2
		case strings.HasPrefix(a, "--config="):
			configPath = strings.TrimPrefix(a, "--config=")
			i++
		case a == "/gummie":
			return true, configPath
		case strings.HasPrefix(a, "-"):
			// Skip unknown global flags; if the next token looks like a value
			// (not a flag and not /gummie), consume it too.
			i++
			if i < len(args) && !strings.HasPrefix(args[i], "-") && args[i] != "/gummie" {
				i++
			}
		default:
			return false, ""
		}
	}
	return false, ""
}

// gummieToggleTarget returns the base URL and display label to switch to.
// Default / production → staging; staging → production. Non-canonical URLs
// are refused so custom endpoints (e.g. localhost) are not destroyed.
func gummieToggleTarget(currentBaseURL string) (nextURL, label string, err error) {
	cur := normalizeGummieBaseURL(currentBaseURL)
	if isGummieStagingURL(cur) {
		return gummieProductionBaseURL, "PRODUCTION", nil
	}
	if isGummieProductionURL(cur) {
		return gummieStagingBaseURL, "STAGING", nil
	}
	return "", "", fmt.Errorf("current base_url %q is neither production nor staging; refusing to overwrite", currentBaseURL)
}

func normalizeGummieBaseURL(baseURL string) string {
	u := strings.TrimRight(strings.TrimSpace(baseURL), "/")
	if u == "" {
		return gummieProductionBaseURL
	}
	return u
}

func isGummieStagingURL(baseURL string) bool {
	u := strings.ToLower(normalizeGummieBaseURL(baseURL))
	return strings.Contains(u, "://staging.datpaq.com/") ||
		strings.HasSuffix(u, "://staging.datpaq.com") ||
		strings.Contains(u, "://staging.datpaq.com?")
}

func isGummieProductionURL(baseURL string) bool {
	u := strings.ToLower(normalizeGummieBaseURL(baseURL))
	switch u {
	case strings.ToLower(gummieProductionBaseURL),
		strings.ToLower(gummieLegacyProductionBaseURL),
		"https://www.datpaq.com",
		"https://datpaq.com":
		return true
	default:
		return false
	}
}
