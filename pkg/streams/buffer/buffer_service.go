package buffer

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/store"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog/log"
)

// BufferOptions controls the buffer runner behaviour.
type BufferOptions struct {
	DeviceBucket    string
	MonitorSubject  string
	RefreshInterval time.Duration
}

// RunBuffer mirrors monitor traffic into device buffers based on device metadata.
func RunBuffer(ctx context.Context, js jetstream.JetStream, opts BufferOptions) error {
	if js == nil {
		return nil
	}

	bucket := opts.DeviceBucket
	if bucket == "" {
		bucket = config.DeviceBucket
	}
	subjectPrefix := opts.MonitorSubject
	if subjectPrefix == "" {
		subjectPrefix = "monitor"
	}
	refresh := opts.RefreshInterval
	if refresh <= 0 {
		refresh = config.BufferRefresh
	}

	devStore, err := store.NewDeviceStore(js, bucket)
	if err != nil {
		return err
	}

	var (
		mu     sync.RWMutex
		active = map[string]time.Duration{}
	)

	updateActive := func() {
		entries, err := devStore.List()
		if err != nil {
			log.Error().Err(err).Msg("buffer: list devices failed")
			return
		}
		mu.Lock()
		defer mu.Unlock()
		active = make(map[string]time.Duration, len(entries))
		for _, entry := range entries {
			if entry.Info.BufferDuration == "" {
				continue
			}
			dur, err := time.ParseDuration(entry.Info.BufferDuration)
			if err != nil || dur <= 0 {
				continue
			}
			active[strings.ToLower(entry.DeviceID)] = dur
		}
	}

	updateActive()

	subjectPattern := subjectPrefix + ".>"
	sub, err := js.Conn().Subscribe(subjectPattern, func(msg *nats.Msg) {
		deviceID := extractDeviceID(subjectPrefix, msg.Subject)
		if deviceID == "" {
			return
		}
		mu.RLock()
		window := active[strings.ToLower(deviceID)]
		mu.RUnlock()
		if window <= 0 {
			return
		}
		err := Append(ctx, js, deviceID, window, msg)
		if err != nil {
			log.Error().Err(err).Str("device", deviceID).Msg("buffer append failed")
		}
	})
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	ticker := time.NewTicker(refresh)
	defer ticker.Stop()

	log.Info().Str("subject", subjectPrefix).Dur("refresh", refresh).Msg("buffer runner started")

	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("buffer runner stopped")
			return ctx.Err()
		case <-ticker.C:
			updateActive()
		}
	}
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
