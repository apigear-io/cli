package msgio

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/apigear-io/cli/pkg/streams/buffer"
	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/natsutil"
	"github.com/apigear-io/cli/pkg/streams/store"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog/log"
)

// TailOptions controls how a monitoring subscription behaves.
type TailOptions struct {
	ServerURL    string
	Subject      string
	DeviceID     string
	Pretty       bool
	Headers      bool
	Verbose      bool
	DeviceBucket string
	BufferWindow time.Duration
}

// Tail subscribes to the specified device stream and writes messages to stdout.
func Tail(ctx context.Context, opts TailOptions) error {
	t, err := newTailer(ctx, opts)
	if err != nil {
		return err
	}
	defer t.close()
	return t.run()
}

type tailer struct {
	ctx          context.Context
	opts         TailOptions
	deviceID     string
	fullSubject  string
	nc           *nats.Conn
	js           jetstream.JetStream
	bufferWindow time.Duration
	ensureBuffer bool
}

func newTailer(ctx context.Context, opts TailOptions) (*tailer, error) {
	baseSubject := strings.TrimSpace(opts.Subject)
	if baseSubject == "" {
		return nil, errors.New("subject cannot be empty")
	}
	deviceID := strings.TrimSpace(opts.DeviceID)
	if deviceID == "" {
		return nil, errors.New("device-id cannot be empty")
	}
	if opts.ServerURL == "" {
		return nil, errors.New("server URL cannot be empty")
	}

	js, err := natsutil.ConnectJetStream(opts.ServerURL)
	if err != nil {
		return nil, err
	}
	nc := js.Conn()

	t := &tailer{
		ctx:          ctx,
		opts:         opts,
		deviceID:     deviceID,
		fullSubject:  config.SubjectJoin(baseSubject, deviceID),
		nc:           nc,
		js:           js,
		bufferWindow: opts.BufferWindow,
	}

	if err := t.determineBufferWindow(); err != nil {
		t.close()
		return nil, err
	}
	if err := t.setupBuffer(); err != nil {
		t.close()
		return nil, err
	}

	return t, nil
}

func (t *tailer) determineBufferWindow() error {
	if t.bufferWindow != 0 || t.opts.DeviceBucket == "" {
		return nil
	}
	devStore, err := store.NewDeviceStore(t.js, t.opts.DeviceBucket)
	if err != nil {
		return nil
	}
	info, err := devStore.Get(t.deviceID)
	if err != nil || info.BufferDuration == "" {
		return nil
	}
	dur, parseErr := time.ParseDuration(info.BufferDuration)
	if parseErr != nil {
		if t.opts.Verbose {
			log.Warn().Str("device", t.deviceID).Err(parseErr).Msg("invalid buffer duration")
		}
		return nil
	}
	t.bufferWindow = dur
	return nil
}

func (t *tailer) setupBuffer() error {
	if t.bufferWindow <= 0 {
		return nil
	}
	_, _, err := buffer.EnsureStream(t.js, t.deviceID, t.bufferWindow)
	if err != nil {
		if t.opts.Verbose {
			log.Warn().Err(err).Msg("buffer disabled")
		}
		return nil
	}
	t.ensureBuffer = true
	return nil
}

func (t *tailer) run() error {
	msgCh := make(chan *nats.Msg, 256)
	sub, err := t.nc.ChanSubscribe(t.fullSubject, msgCh)
	if err != nil {
		return fmt.Errorf("subscribe: %w", err)
	}
	defer func() {
		if drainErr := sub.Drain(); drainErr != nil && !errors.Is(drainErr, nats.ErrConnectionClosed) {
			log.Warn().Err(drainErr).Msg("drain subscription error")
		}
	}()

	if t.opts.Verbose {
		log.Info().Str("subject", t.fullSubject).Msg("monitoring")
	}

	for {
		select {
		case <-t.ctx.Done():
			if t.opts.Verbose {
				log.Info().Err(t.ctx.Err()).Msg("monitor stopped")
			}
			return nil
		case msg, ok := <-msgCh:
			if !ok {
				return nil
			}
			if err := t.handleMessage(msg); err != nil {
				return err
			}
		}
	}
}

func (t *tailer) handleMessage(msg *nats.Msg) error {
	if err := renderMessage(msg, t.opts); err != nil {
		return err
	}
	if t.ensureBuffer {
		if err := buffer.Append(t.ctx, t.js, t.deviceID, t.bufferWindow, msg); err != nil && t.opts.Verbose {
			log.Warn().Err(err).Msg("buffer append failed")
		}
	}
	return nil
}

func (t *tailer) close() {
	if t.nc != nil {
		t.nc.Drain()
	}
}

func renderMessage(msg *nats.Msg, opts TailOptions) error {
	if opts.Headers && len(msg.Header) > 0 {
		keys := make([]string, 0, len(msg.Header))
		for key := range msg.Header {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			for _, value := range msg.Header.Values(key) {
				_, err := fmt.Fprintf(os.Stdout, "# header %s=%s\n", key, value)
				if err != nil {
					return err
				}
			}
		}
	}

	body := msg.Data
	if opts.Pretty {
		var buf bytes.Buffer
		err := json.Indent(&buf, body, "", "  ")
		if err == nil {
			body = buf.Bytes()
		} else if opts.Verbose {
			log.Warn().Err(err).Msg("pretty print failed")
		}
	}

	line := strings.TrimRight(string(body), "\n")
	_, err := fmt.Fprintln(os.Stdout, line)
	if err != nil {
		return err
	}
	return nil
}
