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
	"github.com/apigear-io/cli/pkg/streams/natsutil"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

// TailOptions controls how a monitoring subscription behaves.
type TailOptions struct {
	ServerURL string
	Subject   string
	DeviceID  string
	Pretty    bool
	Headers   bool
	Verbose   bool
	Writer    io.Writer
}

func (o *TailOptions) Validate() error {
	o.ServerURL = strings.TrimSpace(o.ServerURL)
	o.Subject = strings.TrimSpace(o.Subject)
	o.DeviceID = strings.TrimSpace(o.DeviceID)
	if o.ServerURL == "" {
		return errors.New("server URL cannot be empty")
	}
	if o.Subject == "" {
		o.Subject = config.MonitorSubject
	}
	if o.DeviceID == "" {
		o.DeviceID = ">"
	}
	if o.Writer == nil {
		o.Writer = os.Stdout
	}
	return nil
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
	ctx         context.Context
	opts        TailOptions
	deviceID    string
	fullSubject string
	nc          *nats.Conn
}

func newTailer(ctx context.Context, opts TailOptions) (*tailer, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}
	nc, err := natsutil.ConnectNATS(opts.ServerURL)
	if err != nil {
		return nil, err
	}

	t := &tailer{
		ctx:         ctx,
		opts:        opts,
		deviceID:    opts.DeviceID,
		fullSubject: config.SubjectJoin(opts.Subject, opts.DeviceID),
		nc:          nc,
	}

	return t, nil
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
	if err := t.renderMessage(msg); err != nil {
		return err
	}
	return nil
}

func (t *tailer) close() {
	if t.nc != nil {
		t.nc.Drain()
	}
}

func (t *tailer) renderMessage(msg *nats.Msg) error {
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
					return err
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
		return err
	}
	return nil
}
