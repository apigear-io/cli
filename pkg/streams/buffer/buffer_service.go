package buffer

import (
	"context"
	"strings"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog/log"
)

// BufferOptions controls the buffer runner behaviour.
type BufferOptions struct {
	MonitorSubject string
}

// RunBuffer mirrors all monitor traffic into device buffers using the global buffer window.
func RunBuffer(ctx context.Context, js jetstream.JetStream, opts BufferOptions) error {
	if js == nil {
		return nil
	}

	subjectPrefix := opts.MonitorSubject
	if subjectPrefix == "" {
		subjectPrefix = config.MonitorSubject
	}

	subjectPattern := subjectPrefix + ".>"
	sub, err := js.Conn().Subscribe(subjectPattern, func(msg *nats.Msg) {
		deviceID := extractDeviceID(subjectPrefix, msg.Subject)
		if deviceID == "" {
			return
		}
		// Always buffer with the global window
		err := Append(ctx, js, deviceID, config.BufferWindow, msg)
		if err != nil {
			log.Error().Err(err).Str("device", deviceID).Msg("buffer append failed")
		}
	})
	if err != nil {
		return err
	}
	defer func() {
		if err := sub.Unsubscribe(); err != nil {
			log.Warn().Err(err).Msg("buffer: unsubscribe failed")
		}
	}()

	log.Info().
		Str("subject", subjectPrefix).
		Dur("window", config.BufferWindow).
		Msg("buffer runner started (all devices)")

	<-ctx.Done()
	log.Info().Msg("buffer runner stopped")
	return ctx.Err()
}

func extractDeviceID(prefix, subject string) string {
	if !strings.HasPrefix(subject, prefix+".") {
		return ""
	}
	trimmed := strings.TrimPrefix(subject, prefix+".")
	if trimmed == "" {
		return ""
	}
	parts := strings.Split(trimmed, ".")
	return parts[0]
}
