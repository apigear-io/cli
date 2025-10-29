package session

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/natsutil"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog/log"
)

// PlaybackOptions controls replay of a recorded session.
type PlaybackOptions struct {
	ServerURL     string
	SessionID     string
	Bucket        string
	TargetSubject string
	Speed         float64
	Verbose       bool
}

// Playback replays a recorded session at the requested speed.
func Playback(ctx context.Context, opts PlaybackOptions) error {
	if opts.ServerURL == "" {
		return errors.New("server URL cannot be empty")
	}
	if strings.TrimSpace(opts.SessionID) == "" {
		return errors.New("session-id cannot be empty")
	}
	if opts.Speed == 0 {
		opts.Speed = 1
	}
	if opts.Speed <= 0 || math.IsNaN(opts.Speed) || math.IsInf(opts.Speed, 0) {
		return fmt.Errorf("invalid speed %v", opts.Speed)
	}

	nc, err := nats.Connect(opts.ServerURL)
	if err != nil {
		return fmt.Errorf("connect to NATS: %w", err)
	}
	defer func() {
		if drainErr := nc.Drain(); drainErr != nil {
			log.Error().Err(drainErr).Msg("failed to drain NATS connection after playback")
		}
	}()

	js, err := jetstream.New(nc)
	if err != nil {
		return fmt.Errorf("jetstream context: %w", err)
	}

	sessMgr, err := NewSessionStore(js, opts.Bucket)
	if err != nil {
		return err
	}

	meta, err := sessMgr.Info(opts.SessionID)
	if err != nil {
		return fmt.Errorf("load metadata: %w", err)
	}

	targetSubject := strings.TrimSpace(opts.TargetSubject)
	if targetSubject == "" {
		targetSubject = config.PlaybackSubject
	}

	durable := config.PlaybackConsumerName(meta.SessionID)
	consumer, err := js.CreateOrUpdateConsumer(context.Background(), meta.Stream, jetstream.ConsumerConfig{
		Durable:       durable,
		AckPolicy:     jetstream.AckExplicitPolicy,
		DeliverPolicy: jetstream.DeliverAllPolicy,
		FilterSubject: meta.SessionSubject,
	})
	if err != nil {
		return fmt.Errorf("create consumer: %w", err)
	}
	defer func() {
		_ = js.DeleteConsumer(context.Background(), meta.Stream, durable)
	}()

	var (
		prevTime time.Time
		played   int
	)

	for {
		err := ctx.Err()
		if err != nil {
			return err
		}

		batch, err := consumer.Fetch(50, jetstream.FetchContext(ctx))
		if err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return err
			}
			if errors.Is(err, jetstream.ErrNoMessages) {
				break
			}
			return fmt.Errorf("fetch: %w", err)
		}

		received := 0
		for msg := range batch.Messages() {
			if msg == nil {
				continue
			}
			received++
			err := ctx.Err()
			if err != nil {
				return err
			}

			headers := natsutil.CloneHeader(msg.Headers())
			recordedAt := parseRecordedAt(headers.Get(config.HeaderRecordedAt))
			if !prevTime.IsZero() {
				delay := recordedAt.Sub(prevTime)
				if delay < 0 {
					delay = 0
				}
				scaled := time.Duration(float64(delay) / opts.Speed)
				if scaled > 0 {
					select {
					case <-time.After(scaled):
					case <-ctx.Done():
						return ctx.Err()
					}
				}
			}

			publishMsg := &nats.Msg{
				Subject: targetSubject,
				Header:  headers,
				Data:    append([]byte(nil), msg.Data()...),
			}
			publishMsg.Header.Set(config.HeaderReplayedAt, time.Now().UTC().Format(time.RFC3339Nano))

			err = nc.PublishMsg(publishMsg)
			if err != nil {
				return fmt.Errorf("publish replay: %w", err)
			}

			err = msg.Ack()
			if err != nil {
				return fmt.Errorf("ack: %w", err)
			}

			prevTime = recordedAt
			played++
		}

		if batchErr := batch.Error(); batchErr != nil {
			if errors.Is(batchErr, context.Canceled) || errors.Is(batchErr, context.DeadlineExceeded) {
				return batchErr
			}
			if !errors.Is(batchErr, jetstream.ErrNoMessages) {
				return fmt.Errorf("fetch: %w", batchErr)
			}
		}

		if received == 0 {
			break
		}

		if meta.MessageCount > 0 && played >= meta.MessageCount {
			break
		}
	}

	return nil
}

func parseRecordedAt(value string) time.Time {
	if value == "" {
		return time.Now().UTC()
	}
	t, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		return time.Now().UTC()
	}
	return t
}
