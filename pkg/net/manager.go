package net

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/mon"
	"github.com/apigear-io/objectlink-core-go/olink/ws"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

type Options struct {
	NatsHost               string `json:"nats_host"`
	NatsPort               int    `json:"nats_port"`
	NatsDisabled           bool   `json:"nats_disabled"`
	NatsInprocessOnly      bool   `json:"nats_inprocess_only"`
	NatsLeafURL            string `json:"nats_leaf_url"`
	NatsCredentials        string `json:"nats_credentials"`
	HttpAddr               string `json:"http_addr"`
	HttpDisabled           bool   `json:"http_disabled"`
	MonitorDisabled        bool   `json:"monitor_disabled"`
	ObjectAPIDisabled      bool   `json:"object_api_disabled"`
	SimulationProviderFunc SimulationProviderFunc
	Logging                bool `json:"logging"`
}

type NetworkManager struct {
	opts       *Options
	natsServer *NatsServer
	httpServer *HTTPServer
	wsHUB      *ws.Hub
	nc         *nats.Conn
}

var (
	defaultManager *NetworkManager
)

func GetManager() *NetworkManager {
	if defaultManager == nil {
		defaultManager = NewManager()
	}
	return defaultManager
}
func NewManager() *NetworkManager {
	log.Debug().Msg("create network manager")
	return &NetworkManager{}
}

func (s *NetworkManager) Start(opts *Options) error {
	log.Debug().Msg("start network manager")
	s.opts = opts
	if !s.opts.HttpDisabled {
		err := s.StartHTTP()
		if err != nil {
			return err
		}
	}
	if !s.opts.NatsDisabled {
		err := s.StartNATS()
		if err != nil {
			return err
		}
	}
	if !s.opts.MonitorDisabled {
		s.EnableMonitor()
	}
	if !s.opts.ObjectAPIDisabled {
		if s.opts.SimulationProviderFunc == nil {
			return fmt.Errorf("object api is enabled, but simulation is not set")
		}
		s.EnableSimulation(s.opts.SimulationProviderFunc)
	}
	return nil
}

func (s *NetworkManager) Wait(ctx context.Context) error {
	log.Info().Msg("running...")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	defer s.Stop()
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

func (s *NetworkManager) StartNATS() error {
	if s.natsServer != nil {
		return fmt.Errorf("nats server already started")
	}
	server, err := NewNatsServer(s.opts)
	if err != nil {
		return err
	}
	s.natsServer = server
	return s.natsServer.Start()
}

func (s *NetworkManager) StopNATS() error {
	log.Info().Msg("stop nats server")
	if s.nc != nil {
		s.nc.Drain()
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
	if s.nc == nil {
		nc, err := nats.Connect(s.NatsClientURL())
		if err != nil {
			return nil, err
		}
		s.nc = nc
	}
	return s.nc, nil
}

func (s *NetworkManager) StartHTTP() error {
	if s.httpServer != nil {
		log.Info().Msg("stop running http server")
		s.httpServer.Stop()
	}
	log.Info().Msgf("start http server at http://%s", s.opts.HttpAddr)
	server := NewHTTPServer()
	s.httpServer = server
	return s.httpServer.Start(s.opts)
}

func (s *NetworkManager) StopHTTP() error {
	log.Info().Msg("stop http server")
	if s.httpServer != nil {
		s.httpServer.Stop()
	}
	return nil
}

func (s *NetworkManager) EnableMonitor() {
	if s.httpServer == nil {
		log.Error().Msg("http server not started")
		return
	}
	nc, err := s.NatsConnection()
	if err != nil {
		log.Error().Msgf("nats connection: %v", err)
		return
	}
	s.httpServer.Router().HandleFunc("/monitor/{source}", MonitorRequestHandler(nc))
	log.Info().Msgf("handle monitor request on http://%s/monitor/{source}", s.httpServer.Address())
}

func (s *NetworkManager) EnableSimulation(provider SimulationProviderFunc) {
	if s.httpServer == nil {
		log.Error().Msg("http server not started")
		return
	}
	if s.wsHUB != nil {
		log.Info().Msg("simulation ws hub already enabled")
		return
	}
	ctx := context.Background()
	s.wsHUB = NewSimuWSServer(ctx, provider)
	s.httpServer.Router().HandleFunc("/ws", s.wsHUB.ServeHTTP)
	log.Info().Msg("simulation server listen on ws://localhost:8080/ws")
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
func (s *NetworkManager) MonitorEmitter() *helper.Hook[*mon.Event] {
	return &mon.Emitter
}

func (s *NetworkManager) OnMonitorEvent(fn func(event *mon.Event)) {
	nc, err := s.NatsConnection()
	if err != nil {
		log.Error().Msgf("nats connection: %v", err)
		return
	}
	log.Info().Msg("subscribe to monitor events")
	nc.Subscribe(mon.MonitorSubject+".>", func(msg *nats.Msg) {
		var event mon.Event
		err := json.Unmarshal(msg.Data, &event)
		if err != nil {
			log.Error().Msgf("unmarshal event: %v", err)
			return
		}
		fn(&event)
	})
}
