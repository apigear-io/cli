package server

import (
	"context"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/net"
	"github.com/apigear-io/cli/pkg/sim"
	"github.com/apigear-io/cli/pkg/streams"
)

type Options struct {
	NatsHost string
	NatsPort int
	HttpAddr string
	Logging  bool
}

type Server struct {
	opts   Options
	strman *streams.Manager
	siman  *sim.Manager
	netman *net.NetworkManager
}

func New(opts Options) *Server {
	return &Server{
		opts: opts,
		strman: streams.NewManager(streams.ManagerOptions{
			NatsPort: opts.NatsPort,
			AppDir:   cfg.ConfigDir(),
			Logging:  opts.Logging,
		}),
		siman: sim.NewManager(sim.ManagerOptions{}),
		netman: net.NewManager(net.Options{
			NatsServerURL: opts.NatsHost,
			HttpAddr:      opts.HttpAddr,
			Logging:       opts.Logging,
		}),
	}
}

func (s *Server) Start(ctx context.Context) error {
	// start http server
	// start nats server
	// start stream server
	log.Info().Msg("starting stream manager")
	err := s.strman.Start(ctx)
	if err != nil {
		return err
	}
	// network services
	log.Info().Msg("starting network manager")
	err = s.netman.Start(ctx)
	if err != nil {
		return err
	}
	// simulation server
	log.Info().Msg("starting simulation manager")
	err = s.siman.Start(ctx, s.netman)
	if err != nil {
		return err
	}
	log.Info().Msg("server started")
	return nil
}

func (s *Server) Stop() error {
	return nil
}

func (s *Server) NetworkManager() *net.NetworkManager {
	return s.netman
}
func (s *Server) StreamManager() *streams.Manager {
	return s.strman
}
func (s *Server) SimulationManager() *sim.Manager {
	return s.siman
}
