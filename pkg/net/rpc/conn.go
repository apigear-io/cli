package rpc

import (
	"time"

	"github.com/apigear-io/cli/pkg/log"

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
	log.Debugln("writePump")
	ticker := time.NewTicker(10 * time.Second)
	defer func() {
		log.Debugln("writePump: closing")
		c.conn.Close()
		ticker.Stop()
	}()
	for {
		select {
		case m, ok := <-c.send:
			log.Debugf("writePump: %v", m)
			if !ok {
				err := c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					log.Warnf("writePump: %s", err)
				}
				return
			}
			err := c.conn.WriteJSON(m)
			if err != nil {
				return
			}
		case <-ticker.C:
			log.Debug("writePump: ping")
			err := c.conn.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				err := c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					log.Warnf("error sending message in write pump: %s", err)
				}
				return
			}
		}
	}
}

func (c *Connection) readPump() {
	log.Debugln("readPump")
	for {
		var m RpcMessage
		err := c.conn.ReadJSON(&m)
		log.Debugf("readPump: %v", m)
		if err != nil {
			c.hub.unregister <- c
			c.conn.Close()
			return
		}
		c.hub.incoming <- RpcRequest{Msg: m, Conn: c}
	}
}
