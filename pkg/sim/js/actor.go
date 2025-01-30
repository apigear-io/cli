package js

import (
	"errors"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/sim/model"
	"github.com/apigear-io/cli/pkg/tools"
	"github.com/dop251/goja"
)

type Actor struct {
	id             string
	world          *World
	vm             *goja.Runtime
	state          *goja.Object
	stateEmitter   *Emitter
	methods        map[string]goja.Callable
	methodsEmitter *Emitter
	signalsEmitter *Emitter
	hooks          tools.Hook[model.SimEvent]
}

func NewActor(id string, state *goja.Object, world *World) (*Actor, error) {
	if world == nil {
		return nil, errors.New("actor must have a world")
	}
	log.Info().Str("actor", id).Str("world", world.id).Msg("new actor")
	if state == nil {
		state = world.vm.NewObject()
	}
	if id == "" {
		return nil, errors.New("actor id cannot be empty")
	}
	a := &Actor{
		id:             id,
		world:          world,
		vm:             world.vm,
		state:          state,
		stateEmitter:   NewEmitter(world.vm),
		methods:        map[string]goja.Callable{},
		methodsEmitter: NewEmitter(world.vm),
		signalsEmitter: NewEmitter(world.vm),
	}

	// connect actor hooks to world hooks
	world.hooks.Connect(&a.hooks)
	return a, nil
}

// Id_ returns the actor's ID
func (a *Actor) Id() string {
	return a.id
}

///////////////////////////////////////////////////////////////////////////////
// Properties
///////////////////////////////////////////////////////////////////////////////

// SetProperty sets the value of a property
func (a *Actor) SetProperty(key string, value goja.Value) {
	a.state.Set(key, value)
	a.stateEmitter.Emit(key, value)
	a.fireEvent(model.EventActorChanged, key, map[string]any{
		"value": value,
	})
}

// GetProperty gets the value of a property
func (a *Actor) GetProperty(key string) goja.Value {
	return a.state.Get(key)
}

// HasProperty checks if the actor has a value
func (a *Actor) HasProperty(key string) bool {
	value := a.GetProperty(key)
	return value != nil && !goja.IsUndefined(value) && !goja.IsNull(value)
}

// OnProperty registers a handler for a value
func (a *Actor) OnProperty(key string, fn goja.Callable) goja.Callable {
	return a.stateEmitter.On(key, fn)
}

// SetState sets the state of the actor
func (a *Actor) SetState(state *goja.Object) {
	for _, k := range state.Keys() {
		a.SetProperty(k, state.Get(k))
	}
}

// GetState gets the state of the actor
func (a *Actor) GetState() *goja.Object {
	return a.state
}

// Properties returns the list of properties
func (a *Actor) Properties() []string {
	return a.state.Keys()
}

///////////////////////////////////////////////////////////////////////////////
// Methods
///////////////////////////////////////////////////////////////////////////////

// SetMethod sets the method of the actor
func (a *Actor) SetMethod(method string, fn goja.Callable) {
	log.Debug().Str("actor", a.id).Str("method", method).Msg("actor set method")
	a.methods[method] = fn
}

// GetMethod gets the method of the actor
func (a *Actor) GetMethod(method string) goja.Callable {
	return a.methods[method]
}

// HasMethod checks if the actor has a method
func (a *Actor) HasMethod(method string) bool {
	return a.methods[method] != nil
}

// OnMethod registers a handler for a method
func (a *Actor) OnMethod(method string, fn goja.Callable) goja.Callable {
	return a.methodsEmitter.On(method, fn)
}

// Methods returns the list of methods
func (a *Actor) Methods() []string {
	methods := make([]string, 0, len(a.methods))
	for k := range a.methods {
		methods = append(methods, k)
	}
	return methods
}

// CallMethod calls the method of the actor
func (a *Actor) CallMethod(method string, args ...goja.Value) (goja.Value, error) {
	log.Info().Str("actor", a.id).Str("method", method).Interface("args", args).Msg("call method")
	fn, ok := a.methods[method]
	if !ok || fn == nil {
		return goja.Undefined(), errors.New("method not found")
	}
	v, err := fn(a.vm.ToValue(a), args...)
	if err != nil {
		return goja.Undefined(), err
	}
	a.methodsEmitter.Emit(method, args...)
	a.fireEvent(model.EventActorCall, method, map[string]any{
		"args":   args,
		"result": v,
	})
	return v, nil
}

///////////////////////////////////////////////////////////////////////////////
// Signals
///////////////////////////////////////////////////////////////////////////////

// EmitSignal emits a signal to the actor
func (a *Actor) EmitSignal(signal string, args ...goja.Value) {
	a.signalsEmitter.Emit(signal, args...)
	a.fireEvent(model.EventActorSignal, signal, map[string]any{
		"args": args,
	})
}

// OnSignal registers a handler for a signal
func (a *Actor) OnSignal(signal string, fn goja.Callable) goja.Callable {
	return a.signalsEmitter.On(signal, fn)
}

// fireEvent fires an event to the actor
func (a *Actor) fireEvent(mt model.EventType, member string, data any) {
	event := &model.SimEvent{
		Type:   mt,
		World:  a.world.id,
		Actor:  a.id,
		Member: member,
		Data:   data,
	}
	a.hooks.Fire(event)
}
