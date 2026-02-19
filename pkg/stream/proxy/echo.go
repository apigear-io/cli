package proxy

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/apigear-io/cli/pkg/stream/relay"
)

// EchoServer implements a simple echo server that sends back received messages.
type EchoServer struct {
	name  string
	stats *ProxyStats
}

// NewEchoServer creates a new echo server.
func NewEchoServer(name string, stats *ProxyStats) *EchoServer {
	return &EchoServer{
		name:  name,
		stats: stats,
	}
}

// Handle processes a client connection by echoing all messages back.
func (e *EchoServer) Handle(ctx context.Context, conn relay.Connection) error {
	log.Debug().Str("proxy", e.name).Msg("echo server: client connected")

	defer func() {
		log.Debug().Str("proxy", e.name).Msg("echo server: client disconnected")
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-conn.Done():
			return nil
		default:
		}

		// Read message
		msgType, data, err := conn.ReadMessage()
		if err != nil {
			return err
		}

		// Record stats
		if e.stats != nil {
			e.stats.RecordMessageReceived(len(data))
		}

		// Echo it back
		if err := conn.WriteMessage(msgType, data); err != nil {
			return err
		}

		// Record stats
		if e.stats != nil {
			e.stats.RecordMessageSent(len(data))
		}
	}
}
