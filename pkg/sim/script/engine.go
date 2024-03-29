package script

import (
	"context"
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/sim/core"
	"github.com/apigear-io/cli/pkg/sim/ostore"

	js "github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
)

// engine implements core.IEngine interface
var _ core.IEngine = (*Engine)(nil)

type Engine struct {
	store      ostore.IObjectStore
	vm         *js.Runtime
	interfaces map[string]*js.Object
	sequencers map[string]*js.Object
	core.EventNotifier
}

func NewEngine(store ostore.IObjectStore) *Engine {
	e := &Engine{
		store:      store,
		vm:         js.New(),
		interfaces: map[string]*js.Object{},
		sequencers: map[string]*js.Object{},
	}
	e.init()
	return e
}

func (e *Engine) LoadScript(name string, script string) (any, error) {
	v, err := e.vm.RunScript(name, script)
	if err != nil {
		return nil, err
	}
	return v.Export(), nil
}

func (e *Engine) HasInterface(symbol string) bool {
	_, ok := e.interfaces[symbol]
	return ok
}

func (e *Engine) InvokeOperation(symbol, name string, args []any) (any, error) {
	log.Info().Msgf("%s/%s invoke", symbol, name)
	obj := e.interfaces[symbol]
	if obj == nil {
		return nil, fmt.Errorf("interface %s not found", symbol)
	}
	m, ok := js.AssertFunction(obj.Get(name))
	if !ok {
		return nil, fmt.Errorf("operation %s not found", name)
	}
	jsArgs := e.vm.ToValue(args)
	v, err := m(obj, jsArgs)
	if err != nil {
		log.Warn().Msgf("InvokeOperation: %s", err)
		return nil, err
	}
	result := v.Export()
	log.Info().Msgf("%s/%s result: %v", symbol, name, result)
	return result, nil
}

func (e *Engine) SetProperties(symbol string, props map[string]any) error {
	obj := e.interfaces[symbol]
	if obj == nil {
		return fmt.Errorf("interface %s not found", symbol)
	}
	for name, value := range props {
		err := obj.Set(name, e.vm.ToValue(value))
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Engine) GetProperties(symbol string) (map[string]any, error) {
	props := map[string]any{}
	obj := e.interfaces[symbol]
	if obj == nil {
		return props, fmt.Errorf("interface %s not found", symbol)
	}
	for _, name := range obj.Keys() {
		if strings.HasPrefix(name, "_") || strings.HasPrefix(name, "$") {
			// skip private properties and special properties
			continue
		}
		v := obj.Get(name)
		_, ok := js.AssertFunction(v)
		if ok {
			// skip functions
			continue
		}
		props[name] = v.Export()
	}
	return props, nil
}

func (e *Engine) HasSequence(sequenceId string) bool {
	_, ok := e.sequencers[sequenceId]
	return ok
}

func (e *Engine) PlaySequence(ctx context.Context, sequenceId string) error {
	obj := e.sequencers[sequenceId]
	if obj == nil {
		return fmt.Errorf("sequence %s not found", sequenceId)
	}
	jsSteps := obj.Get("steps").Export().([]any)
	log.Info().Msgf("PlaySequencer: %d steps", len(jsSteps))
	return nil
}

func (e *Engine) StopSequence(sequenceId string) error {
	log.Info().Msgf("StopSequencer: %s", sequenceId)
	return nil
}

func (e *Engine) init() {
	registry := new(require.Registry)
	registry.Enable(e.vm)
	console.Enable(e.vm)
	err := e.vm.Set("$registerInterface", e.jsRegisterInterface)
	if err != nil {
		log.Error().Msgf("Set $registerInterface: %s", err)
	}
	err = e.vm.Set("$signal", e.jsSignal)
	if err != nil {
		log.Error().Msgf("Set $signal: %s", err)
	}
	err = e.vm.Set("$change", e.jsChange)
	if err != nil {
		log.Error().Msgf("Set $change: %s", err)
	}
	err = e.vm.Set("$registerSequence", e.jsRegisterSequence)
	if err != nil {
		log.Error().Msgf("Set $registerSequence: %s", err)
	}
}

func (e *Engine) PlayAllSequences(ctx context.Context) error {
	log.Debug().Msgf("script engine play all sequences")
	return nil
}

func (e *Engine) StopAllSequences() {
	log.Debug().Msgf("script engine stop all sequences")
}
