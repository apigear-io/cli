package sim

import (
	"context"
	"net/http"

	"github.com/apigear-io/objectlink-core-go/olink/remote"
	"github.com/apigear-io/objectlink-core-go/olink/ws"
)

type IOlinkServer interface {
	RegisterSource(source remote.IObjectSource)
	UnregisterSource(source remote.IObjectSource)
}

type OlinkServer struct {
	registry *remote.Registry
	hub      *ws.Hub
}

func NewOlinkServer() *OlinkServer {
	registry := remote.NewRegistry()
	hub := ws.NewHub(context.Background(), registry)
	return &OlinkServer{
		hub:      hub,
		registry: registry,
	}
}

func (s *OlinkServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.hub.ServeHTTP(w, r)
}

func (s *OlinkServer) Close() {
	s.hub.Close()
}

func (s *OlinkServer) RegisterSource(source remote.IObjectSource) {
	// make sure source is not registered yet
	s.UnregisterSource(source)
	// register source
	err := s.registry.AddObjectSource(source)
	if err != nil {
		log.Error().Err(err).Msg("Failed to register source")
	}
}

func (s *OlinkServer) UnregisterSource(source remote.IObjectSource) {
	s.registry.RemoveObjectSource(source)
}
