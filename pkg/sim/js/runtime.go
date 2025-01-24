package js

import (
	"fmt"
	"time"

	_ "embed"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/sim/model"
	"github.com/apigear-io/cli/pkg/tools"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/sasha-s/go-deadlock"
)

//go:embed actor.js
var actorJS string

type Runtime struct {
	lock       deadlock.Mutex
	vm         *goja.Runtime
	loop       *eventloop.EventLoop
	world      *World
	LastUpdate time.Time
	IsActive   bool
}

var _ model.SimulationAPI = (*Runtime)(nil)

func NewRuntime(id string) *Runtime {
	if id == "" {
		id = "demo"
	}
	log.Debug().Str("id", id).Msg("new runtime")
	loop := eventloop.NewEventLoop()
	var outerWorld *World
	var outerVm *goja.Runtime
	// initial run wait for result
	loop.Run(func(vm *goja.Runtime) {
		log.Debug().Msg("js event loop initial run")
		world := NewWorld(id, vm)
		vm.SetFieldNameMapper(goja.UncapFieldNameMapper())
		vm.Set("$world", world)
		vm.RunScript("actor.js", actorJS)
		outerWorld = world
		outerVm = vm
	})
	// run in background
	loop.Start()
	r := &Runtime{
		loop:  loop,
		world: outerWorld,
		vm:    outerVm,
	}
	return r
}

func (rt *Runtime) Hooks() *tools.Hook[model.SimEvent] {
	return rt.world.hooks
}

// GetWorld returns the world
func (rt *Runtime) GetWorld() *World {
	rt.lock.Lock()
	defer rt.lock.Unlock()
	return rt.world
}

// GetActor returns the actor
func (rt *Runtime) GetActor(id string) *Actor {
	rt.lock.Lock()
	defer rt.lock.Unlock()
	w := rt.world.GetActor(id)
	if w == nil {
		log.Error().Str("actor", id).Msg("actor not found")
		return nil
	}
	return w
}

func (rt *Runtime) RunScript(script model.Script) (any, error) {
	log.Info().Str("script", script.Name).Msg("running script")
	rt.IsActive = true
	rt.LastUpdate = time.Now()

	// loop is already running, just add the function to the queue
	rt.loop.RunOnLoop(func(vm *goja.Runtime) {
		_, err := rt.vm.RunScript(script.Name, script.Source)
		if err != nil {
			log.Error().Err(err).Msg("run script")
			return
		}
	})

	return nil, nil
}

func (si *Runtime) Interupt() {
	si.IsActive = false
	si.loop.Stop()
	si.vm.Interrupt(fmt.Errorf("interrupted"))
}

// RunFunction runs a function by name in the simulation
func (rt *Runtime) RunFunction(name string, args ...any) (any, error) {
	log.Info().Str("function", name).Interface("args", args).Msg("running function")
	// loop is already running, just add the function to the queue
	rt.loop.RunOnLoop(func(vm *goja.Runtime) {
		fn, ok := goja.AssertFunction(vm.Get(name))
		if !ok {
			log.Error().Str("function", name).Msg("function not found")
		}
		jsArgs := make([]goja.Value, len(args))
		for i, arg := range args {
			jsArgs[i] = vm.ToValue(arg)
		}
		v, err := fn(goja.Undefined(), jsArgs...)
		if err != nil {
			return
		}
		rt.world.fireEvent(model.EventWorldCall, name, "", map[string]any{
			"args":   args,
			"result": v,
		})
	})
	return nil, nil
}

func (rt *Runtime) MapToJsObject(m map[string]any) *goja.Object {
	rt.lock.Lock()
	defer rt.lock.Unlock()
	jsObj := rt.vm.ToValue(m).ToObject(rt.vm)
	return jsObj
}

func (rt *Runtime) ArgsToJsArgs(args []any) []goja.Value {
	jsArgs := make([]goja.Value, len(args))
	rt.lock.Lock()
	defer rt.lock.Unlock()
	for i, arg := range args {
		jsArgs[i] = rt.vm.ToValue(arg)
	}
	return jsArgs
}

func (rt *Runtime) JsObjectToMap(jsObj *goja.Object) map[string]any {
	var m map[string]any
	rt.lock.Lock()
	defer rt.lock.Unlock()
	rt.vm.ExportTo(jsObj, &m)
	return m
}

func (rt *Runtime) ToJsValue(v any) goja.Value {
	rt.lock.Lock()
	defer rt.lock.Unlock()
	return rt.vm.ToValue(v)
}

// implements ISimulation interface

func (rt *Runtime) InvokeOperation(actorId string, method string, args []any) (any, error) {
	log.Info().Str("actor", actorId).Str("method", method).Interface("args", args).Str("id", rt.world.id).Msg("invoking operation")
	rt.lock.Lock()
	defer rt.lock.Unlock()

	actor := rt.world.GetActor(actorId)
	if actor == nil {
		return nil, fmt.Errorf("actor %s not found", actorId)
	}

	jsArgs := make([]goja.Value, len(args))
	for i, arg := range args {
		jsArgs[i] = rt.vm.ToValue(arg)
	}
	jsValue, err := actor.CallMethod(method, jsArgs...)
	if err != nil {
		return nil, err
	}
	return jsValue.Export(), nil
}

// SetProperties sets the state of an actor
func (rt *Runtime) SetProperties(actorId string, state map[string]any) error {
	log.Info().Str("actor", actorId).Msg("setting properties")
	rt.lock.Lock()
	defer rt.lock.Unlock()
	actor := rt.world.GetActor(actorId)
	if actor == nil {
		return fmt.Errorf("actor %s not found", actorId)
	}
	jsObj := rt.vm.ToValue(state).ToObject(rt.vm)
	actor.SetState(jsObj)
	return nil
}

// GetProperties returns the state of an actor
func (rt *Runtime) GetProperties(actorId string) (map[string]any, error) {
	log.Info().Str("actor", actorId).Msg("getting properties")
	var state map[string]any
	rt.lock.Lock()
	defer rt.lock.Unlock()
	actor := rt.world.GetActor(actorId)
	if actor == nil {
		return nil, fmt.Errorf("actor %s not found", actorId)
	}
	rt.vm.ExportTo(actor.GetState(), &state)
	return state, nil
}

// OnEvent registers a callback for simulation events
func (rt *Runtime) OnEvent(fn func(evt *model.SimEvent)) func() {
	log.Info().Msg("on event")
	return rt.Hooks().Add(fn)
}
