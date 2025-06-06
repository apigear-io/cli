package evt

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

const (
	NATS_TIMEOUT = 10 * time.Second
)

// NatsEventBus implements IEventBus using nats
type NatsEventBus struct {
	rw         sync.RWMutex
	subject    string
	nc         *nats.Conn
	handlers   map[string]HandlerFunc
	middleware []HandlerFunc
	sub        *nats.Subscription
}

func NewNatsEventBus(subject string, nc *nats.Conn) *NatsEventBus {
	bus := &NatsEventBus{
		subject:  subject,
		nc:       nc,
		handlers: make(map[string]HandlerFunc),
	}
	bus.setup()
	return bus
}

// Publish sends an event
func (b *NatsEventBus) Publish(e *Event) error {
	data, err := json.Marshal(e)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal event")
		return err
	}
	return b.nc.Publish(b.subject, data)
}

// setup subscription
func (b *NatsEventBus) setup() {
	if b.sub != nil {
		b.sub.Unsubscribe()
	}
	sub, err := b.nc.Subscribe(b.subject, b.handleMsg)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to subscribe")
		return
	}
	b.sub = sub
}

// handleMsg handles a message received on the nats subscription
func (b *NatsEventBus) handleMsg(msg *nats.Msg) {
	if msg.Data == nil {
		log.Warn().Msg("Received empty message")
		return
	}
	// Unmarshal event
	var eIn Event
	err := json.Unmarshal(msg.Data, &eIn)
	if err != nil {
		log.Warn().Err(err).Msg("failed to unmarshal event")
		return
	}

	eMid, err := b.applyMiddleware(&eIn)
	if err != nil {
		return
	}
	var eOut *Event
	if b.hasHandler(eMid) {
		eOut, err = b.applyHandler(eMid)
	}

	if err != nil {
		log.Warn().Err(err).Msg("failed to handle event")
		// sets the error, client should check error
		eOut.Error = err.Error()
	}
	if msg.Reply == "" {
		// nothing to respond to
		return
	}
	if eOut == nil {
		// make sure we have an event to respond with
		// otherwise we have a timeout
		eOut = NewErrorEvent(eIn.Kind, "nil event")
	}
	data, err := json.Marshal(eOut)
	if err != nil {
		log.Warn().Err(err).Msg("failed to marshal event")
		return
	}
	msg.Respond(data)
}

func (b *NatsEventBus) Close() error {
	if b.sub == nil {
		return nil
	}
	return b.sub.Unsubscribe()
}

// Register a handler for a specific event kind
// handler is called when an event arrives.
// To handle publish events, register a handler and don't reply.
// To handle request events, register a handler and reply.
func (b *NatsEventBus) Register(kind string, fn HandlerFunc) {
	b.rw.Lock()
	defer b.rw.Unlock()
	b.handlers[kind] = fn
}

// add middleware, middleware is applied in oder and called for each event
func (b *NatsEventBus) Use(middleware ...HandlerFunc) {
	b.middleware = append(b.middleware, middleware...)
}

// apply middleware to event
func (b *NatsEventBus) applyMiddleware(e *Event) (*Event, error) {
	for _, mw := range b.middleware {
		var err error
		e, err = mw(e)
		if err != nil {
			return nil, err
		}
	}
	return e, nil
}

func (b *NatsEventBus) hasHandler(e *Event) bool {
	b.rw.RLock()
	defer b.rw.RUnlock()
	_, ok := b.handlers[e.Kind]
	return ok
}

func (b *NatsEventBus) applyHandler(e *Event) (*Event, error) {
	b.rw.RLock()
	defer b.rw.RUnlock()
	fn, ok := b.handlers[e.Kind]
	if !ok {
		return nil, nil
	}
	return fn(e)
}

// Request sends an event and waits for a response
func (b *NatsEventBus) Request(e *Event) (*Event, error) {
	data, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	msg, err := b.nc.Request(b.subject, data, NATS_TIMEOUT)
	if err != nil {
		return nil, err
	}
	var eOut Event
	err = json.Unmarshal(msg.Data, &eOut)
	if err != nil {
		return nil, err
	}
	return &eOut, nil
}
