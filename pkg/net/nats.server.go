package net

// import (
// 	"fmt"
// 	"net/url"
// 	"time"

// 	"github.com/apigear-io/cli/pkg/cfg"
// 	"github.com/apigear-io/cli/pkg/log"
// 	"github.com/nats-io/nats-server/v2/server"
// 	"github.com/nats-io/nats.go"
// )

// // Create an embedded NATS server

// const (
// 	NatsTimeout = 30 * time.Second
// )

// type NatsServerOptions struct {
// 	Host        string
// 	Port        int
// 	DontListen  bool
// 	LeafURL     string
// 	Credentials string
// 	Logging     bool
// }

// type NatsServer struct {
// 	opts *NatsServerOptions
// 	srv  *server.Server
// 	nc   *nats.Conn
// }

// func NewNatsServer(opts *NatsServerOptions) (*NatsServer, error) {
// 	if opts.Host == "" {
// 		opts.Host = server.DEFAULT_HOST
// 	}
// 	if opts.Port == 0 {
// 		opts.Port = server.DEFAULT_PORT
// 	}
// 	sopts := &server.Options{
// 		ServerName: "apigear-nats",
// 		Host:       opts.Host,
// 		Port:       opts.Port,
// 		DontListen: opts.DontListen,
// 		JetStream:  true,
// 		StoreDir:   cfg.ConfigDir() + "/nats",
// 	}
// 	if opts.LeafURL != "" {
// 		leafURL, err := url.Parse(opts.LeafURL)
// 		if err != nil {
// 			return nil, err
// 		}
// 		sopts.LeafNode = server.LeafNodeOpts{
// 			Remotes: []*server.RemoteLeafOpts{
// 				{
// 					URLs:        []*url.URL{leafURL},
// 					Credentials: opts.Credentials,
// 				},
// 			},
// 		}
// 	}
// 	server, err := server.NewServer(sopts)
// 	if err != nil {
// 		log.Error().Err(err).Msg("failed to create nats server")
// 		return nil, err
// 	}
// 	if opts.Logging {
// 		server.ConfigureLogger()
// 	}

// 	return &NatsServer{opts: opts, srv: server}, nil
// }

// func (ns *NatsServer) Start() error {
// 	ns.srv.Start()
// 	log.Info().Msg("wait for nats server to be ready")
// 	if !ns.srv.ReadyForConnections(NatsTimeout) {
// 		ns.srv.Shutdown()
// 		return fmt.Errorf("nats server not ready")
// 	}
// 	log.Info().Msgf("nats server started: listen at %s", ns.srv.ClientURL())
// 	nc, err := nats.Connect(ns.srv.ClientURL())
// 	if err != nil {
// 		log.Error().Err(err).Msg("failed to create nats connection")
// 		return err
// 	}
// 	if nc.IsConnected() {
// 		log.Info().Msg("nats connection established")
// 	}
// 	ns.nc = nc
// 	return nil
// }

// func (ns *NatsServer) Shutdown() error {
// 	ns.srv.Shutdown()
// 	return nil
// }

// func (ns *NatsServer) ClientURL() string {
// 	return ns.srv.ClientURL()
// }

// func (ns *NatsServer) Connection() (*nats.Conn, error) {
// 	if ns.nc != nil && ns.nc.IsConnected() {
// 		return ns.nc, nil
// 	}
// 	copts := []nats.Option{}
// 	if ns.opts.DontListen {
// 		copts = append(copts, nats.InProcessServer(ns.srv))
// 	}
// 	nc, err := nats.Connect(ns.srv.ClientURL(), copts...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ns.nc = nc
// 	return ns.nc, nil
// }
