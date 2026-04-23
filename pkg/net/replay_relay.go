package net

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/mon"
	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/nats-io/nats.go"
)

type ReplayOlinkRelay struct {
	nc      *nats.Conn
	sub     *nats.Subscription
	subject string
	factory *PlaybackSourceFactory
	conv    *core.MessageConverter
}

func NewReplayOlinkRelay(nc *nats.Conn, subject string, server IOlinkServer) *ReplayOlinkRelay {
	factory := NewPlaybackSourceFactory()
	server.SetSourceFactory(factory.SourceFactoryFunc())
	return &ReplayOlinkRelay{
		nc:      nc,
		subject: subject,
		factory: factory,
		conv:    core.NewConverter(core.FormatJson),
	}
}

func (r *ReplayOlinkRelay) Start(ctx context.Context) error {
	sub, err := r.nc.Subscribe(r.subject, r.handleMsg)
	if err != nil {
		return err
	}
	r.sub = sub

	go func() {
		<-ctx.Done()
		_ = r.Stop()
	}()

	log.Info().Str("subject", r.subject).Msg("playback relay subscribed")
	return nil
}

func (r *ReplayOlinkRelay) Stop() error {
	if r.sub != nil {
		_ = r.sub.Unsubscribe()
		r.sub = nil
	}
	return nil
}

func (r *ReplayOlinkRelay) handleMsg(msg *nats.Msg) {
	if msg == nil {
		return
	}
	log.Debug().Str("subject", msg.Subject).RawJSON("data", msg.Data).Msg("playback relay: message received")

	// Try to read metadata from NATS headers first (optimized path)
	var event mon.Event
	if msg.Header != nil && msg.Header.Get("X-Monitor-Type") != "" {
		// Headers available - reconstruct event from headers + data payload
		event.Type = mon.ParseEventType(msg.Header.Get("X-Monitor-Type"))
		event.Symbol = msg.Header.Get("X-Monitor-Symbol")
		event.Device = msg.Header.Get("X-Monitor-Device")
		event.Id = msg.Header.Get("X-Monitor-Id")
		// Timestamp parsing optional for routing

		// Unmarshal only the Data payload (not full event)
		var payload mon.Payload
		if err := json.Unmarshal(msg.Data, &payload); err != nil {
			log.Error().Err(err).Msg("playback relay: unmarshal data payload failed")
			return
		}
		event.Data = payload
	} else {
		// Fallback: full event decode (backward compatibility with old messages)
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Error().Err(err).Msg("playback relay: unmarshal event failed")
			return
		}
	}

	// one event can trigger multiple OLink frames (e.g., state with multiple fields)
	frames, err := convertEventToOlinkMessages(&event)
	if err != nil {
		log.Error().Err(err).Msg("playback relay: convert event failed")
		return
	}
	r.factory.Dispatch(frames)
}

// convertEventToOlinkMessages converts a monitor event to one or more OLink messages
func convertEventToOlinkMessages(event *mon.Event) ([]core.Message, error) {
	switch event.Type {
	case mon.TypeCall:
		return []core.Message{core.MakeInvokeMessage(0, event.Symbol, core.AsArgs(nil))}, nil
	case mon.TypeSignal:
		return []core.Message{core.MakeSignalMessage(event.Symbol, core.AsArgs(event.Data))}, nil
	case mon.TypeState:
		// State can have multiple fields, each becomes a separate PropertyChange message
		messages := make([]core.Message, 0, len(event.Data))
		module, object := core.SymbolIdToParts(event.Symbol)
		for field, value := range event.Data {
			// Build the property symbol: objectId/member
			objectId := fmt.Sprintf("%s.%s", module, object)
			propertySymbol := core.MakeSymbolId(objectId, field)
			log.Debug().Str("propertySymbol", propertySymbol).Interface("value", value).Msg("playback relay: state field")
			messages = append(messages, core.MakePropertyChangeMessage(propertySymbol, value))
		}
		return messages, nil
	default:
		return nil, fmt.Errorf("unknown event type: %s", event.Type)
	}
}
