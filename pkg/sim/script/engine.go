package script

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/sim/core"

	js "github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
)

// engine implements core.IEngine interface
var _ core.IEngine = (*Engine)(nil)

type Engine struct {
	vm         *js.Runtime
	interfaces map[string]*js.Object
	sequencers map[string]*js.Object
	core.Notifier
}

func NewEngine() *Engine {
	s := &Engine{
		vm:         js.New(),
		interfaces: map[string]*js.Object{},
		sequencers: map[string]*js.Object{},
	}
	s.init()
	return s
}

func (s *Engine) LoadScript(name string, script string) (any, error) {
	v, err := s.vm.RunScript(name, script)
	if err != nil {
		return nil, err
	}
	return v.Export(), nil
}

func (s *Engine) HasInterface(symbol string) bool {
	_, ok := s.interfaces[symbol]
	return ok
}

func (s *Engine) InvokeOperation(symbol, name string, args map[string]any) (any, error) {
	log.Info().Msgf("%s/%s invoke", symbol, name)
	obj := s.interfaces[symbol]
	if obj == nil {
		return nil, fmt.Errorf("interface %s not found", symbol)
	}
	m, ok := js.AssertFunction(obj.Get(name))
	if !ok {
		return nil, fmt.Errorf("operation %s not found", name)
	}
	jsArgs := s.vm.ToValue(args)
	v, err := m(obj, jsArgs)
	if err != nil {
		log.Warn().Msgf("InvokeOperation: %s", err)
		return nil, err
	}
	result := v.Export()
	log.Info().Msgf("%s/%s result: %v", symbol, name, result)
	return result, nil
}

func (s *Engine) SetProperties(symbol string, props map[string]any) error {
	obj := s.interfaces[symbol]
	if obj == nil {
		return fmt.Errorf("interface %s not found", symbol)
	}
	for name, value := range props {
		err := obj.Set(name, s.vm.ToValue(value))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Engine) GetProperties(symbol string) (map[string]any, error) {
	props := map[string]any{}
	obj := s.interfaces[symbol]
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

func (e *Engine) PlaySequence(sequenceId string) error {
	obj := e.sequencers[sequenceId]
	if obj == nil {
		return fmt.Errorf("sequence %s not found", sequenceId)
	}
	jsSteps := obj.Get("steps").Export().([]any)
	log.Info().Msgf("PlaySequencer: %d steps", len(jsSteps))
	return nil
}

func (e *Engine) StopSequence(sequenceId string) {
	log.Info().Msgf("StopSequencer: %s", sequenceId)
}

func (s *Engine) init() {
	registry := new(require.Registry)
	registry.Enable(s.vm)
	console.Enable(s.vm)
	err := s.vm.Set("$registerInterface", s.jsRegisterInterface)
	if err != nil {
		log.Error().Msgf("Set $registerInterface: %s", err)
	}
	err = s.vm.Set("$signal", s.jsSignal)
	if err != nil {
		log.Error().Msgf("Set $signal: %s", err)
	}
	err = s.vm.Set("$change", s.jsChange)
	if err != nil {
		log.Error().Msgf("Set $change: %s", err)
	}
	err = s.vm.Set("$registerSequence", s.jsRegisterSequence)
	if err != nil {
		log.Error().Msgf("Set $registerSequence: %s", err)
	}
}

func (s *Engine) PlayAllSequences() error {
	log.Info().Msgf("script engine play all sequences")
	return nil
}

func (e *Engine) StopAllSequences() {
	log.Info().Msgf("script engine stop all sequences")
}
