package buffer

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/natsutil"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

// EnsureStream creates or updates the buffer stream for a device.
func EnsureStream(js jetstream.JetStream, deviceID string, window time.Duration) (string, string, error) {
	if window <= 0 {
		return "", "", fmt.Errorf("buffer window must be positive")
	}
	streamName := config.BufferStreamName(deviceID)
	subject := config.BufferSubjectName(deviceID)

	cfg := jetstream.StreamConfig{
		Name:      streamName,
		Subjects:  []string{subject},
		Retention: jetstream.LimitsPolicy,
		MaxAge:    window,
		Storage:   jetstream.FileStorage,
	}

	_, err := js.CreateOrUpdateStream(context.Background(), cfg)
	if err != nil {
		return "", "", err
	}
	return streamName, subject, nil
}

// Append stores a monitor message in the device buffer.
func Append(ctx context.Context, js jetstream.JetStream, deviceID string, window time.Duration, msg *nats.Msg) error {
	if window <= 0 {
		return nil
	}
	_, subject, err := EnsureStream(js, deviceID, window)
	if err != nil {
		return err
	}

	buffered := &nats.Msg{
		Subject: subject,
		Header:  natsutil.CloneHeader(msg.Header),
		Data:    append([]byte(nil), msg.Data...),
	}
	if buffered.Header == nil {
		buffered.Header = nats.Header{}
	}
	buffered.Header.Set(config.HeaderBufferedAt, time.Now().UTC().Format(time.RFC3339Nano))

	if deadline, ok := ctx.Deadline(); ok {
		buffered.Header.Set(config.HeaderDeadline, deadline.Format(time.RFC3339Nano))
	}

	_, err = js.PublishMsg(ctx, buffered)
	return err
}

// Replay streams buffered messages in the given window into the provided publisher function.
func Replay(ctx context.Context, js jetstream.JetStream, deviceID string, since time.Time, until time.Time, publish func(*nats.Msg, time.Time) error) (int, time.Time, error) {
	return replay(ctx, js, deviceID, since, until, publish)
}

func replay(ctx context.Context, js jetstream.JetStream, deviceID string, since, until time.Time, publish func(*nats.Msg, time.Time) error) (int, time.Time, error) {
	stream := config.BufferStreamName(deviceID)
	subject := config.BufferSubjectName(deviceID)

	durable := config.BufferReplayConsumerName(deviceID)
	consumer, err := js.CreateOrUpdateConsumer(context.Background(), stream, jetstream.ConsumerConfig{
		Durable:       durable,
		AckPolicy:     jetstream.AckExplicitPolicy,
		DeliverPolicy: jetstream.DeliverAllPolicy,
		FilterSubject: subject,
	})
	if err != nil {
		if errors.Is(err, jetstream.ErrStreamNotFound) {
			return 0, time.Time{}, nil
		}
		return 0, time.Time{}, err
	}
	defer func() {
		_ = js.DeleteConsumer(context.Background(), stream, durable)
	}()

	count := 0
	var last time.Time

	for {
		err := ctx.Err()
		if err != nil {
			return count, last, err
		}

		batch, err := consumer.Fetch(64, jetstream.FetchMaxWait(250*time.Millisecond))
		if err != nil {
			if errors.Is(err, jetstream.ErrNoMessages) || errors.Is(err, nats.ErrTimeout) {
				break
			}
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return count, last, err
			}
			return count, last, err
		}

		processed := 0
		for msg := range batch.Messages() {
			if msg == nil {
				continue
			}
			processed++

			headers := natsutil.CloneHeader(msg.Headers())
			ts := parseBufferedAt(headers.Get(config.HeaderBufferedAt))
			if !ts.IsZero() {
				if ts.Before(since) || ts.After(until) {
					_ = msg.Ack()
					continue
				}
			}

			buffered := &nats.Msg{
				Subject: subject,
				Header:  headers,
				Data:    append([]byte(nil), msg.Data()...),
			}

			err := publish(buffered, ts)
			if err != nil {
				return count, last, err
			}
			count++
			if !ts.IsZero() {
				last = ts
			}
			err = msg.Ack()
			if err != nil {
				return count, last, err
			}
		}

		err = batch.Error()
		if err != nil {
			if errors.Is(err, jetstream.ErrNoMessages) || errors.Is(err, nats.ErrTimeout) {
				break
			}
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return count, last, err
			}
			return count, last, err
		}

		if processed == 0 {
			break
		}
	}

	return count, last, nil
}

func parseBufferedAt(v string) time.Time {
	if v == "" {
		return time.Time{}
	}
	ts, err := time.Parse(time.RFC3339Nano, v)
	if err != nil {
		return time.Time{}
	}
	return ts
}
