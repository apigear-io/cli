package tpl

import (
	"context"

	"github.com/apigear-io/cli/pkg/repos"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterMCPTools(s *server.MCPServer) {
	registerTemplateListTool(s)
}

func registerTemplateListTool(s *server.MCPServer) {
	templateListTool := mcp.NewTool("templateList",
		mcp.WithDescription("List available templates from registry"),
	)
	s.AddTool(templateListTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		infos, err := repos.Registry.List()
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		output_bytes := repoInfosToCSV(infos)

		return mcp.NewToolResultText(string(output_bytes)), nil
	})
}
