package cli

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/natsutil"
	"github.com/apigear-io/cli/pkg/streams/session"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func withSignalContext(parent context.Context, fn func(context.Context) error) error {
	ctx, cancel := signal.NotifyContext(parent, os.Interrupt, syscall.SIGTERM)
	defer cancel()
	return fn(ctx)
}

func withJetStream(ctx context.Context, fn func(jetstream.JetStream) error) error {
	js, err := natsutil.ConnectJetStream(rootOpts.server)
	if err != nil {
		return err
	}
	retErr := fn(js)
	if drainErr := js.Conn().Drain(); drainErr != nil {
		if retErr == nil {
			retErr = drainErr
		} else {
			retErr = errors.Join(retErr, drainErr)
		}
	}
	return retErr
}

func withSessionManager(ctx context.Context, bucket string, fn func(*session.SessionStore) error) error {
	if bucket == "" {
		bucket = config.SessionBucket
	}
	return withJetStream(ctx, func(js jetstream.JetStream) error {
		mgr, err := session.NewSessionStore(js, bucket)
		if err != nil {
			return err
		}
		return fn(mgr)
	})
}

func withNATS(_ context.Context, fn func(*nats.Conn) error) error {
	nc, err := natsutil.ConnectNATS(rootOpts.server)
	if err != nil {
		return err
	}
	defer nc.Drain()
	return fn(nc)
}
