package rpc

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type Hub struct {
	id          string
	register    chan *Connection
	unregister  chan *Connection
	broadcast   chan []byte
	connections map[*Connection]bool
	ctx         context.Context
	requests    chan *Request
}

var nextHubId = helper.MakeIdGenerator("hub")

func NewHub(ctx context.Context) *Hub {
	h := &Hub{
		id:          nextHubId(),
		register:    make(chan *Connection),
		unregister:  make(chan *Connection),
		broadcast:   make(chan []byte),
		connections: make(map[*Connection]bool),
		requests:    make(chan *Request),
		ctx:         ctx,
	}
	go h.run()
	return h
}

func (h *Hub) run() {
	for {
		select {
		case <-h.ctx.Done():
			return
		case conn := <-h.register:
			log.Info().Msgf("register: %s", conn.Id())
			h.connections[conn] = true
		case conn := <-h.unregister:
			if _, ok := h.connections[conn]; ok {
				delete(h.connections, conn)
				conn.Close()
			}
		case data := <-h.broadcast:
			for conn := range h.connections {
				conn.Write(data)
			}
		}
	}
}

func (h *Hub) Requests() <-chan *Request {
	return h.requests
}

// Broadcast sends data to all connections
func (h *Hub) Broadcast(data []byte) {
	h.broadcast <- data
}

// BroadcastJSON encodes v as JSON and broadcasts it to all connections.
func (h *Hub) BroadcastJSON(v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	h.Broadcast(data)
	return nil
}

func (h *Hub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Info().Msgf("upgrade: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Error().Msgf("error writing response: %v", err)
		}
		return
	}
	conn := NewConnection(h.ctx, socket)
	go func() {
		// unregister connection when it is done
		<-conn.Done()
		log.Info().Msgf("unregister: %s", conn.Id())
		h.unregister <- conn
	}()
	go func() {
		// read requests from connection and stream them to the hub
		defer func() {
			h.unregister <- conn
		}()
		for {
			select {
			case <-h.ctx.Done():
				return
			default:
				data := conn.Read()
				req := &Request{
					data: data,
					conn: conn,
					err:  err,
				}
				h.requests <- req
			}
		}
	}()
	h.register <- conn
}
