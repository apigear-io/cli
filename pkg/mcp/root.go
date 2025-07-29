package mcp

import (
	"context"
	"fmt"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/mcp/gen"
	"github.com/apigear-io/cli/pkg/mcp/spec"
	"github.com/apigear-io/cli/pkg/mcp/tpl"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func RunMCPServer() error {
	// Create MCP server
	s := server.NewMCPServer(
		"apigear",
		retrieveVersion(),
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)

	spec.RegisterMCPTools(s)
	tpl.RegisterMCPTools(s)
	gen.RegisterMCPTools(s)

	addCoreTools(s)

	return server.ServeStdio(s)
}

func addCoreTools(s *server.MCPServer) {
	// Version tool
	versionTool := mcp.NewTool("version",
		mcp.WithDescription("Display version information"),
	)
	s.AddTool(versionTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return mcp.NewToolResultText(retrieveVersion()), nil
	})
}

func retrieveVersion() string {
	bi := cfg.GetBuildInfo("cli")
	version := fmt.Sprintf("%s-%s-%s", bi.Version, bi.Commit, bi.Date)
	return version
}
