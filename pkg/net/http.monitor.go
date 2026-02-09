package net

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/mon"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

var counter = atomic.Uint64{}

// STUB: NATS Removed - Event Broadcasting Disabled
// This handler receives monitor events via HTTP but does not broadcast them.
// Events are still emitted to local hooks via mon.Emitter.FireHook()
//
// To re-enable NATS broadcasting:
// 1. Add *nats.Conn parameter back to this function
// 2. Restore NATS publishing code (nc.Publish)
// 3. Update NetworkManager.EnableMonitor() to pass NATS connection
func MonitorRequestHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		source := chi.URLParam(r, "source")
		log.Debug().Msgf("handle monitor request %s", source)
		if source == "" {
			log.Error().Msg("source id is required")
			http.Error(w, "source id is required", http.StatusBadRequest)
			return
		}
		var events []*mon.Event
		err := json.NewDecoder(r.Body).Decode(&events)
		if err != nil {
			log.Error().Msgf("decode event: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for _, event := range events {
			event.Source = source
			if event.Id == "" {
				event.Id = strconv.FormatUint(counter.Add(1), 10)
			}
			if event.Timestamp.IsZero() {
				event.Timestamp = time.Now()
			}
			// Log event details (NATS broadcasting disabled)
			log.Info().
				Str("source", event.Source).
				Str("type", string(event.Type)).
				Str("id", event.Id).
				Str("subject", event.Subject()).
				Msg("Monitor event received (local only, not broadcast)")

			// Fire local hooks (still works)
			mon.Emitter.FireHook(event)
		}
		w.WriteHeader(http.StatusOK)
	}
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
		event.Source = source
		event.Id = uuid.New().String()
		if event.Timestamp.IsZero() {
			event.Timestamp = time.Now()
		}
		mon.Emitter.FireHook(event)
	}
}
