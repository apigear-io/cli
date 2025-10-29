package natsutil

import (
	"errors"
	"os"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog/log"
)

type EmbeddedServer struct {
	srv      *server.Server
	nc       *nats.Conn
	storeDir string
	js       jetstream.JetStream
}

func NewEmbeddedServer() (*EmbeddedServer, error) {
	tmpDir, err := os.MkdirTemp("", "nats-server-")
	if err != nil {
		return nil, err
	}
	srv, err := server.NewServer(&server.Options{
		JetStream:  true,
		DontListen: true,
		StoreDir:   tmpDir,
	})
	if err != nil {
		return nil, err
	}
	srv.Start()
	if !srv.ReadyForConnections(5 * time.Second) {
		srv.Shutdown()
		return nil, errors.New("nats server not ready in time")
	}
	nc, err := nats.Connect(srv.ClientURL(), nats.InProcessServer(srv))
	if err != nil {
		srv.Shutdown()
		return nil, err
	}
	return &EmbeddedServer{
		srv:      srv,
		nc:       nc,
		storeDir: tmpDir,
	}, nil
}

func (e *EmbeddedServer) Close() {
	if e.nc != nil && !e.nc.IsClosed() {
		e.nc.Close()
	}
	if err := os.RemoveAll(e.storeDir); err != nil {
		log.Warn().Err(err).Str("dir", e.storeDir).Msg("failed to remove embedded store directory")
	}
	e.srv.Shutdown()
}

func (e *EmbeddedServer) NatsConn() *nats.Conn {
	return e.nc
}

func (e *EmbeddedServer) Server() *server.Server {
	return e.srv
}

func (e *EmbeddedServer) StoreDir() string {
	return e.storeDir
}

func (e *EmbeddedServer) JetStream() (jetstream.JetStream, error) {
	if e.js == nil {
		js, err := jetstream.New(e.nc)
		if err != nil {
			return nil, err
		}
		e.js = js
	}
	return e.js, nil
}

func (e *EmbeddedServer) ClientURL() string {
	return e.srv.ClientURL()
}
