package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/gorilla/websocket"
)

// Dial creates a new client connection.
// It tries repeatedly to connect to the server until successful or context is done.
func Dial(ctx context.Context, url string) (*Connection, error) {
	log.Info().Msgf("dial: %s", url)
	socket, _, err := websocket.DefaultDialer.DialContext(ctx, url, nil)
	if err == nil {
		conn := NewConnection(ctx, socket)
		return conn, nil
	}
	ticker := time.NewTicker(reconnectInterval)
	defer ticker.Stop()
	for range ticker.C {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			socket, _, err := websocket.DefaultDialer.DialContext(ctx, url, nil)
			if err == nil {
				log.Debug().Msgf("connected to: %s", url)
				return NewConnection(ctx, socket), nil
			} else {
				log.Debug().Msgf("dial: %s", err)
			}
		}
	}
	return nil, fmt.Errorf("dial error: %s", url)
}
