package stream

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/apigear-io/cli/pkg/foundation/logging"
	"github.com/apigear-io/cli/pkg/stream"
	"github.com/apigear-io/cli/pkg/stream/config"
)

// NewEchoCommand creates the echo server command.
func NewEchoCommand() *cobra.Command {
	opts := &struct {
		Listen  string
		Verbose bool
	}{
		Listen: "ws://localhost:8888/ws",
	}

	cmd := &cobra.Command{
		Use:   "echo",
		Short: "Start a simple echo server",
		Long: `Start a simple WebSocket echo server that sends back all received messages.

This is useful for testing WebSocket clients and proxies.

Examples:
  # Start echo server on default port
  apigear stream echo

  # Start on custom port
  apigear stream echo --listen ws://localhost:9999/ws

  # Enable verbose logging
  apigear stream echo --verbose`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Create a simple proxy in echo mode
			services := stream.NewServices()
			defer services.Close()

			// Create echo proxy
			cfg := config.ProxyConfig{
				Listen: opts.Listen,
				Mode:   "echo",
			}

			if err := services.ProxyManager.AddProxy("echo", cfg); err != nil {
				return fmt.Errorf("failed to create echo server: %w", err)
			}

			if err := services.ProxyManager.StartProxy("echo"); err != nil {
				return fmt.Errorf("failed to start echo server: %w", err)
			}

			logging.Info().Msgf("Echo server started on %s", opts.Listen)
			logging.Info().Msg("Press Ctrl+C to stop")

			// Wait for interrupt signal
			sigCh := make(chan os.Signal, 1)
			signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

			<-sigCh
			logging.Info().Msg("Shutting down...")

			return nil
		},
	}

	cmd.Flags().StringVar(&opts.Listen, "listen", opts.Listen, "listen address")
	cmd.Flags().BoolVarP(&opts.Verbose, "verbose", "v", false, "enable verbose logging")

	return cmd
}
