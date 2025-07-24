package gen

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterMCPTools(s *server.MCPServer) {
	registerGenerateSolutionTool(s)
}

func registerGenerateSolutionTool(s *server.MCPServer) {
	genSolutionTool := mcp.NewTool("generateSolution",
		mcp.WithDescription("Generate SDK based on a solution document"),
		mcp.WithString("solution", mcp.Required(), mcp.Description("Path to solution file")),
		mcp.WithString("force", mcp.Description("Force overwrite (true/false)")),
		mcp.WithString("watch", mcp.Description("Watch for changes (true/false). This keeps the process running.")),
	)
	s.AddTool(genSolutionTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		solutionPath, err := request.RequireString("solution")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		forceEnabled := false
		if force, err := request.RequireString("force"); err == nil && force == "true" {
			forceEnabled = true
		}
		watchEnabled := false
		if watch, err := request.RequireString("watch"); err == nil && watch == "true" {
			watchEnabled = true
		}

		err = runGenerateSolution(solutionPath, watchEnabled, forceEnabled)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(fmt.Sprintf("Successfully ran code generation for solution file: %s", solutionPath)), nil
	})
}
