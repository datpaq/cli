// Copyright 2026 datpaq. Licensed under Apache-2.0. See LICENSE.
// Hand-authored: verifies datpaq.com active APIs appear in the curated
// discovery manifest so `datpaq api` / `datpaq sample` surface them.

package cli

import "testing"

func TestSupportedWebsiteActiveAPIsAreActive(t *testing.T) {
	// Mirrored from ProApi/api_list_documentation_upload.sql IsActive=true,
	// using CLI command/interface names where website slugs differ:
	// dictionary -> define, dns-lookup -> dns.
	for _, slug := range []string{
		"aircraft",
		"company-enrichment",
		"convert-time",
		"country-codes",
		"current-time",
		"define",
		"dns",
		"domain-lookup",
		"email-validation",
		"ev-charger",
		"exchange-rates-and-currency",
		"helicopter",
		"image-processing",
		"ip-geolocation",
		"ip-intelligence",
		"mac-address",
		"mx-lookup",
		"phone-validation",
		"precious-metals",
		"profanity",
		"public-holidays",
		"sample-data",
		"secure-relay",
		"spell-check",
		"states",
		"text-language",
		"thesaurus",
		"unit-conversion",
		"user-avatar",
		"validate-ip",
		"vin-lookup",
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
