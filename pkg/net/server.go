package net

import (
	"encoding/json"
	"net/http"
	"objectapi/pkg/logger"
	"objectapi/pkg/mon"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

var log = logger.Get()

type Server struct {
	r chi.Router
}

func NewServer() *Server {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	return &Server{
		r: r,
	}
}

func (s *Server) Router() chi.Router {
	return s.r
}

func (s *Server) HandleMonitor(w http.ResponseWriter, r *http.Request) {
	log.Debug("handle monitor")
	deviceId := chi.URLParam(r, "device")
	if deviceId == "" {
		http.Error(w, "device id is required", http.StatusBadRequest)
		return
	}
	event := &mon.Event{}
	if err := json.NewDecoder(r.Body).Decode(event); err != nil {
		log.Infof("failed to decode event: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	event.DeviceId = deviceId
	event.Id = uuid.New().String()
	if event.Timestamp.IsZero() {
		event.Timestamp = time.Now()
	}
	log.Debugf("emit event: %+v", event)
	mon.EmitEvent(event)
}

func (s *Server) Start(addr string) error {
	s.r.Post("/monitor/{device}/", s.HandleMonitor)
	return http.ListenAndServe(addr, s.r)
}
