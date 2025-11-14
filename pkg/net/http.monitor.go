package net

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/mon"
	"github.com/apigear-io/cli/pkg/streams"
	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/controller"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

var counter = atomic.Uint64{}

// deviceTracker tracks seen device IDs and triggers on-demand recording for new devices.
type deviceTracker struct {
	devices sync.Map // deviceId (string) -> true (bool)
}

// isNewDevice checks if a device is new and marks it as seen atomically.
// Returns true if the device was newly added, false if already seen.
func (dt *deviceTracker) isNewDevice(deviceId string) bool {
	_, loaded := dt.devices.LoadOrStore(deviceId, true)
	return !loaded // true if newly stored, false if already existed
}

func MonitorRequestHandler(nc *nats.Conn) http.HandlerFunc {
	// Create device tracker for auto on-demand recording
	tracker := &deviceTracker{}

	return func(w http.ResponseWriter, r *http.Request) {
		deviceId := chi.URLParam(r, "source")
		log.Debug().Msgf("handle monitor request %s", deviceId)
		if deviceId == "" {
			log.Error().Msg("source id is required")
			http.Error(w, "source id is required", http.StatusBadRequest)
			return
		}

		// Check if this is a new device and auto-start recording if needed
		if tracker.isNewDevice(deviceId) {
			log.Info().Msgf("new device detected: %s, starting recording", deviceId)
			go autoStartRecording(nc, deviceId)
		}

		var events []*mon.Event
		err := json.NewDecoder(r.Body).Decode(&events)
		if err != nil {
			log.Error().Msgf("decode event: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Prepare all events (set metadata, fire hooks)
		for _, event := range events {
			event.Device = deviceId
			if event.Id == "" {
				event.Id = strconv.FormatUint(counter.Add(1), 10)
			}
			if event.Timestamp.IsZero() {
				event.Timestamp = time.Now()
			}
			mon.Emitter.FireHook(event)
		}

		// Bulk publish all events with single flush
		err = streams.PublishMonitorMessageBulk(nc, events)
		if err != nil {
			log.Error().Msgf("bulk publish events: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// autoStartRecording sends an RPC command to start recording for a device.
// This runs in a background goroutine to avoid blocking the HTTP response.
func autoStartRecording(nc *nats.Conn, deviceId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Check if a recording is already active for this device
	js, err := jetstream.New(nc)
	if err != nil {
		log.Warn().Msgf("auto-start recording: failed to get jetstream for device %s: %v", deviceId, err)
		// Continue anyway - controller will reject if recording exists
	} else {
		states, err := controller.ListStates(js, config.StateBucket)
		if err != nil {
			log.Debug().Msgf("auto-start recording: failed to list states for device %s: %v", deviceId, err)
			// Continue anyway - controller will reject if recording exists
		} else {
			// Check if any active recording exists for this device
			for _, state := range states {
				if state.DeviceID == deviceId && state.Status == "running" {
					log.Debug().Msgf("auto-start recording skipped: device %s already has active session %s", deviceId, state.SessionID)
					return
				}
			}
		}
	}

	request := controller.RpcRequest{
		Action:   controller.ActionStart,
		Subject:  config.MonitorSubject,
		DeviceID: deviceId,
		// SessionID will be auto-generated
		// Retention, PreRoll, etc. use defaults
	}

	resp, err := controller.SendCommand(ctx, nc, config.RecordRpcSubject, request)
	if err != nil {
		log.Warn().Msgf("auto-start recording failed for device %s: %v", deviceId, err)
		return
	}

	if !resp.OK {
		// Controller rejected (e.g., recording already exists)
		log.Debug().Msgf("auto-start recording response for device %s: %s", deviceId, resp.Message)
		return
	}

	log.Info().Msgf("auto-started recording for device %s, session=%s", deviceId, resp.SessionID)
}

// HandleMonitorRequest handles the monitor http request.
// events are emitted to the monitor event channel.
func HandleMonitorRequest(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("handle monitor request")
	source := chi.URLParam(r, "source")
	if source == "" {
		log.Error().Msg("source id is required")
		http.Error(w, "source id is required", http.StatusBadRequest)
		return
	}
	// monitor events are sent as an array of json objects
	var events []*mon.Event
	err := json.NewDecoder(r.Body).Decode(&events)
	if err != nil {
		log.Error().Msgf("decode event: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// set source and id for each event
	for _, event := range events {
		event.Device = source
		event.Id = uuid.New().String()
		if event.Timestamp.IsZero() {
			event.Timestamp = time.Now()
		}
		mon.Emitter.FireHook(event)
	}
}
