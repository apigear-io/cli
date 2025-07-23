package tpl

import (
	"context"
	"os/exec"

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
		cmd := exec.Command("apigear", "template", "list")
		output_bytes, err := cmd.CombinedOutput()
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(string(output_bytes)), nil
	})
}
