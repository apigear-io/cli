package rpc

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Connection struct {
	hub  *Hub
	conn *websocket.Conn
	send chan RpcMessage
}

func NewConnection(server *Hub, conn *websocket.Conn) *Connection {
	c := &Connection{hub: server, conn: conn}
	server.register <- c
	go c.writePump()
	go c.readPump()
	return c
}

func (c *Connection) Send(m RpcMessage) {
	c.send <- m
}

func (c *Connection) BroadCast(m RpcMessage) {
	c.hub.broadcast <- m
}

func (c *Connection) writePump() {
	log.Println("writePump")
	ticker := time.NewTicker(10 * time.Second)
	defer func() {
		log.Println("writePump: close")
		c.conn.Close()
		ticker.Stop()
	}()
	for {
		select {
		case m, ok := <-c.send:
			log.Printf("writePump: %v", m)
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			err := c.conn.WriteJSON(m)
			if err != nil {
				return
			}
		case <-ticker.C:
			log.Printf("writePump: ping")
			err := c.conn.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
		}
	}
}

func (c *Connection) readPump() {
	log.Println("readPump")
	for {
		var m RpcMessage
		err := c.conn.ReadJSON(&m)
		log.Printf("readPump: %v", m)
		if err != nil {
			c.hub.unregister <- c
			c.conn.Close()
			return
		}
		c.hub.incoming <- RpcRequest{Msg: m, Conn: c}
	}
}
