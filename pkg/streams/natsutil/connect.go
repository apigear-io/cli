package natsutil

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

// ConnectJetStream establishes a NATS connection and JetStream context.
func ConnectJetStream(server string, opt ...nats.Option) (jetstream.JetStream, error) {
	nc, err := nats.Connect(server, opt...)
	if err != nil {
		return nil, fmt.Errorf("connect to NATS: %w", err)
	}
	js, err := jetstream.New(nc)
	if err != nil {
		nc.Drain()
		return nil, fmt.Errorf("jetstream context: %w", err)
	}
	return js, nil
}

func ConnectNATS(server string, opt ...nats.Option) (*nats.Conn, error) {
	nc, err := nats.Connect(server, opt...)
	if err != nil {
		return nil, fmt.Errorf("connect to NATS: %w", err)
	}
	return nc, nil
}
