package net

import (
	"context"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/sim"
)

func RunSimuServer(ctx context.Context, simu *sim.Simulation, addr string) *HTTPServer {
	hub := NewSimuWSServer(ctx, simu)
	s := NewHTTPServer()
	s.Router().HandleFunc("/ws", hub.ServeHTTP)
	log.Info().Msgf("simulation server listens on ws://%s/ws", addr)
	return s
}
