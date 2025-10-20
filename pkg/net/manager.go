package net

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/mon"
	"github.com/nats-io/nats.go"
)

type Options struct {
	NatsServerURL string `json:"nats_server_url"`
	HttpAddr      string `json:"http_addr"`
	Logging       bool   `json:"logging"`
}

func (o *Options) Validate() error {
	if o.NatsServerURL == "" {
		o.NatsServerURL = nats.DefaultURL
		log.Info().Msgf("nats server URL not set, using default: %s", o.NatsServerURL)
	}
	if o.HttpAddr == "" {
		o.HttpAddr = "127.0.0.1:5555"
		log.Info().Msgf("http address not set, using default: %s", o.HttpAddr)
	}
	return nil
}

type NetworkManager struct {
	opts       Options
	httpServer *HTTPServer
	nc         *nats.Conn
}

func NewManager() *NetworkManager {
	log.Debug().Msg("net.NewManager")
	return &NetworkManager{}
}

func (s *NetworkManager) NatsConnection() (*nats.Conn, error) {
	if s.nc != nil && !s.nc.IsClosed() {
		return s.nc, nil
	}
	if s.opts.NatsServerURL == "" {
		return nil, fmt.Errorf("nats server URL not set")
	}
	nc, err := nats.Connect(s.opts.NatsServerURL)
	if err != nil {
		return nil, err
	}
	s.nc = nc
	return s.nc, nil
}

func (s *NetworkManager) Start(opts Options) error {
	err := opts.Validate()
	if err != nil {
		return err
	}
	s.opts = opts
	log.Debug().Msg("start network manager")
	err = s.StartHTTP(s.opts.HttpAddr)
	if err != nil {
		log.Error().Err(err).Msg("failed to start http server")
		return err
	}
	err = s.EnableMonitor()
	if err != nil {
		log.Error().Err(err).Msg("failed to enable monitor")
		return err
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
	log.Info().Msg("enable monitor endpoint")
	if s.httpServer == nil {
		log.Error().Msg("http server not started")
		return fmt.Errorf("http server not started")
	}
	nc, err := s.NatsConnection()
	if err != nil {
		log.Error().Err(err).Msg("nats connection")
		return err
	}
	s.httpServer.Router().HandleFunc("/monitor/{source}", MonitorRequestHandler(nc))
	log.Info().Msgf("start http monitor endpoint on http://%s/monitor/{source}", s.httpServer.Address())
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

func (s *NetworkManager) OnMonitorEvent(fn func(event *mon.Event)) {
	nc, err := s.NatsConnection()
	if err != nil {
		log.Error().Msgf("nats connection: %v", err)
		return
	}
	log.Debug().Msg("subscribe to monitor events")
	_, err = nc.Subscribe(mon.MonitorSubject+".>", func(msg *nats.Msg) {
		var event mon.Event
		err := json.Unmarshal(msg.Data, &event)
		if err != nil {
			log.Error().Msgf("unmarshal event: %v", err)
			return
		}
		fn(&event)
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to subscribe to monitor events")
	}
}
