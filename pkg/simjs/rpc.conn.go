package simjs

import (
	"encoding/json"
	"time"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/nats-io/nats.go"
)

const (
	timeout = 10*time.Second + 1*time.Millisecond
)

type SimuHandler func(msg *SimuMessage) *SimuMessage

type SimuMiddleware func(SimuHandler) SimuHandler

type SimuPublisher interface {
	Publish(subject string, msg *SimuMessage) error
}

// ChainHandlers applies middlewares to a handler in reverse order (right to left)
func ChainHandlers(handler SimuHandler, middlewares ...SimuMiddleware) SimuHandler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

type Conn struct {
	nc          *nats.Conn
	sub         *nats.Subscription
	middlewares []SimuMiddleware
	handler     SimuHandler
}

func NewConn(nc *nats.Conn) (*Conn, error) {
	c := &Conn{
		nc:          nc,
		middlewares: make([]SimuMiddleware, 0),
		handler:     func(msg *SimuMessage) *SimuMessage { return nil },
	}
	err := c.Subscribe()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Conn) Use(middleware ...SimuMiddleware) {
	c.middlewares = append(c.middlewares, middleware...)
}

func (c *Conn) SetHandler(handler SimuHandler) {
	c.handler = handler
}

func (c *Conn) Subscribe() error {
	// Subscribe to all messages
	sub, err := c.nc.Subscribe("simu.>", func(msg *nats.Msg) {
		var simuMsg SimuMessage
		err := json.Unmarshal(msg.Data, &simuMsg)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal message")
			return
		}
		// handle the message
		reply := c.HandleMessage(&simuMsg)
		if msg.Reply != "" { // an answer is expected
			if reply == nil { // no answer
				msg.Respond(nil)
			} else { // answer with a simu message
				data, err := json.Marshal(reply)
				if err != nil {
					log.Error().Err(err).Msg("failed to marshal message")
					return
				}
				msg.Respond(data)
			}
		}
	})
	if err != nil { // subscribe failed
		log.Error().Err(err).Msg("failed to subscribe")
		return err
	}
	// store the subscription
	c.sub = sub
	return nil
}

func (c *Conn) Unsubscribe() {
	if c.sub == nil {
		return
	}
	err := c.sub.Unsubscribe()
	if err != nil {
		log.Error().Err(err).Msg("failed to unsubscribe")
	}
}

func (c *Conn) HandleMessage(msg *SimuMessage) *SimuMessage {
	handler := ChainHandlers(c.handler, c.middlewares...)
	return handler(msg)
}

func (c *Conn) Publish(subject string, msg *SimuMessage) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return c.nc.Publish(subject, data)
}

func (c *Conn) Request(subject string, msg *SimuMessage) (*SimuMessage, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	reply, err := c.nc.Request(subject, data, timeout)
	if err != nil {
		return nil, err
	}
	var re SimuMessage
	err = json.Unmarshal(reply.Data, &re)
	if err != nil {
		return nil, err
	}
	return &re, nil
}
