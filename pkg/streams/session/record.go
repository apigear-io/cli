package session

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/apigear-io/cli/pkg/streams/buffer"
	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/natsutil"
	"github.com/apigear-io/cli/pkg/streams/store"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog/log"
)

// RecordOptions controls how a live device stream is captured into JetStream.
type RecordOptions struct {
	ServerURL     string
	Subject       string
	DeviceID      string
	SessionID     string
	Retention     time.Duration
	SessionBucket string
	DeviceBucket  string
	Device        store.DeviceInfo
	Verbose       bool
	Progress      func(Metadata)
	PreRoll       time.Duration
}

// Record subscribes to subject.deviceID and persists messages into a dedicated JetStream stream, tracking metadata in KV.
func Record(ctx context.Context, opts RecordOptions) (*Metadata, error) {
	if opts.ServerURL == "" {
		return nil, errors.New("server URL cannot be empty")
	}
	baseSubject := strings.TrimSpace(opts.Subject)
	if baseSubject == "" {
		return nil, errors.New("subject cannot be empty")
	}
	opts.DeviceID = strings.TrimSpace(opts.DeviceID)
	if opts.DeviceID == "" {
		return nil, errors.New("device-id cannot be empty")
	}
	sessionID := strings.TrimSpace(opts.SessionID)
	if sessionID == "" {
		sessionID = uuid.NewString()
	}
	sessionBucket := strings.TrimSpace(opts.SessionBucket)
	if sessionBucket == "" {
		sessionBucket = config.SessionBucket
	}

	nc, err := nats.Connect(opts.ServerURL)
	if err != nil {
		return nil, fmt.Errorf("connect to NATS: %w", err)
	}
	defer nc.Drain()

	js, err := jetstream.New(nc)
	if err != nil {
		return nil, fmt.Errorf("jetstream context: %w", err)
	}

	sessMgr, err := NewSessionStore(js, sessionBucket)
	if err != nil {
		return nil, err
	}
	_, _, err = sessMgr.Load(sessionID)
	if err == nil {
		log.Warn().Str("session", sessionID).Msg("session already exists")
		return nil, fmt.Errorf("session %s already exists", sessionID)
	} else if !errors.Is(err, jetstream.ErrKeyNotFound) {
		return nil, err
	}

	devStore, err := store.NewDeviceStore(js, opts.DeviceBucket)
	if err != nil {
		return nil, err
	}
	if infoIsEmpty(opts.Device) {
		err := devStore.Ensure(opts.DeviceID)
		if err != nil {
			return nil, fmt.Errorf("ensure device: %w", err)
		}
	} else {
		err := devStore.Upsert(opts.DeviceID, opts.Device)
		if err != nil {
			return nil, fmt.Errorf("upsert device: %w", err)
		}
	}

	sourceSubject := config.DeviceSubject(baseSubject, opts.DeviceID)
	sessionSubject := config.SessionSubject(sessionID)
	streamName := StreamName(sessionID)

	streamCfg := jetstream.StreamConfig{
		Name:      streamName,
		Subjects:  []string{sessionSubject},
		Retention: jetstream.LimitsPolicy,
		Storage:   jetstream.FileStorage,
	}
	if opts.Retention > 0 {
		streamCfg.MaxAge = opts.Retention
	}

	_, err = js.CreateStream(ctx, streamCfg)
	if err != nil {
		return nil, fmt.Errorf("add stream: %w", err)
	}

	log.Info().Str("session", sessionID).Str("device", opts.DeviceID).Msg("record stream created")

	meta := &Metadata{
		SessionID:      sessionID,
		DeviceID:       opts.DeviceID,
		SourceSubject:  sourceSubject,
		SessionSubject: sessionSubject,
		Stream:         streamName,
		Bucket:         sessionBucket,
		Start:          time.Now().UTC(),
		End:            time.Now().UTC(),
	}
	if opts.Retention > 0 {
		meta.Retention = opts.Retention.String()
	}

	if opts.PreRoll > 0 {
		replayCtx, cancelReplay := context.WithTimeout(context.Background(), opts.PreRoll+time.Second)
		defer cancelReplay()
		since := time.Now().Add(-opts.PreRoll)
		until := time.Now()
		count, last, err := buffer.Replay(replayCtx, js, opts.DeviceID, since, until, func(bufMsg *nats.Msg, bufferedAt time.Time) error {
			recordedAt := bufferedAt
			if recordedAt.IsZero() {
				recordedAt = time.Now().UTC()
			}
			replayed := &nats.Msg{
				Subject: sessionSubject,
				Header:  nats.Header{},
				Data:    append([]byte(nil), bufMsg.Data...),
			}
			replayed.Header.Set("Content-Type", "application/json")
			replayed.Header.Set(config.HeaderDevice, opts.DeviceID)
			replayed.Header.Set(config.HeaderSession, sessionID)
			replayed.Header.Set(config.HeaderRecordedAt, recordedAt.Format(time.RFC3339Nano))
			replayed.Header.Set(config.HeaderPreRoll, "true")
			return publishToStream(replayCtx, js, replayed)
		})
		if err != nil {
			log.Error().Err(err).Str("session", sessionID).Msg("pre-roll replay failed")
		} else if count > 0 {
			meta.MessageCount = count
			if !last.IsZero() {
				meta.End = last
			}
		}
	}

	revision, err := sessMgr.Put(meta, 0)
	if err != nil {
		return nil, err
	}
	if opts.Progress != nil {
		opts.Progress(*meta)
	}

	msgCh := make(chan *nats.Msg, 1024)
	sub, err := nc.ChanSubscribe(sourceSubject, msgCh)
	if err != nil {
		return nil, fmt.Errorf("subscribe source: %w", err)
	}
	defer sub.Drain()

	var mu sync.Mutex

	updateMeta := func(update func(*Metadata)) error {
		mu.Lock()
		defer mu.Unlock()
		update(meta)
		rev, err := sessMgr.Put(meta, revision)
		if err != nil {
			return err
		}
		revision = rev
		if opts.Progress != nil {
			copy := *meta
			opts.Progress(copy)
		}
		return nil
	}

	for {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			_ = updateMeta(func(m *Metadata) {
				m.End = time.Now().UTC()
			})
			if errors.Is(err, context.Canceled) {
				log.Info().Str("session", sessionID).Msg("record context canceled")
				return meta, nil
			}
			return meta, err
		case msg, ok := <-msgCh:
			if !ok {
				log.Info().Str("session", sessionID).Msg("record channel closed")
				return meta, nil
			}

			recordedAt := time.Now().UTC()
			stored := &nats.Msg{
				Subject: sessionSubject,
				Header:  natsutil.CloneHeader(msg.Header),
				Data:    append([]byte(nil), msg.Data...),
			}
			stored.Header.Set("Content-Type", "application/json")
			stored.Header.Set(config.HeaderDevice, opts.DeviceID)
			stored.Header.Set(config.HeaderSession, sessionID)
			stored.Header.Set(config.HeaderRecordedAt, recordedAt.Format(time.RFC3339Nano))

			err := publishToStream(ctx, js, stored)
			if err != nil {
				log.Error().Err(err).Str("session", sessionID).Msg("publish to stream failed")
				return meta, err
			}

			err = updateMeta(func(m *Metadata) {
				m.MessageCount++
				m.End = recordedAt
			})
			if err != nil {
				log.Error().Err(err).Str("session", sessionID).Msg("update metadata failed")
				return meta, err
			}
		}
	}
}

func publishToStream(ctx context.Context, js jetstream.JetStream, msg *nats.Msg) error {
	err := ctx.Err()
	if err != nil {
		return err
	}
	_, err = js.PublishMsg(ctx, msg)
	return err
}

func infoIsEmpty(info store.DeviceInfo) bool {
	return info.IsZero()
}
