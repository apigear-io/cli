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
	NatsHost          string `json:"nats_host"`
	NatsPort          int    `json:"nats_port"`
	NatsDisabled      bool   `json:"nats_disabled"`
	NatsListen        bool   `json:"nats_inprocess_only"`
	NatsLeafURL       string `json:"nats_leaf_url"`
	NatsCredentials   string `json:"nats_credentials"`
	HttpAddr          string `json:"http_addr"`
	HttpDisabled      bool   `json:"http_disabled"`
	MonitorDisabled   bool   `json:"monitor_disabled"`
	ObjectAPIDisabled bool   `json:"object_api_disabled"`
	Logging           bool   `json:"logging"`
}

var DefaultOptions = &Options{
	NatsHost:          "localhost",
	NatsPort:          4222,
	NatsDisabled:      false,
	NatsListen:        false,
	NatsLeafURL:       "",
	NatsCredentials:   "",
	HttpAddr:          "localhost:5555",
	HttpDisabled:      false,
	MonitorDisabled:   false,
	ObjectAPIDisabled: false,
	Logging:           false,
}

type NetworkManager struct {
	opts       *Options
	natsServer *NatsServer
	httpServer *HTTPServer
	nc         *nats.Conn
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
	if !s.opts.NatsDisabled {
		err := s.StartNATS(&NatsServerOptions{
			Host:        s.opts.NatsHost,
			Port:        s.opts.NatsPort,
			NatsListen:  s.opts.NatsListen,
			LeafURL:     s.opts.NatsLeafURL,
			Credentials: s.opts.NatsCredentials,
		})
		if err != nil {
			log.Error().Err(err).Msg("failed to start nats server")
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
	err = s.StopNATS()
	if err != nil {
		return err
	}
	return nil
}

func (s *NetworkManager) StartNATS(opts *NatsServerOptions) error {
	if s.natsServer != nil {
		return fmt.Errorf("nats server already started")
	}
	server, err := NewNatsServer(opts)
	if err != nil {
		return err
	}
	s.natsServer = server
	return s.natsServer.Start()
}

func (s *NetworkManager) StopNATS() error {
	log.Info().Msg("stop nats server")
	if s.nc != nil {
		err := s.nc.Drain()
		if err != nil {
			return err
		}
	}
	if s.natsServer != nil {
		return s.natsServer.Shutdown()
	}
	return nil
}

func (s *NetworkManager) NatsClientURL() string {
	if s.natsServer != nil {
		return s.natsServer.ClientURL()
	}
	return ""
}

func (s *NetworkManager) NatsConnection() (*nats.Conn, error) {
	if s.natsServer == nil {
		return nil, fmt.Errorf("nats server not started")
	}
	return s.natsServer.Connection()
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
	nc, err := s.NatsConnection()
	if err != nil {
		log.Error().Msgf("nats connection: %v", err)
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
