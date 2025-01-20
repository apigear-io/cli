package net

import (
	"context"

	"github.com/apigear-io/cli/pkg/log"
)

func RunSimuServer(ctx context.Context, provider SimulationProviderFunc, addr string) *HTTPServer {
	hub := NewSimuWSServer(ctx, provider)
	s := NewHTTPServer()
	s.Router().HandleFunc("/ws", hub.ServeHTTP)
	log.Info().Msgf("simulation server listens on ws://%s/ws", addr)
	return s
}
