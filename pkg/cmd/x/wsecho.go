package x

import (
	"github.com/apigear-io/cli/pkg/net"
	"github.com/spf13/cobra"
)

func NewWSEchoCommand() *cobra.Command {
	var opts net.WSEchoOptions
	var cmd = &cobra.Command{
		Use:     "wsecho",
		Aliases: []string{"wse", "websocket-echo"},
		Short:   "Run the WebSocket echo server",
		Long:    `The WebSocket echo server echoes back any message it receives from clients.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return net.RunWSEcho(cmd.Context(), opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Addr, "address", "a", ":8080", "WebSocket server address")

	return cmd
}
