package net

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	router chi.Router
	server *http.Server
}

func NewHTTPServer() *Server {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(NewHttpLogger())
	// r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	return &Server{
		router: r,
	}
}

func (s *Server) Router() chi.Router {
	return s.router
}

func (s *Server) Start(addr string) error {
	if s.server != nil {
		return fmt.Errorf("server already started")
	}
	log.Debugf("start http server at %s", addr)
	server := &http.Server{Addr: addr, Handler: s.router}
	s.server = server
	return server.ListenAndServe()
}

func (s Server) Address() string {
	if s.server == nil {
		return ""
	}
	return s.server.Addr
}

func (s *Server) Restart(ctx context.Context, addr string) error {
	if s.server == nil {
		return fmt.Errorf("server not started")
	}
	log.Debugf("restart http server at %s", addr)
	err := s.server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("error shutting down server: %s", err)
	}
	s.server.Addr = addr
	return s.server.ListenAndServe()
}
