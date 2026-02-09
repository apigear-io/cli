package mon

import (
	"github.com/apigear-io/cli/pkg/foundation/logging"
	"github.com/apigear-io/cli/pkg/runtime/monitoring"
	"github.com/apigear-io/cli/pkg/runtime/network"
	"github.com/spf13/cobra"
)

func NewServerCommand() *cobra.Command {
	var addr string
	var cmd = &cobra.Command{
		Use:     "run",
		Aliases: []string{"r", "start"},
		Short:   "Run the monitor server",
		Long:    `The monitor server runs on a HTTP port and listens for API calls.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			netman := network.NewManager()
			opts := network.Options{
				HttpAddr: addr,
			}
			err := netman.Start(&opts)
			if err != nil {
				return err
			}
			netman.MonitorEmitter().AddHook(func(e *monitoring.Event) {
				logging.Info().Msgf("event: %s %s %v", e.Type.String(), e.Source, e.Data)
			})
			// Note: NATS-based OnMonitorEvent removed. Only local hooks work now.
			// Events received via HTTP /monitor/{source} will trigger the hook above.
			return netman.Wait(cmd.Context())
		},
	}
	cmd.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1:5555", "address to listen on")
	return cmd
}
