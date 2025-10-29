package session

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog/log"
)

// ExportOptions controls exporting a recorded session to JSONL.
type ExportOptions struct {
	ServerURL  string
	SessionID  string
	Bucket     string
	Writer     io.Writer
	OutputPath string // optional destination path for messaging purposes
	Verbose    bool
}

// Export writes the messages of a recorded session to the provided writer as JSONL.
func Export(ctx context.Context, opts ExportOptions) error {
	if opts.ServerURL == "" {
		return errors.New("server URL cannot be empty")
	}
	if strings.TrimSpace(opts.SessionID) == "" {
		return errors.New("session-id cannot be empty")
	}
	if opts.Writer == nil {
		return errors.New("writer cannot be nil")
	}

	nc, err := nats.Connect(opts.ServerURL)
	if err != nil {
		return fmt.Errorf("connect to NATS: %w", err)
	}
	defer func() {
		if drainErr := nc.Drain(); drainErr != nil {
			log.Error().Err(drainErr).Msg("failed to drain NATS connection after export")
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

	durable := config.ExportConsumerName(meta.SessionID)
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

	written := 0
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		batch, err := consumer.Fetch(128, jetstream.FetchContext(ctx))
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
			if _, err := opts.Writer.Write(msg.Data()); err != nil {
				return fmt.Errorf("write message: %w", err)
			}
			if _, err := opts.Writer.Write([]byte("\n")); err != nil {
				return fmt.Errorf("write message: %w", err)
			}
			written++
			_ = msg.Ack()
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

		if meta.MessageCount > 0 && written >= meta.MessageCount {
			break
		}
	}

	return nil
}
