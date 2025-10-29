package net

import (
	"context"

	"github.com/apigear-io/cli/pkg/log"
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
	frame, err := r.conv.FromData(msg.Data)
	if err != nil {
		log.Error().Err(err).Msg("playback relay: decode failed")
		return
	}
	r.factory.Dispatch(frame)
}
