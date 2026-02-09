package net

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/mon"
)

type Options struct {
	HttpAddr          string `json:"http_addr"`
	HttpDisabled      bool   `json:"http_disabled"`
	MonitorDisabled   bool   `json:"monitor_disabled"`
	ObjectAPIDisabled bool   `json:"object_api_disabled"`
	Logging           bool   `json:"logging"`
}

var DefaultOptions = &Options{
	HttpAddr:          "localhost:5555",
	HttpDisabled:      false,
	MonitorDisabled:   false,
	ObjectAPIDisabled: false,
	Logging:           false,
}

type NetworkManager struct {
	opts       *Options
	httpServer *HTTPServer
}

func NewManager() *NetworkManager {
	log.Debug().Msg("net.NewManager")
	return &NetworkManager{}
}

func (s *NetworkManager) Start(opts *Options) error {
	s.opts = opts
	log.Debug().Msg("start network manager")
	if !s.opts.HttpDisabled {
		err := s.StartHTTP(s.opts.HttpAddr)
		if err != nil {
			log.Error().Err(err).Msg("failed to start http server")
			return err
		}
	}
	if !s.opts.MonitorDisabled {
		err := s.EnableMonitor()
		if err != nil {
			log.Error().Err(err).Msg("failed to enable monitor")
			return err
		}
	}
	return nil
}

func (s *NetworkManager) Wait(ctx context.Context) error {
	log.Info().Msg("services running...")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		err := s.Stop()
		if err != nil {
			log.Error().Err(err).Msg("failed to stop services")
		}
		log.Info().Msg("services stopped")
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-sig:
		return nil
	}
}

func (s *NetworkManager) Stop() error {
	log.Info().Msg("stop network manager")
	err := s.StopHTTP()
	if err != nil {
		return err
	}
	return nil
}

func (s *NetworkManager) StartHTTP(addr string) error {
	if s.httpServer != nil {
		log.Info().Msg("stop running http server")
		s.httpServer.Stop()
	}
	log.Info().Msg("start http server")
	s.httpServer = NewHTTPServer(&HttpServerOptions{Addr: addr})
	err := s.httpServer.Start()
	if err != nil {
		log.Error().Err(err).Msg("failed to start http server")
	}
	log.Info().Msgf("http server started at http://%s", addr)
	return err
}

func (s *NetworkManager) StopHTTP() error {
	log.Info().Msg("stop http server")
	if s.httpServer != nil {
		s.httpServer.Stop()
	}
	return nil
}

func (s *NetworkManager) HttpServer() *HTTPServer {
	return s.httpServer
}

func (s *NetworkManager) EnableMonitor() error {
	if s.httpServer == nil {
		log.Error().Msg("http server not started")
		return fmt.Errorf("http server not started")
	}
	s.httpServer.Router().HandleFunc("/monitor/{source}", MonitorRequestHandler())
	log.Info().Msgf("start http monitor endpoint on http://%s/monitor/{source}", s.httpServer.Address())
	log.Warn().Msg("NATS disabled: monitor events will be logged locally but not broadcast")
	return nil
}

func (s *NetworkManager) GetMonitorAddress() (string, error) {
	log.Info().Msg("get monitor address")
	if s.httpServer == nil {
		return "", fmt.Errorf("http server not started")
	}
	return fmt.Sprintf("http://%s/monitor/${source}", s.httpServer.Address()), nil
}

func (s *NetworkManager) GetSimulationAddress() (string, error) {
	log.Info().Msg("get simulation address")
	if s.httpServer == nil {
		return "", fmt.Errorf("http server not started")
	}
	return fmt.Sprintf("ws://%s/ws", s.httpServer.Address()), nil
}

// MonitorEmitter return the monitor event emitter.
func (s *NetworkManager) MonitorEmitter() *helper.Hook[mon.Event] {
	return &mon.Emitter
}
