package x

import (
	"time"

	"github.com/apigear-io/cli/pkg/net"
	"github.com/spf13/cobra"
)

func NewWSCatCommand() *cobra.Command {
	var opts net.WSClientOptions
	var cmd = &cobra.Command{
		Use:     "wscat",
		Aliases: []string{"ws", "websocket"},
		Short:   "Run the WebSocket cat client",
		Long:    `The WebSocket cat client connects to the WebSocket proxy and allows sending and receiving messages.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return net.RunWSClient(cmd.Context(), opts)
		},
	}

	cmd.Flags().StringVarP(&opts.URL, "url", "u", "", "WebSocket server URL")
	cmd.Flags().DurationVarP(&opts.Interval, "interval", "i", 100*time.Millisecond, "Interval between messages")
	cmd.Flags().IntVarP(&opts.Repeat, "repeat", "r", 1, "Number of times to repeat the messages")
	cmd.Flags().BoolVarP(&opts.DecodeJSON, "decode-json", "d", false, "Decode JSON messages")

	return cmd
}
