package mon

import (
	"github.com/apigear-io/cli/internal/handler"
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
			// Create HTTP server
			httpServer := network.NewHTTPServer(&network.HttpServerOptions{
				Addr: addr,
			})

			// Register monitor routes
			handler.RegisterMonitorRoutes(httpServer.Router())

			// Start server
			if err := httpServer.Start(); err != nil {
				return err
			}

			// Register event hook (directly on global emitter)
			monitoring.Emitter.AddHook(func(e *monitoring.Event) {
				logging.Info().Msgf("event: %s %s %v", e.Type.String(), e.Source, e.Data)
			})

			logging.Info().Msgf("Monitor server started on http://%s", addr)
			logging.Info().Msgf("Monitor endpoint: POST http://%s/monitor/{{source}}", addr)

			// Wait for shutdown signal
			return network.WaitForShutdown(cmd.Context(), func() {
				logging.Info().Msg("stopping monitor server...")
				httpServer.Stop()
			})
		},
	}
	cmd.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1:5555", "address to listen on")
	return cmd
}
