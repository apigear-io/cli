package streams

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/mon"
	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/nats-io/nats.go"
)

func PublishMonitorMessage(nc *nats.Conn, deviceId string, data []byte) error {
	log.Debug().Msgf("publish monitor message from device %s", deviceId)
	return nc.Publish(config.DeviceSubject(config.MonitorSubject, deviceId), data)
}

// PublishMonitorMessageBulk publishes multiple monitor events efficiently using NATS headers
// and bulk publishing with a single flush operation.
func PublishMonitorMessageBulk(nc *nats.Conn, events []*mon.Event) error {
	if nc == nil {
		return fmt.Errorf("nats connection is nil")
	}
	if len(events) == 0 {
		return nil
	}

	log.Debug().Msgf("bulk publish %d monitor messages", len(events))

	// Publish all messages (client buffers automatically)
	for _, event := range events {
		// Marshal only the Data payload (not the full event)
		data, err := json.Marshal(event.Data)
		if err != nil {
			return fmt.Errorf("marshal event data: %w", err)
		}

		// Create message with headers for metadata
		subject := config.DeviceSubject(config.MonitorSubject, event.Device)
		msg := &nats.Msg{
			Subject: subject,
			Header:  nats.Header{},
			Data:    data,
		}

		// Add metadata as NATS headers
		msg.Header.Set("X-Monitor-Device", event.Device)
		msg.Header.Set("X-Monitor-Id", event.Id)
		msg.Header.Set("X-Monitor-Type", string(event.Type))
		msg.Header.Set("X-Monitor-Timestamp", event.Timestamp.Format("2006-01-02T15:04:05.999999999Z07:00"))
		msg.Header.Set("X-Monitor-Symbol", event.Symbol)

		// Publish (buffered by client)
		if err := nc.PublishMsg(msg); err != nil {
			return fmt.Errorf("publish event %s: %w", event.Id, err)
		}
	}

	// FlushTimeout ensures all buffered messages are sent AND confirmed by server
	// This waits for a PING/PONG roundtrip, guaranteeing the server has received all messages
	if err := nc.FlushTimeout(5 * time.Second); err != nil {
		return fmt.Errorf("flush timeout: %w", err)
	}

	// Check for any async publish errors
	if err := nc.LastError(); err != nil {
		return fmt.Errorf("nats error: %w", err)
	}

	return nil
}
