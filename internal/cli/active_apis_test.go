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
		"color-convert",
		"convert-time",
		"country-codes",
		"current-time",
		"data-format-convert",
		"define",
		"deslop",
		"diffy",
		"dns",
		"document-convert",
		"document-generation",
		"domain-lookup",
		"economic-indicators",
		"encoding-convert",
		"ev-charger",
		"gdp",
		"geocoding",
		"helicopter",
		"iban",
		"image-processing",
		"income-tax",
		"interest-rate",
		"ip-geolocation",
		"ip-intelligence",
		"libor",
		"lorem-ipsum",
		"mac-address",
		"markdown-convert",
		"market-cap",
		"mortgage-calculator",
		"mx-lookup",
		"ocr",
		"pdf-generation",
		"pdf-toolkit",
		"profanity",
		"public-holidays",
		"qr-code",
		"sample-data",
		"sec-filings",
		"spell-check",
		"translate",
		"unemployment",
		"us-states",
		"text-language",
		"thesaurus",
		"unit-conversion",
		"user-avatar",
		"validate-ip",
		"vat-validation-and-rates",
		"vin-lookup",
		"weather",
		"web-scraping",
		"web-screenshot",
		"web-search",
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
	if got != 58 {
		t.Errorf("activeAPICount() = %d, want 58 supported active APIs", got)
	}
}

func TestInactiveWebsiteCatalogAPIsAreNotActive(t *testing.T) {
	// Previously active in the CLI manifest but no longer in website API_ACTIVE_SLUGS.
	for _, slug := range []string{
		"company-enrichment",
		"email-validation",
		"exchange-rates-and-currency",
		"phone-validation",
		"precious-metals",
		"secure-relay",
	} {
		if isActiveInterface(slug) {
			t.Errorf("did not expect inactive %q in active-apis.json", slug)
		}
	}
}
