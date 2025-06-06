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
	NatsTimeout = 30 * time.Second
)

type NatsServerOptions struct {
	Host        string
	Port        int
	NatsListen  bool
	LeafURL     string
	Credentials string
	Logging     bool
}

type NatsServer struct {
	opts *NatsServerOptions
	ns   *server.Server
	nc   *nats.Conn
}

func NewNatsServer(opts *NatsServerOptions) (*NatsServer, error) {
	if opts.Host == "" {
		opts.Host = "localhost"
	}
	if opts.Port == 0 {
		opts.Port = 4222
	}
	sopts := &server.Options{
		ServerName:      "apigear_server",
		Host:            opts.Host,
		Port:            opts.Port,
		DontListen:      !opts.NatsListen,
		JetStream:       true,
		JetStreamDomain: "apigear",
		StoreDir:        cfg.ConfigDir() + "/nats",
	}
	if opts.LeafURL != "" {
		leafURL, err := url.Parse(opts.LeafURL)
		if err != nil {
			return nil, err
		}
		sopts.LeafNode = server.LeafNodeOpts{
			Remotes: []*server.RemoteLeafOpts{
				{
					URLs:        []*url.URL{leafURL},
					Credentials: opts.Credentials,
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

	return &NatsServer{opts: opts, ns: server}, nil
}

func (ns *NatsServer) Start() error {
	log.Info().Msg("start nats server")
	ns.ns.Start()
	log.Info().Msg("wait for nats server to be ready")
	if !ns.ns.ReadyForConnections(NatsTimeout) {
		return fmt.Errorf("nats server not ready")
	}
	log.Info().Msgf("start nats server listen at %s", ns.ns.ClientURL())
	return nil
}

func (ns *NatsServer) Shutdown() error {
	ns.ns.Shutdown()
	return nil
}

func (ns *NatsServer) ClientURL() string {
	return ns.ns.ClientURL()
}

func (ns *NatsServer) Connection() (*nats.Conn, error) {
	if ns.nc == nil {
		copts := []nats.Option{}
		if ns.opts.NatsListen {
			copts = append(copts, nats.InProcessServer(ns.ns))
		}
		nc, err := nats.Connect(ns.ns.ClientURL(), copts...)
		if err != nil {
			return nil, err
		}
		ns.nc = nc
	}
	return ns.nc, nil
}
