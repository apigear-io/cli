package helper

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func Wait(ctx context.Context, cleanup func()) error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	if cleanup != nil {
		defer cleanup()
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-sig:
		return nil
	}
}
