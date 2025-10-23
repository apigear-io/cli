package app

import (
	"context"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/server"
	"github.com/apigear-io/cli/pkg/sim"
	"github.com/apigear-io/cli/pkg/streams/natsutil"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func WithNATS(ctx context.Context, addr string, fn func(*nats.Conn) error) error {
	nc, err := natsutil.ConnectNATS(addr)
	if err != nil {
		log.Info().Msg("NATS server not available, starting temporary server")
		WithServer(ctx, server.Options{
			NatsHost: "localhost",
			NatsPort: 4222,
			HttpAddr: "localhost:5555",
		}, func(s *server.Server) error {
			nc, err = s.NetworkManager().NatsConnection()
			return err
		})
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	defer nc.Drain()
	log.Info().Msg("NATS server available")
	return fn(nc)
}

func WithJetstream(server string, fn func(js jetstream.JetStream) error, opt ...nats.Option) (err error) {
	js, err := natsutil.ConnectJetStream(server, opt...)
	if err != nil {
		return err
	}
	err = fn(js)
	js.Conn().Drain()
	return err
}

func WithServer(ctx context.Context, opts server.Options, fn func(*server.Server) error) error {
	server := server.New(opts)
	err := server.Start(ctx)
	if err != nil {
		return err
	}
	defer server.Stop()
	return fn(server)
}

func WithSimuClient(ctx context.Context, natsServer string, action func(ctx context.Context, client *sim.Client) error) error {
	nc, err := nats.Connect(natsServer)
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to nats server")
		return err
	}
	defer func() {
		nc.Drain()
		nc.Close()
	}()
	client := sim.NewClient(nc)
	return action(ctx, client)
}
