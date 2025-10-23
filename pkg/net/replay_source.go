package net

import (
	"context"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/nats-io/nats.go"
)

type ReplayOlinkRelay struct {
	nc          *nats.Conn
	sub         *nats.Subscription
	subject     string
	olinkServer IOlinkServer
}

func NewReplayOlinkRelay(nc *nats.Conn, subject string, olinkServer IOlinkServer) *ReplayOlinkRelay {
	return &ReplayOlinkRelay{
		nc:          nc,
		subject:     subject,
		olinkServer: olinkServer,
	}
}

func (r *ReplayOlinkRelay) Start(ctx context.Context) error {
	sub, err := r.nc.Subscribe(r.subject, r.handleMsg)
	if err != nil {
		return err
	}
	r.sub = sub
	return nil
}

func (r *ReplayOlinkRelay) Stop() error {
	if r.sub != nil {
		r.sub.Unsubscribe()
	}
	return nil
}

func (r *ReplayOlinkRelay) handleMsg(msg *nats.Msg) {
	log.Info().Msgf("ReplayOlinkSource received message on subject %s: %s", msg.Subject, string(msg.Data))
}
