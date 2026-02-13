package serve

import (
	"fmt"
	"net/http"

	"github.com/apigear-io/cli/internal/handler"
	_ "github.com/apigear-io/cli/docs"
	"github.com/apigear-io/cli/pkg/foundation/logging"
	"github.com/apigear-io/cli/pkg/runtime/network"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"
	httpSwagger "github.com/swaggo/http-swagger"
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

			// Create and start NetworkManager
			netman := network.NewManager()
			netOpts := &network.Options{
				HttpAddr:          opts.Addr,
				HttpDisabled:      false,
				MonitorDisabled:   true,
				ObjectAPIDisabled: true,
			}

			err := netman.Start(netOpts)
			if err != nil {
				return fmt.Errorf("failed to start HTTP server: %w", err)
			}

			// Register API routes
			router := netman.HttpServer().Router()

			// API v1 routes
			router.Route("/api/v1", func(r chi.Router) {
				r.Get("/health", handler.Health())
				r.Get("/status", handler.Status())
			})

			// Swagger documentation
			router.Get("/swagger/*", httpSwagger.WrapHandler)

			// Root redirect to Swagger UI
			router.Get("/", func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
			})

			logging.Info().Msgf("Server starting on %s", opts.Addr)
			logging.Info().Msgf("API endpoints available at http://%s/api/v1", opts.Addr)
			logging.Info().Msgf("Swagger UI available at http://%s/swagger/index.html", opts.Addr)

			// Wait for shutdown signal
			return netman.Wait(cmd.Context())
		},
	}

	cmd.Flags().StringVar(&opts.Addr, "addr", "", "address to listen on (overrides host:port)")
	cmd.Flags().StringVar(&opts.Host, "host", "localhost", "host to listen on")
	cmd.Flags().IntVar(&opts.Port, "port", 8080, "port to listen on")

	return cmd
}
