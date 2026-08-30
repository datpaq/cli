// Copyright 2026 datpaq. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"net/url"
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
