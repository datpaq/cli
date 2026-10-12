// Copyright 2026 datpaq. Licensed under Apache-2.0. See LICENSE.

package mcp

import (
	"testing"

	"github.com/datpaq/cli/internal/cli"
	"github.com/mark3labs/mcp-go/server"
)

func TestDeclaredAPIToolsRegisteredWithParams(t *testing.T) {
	s := server.NewMCPServer("test", "0")
	RegisterTools(s)
	for _, g := range cli.DeclaredAPIGroups() {
		for _, ep := range g.Endpoints {
			name := declaredToolName(g.Name, ep)
			tool := s.GetTool(name)
			if tool == nil {
				t.Errorf("MCP tool %q not registered", name)
				continue
			}
			for _, p := range ep.Params {
				if _, ok := tool.Tool.InputSchema.Properties[declaredPublicName(p)]; !ok {
					t.Errorf("MCP tool %q missing param %q", name, declaredPublicName(p))
				}
			}
		}
	}
}
