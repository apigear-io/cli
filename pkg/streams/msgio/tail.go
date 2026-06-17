package msgio

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

// TailOptions controls how a monitoring subscription behaves.
type TailOptions struct {
	Subject  string
	DeviceID string
	Pretty   bool
	Headers  bool
	Verbose  bool
	Writer   io.Writer
}

// Tailer handles streaming monitor messages from NATS.
type Tailer struct {
	opts        TailOptions
	deviceID    string
	fullSubject string
	nc          *nats.Conn
}

// NewTailer prepares a Tailer instance for the provided context and options.
func NewTailer(nc *nats.Conn, opts TailOptions) *Tailer {
	opts.Subject = strings.TrimSpace(opts.Subject)
	opts.DeviceID = strings.TrimSpace(opts.DeviceID)

	if opts.DeviceID == "" {
		opts.DeviceID = ">"
	}
	if opts.Writer == nil {
		opts.Writer = os.Stdout
	}
	if opts.Subject == "" {
		opts.Subject = config.MonitorSubject
	}

	t := &Tailer{
		opts:        opts,
		deviceID:    opts.DeviceID,
		fullSubject: config.SubjectJoin(opts.Subject, opts.DeviceID),
		nc:          nc,
	}

	return t
}

// Run subscribes to the specified device stream and processes incoming messages.
func (t *Tailer) Run(ctx context.Context) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	sub, err := t.nc.Subscribe(t.fullSubject, func(msg *nats.Msg) {
		t.renderMessage(msg)
	})
	if err != nil {
		return fmt.Errorf("subscribe: %w", err)
	}
	defer func() {
		if drainErr := sub.Drain(); drainErr != nil && !errors.Is(drainErr, nats.ErrConnectionClosed) {
			log.Warn().Err(drainErr).Msg("drain subscription error")
		}
	}()
	err = t.nc.Flush()
	if err != nil {
		return fmt.Errorf("flush: %w", err)
	}
	if t.opts.Verbose {
		log.Info().Str("subject", t.fullSubject).Msg("monitoring")
	}
	<-ctx.Done()
	return nil
}

func (t *Tailer) renderMessage(msg *nats.Msg) {
	if t.opts.Headers && len(msg.Header) > 0 {
		keys := make([]string, 0, len(msg.Header))
		for key := range msg.Header {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			for _, value := range msg.Header.Values(key) {
				_, err := fmt.Fprintf(t.opts.Writer, "# header %s=%s\n", key, value)
				if err != nil {
					log.Error().Err(err).Msg("failed to write header")
				}
			}
		}
	}

	body := msg.Data
	if t.opts.Pretty {
		var buf bytes.Buffer
		err := json.Indent(&buf, body, "", "  ")
		if err == nil {
			body = buf.Bytes()
		} else if t.opts.Verbose {
			log.Warn().Err(err).Msg("pretty print failed")
		}
	}

	line := strings.TrimRight(string(body), "\n")
	_, err := fmt.Fprintln(os.Stdout, line)
	if err != nil {
		log.Error().Err(err).Msg("failed to write message")
	}
}
