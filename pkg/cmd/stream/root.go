// Package stream provides commands for WebSocket streaming and proxy functionality.
package stream

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/apigear-io/cli/pkg/foundation/logging"
	"github.com/apigear-io/cli/pkg/stream"
	"github.com/apigear-io/cli/pkg/stream/config"
)

// Options holds configuration for the stream command.
type Options struct {
	ConfigFile string
	Verbose    bool
	Trace      bool
	LogLevel   string
	Watch      bool
}

// NewRootCommand creates the stream root command.
func NewRootCommand() *cobra.Command {
	opts := &Options{}

	cmd := &cobra.Command{
		Use:   "stream [config.yaml]",
		Short: "WebSocket streaming and proxy server",
		Long: `Start a WebSocket streaming server with proxy, client management,
and real-time message tracing capabilities.

The stream command provides:
- WebSocket proxy with multiple modes (proxy, echo, backend, inbound-only)
- ObjectLink client management with auto-reconnect
- Message tracing and replay (JSONL format)
- Real-time monitoring and statistics

Examples:
  # Start with default configuration
  apigear stream

  # Start with custom configuration
  apigear stream config.yaml

  # Start with verbose logging
  apigear stream --verbose

  # Enable trace logging
  apigear stream --trace

  # Watch config file for changes
  apigear stream --watch`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Note: Log level is set via DEBUG environment variable
			// DEBUG=1 for debug, DEBUG=2 for trace

			// Determine config file
			configFile := opts.ConfigFile
			if len(args) > 0 {
				configFile = args[0]
			}
			if configFile == "" {
				configFile = "stream.yaml"
			}

			// Load or create config
			cfg, created, err := config.LoadOrCreateConfig(configFile)
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}

			if created {
				logging.Info().Msgf("Created default config file: %s", configFile)
				logging.Info().Msg("Edit the config file and restart to configure proxies and clients")
			}

			// Validate config
			if err := cfg.Validate(); err != nil {
				return fmt.Errorf("invalid config: %w", err)
			}

			// Override config with flags
			if opts.Verbose {
				cfg.Verbose = true
			}
			if opts.Trace {
				cfg.Trace = true
			}

			// Initialize services
			services := stream.NewServices()
			defer services.Close()

			// Load proxies from config
			if len(cfg.Proxies) > 0 {
				logging.Info().Msgf("Loading %d proxies from config", len(cfg.Proxies))
				if err := services.ProxyManager.LoadFromConfig(cfg.Proxies); err != nil {
					log.Warn().Err(err).Msg("failed to load some proxies")
				}
			}

			// Load clients from config
			if len(cfg.Clients) > 0 {
				logging.Info().Msgf("Loading %d clients from config", len(cfg.Clients))
				if err := services.ClientManager.LoadFromConfig(cfg.Clients); err != nil {
					log.Warn().Err(err).Msg("failed to load some clients")
				}
			}

			// Print summary
			proxies := services.ProxyManager.ListProxies()
			clients := services.ClientManager.ListClients()

			logging.Info().Msgf("Stream server started with %d proxies and %d clients",
				len(proxies), len(clients))

			if len(proxies) > 0 {
				logging.Info().Msg("Active proxies:")
				for _, p := range proxies {
					logging.Info().Msgf("  - %s: %s -> %s (%s, %s)",
						p.Name, p.Listen, p.Backend, p.Mode, p.Status)
				}
			}

			if len(clients) > 0 {
				logging.Info().Msg("Active clients:")
				for _, c := range clients {
					logging.Info().Msgf("  - %s: %s (%s)",
						c.Name, c.URL, c.Status)
				}
			}

			if len(proxies) == 0 && len(clients) == 0 {
				logging.Info().Msg("No proxies or clients configured. Edit the config file to add them.")
			}

			// Wait for interrupt signal
			sigCh := make(chan os.Signal, 1)
			signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

			logging.Info().Msg("Press Ctrl+C to stop")

			<-sigCh
			logging.Info().Msg("Shutting down...")

			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.ConfigFile, "config", "c", "", "config file (default: stream.yaml)")
	cmd.Flags().BoolVarP(&opts.Verbose, "verbose", "v", false, "enable verbose logging")
	cmd.Flags().BoolVarP(&opts.Trace, "trace", "t", false, "enable trace logging to files")
	cmd.Flags().StringVar(&opts.LogLevel, "log-level", "", "log level (trace, debug, info, warn, error)")
	cmd.Flags().BoolVarP(&opts.Watch, "watch", "w", false, "watch config file for changes")

	// Add subcommands
	cmd.AddCommand(NewProxyCommand())
	cmd.AddCommand(NewClientCommand())
	cmd.AddCommand(NewEchoCommand())

	return cmd
}
