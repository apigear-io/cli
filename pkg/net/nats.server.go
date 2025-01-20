package net

import (
	"fmt"
	"net/url"
	"time"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

// Create an embedded NATS server

const (
	NatsTimeout = 5 * time.Second
)

type NatsServer struct {
	ns   *server.Server
	nc   *nats.Conn
	nipc *nats.Conn
}

func NewNatsServer(opts *Options) (*NatsServer, error) {
	if opts.NatsHost == "" {
		opts.NatsHost = "localhost"
	}
	if opts.NatsPort == 0 {
		opts.NatsPort = 4222
	}
	sopts := &server.Options{
		ServerName:      "apigear_server ",
		Host:            opts.NatsHost,
		Port:            opts.NatsPort,
		DontListen:      opts.NatsInprocessOnly,
		JetStream:       true,
		JetStreamDomain: "apigear",
		StoreDir:        cfg.ConfigDir() + "/nats",
	}
	if opts.NatsLeafURL != "" {
		leafURL, err := url.Parse(opts.NatsLeafURL)
		if err != nil {
			return nil, err
		}
		sopts.LeafNode = server.LeafNodeOpts{
			Remotes: []*server.RemoteLeafOpts{
				{
					URLs:        []*url.URL{leafURL},
					Credentials: opts.NatsCredentials,
				},
			},
		}
	}
	server, err := server.NewServer(sopts)
	if err != nil {
		log.Error().Err(err).Msg("failed to create nats server")
		return nil, err
	}
	if opts.Logging {
		server.ConfigureLogger()
	}
	return &NatsServer{ns: server}, nil
}

func (ns *NatsServer) Start() error {
	go ns.ns.Start()
	if !ns.ns.ReadyForConnections(NatsTimeout) {
		return fmt.Errorf("nats server not ready")
	}
	log.Info().Msgf("nats server started at %s", ns.ns.ClientURL())
	return nil
}

func (ns *NatsServer) Shutdown() error {
	ns.ns.WaitForShutdown()
	return nil
}

func (ns *NatsServer) ClientURL() string {
	return ns.ns.ClientURL()
}

func (ns *NatsServer) Connection() (*nats.Conn, error) {
	if ns.nc != nil {
		return ns.nc, nil
	}
	nc, err := nats.Connect(ns.ns.ClientURL())
	if err != nil {
		return nil, err
	}
	ns.nc = nc
	return nc, nil
}

func (ns *NatsServer) InProcessConnection() (*nats.Conn, error) {
	if ns.nipc != nil {
		return ns.nipc, nil
	}
	opts := []nats.Option{}
	opts = append(opts, nats.InProcessServer(ns.ns))
	nc, err := nats.Connect(ns.ns.ClientURL(), opts...)
	if err != nil {
		return nil, err
	}
	ns.nc = nc
	return ns.nc, nil
}
