// Copyright 2026 datpaq. Licensed under Apache-2.0. See LICENSE.
// Hand-authored: verifies newly-shipped datpaq.com APIs appear in the
// curated discovery manifest so `datpaq api` / `datpaq sample` surface them.

package cli

import "testing"

func TestNewLookupAPIsAreActive(t *testing.T) {
	for _, slug := range []string{"dns", "domain-lookup", "mac-address", "mx-lookup"} {
		if !isActiveInterface(slug) {
			t.Errorf("expected %q in active-apis.json", slug)
		}
	}
}

func TestActiveAPICountIncludesNewLookups(t *testing.T) {
	got := activeAPICount()
	if got < 17 {
		t.Errorf("activeAPICount() = %d, want at least 17 (13 prior + 4 new lookups)", got)
	}
}
