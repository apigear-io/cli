package mon

import (
	"github.com/apigear-io/cli/pkg/streams/msgio"
	"github.com/spf13/cobra"
)

func NewServerCommand() *cobra.Command {
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
				ServerURL: natsURL,
				Verbose:   verbose,
				Pretty:    pretty,
				Headers:   headers,
			}
			if deviceID != "" {
				opts.DeviceID = deviceID
			}
			return msgio.Tail(cmd.Context(), opts)
		},
	}

	cmd.Flags().StringVarP(&natsURL, "nats-url", "n", "nats://127.0.0.1:4222", "NATS server URL")
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logging")
	cmd.Flags().StringVarP(&deviceID, "device-id", "d", "", "device ID to monitor")
	cmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "pretty print JSON output")
	cmd.Flags().BoolVarP(&headers, "headers", "H", false, "include headers in output")
	return cmd
}
