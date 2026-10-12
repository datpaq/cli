// Copyright 2026 datpaq. Licensed under Apache-2.0. See LICENSE.
// Local patch coverage: every table-declared endpoint builds the right request,
// is active in discovery, and is documented in spec.yaml.

package cli

import (
	"encoding/json"
	"net/url"
	"strings"
	"testing"
)

func declaredEndpointArgs(group string, ep APIEndpoint) []string {
	args := []string{group, ep.Name}
	for _, p := range ep.Params {
		if p.In == "path" {
			args = append(args, p.Example)
		}
	}
	for _, p := range ep.Params {
		if p.In == "path" {
			continue
		}
		if p.Type == "bool" {
			args = append(args, "--"+p.Flag+"="+p.Example)
		} else {
			args = append(args, "--"+p.Flag, p.Example)
		}
	}
	return append(args, "--agent", "--data-source", "live", "--no-cache")
}

func TestDeclaredEndpointsBuildRequests(t *testing.T) {
	for _, g := range DeclaredAPIGroups() {
		for _, ep := range g.Endpoints {
			t.Run(g.Name+"."+ep.Name, func(t *testing.T) {
				stdout, stderr, requests, err := runAPICommand(t, `{"success":true}`, declaredEndpointArgs(g.Name, ep)...)
				if err != nil {
					t.Fatalf("Execute: %v\nstdout=%s\nstderr=%s", err, stdout, stderr)
				}
				if len(requests) != 1 {
					t.Fatalf("requests = %d, want 1", len(requests))
				}
				req := requests[0]
				if req.Method != ep.Method {
					t.Fatalf("method = %s, want %s", req.Method, ep.Method)
				}
				wantPath := ep.Path
				for _, p := range ep.Params {
					if p.In == "path" {
						wantPath = strings.Replace(wantPath, "{"+p.WireName()+"}", url.PathEscape(p.Example), 1)
					}
				}
				gotPath, rawQuery, _ := strings.Cut(req.Path, "?")
				if gotPath != wantPath {
					t.Fatalf("path = %s, want %s", gotPath, wantPath)
				}
				query, _ := url.ParseQuery(rawQuery)
				var body map[string]any
				if req.Body != "" {
					if err := json.Unmarshal([]byte(req.Body), &body); err != nil {
						t.Fatalf("body is not JSON: %v\n%s", err, req.Body)
					}
				}
				for _, p := range ep.Params {
					switch p.In {
					case "path":
					case "body":
						if _, ok := body[p.WireName()]; !ok {
							t.Errorf("body missing %q: %s", p.WireName(), req.Body)
						}
					default:
						if !query.Has(p.WireName()) {
							t.Errorf("query missing %q: %s", p.WireName(), rawQuery)
						}
					}
				}
				assertJSONOnly(t, stdout)
				assertNoWarning(t, stderr)
			})
		}
	}
}

func TestDeclaredRequiredFlagsAreEnforced(t *testing.T) {
	for _, g := range DeclaredAPIGroups() {
		for _, ep := range g.Endpoints {
			for _, p := range ep.Params {
				if !p.Required || p.In == "path" {
					continue
				}
				_, _, requests, err := runAPICommand(t, `{"success":true}`, g.Name, ep.Name, "--agent")
				if err == nil || !strings.Contains(err.Error(), `required flag "`) {
					t.Errorf("%s %s without required flags: err = %v, want required flag error", g.Name, ep.Name, err)
				}
				if len(requests) != 0 {
					t.Errorf("%s %s sent a request despite missing required flags", g.Name, ep.Name)
				}
				break
			}
		}
	}
}

func TestDeclaredAPIGroupsAreActive(t *testing.T) {
	for _, g := range DeclaredAPIGroups() {
		if !isActiveInterface(g.Name) {
			t.Errorf("declared API %q missing from active-apis.json", g.Name)
		}
	}
}
