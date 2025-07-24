package spec

import (
	"context"
	"fmt"

	"github.com/apigear-io/cli/pkg/spec"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterMCPTools(s *server.MCPServer) {
	registerSpecCheckTool(s)
}

func registerSpecCheckTool(s *server.MCPServer) {
	specCheckTool := mcp.NewTool("specificationCheck",
		mcp.WithDescription("Load and validate API specification(module, solution, rules) files"),
		mcp.WithString("file", mcp.Required(), mcp.Description("Path to specification file")),
	)
	s.AddTool(specCheckTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		file, err := request.RequireString("file")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		result := ""
		spec, err := spec.CheckFile(file)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		if spec.Valid() {
			result = fmt.Sprintf("valid: %s\n", file)
		} else {
			result = fmt.Sprintf("invalid: %s\n", file)
			for _, desc := range spec.Errors {
				result += fmt.Sprintf("file: %s\n", file)
				result += fmt.Sprintf("%s\n", desc.String())
			}
		}
		return mcp.NewToolResultText(string(result)), nil
	})
}
