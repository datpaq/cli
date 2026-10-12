// Copyright 2026 datpaq. Licensed under Apache-2.0. See LICENSE.

package mcp

import (
	"encoding/base64"
	"os"
	"path/filepath"
	"testing"

	"github.com/datpaq/cli/internal/cli"
	"github.com/mark3labs/mcp-go/server"
)

func TestEncodeDeclaredFileArgsReadsLocalPaths(t *testing.T) {
	path := filepath.Join(t.TempDir(), "doc.pdf")
	if err := os.WriteFile(path, []byte("hello"), 0o600); err != nil {
		t.Fatal(err)
	}
	ep := cli.APIEndpoint{Params: []cli.APIParam{
		{Flag: "file", Type: "file"},
		{Flag: "files", Type: "files"},
	}}
	args := map[string]any{"file": path, "files": []any{path, "aGk="}}
	if err := encodeDeclaredFileArgs(ep, args); err != nil {
		t.Fatal(err)
	}
	want := base64.StdEncoding.EncodeToString([]byte("hello"))
	if args["file"] != want {
		t.Errorf("file = %v, want base64 of file content", args["file"])
	}
	if list := args["files"].([]any); list[0] != want || list[1] != "aGk=" {
		t.Errorf("files = %v, want [encoded path, unchanged base64]", list)
	}
}

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
