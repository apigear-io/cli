package cmd

import (
	"github.com/apigear-io/cli/pkg/mcp"
	"github.com/spf13/cobra"
)

func NewMCPCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mcp",
		Short: "Start MCP server exposing apigear CLI commands",
		Long:  `Start a Model Context Protocol (MCP) server that exposes selected apigear CLI commands as tools for AI assistants.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return mcp.RunMCPServer()
		},
	}
	return cmd
}
