package simjs

import (
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/nats-io/nats.go"
)

type SimuClient struct {
	conn    *Conn
	worldID string
	events  helper.Hook[*SimuMessage]
}

func NewSimuClient(nc *nats.Conn, worldID string) (*SimuClient, error) {
	c, err := NewConn(nc)
	if err != nil {
		return nil, err
	}
	sc := &SimuClient{
		conn:    c,
		worldID: worldID,
		events:  helper.Hook[*SimuMessage]{},
	}
	return sc, nil
}

// HandleMessage handles incoming messages
func (sc *SimuClient) HandleMessage(msg *SimuMessage) *SimuMessage {
	if msg.WorldID != sc.worldID {
		return nil
	}
	switch msg.Event {
	case EActorPropertyChanged, EActorStateChanged, EWorldChanged:
		sc.events.FireHook(msg)
	}
	return nil
}

func (sc *SimuClient) PropertySet(actorID string, property string, value any) error {
	msg := &SimuMessage{
		Event:   EActorPropertySet,
		WorldID: sc.worldID,
		ActorID: actorID,
		Member:  property,
		Value:   value,
	}
	return sc.conn.Publish(msg.ActorSubject(), msg)
}

func (sc *SimuClient) Call(actorID string, method string, args ...any) (any, error) {
	msg := &SimuMessage{
		Event:   EActorRPCCall,
		WorldID: sc.worldID,
		ActorID: actorID,
		Member:  method,
		Args:    args,
	}
	reply, err := sc.conn.Request(msg.ActorSubject(), msg)
	if err != nil {
		return nil, err
	}
	return reply.Value, nil
}

// PropertyGet returns the value of a property of an actor
func (sc *SimuClient) PropertyGet(actorID string, property string) (any, error) {
	msg := &SimuMessage{
		Event:   EActorPropertyGet,
		WorldID: sc.worldID,
		ActorID: actorID,
		Member:  property,
	}
	reply, err := sc.conn.Request(msg.ActorSubject(), msg)
	if err != nil {
		return nil, err
	}
	return reply.Value, nil
}

func (sc *SimuClient) OnChange(fn func(msg *SimuMessage)) {
	sc.events.AddHook(fn)
}

// Run a world
func (sc *SimuClient) WorldRun(worldID string, name string, content string) error {
	msg := &SimuMessage{
		Event:   EWorldRun,
		WorldID: worldID,
		Meta: map[string]any{
			"name":    name,
			"content": content,
		},
	}
	return sc.conn.Publish(msg.WorldSubject(), msg)
}

// WorldStop stops a world
func (sc *SimuClient) WorldStop(worldID string) error {
	msg := &SimuMessage{
		Event:   EWorldStop,
		WorldID: worldID,
	}
	return sc.conn.Publish(msg.WorldSubject(), msg)
}
