package serve

import (
	"fmt"

	"github.com/apigear-io/cli/internal/handler"
	_ "github.com/apigear-io/cli/internal/swagger"
	"github.com/apigear-io/cli/pkg/foundation/logging"
	"github.com/apigear-io/cli/pkg/runtime/network"
	"github.com/spf13/cobra"
)

// ServeOptions holds the configuration for the serve command
type ServeOptions struct {
	Host string
	Port int
	Addr string
}

// NewServeCommand creates a new serve command
func NewServeCommand() *cobra.Command {
	opts := &ServeOptions{
		Host: "localhost",
		Port: 8080,
	}

	cmd := &cobra.Command{
		Use:     "serve",
		Aliases: []string{"server", "s"},
		Short:   "Start the HTTP REST API server",
		Long:    `Start the HTTP REST API server with health, status, and Swagger documentation endpoints.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			// Set Addr from host and port if not explicitly provided
			if opts.Addr == "" {
				opts.Addr = fmt.Sprintf("%s:%d", opts.Host, opts.Port)
			}

			// Create HTTP server
			httpServer := network.NewHTTPServer(&network.HttpServerOptions{
				Addr: opts.Addr,
			})

			// Register routes
			router := httpServer.Router()
			handler.RegisterAPIRoutes(router)
			handler.RegisterSwaggerRoutes(router)

			// Start server
			if err := httpServer.Start(); err != nil {
				return fmt.Errorf("failed to start HTTP server: %w", err)
			}

			logging.Info().Msgf("Server starting on %s", opts.Addr)
			logging.Info().Msgf("API endpoints available at http://%s/api/v1", opts.Addr)
			logging.Info().Msgf("Swagger UI available at http://%s/swagger/index.html", opts.Addr)

			// Wait for shutdown signal
			return network.WaitForShutdown(cmd.Context(), func() {
				logging.Info().Msg("stopping HTTP server...")
				httpServer.Stop()
			})
		},
	}

	cmd.Flags().StringVar(&opts.Addr, "addr", "", "address to listen on (overrides host:port)")
	cmd.Flags().StringVar(&opts.Host, "host", "localhost", "host to listen on")
	cmd.Flags().IntVar(&opts.Port, "port", 8080, "port to listen on")

	return cmd
}
