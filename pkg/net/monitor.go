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
		http.Error(w, "source id is required", http.StatusBadRequest)
		return
	}
	event := &mon.Event{}
	if err := json.NewDecoder(r.Body).Decode(event); err != nil {
		log.Info().Msgf("decode event: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	event.Source = source
	event.Id = uuid.New().String()
	if event.Timestamp.IsZero() {
		event.Timestamp = time.Now()
	}
	log.Debug().Msgf("emit event: %+v", event)
	mon.EmitEvent(event)
}
