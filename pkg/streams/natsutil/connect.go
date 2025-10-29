package natsutil

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog/log"
)

// ConnectJetStream establishes a NATS connection and JetStream context.
func ConnectJetStream(server string, opt ...nats.Option) (jetstream.JetStream, error) {
	nc, err := nats.Connect(server, opt...)
	if err != nil {
		return nil, fmt.Errorf("connect to NATS: %w", err)
	}
	js, err := jetstream.New(nc)
	if err != nil {
		if drainErr := nc.Drain(); drainErr != nil {
			log.Warn().Err(drainErr).Msg("failed to drain NATS connection after jetstream error")
		}
		return nil, fmt.Errorf("jetstream context: %w", err)
	}
	inProcess := nc.ConnectedAddr() == "pipe"
	log.Debug().Str("url", nc.ConnectedUrl()).Str("addr", nc.ConnectedAddr()).Bool("in_process", inProcess).Msg("JetStream connection established")
	return js, nil
}

func ConnectNATS(server string, opt ...nats.Option) (*nats.Conn, error) {
	nc, err := nats.Connect(server, opt...)
	if err != nil {
		return nil, fmt.Errorf("connect to NATS: %w", err)
	}
	inProcess := nc.ConnectedAddr() == "pipe"
	log.Debug().Str("url", nc.ConnectedUrl()).Str("addr", nc.ConnectedAddr()).Bool("in_process", inProcess).Msg("NATS connection established")
	return nc, nil
}
