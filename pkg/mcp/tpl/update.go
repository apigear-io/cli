package tpl

import (
	"context"

	"github.com/apigear-io/cli/pkg/repos"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerTemplateUpdateTool(s *server.MCPServer) {
	templateUpdateTool := mcp.NewTool("templateUpdate",
		mcp.WithDescription("Update the template registry from remote sources. Use this to refresh the local cache with latest templates and versions."),
		mcp.WithReadOnlyHintAnnotation(false),
		mcp.WithDestructiveHintAnnotation(false),
		mcp.WithIdempotentHintAnnotation(true),
	)
	s.AddTool(templateUpdateTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		err := repos.Registry.Update()
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText("template registry updated"), nil
	})
}
