package simjs

import (
	"errors"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
	"github.com/rs/zerolog/log"
)

type LogPrinter struct {
}

var _ console.Printer = (*LogPrinter)(nil)

func (lp *LogPrinter) Log(s string) {
	log.Info().Msg(s)
}

func (lp *LogPrinter) Warn(s string) {
	log.Warn().Msg(s)
}

func (lp *LogPrinter) Error(s string) {
	log.Error().Msg(s)
}

type World struct {
	ID       string
	rt       *goja.Runtime
	registry *require.Registry
	actors   map[string]*Actor
	pub      SimuPublisher
}

func NewWorld(id string, pub SimuPublisher) (*World, error) {
	log.Info().Msgf("creating world %s", id)
	w := &World{
		ID:       id,
		rt:       goja.New(),
		registry: require.NewRegistry(),
		actors:   make(map[string]*Actor),
		pub:      pub,
	}
	w.registry.Enable(w.rt)
	require.RegisterNativeModule(console.ModuleName, func(rt *goja.Runtime, m *goja.Object) {
		console.RequireWithPrinter(&LogPrinter{})(rt, m)
	})
	console.Enable(w.rt)
	w.rt.SetFieldNameMapper(goja.UncapFieldNameMapper())
	err := w.rt.Set("$world", w)
	if err != nil {
		return nil, err
	}
	w.WorldChanged("created")
	return w, nil
}

func (w *World) Publish(subject string, msg *SimuMessage) error {
	return w.pub.Publish(subject, msg)
}

func (w *World) RunScript(name, script string) goja.Value {
	log.Info().Msgf("script %s: %s", name, script)
	log.Info().Msgf("world %s run script %s", w.ID, name)
	v, err := w.rt.RunString(script)
	if err != nil {
		log.Error().Err(err).Msgf("script %s error", name)
		return nil
	}
	log.Info().Msgf("script value %s: %v", name, v)
	w.WorldChanged("running")
	return v
}

func (w *World) StartScript() {
	log.Info().Msgf("world %s start script", w.ID)
	w.WorldChanged("started")
}

func (w *World) StopScript() {
	log.Info().Msgf("world %s stop script", w.ID)
	w.rt.Interrupt(errors.New("world stopped"))
	w.WorldChanged("stopped")
}

func (w *World) CreateActor(actorID string, state map[string]any) *Actor {
	log.Info().Msgf("create actor %s", actorID)
	a, ok := w.actors[actorID]
	if ok {
		return a
	}
	a = NewActor(actorID, state, w)
	w.actors[actorID] = a
	return a
}

func (w *World) GetActor(actorID string) *Actor {
	a, ok := w.actors[actorID]
	if !ok {
		return w.CreateActor(actorID, nil)
	}
	return a
}

func (w *World) RemoveActor(actorID string) {
	delete(w.actors, actorID)
}

func (w *World) Call(name string, args ...any) any {
	fn, ok := goja.AssertFunction(w.rt.Get(name))
	if !ok {
		log.Error().Msgf("call %s not found", name)
		return nil
	}
	values := make([]goja.Value, len(args))
	for i, arg := range args {
		values[i] = w.rt.ToValue(arg)
	}
	result, err := fn(goja.Undefined(), values...)
	if err != nil {
		log.Error().Err(err).Msgf("call %s error", name)
		return nil
	}
	return result.Export()
}

func (w *World) GetFunction(name string) goja.Callable {
	fn, ok := goja.AssertFunction(w.rt.Get(name))
	if !ok {
		return nil
	}
	return fn
}

func (w *World) GetValue(name string) any {
	v := w.rt.Get(name)
	if v == nil {
		return nil
	}
	return v.Export()
}

func (w *World) WorldChanged(value string) {
	msg := &SimuMessage{
		Event:   EWorldChanged,
		WorldID: w.ID,
		Value:   value,
	}
	w.Publish(msg.WorldSubject(), msg)
}

func (w *World) HandleMessage(msg *SimuMessage) *SimuMessage {
	if w.ID != msg.WorldID {
		return nil
	}
	switch msg.Event {
	case EWorldRun:
		name := msg.Meta["name"].(string)
		script := msg.Meta["script"].(string)
		v := w.RunScript(name, script)
		log.Info().Msgf("script %s value %v", name, v)
	case EWorldStop:
		w.StopScript()
	}
	return nil
}
