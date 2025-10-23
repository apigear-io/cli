package net

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/mon"
	"github.com/nats-io/nats.go"
)

type Options struct {
	NatsServerURL string         `json:"nats_server_url"`
	HttpAddr      string         `json:"http_addr"`
	Logging       bool           `json:"logging"`
	WSProxy       *WSProxyConfig `json:"ws_proxy,omitempty"`
}

type WSProxyConfig struct {
	Enabled           bool          `json:"enabled"`
	BasePath          string        `json:"base_path"`
	Routes            []RouteConfig `json:"routes"`
	ReconnectAttempts int           `json:"reconnect_attempts"`
	ReconnectBackoff  time.Duration `json:"reconnect_backoff"`
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
	if o.WSProxy != nil {
		if o.WSProxy.BasePath == "" {
			o.WSProxy.BasePath = "/ws"
		}
		if o.WSProxy.ReconnectAttempts <= 0 {
			o.WSProxy.ReconnectAttempts = 3
		}
		if o.WSProxy.ReconnectBackoff <= 0 {
			o.WSProxy.ReconnectBackoff = 500 * time.Millisecond
		}
	}
	return nil
}

type NetworkManager struct {
	opts       Options
	httpServer *HTTPServer
	nc         *nats.Conn
	wsProxy    *WSProxy
	olnkServer *OlinkServer
	olnkRelay  *ReplayOlinkRelay
}

func NewManager(opts Options) *NetworkManager {
	log.Debug().Msg("net.NewManager")
	opts.Validate()
	return &NetworkManager{
		opts:       opts,
		olnkServer: NewOlinkServer(),
	}
}

func (m *NetworkManager) NatsConnection() (*nats.Conn, error) {
	if m.nc != nil && !m.nc.IsClosed() {
		return m.nc, nil
	}
	if m.opts.NatsServerURL == "" {
		return nil, fmt.Errorf("nats server URL not set")
	}
	nc, err := nats.Connect(m.opts.NatsServerURL)
	if err != nil {
		return nil, err
	}
	m.nc = nc
	return m.nc, nil
}

func (m *NetworkManager) Start(ctx context.Context) error {
	log.Debug().Msg("start network manager")
	err := m.StartHTTP(m.opts.HttpAddr)
	if err != nil {
		log.Error().Err(err).Msg("failed to start http server")
		return err
	}
	err = m.EnableMonitor()
	if err != nil {
		log.Error().Err(err).Msg("failed to enable monitor")
		return err
	}
	if err := m.EnableWSProxy(); err != nil {
		log.Error().Err(err).Msg("failed to enable ws proxy")
		return err
	}
	err = m.enableOlinkServer()
	if err != nil {
		log.Error().Err(err).Msg("failed to enable olink server")
		return err
	}
	err = m.enableReplayRelay()
	if err != nil {
		log.Error().Err(err).Msg("failed to enable replay relay")
		return err
	}
	return nil
}

func (m *NetworkManager) Wait(ctx context.Context) error {
	log.Info().Msg("services running...")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		err := m.Stop()
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

func (m *NetworkManager) Stop() error {
	log.Info().Msg("stop network manager")
	err := m.StopHTTP()
	if err != nil {
		return err
	}
	if m.olnkRelay != nil {
		log.Info().Msg("stop olink replay relay")
		err = m.olnkRelay.Stop()
		if err != nil {
			log.Error().Err(err).Msg("failed to stop olink replay relay")
		}
	}
	return nil
}

func (m *NetworkManager) StartHTTP(addr string) error {
	if m.httpServer != nil {
		log.Info().Msg("stop running http server")
		m.httpServer.Stop()
	}
	log.Info().Msg("start http server")
	m.httpServer = NewHTTPServer(&HttpServerOptions{Addr: addr})
	err := m.httpServer.Start()
	if err != nil {
		log.Error().Err(err).Msg("failed to start http server")
	}
	log.Info().Msgf("http server started at http://%s", addr)
	return err
}

func (m *NetworkManager) StopHTTP() error {
	log.Info().Msg("stop http server")
	if m.httpServer != nil {
		m.httpServer.Stop()
	}
	return nil
}

func (m *NetworkManager) HttpServer() *HTTPServer {
	return m.httpServer
}

func (m *NetworkManager) EnableMonitor() error {
	log.Info().Msg("enable monitor endpoint")
	if m.httpServer == nil {
		log.Error().Msg("http server not started")
		return fmt.Errorf("http server not started")
	}
	nc, err := m.NatsConnection()
	if err != nil {
		log.Error().Err(err).Msg("nats connection")
		return err
	}
	m.httpServer.Router().HandleFunc("/monitor/{source}", MonitorRequestHandler(nc))
	log.Info().Msgf("start http monitor endpoint on http://%s/monitor/{source}", m.httpServer.Address())
	return nil
}

func (m *NetworkManager) EnableWSProxy() error {
	cfg := m.opts.WSProxy
	if cfg == nil || !cfg.Enabled {
		return nil
	}
	if m.httpServer == nil {
		return fmt.Errorf("http server not started")
	}

	opts := ProxyOptions{
		BasePath:          cfg.BasePath,
		Routes:            cfg.Routes,
		ReconnectAttempts: cfg.ReconnectAttempts,
		ReconnectBackoff:  cfg.ReconnectBackoff,
		OnConnect: func(ctx context.Context, info *ConnectionInfo) error {
			log.Info().Str("target", info.TargetURL).Str("path", info.Route.Path).Msg("ws proxy connection accepted")
			return nil
		},
		OnDisconnect: func(ctx context.Context, info *ConnectionInfo, err error) {
			event := log.Info()
			if err != nil && !errors.Is(err, context.Canceled) {
				event = log.Warn().Err(err)
			}
			event.Str("target", info.TargetURL).Str("path", info.Route.Path).Msg("ws proxy connection closed")
		},
	}

	proxy, err := NewWSProxy(opts)
	if err != nil {
		return fmt.Errorf("ws proxy init: %w", err)
	}
	m.wsProxy = proxy

	m.httpServer.Router().Mount("/", proxy)
	log.Info().Msgf("ws proxy enabled at %s", cfg.BasePath)
	return nil
}

func (m *NetworkManager) GetMonitorAddress() (string, error) {
	log.Info().Msg("get monitor address")
	if m.httpServer == nil {
		return "", fmt.Errorf("http server not started")
	}
	return fmt.Sprintf("http://%s/monitor/${source}", m.httpServer.Address()), nil
}

func (m *NetworkManager) GetSimulationAddress() (string, error) {
	log.Info().Msg("get simulation address")
	if m.httpServer == nil {
		return "", fmt.Errorf("http server not started")
	}
	return fmt.Sprintf("ws://%s/ws", m.httpServer.Address()), nil
}

// MonitorEmitter return the monitor event emitter.
func (m *NetworkManager) MonitorEmitter() *helper.Hook[mon.Event] {
	return &mon.Emitter
}

func (m *NetworkManager) OnMonitorEvent(fn func(event *mon.Event)) {
	nc, err := m.NatsConnection()
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

func (m *NetworkManager) enableOlinkServer() error {
	if m.httpServer == nil {
		return fmt.Errorf("http server not started")
	}
	addr := m.HttpServer().Address()
	log.Info().Msgf("starting Olink server at ws://%s/ws", addr)
	m.HttpServer().Router().Handle("/ws", m.olnkServer)
	return nil
}

func (m *NetworkManager) OlinkServer() *OlinkServer {
	return m.olnkServer
}

func (m *NetworkManager) enableReplayRelay() error {
	log.Info().Msg("enable olink replay relay")
	nc, err := m.NatsConnection()
	if err != nil {
		log.Error().Err(err).Msg("failed to get nats connection for replay relay")
		return err
	}
	m.olnkRelay = NewReplayOlinkRelay(nc, "replay.olink", m.OlinkServer())
	return nil
}
