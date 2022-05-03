package rpc

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Hub struct {
	writer     RpcRequestHandler
	upgrader   *websocket.Upgrader
	conns      map[*Connection]bool
	register   chan *Connection
	unregister chan *Connection
	incoming   chan RpcRequest
	broadcast  chan RpcMessage
}

func NewHub(writer RpcRequestHandler) *Hub {
	h := &Hub{
		writer:     writer,
		upgrader:   &websocket.Upgrader{},
		conns:      make(map[*Connection]bool),
		register:   make(chan *Connection),
		unregister: make(chan *Connection),
		incoming:   make(chan RpcRequest),
		broadcast:  make(chan RpcMessage),
	}
	go h.run()
	return h
}

func (h *Hub) HandleWebsocketRequest(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	NewConnection(h, conn)
}

func (h *Hub) run() {
	for {
		select {
		case c := <-h.register:
			h.conns[c] = true
		case c := <-h.unregister:
			if _, ok := h.conns[c]; ok {
				close(c.send)
				delete(h.conns, c)
			}
		case r := <-h.incoming:
			if h.writer != nil {
				err := h.writer.HandleMessage(r)
				if err != nil {

					log.Printf("hub: write error: %v", err)
				}
			}
		case m := <-h.broadcast:
			for c := range h.conns {
				select {
				case c.send <- m:
				default:
					h.unregister <- c
				}
			}
		}
	}
}

func (h *Hub) SendMessage(m RpcMessage) {
	h.broadcast <- m
}
