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
	var event mon.Event
	err := json.Unmarshal(msg.Data, &event)
	if err != nil {
		log.Error().Err(err).Msg("playback relay: unmarshal failed")
		return
	}
	frame, err := convertEventToOlinkMessage(&event)
	if err != nil {
		log.Error().Err(err).Msg("playback relay: convert event failed")
		return
	}
	r.factory.Dispatch(frame)
}

func convertEventToOlinkMessage(event *mon.Event) (core.Message, error) {
	switch event.Type {
	case mon.TypeCall:
		return core.MakeInvokeMessage(0, event.Symbol, core.AsArgs(nil)), nil
	case mon.TypeSignal:
		return core.MakeSignalMessage(event.Symbol, core.AsArgs(event.Data)), nil
	case mon.TypeState:
		for _, v := range event.Data {
			return core.MakePropertyChangeMessage(event.Symbol, v), nil
		}
	default:
		return core.Message{}, fmt.Errorf("unknown event type: %s", event.Type)
	}
	return core.Message{}, nil
}
