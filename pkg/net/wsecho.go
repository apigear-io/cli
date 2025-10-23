package net

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// WSEchoOptions configure the echo server behaviour.
type WSEchoOptions struct {
	Addr     string
	Path     string
	Headers  http.Header
	Upgrader *websocket.Upgrader
}

// RunWSEcho starts a WebSocket echo server until the context is cancelled.
func RunWSEcho(ctx context.Context, opts WSEchoOptions) error {
	path := NormalizePath(opts.Path)
	if path == "" || path == "/" {
		path = "/ws"
	}

	upgrader := opts.Upgrader
	if upgrader == nil {
		upgrader = &websocket.Upgrader{
			CheckOrigin: func(*http.Request) bool { return true },
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, opts.Headers)
		if err != nil {
			log.Printf("wsecho: upgrade error: %v", err)
			return
		}
		defer conn.Close()

		log.Printf("wsecho: client connected %s", r.RemoteAddr)
		for {
			typ, payload, err := conn.ReadMessage()
			if err != nil {
				log.Printf("wsecho: read error: %v", err)
				return
			}
			if err := conn.WriteMessage(typ, payload); err != nil {
				log.Printf("wsecho: write error: %v", err)
				return
			}
		}
	})

	server := &http.Server{
		Addr:    opts.Addr,
		Handler: mux,
	}

	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("wsecho: shutdown error: %v", err)
		}
		close(done)
	}()

	log.Printf("wsecho: listening on %s%s", opts.Addr, path)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	<-done
	return nil
}
