package net

import (
	"fmt"
	"time"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/nats-io/nats-server/v2/server"
)

// Create an embedded NATS server

type NatsServerOptions struct {
	Port int
}

type NatsServer struct {
	server *server.Server
}

func NewNatsServer(sopts NatsServerOptions) (*NatsServer, error) {
	opts := server.Options{
		Port: sopts.Port,
	}
	server, err := server.NewServer(&opts)
	if err != nil {
		log.Error().Err(err).Msg("failed to create nats server")
		return nil, err
	}
	return &NatsServer{server: server}, nil
}

func (ns *NatsServer) Start() error {
	go ns.server.Start()
	if !ns.server.ReadyForConnections(4 * time.Second) {
		return fmt.Errorf("nats server not ready")
	}
	return nil
}

func (ns *NatsServer) Stop() error {
	ns.server.WaitForShutdown()
	return nil
}

func (ns *NatsServer) ClientURL() string {
	return ns.server.ClientURL()
}
