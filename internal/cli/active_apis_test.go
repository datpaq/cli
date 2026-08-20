// Copyright 2026 datpaq. Licensed under Apache-2.0. See LICENSE.
// Hand-authored: verifies datpaq.com active APIs appear in the curated
// discovery manifest so `datpaq api` / `datpaq sample` surface them.

package cli

import "testing"

func TestSupportedWebsiteActiveAPIsAreActive(t *testing.T) {
	// Mirrored from website API_ACTIVE_SLUGS (35), using CLI command/interface
	// names where website slugs differ:
	// dictionary -> define, dns-lookup -> dns. us-states matches the website slug.
	for _, slug := range []string{
		"aircraft",
		"calendar",
		"convert-time",
		"country-codes",
		"current-time",
		"define",
		"dns",
		"domain-lookup",
		"ev-charger",
		"exchange-rates-and-currency",
		"geocoding",
		"helicopter",
		"image-processing",
		"ip-geolocation",
		"ip-intelligence",
		"mac-address",
		"mx-lookup",
		"pdf-generation",
		"profanity",
		"public-holidays",
		"qr-code",
		"sample-data",
		"spell-check",
		"us-states",
		"text-language",
		"thesaurus",
		"unit-conversion",
		"user-avatar",
		"validate-ip",
		"vin-lookup",
		"weather",
		"web-scraping",
		"web-screenshot",
		"whois",
		"working-days",
	} {
		if !isActiveInterface(slug) {
			t.Errorf("expected %q in active-apis.json", slug)
		}
	}
}

func TestActiveAPICountMatchesSupportedWebsiteActiveAPIs(t *testing.T) {
	got := activeAPICount()
	if got != 35 {
		t.Errorf("activeAPICount() = %d, want 35 supported active APIs", got)
	}
}

func TestInactiveWebsiteCatalogAPIsAreNotActive(t *testing.T) {
	// Previously active in the CLI manifest but no longer in website API_ACTIVE_SLUGS.
	for _, slug := range []string{
		"company-enrichment",
		"email-validation",
		"phone-validation",
		"precious-metals",
		"secure-relay",
		"web-search",
	} {
		if isActiveInterface(slug) {
			t.Errorf("did not expect inactive %q in active-apis.json", slug)
		}
	}
}
