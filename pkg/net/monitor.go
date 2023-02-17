package net

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/mon"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

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
		log.Debug().Msgf("emit event: %+v", event)
		mon.Emitter.Emit(event)
	}
}
