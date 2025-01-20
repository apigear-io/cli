package js

import (
	"errors"
	"fmt"
	"sync"
	"time"

	_ "embed"

	"github.com/apigear-io/cli/pkg/sim/model"
	"github.com/apigear-io/cli/pkg/sim/tools"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
	"github.com/rs/zerolog/log"
)

//go:embed actor.js
var actorJS string

type Runtime struct {
	lock       sync.RWMutex
	vm         *goja.Runtime
	world      *World
	LastUpdate time.Time
	IsActive   bool
}

func NewRuntime(id string) *Runtime {
	vm := goja.New()
	world := NewWorld(id, vm)
	r := &Runtime{
		vm:    vm,
		world: world,
	}
	r.register()
	return r
}

func (si *Runtime) Hooks() *tools.Hook[model.SimEvent] {
	return &si.world.hooks
}

func (si *Runtime) register() {
	si.lock.Lock()
	defer si.lock.Unlock()

	registry := require.NewRegistry()
	registry.Enable(si.vm)

	require.RegisterNativeModule("console", console.RequireWithPrinter(&LogPrinter{}))
	console.Enable(si.vm)
	si.vm.SetFieldNameMapper(goja.UncapFieldNameMapper())

	err := si.vm.GlobalObject().Set("$world", si.world)
	if err != nil {
		log.Error().Err(err).Msg("failed to set world")
	}
	_, err = si.vm.RunScript("actor.js", actorJS)
	if err != nil {
		log.Error().Err(err).Msg("failed to run actor.js")
	}
}

// GetWorld returns the world
func (si *Runtime) GetWorld() *World {
	si.lock.RLock()
	defer si.lock.RUnlock()
	return si.world
}

// GetActor returns the actor
func (si *Runtime) GetActor(id string) *Actor {
	si.lock.RLock()
	defer si.lock.RUnlock()
	w := si.world.GetActor(id)
	if w == nil {
		log.Error().Str("actor", id).Msg("actor not found")
		return nil
	}
	return w
}

func (si *Runtime) RunScript(script model.Script) (any, error) {
	log.Info().Str("script", script.Name).Msg("running script")
	si.lock.Lock()
	defer si.lock.Unlock()
	si.IsActive = true
	si.LastUpdate = time.Now()

	v, err := si.vm.RunScript(script.Name, script.Source)
	if err != nil {
		return nil, err
	}
	_, ok := goja.AssertFunction(v)
	if ok {
		v = goja.Undefined()
	}
	return v.Export(), nil
}

func (si *Runtime) Interupt() {
	si.lock.Lock()
	defer si.lock.Unlock()
	si.IsActive = false
	si.vm.Interrupt(errors.New("interrupted"))
}

// RunFunction runs a function by name in the simulation
func (si *Runtime) RunFunction(name string, args ...any) (any, error) {
	log.Info().Str("function", name).Interface("args", args).Msg("running function")
	si.lock.Lock()
	defer si.lock.Unlock()
	log.Info().Interface("functions", si.vm.GlobalObject().Keys()).Msg("functions")

	log.Info().Any("fn", si.vm.GlobalObject().Get(name)).Msg("function")
	fn, ok := goja.AssertFunction(si.vm.Get(name))
	if !ok {
		log.Error().Str("function", name).Msg("function not found")
		return nil, fmt.Errorf("function %s not found", name)
	}
	jsArgs := make([]goja.Value, len(args))
	for i, arg := range args {
		jsArgs[i] = si.vm.ToValue(arg)
	}
	v, err := fn(goja.Undefined(), jsArgs...)
	if err != nil {
		return nil, err
	}
	si.world.fireEvent(model.EventWorldCall, name, "", map[string]any{
		"args":   args,
		"result": v,
	})
	return v.Export(), nil
}

func (si *Runtime) MapToJsObject(m map[string]any) *goja.Object {
	si.lock.Lock()
	defer si.lock.Unlock()
	jsObj := si.vm.ToValue(m).ToObject(si.vm)
	return jsObj
}

func (si *Runtime) ArgsToJsArgs(args []any) []goja.Value {
	si.lock.Lock()
	defer si.lock.Unlock()
	jsArgs := make([]goja.Value, len(args))
	for i, arg := range args {
		jsArgs[i] = si.vm.ToValue(arg)
	}
	return jsArgs
}

func (si *Runtime) JsObjectToMap(jsObj *goja.Object) map[string]any {
	si.lock.Lock()
	defer si.lock.Unlock()
	var m map[string]any
	si.vm.ExportTo(jsObj, &m)
	return m
}

func (rt *Runtime) ToJsValue(v any) goja.Value {
	rt.lock.Lock()
	defer rt.lock.Unlock()
	return rt.vm.ToValue(v)
}

// implements ISimulation interface

func (rt *Runtime) InvokeOperation(actorId string, method string, args []any) (any, error) {
	rt.lock.Lock()
	defer rt.lock.Unlock()
	actor := rt.GetActor(actorId)
	if actor == nil {
		return nil, fmt.Errorf("actor %s not found", actorId)
	}

	jsValue, err := actor.CallMethod(method, rt.ArgsToJsArgs(args)...)
	if err != nil {
		return nil, err
	}
	return jsValue.Export(), nil
}

// SetProperties sets the state of an actor
func (rt *Runtime) SetProperties(actorId string, state map[string]any) error {
	rt.lock.Lock()
	defer rt.lock.Unlock()
	actor := rt.GetActor(actorId)
	if actor == nil {
		return fmt.Errorf("actor %s not found", actorId)
	}
	actor.SetState(rt.MapToJsObject(state))
	return nil
}

// GetProperties returns the state of an actor
func (rt *Runtime) GetProperties(actorId string) (map[string]any, error) {
	rt.lock.Lock()
	defer rt.lock.Unlock()
	actor := rt.GetActor(actorId)
	if actor == nil {
		return nil, fmt.Errorf("actor %s not found", actorId)
	}
	return rt.JsObjectToMap(actor.GetState()), nil
}

// OnEvent registers a callback for simulation events
func (rt *Runtime) OnEvent(fn func(evt *model.SimEvent)) {
	rt.Hooks().Add(fn)
}
