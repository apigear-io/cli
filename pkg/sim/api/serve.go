package api

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

type ServeOptions struct {
	NatsUrl string
	manager SimulationManager
}

func Serve(opts ServeOptions) error {
	_, err := NewService(opts.NatsUrl, opts.manager)
	if err != nil {
		return err
	}

	log.Info().Str("nats", opts.NatsUrl).Msg("service registered with nats")

	// wait for SIGINT or SIGTERM
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	log.Info().Msg("shutting down")
	return nil
}
