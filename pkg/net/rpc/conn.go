package rpc

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/gorilla/websocket"
)

var nextConnId = helper.MakeIdGenerator("conn")

const (
	reconnectInterval = 1 * time.Second
	// max message size in bytes
	maxMessageSize = 512
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

// Connection represents a cancelable websocket connection.
type Connection struct {
	id     string
	socket *websocket.Conn
	mu     sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc
	send   chan []byte
	recv   chan []byte
}

// NewConnection handles a new websocket
func NewConnection(ctx context.Context, socket *websocket.Conn) *Connection {
	c := &Connection{
		id:     nextConnId(),
		socket: socket,
		mu:     sync.Mutex{},
		send:   make(chan []byte),
		recv:   make(chan []byte),
	}
	c.ctx, c.cancel = context.WithCancel(ctx)
	go c.readPump()
	go c.writePump()
	return c
}

// Id returns the connection id.
func (c *Connection) Id() string {
	return c.id
}

func (c *Connection) ReadJSON(v interface{}) error {
	select {
	case <-c.ctx.Done():
		return c.ctx.Err()
	case data := <-c.recv:
		return json.Unmarshal(data, v)
	}
}

func (c *Connection) Read() []byte {
	return <-c.recv
}

func (c *Connection) Write(data []byte) {
	c.send <- data
}

func (c *Connection) WriteJSON(v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	c.send <- data
	return nil
}

func (c *Connection) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Close()
	}()
	for {
		select {
		case <-c.ctx.Done():
			return
		case <-ticker.C:
			err := c.socket.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				log.Errorf("conn: write ping error %v", err)
				return
			}
		case data, ok := <-c.send:
			if !ok {
				err := c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					log.Errorf("conn: write close message: %v", err)
				}
				return
			}
			err := c.socket.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Errorf("conn: write error %v", err)
				return
			}
		}
	}
}

func (c *Connection) readPump() {
	c.socket.SetReadLimit(maxMessageSize)
	c.socket.SetPongHandler(func(string) error {
		deadline := time.Now().Add(pongWait)
		return c.socket.SetReadDeadline(deadline)
	})
	c.socket.SetCloseHandler(func(code int, text string) error {
		// close connection and let write pump handle it
		c.Close()
		return nil
	})
	defer func() {
		c.Close()
	}()
	for {
		_, data, err := c.socket.ReadMessage()
		if err != nil {
			log.Errorf("conn: read error %v", err)
			return
		}
		c.recv <- data
	}
}

// Done returns a channel that is closed when the connection is closed.
func (c *Connection) Done() <-chan struct{} {
	return c.ctx.Done()
}

// Close closes the connection and the underlying socket.
func (c *Connection) Close() {
	log.Debugf("%s: close", c.id)
	c.cancel()
	c.socket.Close()
}
