package net

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/mon"
	"github.com/nats-io/nats.go"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

var counter = atomic.Uint64{}

func MonitorRequestHandler(nc *nats.Conn) http.HandlerFunc {
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
			data, err := json.Marshal(event)
			if err != nil {
				log.Error().Msgf("marshal event: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			mon.Emitter.FireHook(event)
			subject := event.Subject()
			err = nc.Publish(subject, data)
			if err != nil {
				log.Error().Msgf("publish event: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
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
