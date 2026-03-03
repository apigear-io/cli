// Package stream provides commands for WebSocket streaming and proxy functionality.
package stream

import (
	"github.com/spf13/cobra"
)

// NewRootCommand creates the stream root command.
func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stream",
		Short: "WebSocket streaming tools",
		Long: `Tools for WebSocket streaming, proxying, and message debugging.

Available subcommands:
  proxy      Run proxy server or manage proxy configuration
  client     Manage ObjectLink clients
  publish    Send messages to a WebSocket server
  subscribe  Start a WebSocket server and print received messages`,
	}

	// Add subcommands
	cmd.AddCommand(NewProxyCommand())
	cmd.AddCommand(NewClientCommand())
	cmd.AddCommand(NewPublishCommand())
	cmd.AddCommand(NewSubscribeCommand())

	return cmd
}
