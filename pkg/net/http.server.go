package net

import (
	"context"
	"fmt"
	"net/http"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HTTPServer struct {
	router chi.Router
	server *http.Server
}

func NewHTTPServer() *HTTPServer {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	return &HTTPServer{
		router: r,
	}
}

func (s *HTTPServer) Router() chi.Router {
	return s.router
}

func (s *HTTPServer) Start(addr string) error {
	if s.server != nil {
		log.Info().Msgf("http server already started at %s", s.server.Addr)
		return nil
	}
	log.Debug().Msgf("start http server at %s", addr)
	server := &http.Server{Addr: addr, Handler: s.router}
	s.server = server
	return server.ListenAndServe()
}

func (s HTTPServer) Address() string {
	if s.server == nil {
		return ""
	}
	return s.server.Addr
}

func (s *HTTPServer) Restart(ctx context.Context, addr string) error {
	if s.server == nil {
		return fmt.Errorf("server not started")
	}
	log.Debug().Msgf("restart http server at %s", addr)
	s.server.RegisterOnShutdown(func() {
		log.Info().Msgf("shutdown http server at %s", addr)
	})
	err := s.server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("error shutting down server: %s", err)
	}
	s.server.Addr = addr
	return s.server.ListenAndServe()
}

func (s *HTTPServer) Stop() {
	if s.server == nil {
		return
	}
	log.Debug().Msgf("stop http server at %s", s.server.Addr)
	err := s.server.Shutdown(context.Background())
	if err != nil {
		log.Error().Msgf("error shutting down server: %s", err)
	}
}
