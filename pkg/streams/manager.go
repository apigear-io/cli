package streams

import (
	"context"
	"errors"
	"path"
	"time"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/streams/buffer"
	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/controller"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

const (
	NatsTimeout = 10 * time.Second
)

type ManagerOptions struct {
	NatsPort int
	AppDir   string
	Logging  bool
}

type Manager struct {
	js         jetstream.JetStream
	srv        *server.Server
	nc         *nats.Conn
	opts       ManagerOptions
	controller *controller.Controller
}

func NewManager(opts ManagerOptions) *Manager {
	return &Manager{
		opts: opts,
	}
}

func (m *Manager) Start(ctx context.Context) error {
	log.Info().Msg("starting streams manager")
	err := m.runServer()
	if err != nil {
		return err
	}
	err = m.runServices(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) runServer() error {
	if m.srv != nil && m.srv.ReadyForConnections(0) {
		return nil
	}
	m.srv = server.New(&server.Options{
		Port:      m.opts.NatsPort,
		JetStream: true,
		StoreDir:  path.Join(m.opts.AppDir, "nats"),
	})
	if m.opts.Logging {
		m.srv.ConfigureLogger()
	}
	m.srv.Start()
	// wait for server to be ready
	if !m.srv.ReadyForConnections(NatsTimeout) {
		m.srv.Shutdown()
		return errors.New("nats server not ready in time")
	}
	log.Info().Msgf("NATS server started at %s", m.srv.ClientURL())
	// connect to server
	nc, err := nats.Connect(m.srv.ClientURL(), nats.InProcessServer(m.srv))
	if err != nil {
		m.srv.Shutdown()
		return err
	}
	m.nc = nc
	// create jetstream context
	js, err := jetstream.New(m.nc)
	if err != nil {
		m.nc.Close()
		m.srv.Shutdown()
		return err
	}
	m.js = js
	log.Info().Msgf("NATS server running at %s", js.Conn().ConnectedUrl())
	return nil
}

func (m *Manager) runServices(ctx context.Context) error {
	// Create and start controller
	ctrl, err := controller.NewController(m.js, controller.Options{
		ServerURL:        m.js.Conn().ConnectedAddr(),
		RecordRpcSubject: config.RecordRpcSubject,
		StateBucket:      config.StateBucket,
	})
	if err != nil {
		return err
	}
	if err := ctrl.Start(); err != nil {
		return err
	}
	m.controller = ctrl

	// Start buffer service in background
	go func() {
		err := buffer.RunBuffer(ctx, m.js, buffer.BufferOptions{
			DeviceBucket:    config.DeviceBucket,
			MonitorSubject:  config.MonitorSubject,
			RefreshInterval: config.BufferRefresh,
		})
		if err != nil && !errors.Is(err, context.Canceled) {
			log.Error().Err(err).Msg("buffer service error")
		}
	}()

	return nil
}

func (m *Manager) ClientURL() string {
	if m.srv != nil {
		return m.srv.ClientURL()
	}
	return nats.DefaultURL
}

func (m *Manager) JetStream() jetstream.JetStream {
	return m.js
}

func (m *Manager) Shutdown() error {
	if m.controller != nil {
		m.controller.Close()
	}
	if m.nc != nil && !m.nc.IsClosed() {
		m.nc.Close()
	}
	if m.srv != nil {
		m.srv.Shutdown()
	}
	return nil
}

func (m *Manager) Connection() (*nats.Conn, error) {
	if m.nc == nil || m.nc.IsClosed() {
		return nil, errors.New("nats server not started")
	}
	if m.srv == nil || !m.srv.ReadyForConnections(0) {
		return nil, errors.New("nats server not ready")
	}
	return m.nc, nil
}
