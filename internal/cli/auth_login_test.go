// Copyright 2026 datpaq. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"net/url"
	"strings"
	"testing"
)

func TestAuthLoginExchangeURLUsesMainsiteNamespace(t *testing.T) {
	t.Parallel()

	u, err := url.Parse(authLoginExchangeURL)
	if err != nil {
		t.Fatalf("parse authLoginExchangeURL: %v", err)
	}
	if got, want := u.Scheme, "https"; got != want {
		t.Fatalf("scheme = %q, want %q", got, want)
	}
	if got, want := u.Host, "datpaq.com"; got != want {
		t.Fatalf("host = %q, want %q", got, want)
	}
	if got, want := u.Path, "/api/internal/cli/auth/exchange"; got != want {
		t.Fatalf("path = %q, want %q", got, want)
	}
}

func TestAuthLoginCallbackUsesDatpaqBrandShell(t *testing.T) {
	t.Parallel()

	required := []string{
		`<header class="site-header">`,
		`class="wordmark"`,
		`justify-content: flex-start`,
		`class="success-icon"`,
		`<footer class="site-footer">`,
		`&copy; 2026 DATPAQ, Inc. All rights reserved.`,
	}
	for _, want := range required {
		if !strings.Contains(callbackSuccessHTML, want) {
			t.Errorf("callbackSuccessHTML missing %q", want)
		}
	}

	forbidden := []string{"#0b0118", "#6b21a8", "Datpaq CLI — Signed in", "footer-mark", "concentric-ring"}
	for _, value := range forbidden {
		if strings.Contains(callbackSuccessHTML, value) {
			t.Errorf("callbackSuccessHTML still contains retired styling %q", value)
		}
	}
}
