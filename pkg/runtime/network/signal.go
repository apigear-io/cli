package network

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/apigear-io/cli/pkg/foundation/logging"
)

// WaitForShutdown blocks until shutdown signal (SIGINT/SIGTERM) or context cancellation.
// Executes onShutdown callback before returning.
func WaitForShutdown(ctx context.Context, onShutdown func()) error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sig)

	select {
	case <-ctx.Done():
		logging.Info().Msg("context cancelled, shutting down...")
		onShutdown()
		return ctx.Err()
	case <-sig:
		logging.Info().Msg("shutdown signal received...")
		onShutdown()
		logging.Info().Msg("shutdown complete")
		return nil
	}
}
