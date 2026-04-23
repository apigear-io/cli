package natsutil

import (
	"errors"
	"os"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

// ServerConfig controls how a test or embedded NATS server should be started.
type ServerConfig struct {
	Options  *server.Options
	Embedded bool
	TempDir  string
}

// ServerHandle wraps a running NATS server instance.
type ServerHandle struct {
	srv      *server.Server
	storeDir string
	embedded bool
}

// StartServer boots a NATS server suitable for tests or local tooling.
func StartServer(cfg ServerConfig) (*ServerHandle, error) {
	opts := cloneOptions(cfg.Options)
	if opts.StoreDir == "" {
		dir := cfg.TempDir
		if dir == "" {
			tmp, err := os.MkdirTemp("", "streams-nats-")
			if err != nil {
				return nil, err
			}
			dir = tmp
		}
		opts.StoreDir = dir
	}
	if !opts.JetStream {
		opts.JetStream = true
	}
	if opts.Host == "" {
		opts.Host = "127.0.0.1"
	}
	if cfg.Embedded {
		// When running embedded, callers should connect via nats.InProcessServer.
		opts.Port = 0
	} else if opts.Port == 0 {
		opts.Port = -1 // auto-select free port
	}
	opts.NoSigs = true
	opts.NoLog = true
	log.Debug().Str("host", opts.Host).Int("port", opts.Port).Bool("embedded", cfg.Embedded).Str("store", opts.StoreDir).Msg("starting embedded NATS server")

	srv, err := server.NewServer(opts)
	if err != nil {
		return nil, err
	}
	srv.Start()
	if !srv.ReadyForConnections(5 * time.Second) {
		srv.Shutdown()
		return nil, errors.New("nats server not ready in time")
	}
	log.Debug().Str("url", srv.ClientURL()).Bool("embedded", cfg.Embedded).Msg("embedded NATS server ready")
	return &ServerHandle{srv: srv, storeDir: opts.StoreDir, embedded: cfg.Embedded}, nil
}

// Shutdown stops the server and cleans up temporary resources.
func (h *ServerHandle) Shutdown() {
	if h == nil {
		return
	}
	if h.srv != nil {
		log.Debug().Str("url", h.srv.ClientURL()).Msg("shutting down embedded NATS server")
		h.srv.Shutdown()
	}
	if h.storeDir != "" {
		_ = os.RemoveAll(h.storeDir)
	}
}

// ClientURL returns the URL clients can use to connect (only valid when not embedded).
func (h *ServerHandle) ClientURL() string {
	if h == nil || h.srv == nil {
		return ""
	}
	return h.srv.ClientURL()
}

func (h *ServerHandle) NatsConn() (*nats.Conn, error) {
	if h == nil || h.srv == nil {
		return nil, errors.New("server not running")
	}
	return nats.Connect(h.srv.ClientURL())
}

// InProcessOption returns a connection option for embedded servers.
func (h *ServerHandle) InProcessOption() nats.Option {
	if h == nil || h.srv == nil {
		log.Debug().Msg("in-process option requested without running server")
		return nil
	}
	log.Debug().Str("url", h.srv.ClientURL()).Msg("providing in-process NATS option")
	inner := nats.InProcessServer(h.srv)
	return func(o *nats.Options) error {
		log.Debug().Msg("applying in-process NATS option")
		return inner(o)
	}
}

func cloneOptions(opts *server.Options) *server.Options {
	if opts == nil {
		return &server.Options{}
	}
	out := *opts
	return &out
}

func WithServer(cfg ServerConfig, fn func(h *ServerHandle) error) error {
	srv, err := StartServer(cfg)
	if err != nil {
		return err
	}
	defer srv.Shutdown()
	return fn(srv)
}
