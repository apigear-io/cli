package helper

import (
	"os"
	"os/signal"
	"syscall"
)

// WaitForInterrupt handles SIGINT and SIGTERM signals.
// It cancels the context when a signal is received.
func WaitForInterrupt(cancel func()) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	cancel()
}
