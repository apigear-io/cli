package simjs

import (
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/objectlink-core-go/log"
)

type Actor struct {
	id      string
	world   *World
	state   map[string]any
	changes helper.Hook[any]
	methods map[string]func(args ...any) any
	signals map[string]func(args ...any)
	events  helper.Hook[SimuMessage]
}

func NewActor(id string, state map[string]any, world *World) *Actor {
	if state == nil {
		state = make(map[string]any)
	}
	return &Actor{
		id:      id,
		state:   state,
		world:   world,
		changes: helper.Hook[any]{},
		methods: make(map[string]func(args ...any) any),
		signals: make(map[string]func(args ...any)),
		events:  helper.Hook[SimuMessage]{},
	}
}

func (a *Actor) ID() string {
	return a.id
}

func (a *Actor) World() *World {
	return a.world
}

func (a *Actor) OnEvent(fn func(SimuMessage)) {
	a.events.AddHook(fn)
}

func (a *Actor) emitMessage(e SimuEvent, member string, value any) {
	a.events.FireHook(SimuMessage{
		Event:   e,
		ActorID: a.id,
		WorldID: a.world.ID,
		Member:  member,
		Value:   value,
	})
}

func (a *Actor) Set(prop string, value any) {
	a.state[prop] = value
	a.emitMessage(EActorPropertySet, prop, value)
	a.EmitChange(prop, value)
}

func (a *Actor) Get(prop string) any {
	return a.state[prop]
}

func (a *Actor) OnChange(prop string, fn func(value any)) {
	a.changes.AddHook(fn)
}

func (a *Actor) EmitChange(prop string, value any) {
	a.changes.FireHook(value)
	a.emitMessage(EActorPropertyChanged, prop, value)
}

// state handling
func (a *Actor) SetState(properties map[string]any) {
	for k, v := range properties {
		a.Set(k, v)
	}
}

func (a *Actor) State() map[string]any {
	return a.state
}

// Method Handling

func (a *Actor) Call(name string, args ...any) any {
	fn := a.methods[name]
	if fn == nil {
		return nil
	}
	return fn(args)
}

func (a *Actor) OnMethod(name string, fn func(args ...any) any) {
	a.methods[name] = fn
}

func (a *Actor) Method(name string) func(args ...any) any {
	return a.methods[name]
}

// Signal Handling

func (a *Actor) OnSignal(name string, fn func(args ...any)) {
	a.signals[name] = fn
}

func (a *Actor) EmitSignal(name string, args ...any) {
	fn := a.signals[name]
	if fn == nil {
		return
	}
	fn(args)
}

func (a *Actor) HandleMessage(msg *SimuMessage) *SimuMessage {
	if msg.WorldID != a.world.ID || msg.ActorID != a.id {
		log.Warn().Msgf("message not for world %s or actor %s", a.world.ID, a.id)
		return nil
	}
	switch msg.Event {
	case EActorPropertySet:
		a.Set(msg.Member, msg.Value)
		a.EmitChange(msg.Member, msg.Value)
		return nil
	case EActorPropertyGet:
		msg.Value = a.Get(msg.Member)
		return msg
	case EActorPropertyChanged:
		// DO NOTHING ON SERVICE SIDE
		return nil
	case EActorStateSet:
		a.SetState(msg.KWArgs)
		return nil
	case EActorStateGet:
		msg.KWArgs = a.State()
		return msg
	case EActorStateChanged:
		// DO NOTHING ON SERVICE SIDE
		return nil
	case EActorSignal:
		// DO NOTHING ON SERVICE SIDE
		return nil
	case EActorRPCCall:
		msg.Value = a.Call(msg.Member, msg.Args...)
		return msg
	default:
		log.Warn().Msgf("unknown message event %s", msg.Event)
		return nil
	}

}
