package net

import (
	"context"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/sim/model"
)

func RunSimuServer(ctx context.Context, provider model.SimulationProvider, addr string) *HTTPServer {
	hub := NewSimuWSServer(ctx, provider, "demo")
	s := NewHTTPServer()
	s.Router().HandleFunc("/ws", hub.ServeHTTP)
	log.Info().Msgf("simulation server listens on ws://%s/ws", addr)
	return s
}
