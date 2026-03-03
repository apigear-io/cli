package serve

import (
	"fmt"
	"os"

	"github.com/apigear-io/cli/internal/handler"
	_ "github.com/apigear-io/cli/internal/swagger"
	"github.com/apigear-io/cli/pkg/foundation/logging"
	"github.com/apigear-io/cli/pkg/runtime/network"
	"github.com/apigear-io/cli/pkg/stream/config"
	"github.com/apigear-io/cli/web"
	"github.com/spf13/cobra"
)

// ServeOptions holds the configuration for the serve command
type ServeOptions struct {
	Host   string
	Port   int
	Addr   string
	WebDir string
	UI     bool
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

			// Initialize stream config path for persistence
			configPath := "./stream.yaml"
			handler.SetStreamConfigPath(configPath)

			// Load existing stream config if it exists
			if _, err := os.Stat(configPath); err == nil {
				cfg, err := config.LoadConfig(configPath)
				if err != nil {
					logging.Warn().Err(err).Msg("Failed to load stream config, starting with empty state")
				} else {
					// Initialize stream services and load config
					services := handler.GetStreamServices()

					if len(cfg.Proxies) > 0 {
						logging.Info().Msgf("Loading %d proxies from config", len(cfg.Proxies))
						if err := services.ProxyManager.LoadFromConfig(cfg.Proxies); err != nil {
							logging.Warn().Err(err).Msg("Failed to load some proxies")
						}
					}

					if len(cfg.Clients) > 0 {
						logging.Info().Msgf("Loading %d clients from config", len(cfg.Clients))
						if err := services.ClientManager.LoadFromConfig(cfg.Clients); err != nil {
							logging.Warn().Err(err).Msg("Failed to load some clients")
						}
					}
				}
			}

			// Create HTTP server
			httpServer := network.NewHTTPServer(&network.HttpServerOptions{
				Addr: opts.Addr,
			})

			// Register routes
			router := httpServer.Router()
			handler.RegisterAPIRoutes(router)

			// Determine which UI to serve (priority: custom dir > embedded > swagger)
			hasWebUI := false

			// Check if custom Web UI directory is specified and exists
			if opts.WebDir != "" {
				if _, err := os.Stat(opts.WebDir); err == nil {
					// Custom Web UI directory exists, serve from filesystem
					handler.RegisterWebUIRoutes(router, opts.WebDir)
					logging.Info().Msgf("Web UI directory found at: %s", opts.WebDir)
					hasWebUI = true
				}
			}

			// If no custom directory, try embedded Web UI
			if !hasWebUI && web.Available() {
				webFS, err := web.FS()
				if err == nil {
					handler.RegisterEmbeddedWebUIRoutes(router, webFS)
					logging.Info().Msg("Serving embedded Web UI")
					hasWebUI = true
				} else {
					logging.Warn().Err(err).Msg("Failed to load embedded Web UI")
				}
			}

			// If no Web UI available, serve Swagger at root
			if !hasWebUI {
				handler.RegisterSwaggerRoutes(router)
				logging.Info().Msg("Web UI not available, serving Swagger at root")
			}

			// Start server
			if err := httpServer.Start(); err != nil {
				return fmt.Errorf("failed to start HTTP server: %w", err)
			}

			logging.Info().Msgf("Server starting on %s", opts.Addr)
			logging.Info().Msgf("API endpoints available at http://%s/api/v1", opts.Addr)

			// Log appropriate UI location
			if hasWebUI {
				logging.Info().Msgf("Web UI available at http://%s/", opts.Addr)
				logging.Info().Msgf("Swagger UI available at http://%s/swagger/index.html", opts.Addr)
			} else {
				logging.Info().Msgf("Swagger UI available at http://%s/swagger/index.html", opts.Addr)
			}

			// Open browser if --ui flag is set
			if opts.UI {
				go func() {
					var url string
					if hasWebUI {
						url = fmt.Sprintf("http://%s/", opts.Addr)
					} else {
						url = fmt.Sprintf("http://%s/swagger/index.html", opts.Addr)
					}
					openBrowser(url)
				}()
			}

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
	cmd.Flags().StringVar(&opts.WebDir, "web-dir", "", "directory containing web UI static files (overrides embedded UI)")
	cmd.Flags().BoolVar(&opts.UI, "ui", false, "automatically open the UI in your default browser")

	return cmd
}
