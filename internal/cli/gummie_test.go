// Copyright 2026 datpaq. Licensed under Apache-2.0. See LICENSE.
// PATCH: tests for the private "/gummie" env toggle. Do not surface in docs.

package cli

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/datpaq/cli/internal/config"
	"github.com/pelletier/go-toml/v2"
)

func TestParseGummieArgs(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		args       []string
		wantOK     bool
		wantConfig string
	}{
		{"plain", []string{"datpaq", "/gummie"}, true, ""},
		{"with config flag", []string{"datpaq", "--config", "/tmp/c.toml", "/gummie"}, true, "/tmp/c.toml"},
		{"with config equals", []string{"datpaq", "--config=/tmp/c.toml", "/gummie"}, true, "/tmp/c.toml"},
		{"other command", []string{"datpaq", "doctor"}, false, ""},
		{"help gummie is not toggle", []string{"datpaq", "help", "/gummie"}, false, ""},
		{"empty", []string{"datpaq"}, false, ""},
		{"json then gummie", []string{"datpaq", "--json", "/gummie"}, true, ""},
		{"timeout then gummie", []string{"datpaq", "--timeout", "30s", "/gummie"}, true, ""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ok, cfg := parseGummieArgs(tc.args)
			if ok != tc.wantOK || cfg != tc.wantConfig {
				t.Fatalf("parseGummieArgs(%v) = (%v, %q), want (%v, %q)", tc.args, ok, cfg, tc.wantOK, tc.wantConfig)
			}
		})
	}
}

func TestGummieToggleTarget(t *testing.T) {
	t.Parallel()
	url, label, err := gummieToggleTarget(gummieProductionBaseURL)
	if err != nil || url != gummieStagingBaseURL || label != "STAGING" {
		t.Fatalf("from production: got %s %s err=%v", url, label, err)
	}
	url, label, err = gummieToggleTarget(gummieLegacyProductionBaseURL)
	if err != nil || url != gummieStagingBaseURL || label != "STAGING" {
		t.Fatalf("from legacy production: got %s %s err=%v", url, label, err)
	}
	url, label, err = gummieToggleTarget("")
	if err != nil || url != gummieStagingBaseURL || label != "STAGING" {
		t.Fatalf("from empty/default: got %s %s err=%v", url, label, err)
	}
	url, label, err = gummieToggleTarget(gummieStagingBaseURL)
	if err != nil || url != gummieProductionBaseURL || label != "PRODUCTION" {
		t.Fatalf("from staging: got %s %s err=%v", url, label, err)
	}
	url, label, err = gummieToggleTarget("https://staging.datpaq.com")
	if err != nil || url != gummieProductionBaseURL || label != "PRODUCTION" {
		t.Fatalf("from staging host: got %s %s err=%v", url, label, err)
	}
	_, _, err = gummieToggleTarget("http://localhost:8080/api/v1")
	if err == nil {
		t.Fatal("custom base_url must be refused")
	}
}

func TestTryGummieDefaultsToStagingFirst(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.toml")
	t.Setenv("DATPAQ_BASE_URL", "")
	t.Setenv("DATPAQ_API_KEY", "")
	t.Setenv("DATPAQ_API_KEY_HEADER", "")

	var out bytes.Buffer
	handled, err := tryGummie([]string{"datpaq", "--config", cfgPath, "/gummie"}, &out)
	if err != nil {
		t.Fatal(err)
	}
	if !handled {
		t.Fatal("expected handled")
	}
	if got := strings.TrimSpace(out.String()); got != "Environment is set to: STAGING ("+gummieStagingBaseURL+")" {
		t.Fatalf("output = %q", got)
	}

	cfg, err := config.LoadDisk(cfgPath)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.BaseURL != gummieStagingBaseURL {
		t.Fatalf("base_url = %q, want staging", cfg.BaseURL)
	}

	out.Reset()
	handled, err = tryGummie([]string{"datpaq", "--config", cfgPath, "/gummie"}, &out)
	if err != nil {
		t.Fatal(err)
	}
	if !handled {
		t.Fatal("expected handled on second toggle")
	}
	if got := strings.TrimSpace(out.String()); got != "Environment is set to: PRODUCTION ("+gummieProductionBaseURL+")" {
		t.Fatalf("second output = %q", got)
	}

	cfg, err = config.LoadDisk(cfgPath)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.BaseURL != gummieProductionBaseURL {
		t.Fatalf("second toggle base_url = %q, want %q", cfg.BaseURL, gummieProductionBaseURL)
	}
}

func TestTryGummiePreservesFileAPIKeyWhenEnvKeySet(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.toml")
	if err := os.WriteFile(cfgPath, []byte("base_url = 'https://datpaq.com/api/v1'\napi_key_header = 'file-secret-key'\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("DATPAQ_BASE_URL", "")
	t.Setenv("DATPAQ_API_KEY", "env-secret-key")
	t.Setenv("DATPAQ_API_KEY_HEADER", "")

	var out bytes.Buffer
	handled, err := tryGummie([]string{"datpaq", "--config", cfgPath, "/gummie"}, &out)
	if err != nil {
		t.Fatal(err)
	}
	if !handled {
		t.Fatal("expected handled")
	}
	if got := strings.TrimSpace(out.String()); got != "Environment is set to: STAGING ("+gummieStagingBaseURL+")" {
		t.Fatalf("output = %q", got)
	}

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		t.Fatal(err)
	}
	var raw struct {
		BaseURL    string `toml:"base_url"`
		APIKeyHdr  string `toml:"api_key_header"`
	}
	if err := toml.Unmarshal(data, &raw); err != nil {
		t.Fatal(err)
	}
	if raw.BaseURL != gummieStagingBaseURL {
		t.Fatalf("base_url = %q, want staging", raw.BaseURL)
	}
	if raw.APIKeyHdr != "file-secret-key" {
		t.Fatalf("api_key_header = %q, want file-secret-key (must not persist env key)", raw.APIKeyHdr)
	}

	// Effective Load still prefers env key for runtime auth.
	loaded, err := config.Load(cfgPath)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.DatpaqApiKeyHeader != "env-secret-key" {
		t.Fatalf("runtime key = %q, want env-secret-key", loaded.DatpaqApiKeyHeader)
	}
}

func TestTryGummieRefusesWhenBaseURLEnvSet(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.toml")
	if err := os.WriteFile(cfgPath, []byte("base_url = 'https://datpaq.com/api/v1'\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("DATPAQ_BASE_URL", "https://datpaq.com/api/v1")
	t.Setenv("DATPAQ_API_KEY", "")

	handled, err := tryGummie([]string{"datpaq", "--config", cfgPath, "/gummie"}, io.Discard)
	if !handled {
		t.Fatal("expected handled")
	}
	if err == nil || !strings.Contains(err.Error(), "DATPAQ_BASE_URL") {
		t.Fatalf("want DATPAQ_BASE_URL error, got %v", err)
	}

	disk, err := config.LoadDisk(cfgPath)
	if err != nil {
		t.Fatal(err)
	}
	if disk.BaseURL != gummieLegacyProductionBaseURL {
		t.Fatalf("file must be unchanged, got %q", disk.BaseURL)
	}
}

func TestTryGummieRefusesCustomBaseURL(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.toml")
	if err := os.WriteFile(cfgPath, []byte("base_url = 'http://localhost:8080/api/v1'\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("DATPAQ_BASE_URL", "")
	t.Setenv("DATPAQ_API_KEY", "")

	handled, err := tryGummie([]string{"datpaq", "--config", cfgPath, "/gummie"}, io.Discard)
	if !handled {
		t.Fatal("expected handled")
	}
	if err == nil || !strings.Contains(err.Error(), "refusing to overwrite") {
		t.Fatalf("want refuse error, got %v", err)
	}
}

func TestTryGummieNotHandledForOtherCommands(t *testing.T) {
	handled, err := tryGummie([]string{"datpaq", "doctor"}, io.Discard)
	if err != nil {
		t.Fatal(err)
	}
	if handled {
		t.Fatal("doctor must not trigger gummie")
	}
}

func TestGummieAbsentFromHelpAndAgentContext(t *testing.T) {
	root := RootCmd()
	var helpBuf bytes.Buffer
	root.SetOut(&helpBuf)
	root.SetArgs([]string{"--help"})
	if err := root.Execute(); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(helpBuf.String(), "gummie") {
		t.Fatal("gummie must not appear in --help")
	}

	ctx := buildAgentContext(root)
	for _, cmd := range ctx.Commands {
		if strings.Contains(strings.ToLower(cmd.Name), "gummie") ||
			strings.Contains(strings.ToLower(cmd.Use), "gummie") {
			t.Fatalf("gummie leaked into agent-context: %+v", cmd)
		}
	}
}
