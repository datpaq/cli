// Copyright 2026 datpaq. Licensed under Apache-2.0. See LICENSE.
// Local patch: typed MCP tools for the table-declared APIs in
// internal/cli/declared_api.go, so MCP parameters always match CLI wiring.

package mcp

import (
	"context"
	"encoding/base64"
	"os"
	"strings"

	"github.com/datpaq/cli/internal/cli"
	mcplib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerDeclaredAPITools(s *server.MCPServer) {
	for _, g := range cli.DeclaredAPIGroups() {
		for _, ep := range g.Endpoints {
			s.AddTool(declaredTool(g.Name, ep), declaredHandler(ep))
		}
	}
}

func declaredToolName(group string, ep cli.APIEndpoint) string {
	return group + "_" + ep.Name
}

func declaredPublicName(p cli.APIParam) string {
	return strings.ReplaceAll(p.Flag, "-", "_")
}

func declaredTool(group string, ep cli.APIEndpoint) mcplib.Tool {
	var required, optional []string
	for _, p := range ep.Params {
		if p.Required || p.In == "path" {
			required = append(required, declaredPublicName(p))
		} else {
			optional = append(optional, declaredPublicName(p))
		}
	}
	desc := ep.Short + "."
	if len(required) > 0 {
		desc += " Required: " + strings.Join(required, ", ") + "."
	}
	if len(optional) > 0 {
		desc += " Optional: " + strings.Join(optional, ", ") + "."
	}

	opts := []mcplib.ToolOption{mcplib.WithDescription(desc)}
	for _, p := range ep.Params {
		usage := p.Usage
		if p.Type == "file" || p.Type == "files" {
			usage += " (local path(s) readable by this server, or base64 content)"
		}
		popts := []mcplib.PropertyOption{mcplib.Description(usage)}
		if p.Required || p.In == "path" {
			popts = append(popts, mcplib.Required())
		}
		name := declaredPublicName(p)
		switch p.Type {
		case "int", "float":
			opts = append(opts, mcplib.WithNumber(name, popts...))
		case "bool":
			opts = append(opts, mcplib.WithBoolean(name, popts...))
		case "array", "files":
			opts = append(opts, mcplib.WithArray(name, popts...))
		case "object":
			opts = append(opts, mcplib.WithObject(name, popts...))
		default:
			opts = append(opts, mcplib.WithString(name, popts...))
		}
	}
	if ep.Method == "GET" {
		opts = append(opts, mcplib.WithReadOnlyHintAnnotation(true))
	}
	opts = append(opts,
		mcplib.WithDestructiveHintAnnotation(ep.Method == "DELETE"),
		mcplib.WithOpenWorldHintAnnotation(true),
	)
	return mcplib.NewTool(declaredToolName(group, ep), opts...)
}

func declaredHandler(ep cli.APIEndpoint) server.ToolHandlerFunc {
	var bindings []mcpParamBinding
	for _, p := range ep.Params {
		loc := p.In
		if loc == "" {
			loc = "query"
		}
		bindings = append(bindings, mcpParamBinding{PublicName: declaredPublicName(p), WireName: p.WireName(), Location: loc})
	}
	inner := makeAPIHandler(ep.Method, ep.Path, false, nil, bindings, nil)
	return func(ctx context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
		if err := encodeDeclaredFileArgs(ep, req.GetArguments()); err != nil {
			return mcplib.NewToolResultError(err.Error()), nil
		}
		return inner(ctx, req)
	}
}

// encodeDeclaredFileArgs mirrors the CLI's --file handling: values naming a
// readable local file are replaced with base64 content; other values are
// assumed to already be base64 and pass through unchanged.
func encodeDeclaredFileArgs(ep cli.APIEndpoint, args map[string]any) error {
	encode := func(v any) (any, error) {
		s, ok := v.(string)
		if !ok {
			return v, nil
		}
		if info, err := os.Stat(s); err != nil || info.IsDir() {
			return v, nil
		}
		data, err := os.ReadFile(s)
		if err != nil {
			return nil, err
		}
		return base64.StdEncoding.EncodeToString(data), nil
	}
	for _, p := range ep.Params {
		name := declaredPublicName(p)
		v, ok := args[name]
		if !ok {
			continue
		}
		switch p.Type {
		case "file":
			enc, err := encode(v)
			if err != nil {
				return err
			}
			args[name] = enc
		case "files":
			list, ok := v.([]any)
			if !ok {
				continue
			}
			for i, item := range list {
				enc, err := encode(item)
				if err != nil {
					return err
				}
				list[i] = enc
			}
		}
	}
	return nil
}

// declaredContextResources lists declared groups for the MCP context tool.
func declaredContextResources() []map[string]any {
	var out []map[string]any
	for _, g := range cli.DeclaredAPIGroups() {
		names := make([]string, 0, len(g.Endpoints))
		for _, ep := range g.Endpoints {
			names = append(names, ep.Name)
		}
		out = append(out, map[string]any{
			"name":        g.Name,
			"description": g.Short,
			"endpoints":   names,
			"searchable":  true,
		})
	}
	return out
}
