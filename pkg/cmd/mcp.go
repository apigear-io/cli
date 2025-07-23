package cmd

import (
	"context"
	"os/exec"
	"fmt"

	"github.com/apigear-io/cli/pkg/cfg"

	"github.com/apigear-io/cli/pkg/cmd/tpl"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
)

func NewMCPCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mcp",
		Short: "Start MCP server exposing apigear CLI commands",
		Long:  `Start a Model Context Protocol (MCP) server that exposes selected apigear CLI commands as tools for AI assistants.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runMCPServer()
		},
	}
	return cmd
}

func runMCPServer() error {
	bi := cfg.GetBuildInfo("cli")
	version := fmt.Sprintf("%s-%s-%s", bi.Version, bi.Commit, bi.Date)
	// Create MCP server
	s := server.NewMCPServer(
		"apigear-cli",
		version,
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)

	tpl.RegisterMCPTools(s)

	addCoreTools(s)

	return server.ServeStdio(s)
}

func addCoreTools(s *server.MCPServer) {
	// Version tool
	versionTool := mcp.NewTool("version",
		mcp.WithDescription("Display version information"),
	)
	s.AddTool(versionTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		cmd := exec.Command("apigear", "version")
		output_bytes, err := cmd.CombinedOutput()
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(string(output_bytes)), nil
	})
}