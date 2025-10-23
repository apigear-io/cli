package mon

import (
	"github.com/apigear-io/cli/pkg/app"
	"github.com/apigear-io/cli/pkg/streams/msgio"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

func NewRunCommand() *cobra.Command {
	var natsURL string
	var verbose bool
	var deviceID string
	var pretty bool
	var headers bool
	var cmd = &cobra.Command{
		Use:     "run",
		Aliases: []string{"r", "start"},
		Short:   "Run the monitor server",
		Long:    `The monitor server runs on a HTTP port and listens for API calls.`,
		RunE: func(cmd *cobra.Command, _ []string) error {

			opts := msgio.TailOptions{
				Verbose:  verbose,
				Pretty:   pretty,
				Headers:  headers,
				DeviceID: deviceID,
			}

			return app.WithNATS(cmd.Context(), natsURL, func(nc *nats.Conn) error {
				tailer := msgio.NewTailer(nc, opts)
				return tailer.Run(cmd.Context())
			})
		},
	}

	cmd.Flags().StringVarP(&natsURL, "nats-url", "n", nats.DefaultURL, "NATS server URL")
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logging")
	cmd.Flags().StringVarP(&deviceID, "device-id", "d", "", "device ID to monitor")
	cmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "pretty print JSON output")
	cmd.Flags().BoolVarP(&headers, "headers", "H", false, "include headers in output")
	return cmd
}
