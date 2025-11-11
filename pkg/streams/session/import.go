package session

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/store"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog/log"
)

// ImportOptions controls importing a recorded session from JSONL.
type ImportOptions struct {
	ServerURL    string
	Reader       io.Reader
	InputPath    string // optional source path for messaging purposes
	DeviceID     string // defaults to "123" if not specified
	SessionBucket string
	DeviceBucket string
	Verbose      bool
}

// Import reads messages from a JSONL file and recreates the session in JetStream.
func Import(ctx context.Context, opts ImportOptions) error {
	if opts.ServerURL == "" {
		return errors.New("server URL cannot be empty")
	}
	if opts.Reader == nil {
		return errors.New("reader cannot be nil")
	}

	// Default device ID to "123" if not specified
	if strings.TrimSpace(opts.DeviceID) == "" {
		opts.DeviceID = "123"
	}

	// Default buckets
	if opts.SessionBucket == "" {
		opts.SessionBucket = config.SessionBucket
	}
	if opts.DeviceBucket == "" {
		opts.DeviceBucket = config.DeviceBucket
	}

	nc, err := nats.Connect(opts.ServerURL)
	if err != nil {
		return fmt.Errorf("connect to NATS: %w", err)
	}
	defer func() {
		if drainErr := nc.Drain(); drainErr != nil {
			log.Error().Err(drainErr).Msg("failed to drain NATS connection after import")
		}
	}()

	js, err := jetstream.New(nc)
	if err != nil {
		return fmt.Errorf("jetstream context: %w", err)
	}

	scanner := bufio.NewScanner(opts.Reader)

	// Read first line (metadata)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("read metadata line: %w", err)
		}
		return errors.New("empty input file")
	}

	var metaLine MetadataLine
	if err := json.Unmarshal(scanner.Bytes(), &metaLine); err != nil {
		return fmt.Errorf("parse metadata: %w", err)
	}

	meta := metaLine.Metadata

	// Override device ID if provided
	if opts.DeviceID != "" {
		meta.DeviceID = opts.DeviceID
	}

	// Create device store and ensure device exists
	deviceStore, err := store.NewDeviceStore(js, opts.DeviceBucket)
	if err != nil {
		return fmt.Errorf("create device store: %w", err)
	}
	if _, err := deviceStore.Get(meta.DeviceID); err != nil {
		// Device doesn't exist, create it
		if err := deviceStore.Upsert(meta.DeviceID, store.DeviceInfo{
			Description: fmt.Sprintf("Auto-created during import of session %s", meta.SessionID),
			Updated:     time.Now(),
		}); err != nil {
			return fmt.Errorf("register device: %w", err)
		}
		if opts.Verbose {
			log.Info().Str("device_id", meta.DeviceID).Msg("created device")
		}
	}

	// Create session manager
	sessMgr, err := NewSessionStore(js, opts.SessionBucket)
	if err != nil {
		return fmt.Errorf("create session store: %w", err)
	}

	// Create session stream
	streamName := StreamName(meta.SessionID)
	stream, err := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:      streamName,
		Subjects:  []string{meta.SessionSubject},
		Retention: jetstream.LimitsPolicy,
		Storage:   jetstream.FileStorage,
		MaxAge:    72 * time.Hour, // Default retention
	})
	if err != nil {
		return fmt.Errorf("create stream: %w", err)
	}

	// Reset metadata counters for import
	meta.Start = time.Now()
	meta.MessageCount = 0
	meta.Stream = streamName

	// Store initial metadata
	if _, err := sessMgr.Put(&meta, 0); err != nil {
		return fmt.Errorf("store metadata: %w", err)
	}

	// Process messages
	messageCount := 0
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		var envelope Envelope
		if err := json.Unmarshal(scanner.Bytes(), &envelope); err != nil {
			return fmt.Errorf("parse envelope at line %d: %w", messageCount+2, err)
		}

		// Create NATS message with headers
		msg := nats.NewMsg(meta.SessionSubject)
		msg.Data = envelope.Data

		// Add headers from envelope
		for key, value := range envelope.Headers {
			msg.Header.Add(key, value)
		}

		// Override device and session IDs if they differ
		msg.Header.Set(config.HeaderDevice, meta.DeviceID)
		msg.Header.Set(config.HeaderSession, meta.SessionID)

		// Publish to stream
		if _, err := js.PublishMsg(ctx, msg); err != nil {
			return fmt.Errorf("publish message %d: %w", messageCount+1, err)
		}

		messageCount++

		if opts.Verbose && messageCount%100 == 0 {
			log.Info().Int("count", messageCount).Msg("imported messages")
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("read input: %w", err)
	}

	// Update metadata with final counts
	meta.End = time.Now()
	meta.MessageCount = messageCount

	// Load current revision and update
	_, rev, err := sessMgr.Load(meta.SessionID)
	if err != nil {
		return fmt.Errorf("load metadata for update: %w", err)
	}
	if _, err := sessMgr.Put(&meta, rev); err != nil {
		return fmt.Errorf("update metadata: %w", err)
	}

	// Verify stream info
	info, err := stream.Info(ctx)
	if err != nil {
		return fmt.Errorf("verify stream: %w", err)
	}

	if opts.Verbose {
		log.Info().
			Str("session_id", meta.SessionID).
			Str("device_id", meta.DeviceID).
			Int("messages", messageCount).
			Uint64("stream_messages", info.State.Msgs).
			Msg("import complete")
	}

	return nil
}
