package actions

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/sim/core"
	"github.com/apigear-io/cli/pkg/spec"
)

// engine implements core.IEngine interface
var _ core.IEngine = (*Engine)(nil)

type Engine struct {
	eval *eval
	docs map[string]*spec.ScenarioDoc
	core.Notifier
}

func NewEngine() *Engine {
	e := &Engine{
		eval: NewEval(),
		docs: make(map[string]*spec.ScenarioDoc),
	}
	e.init()
	return e
}

func (e *Engine) init() {
	e.eval.OnChange(func(symbol string, name string, value any) {
		e.EmitOnChange(symbol, name, value)
	})
	e.eval.OnSignal(func(symbol string, name string, args map[string]any) {
		e.EmitOnSignal(symbol, name, args)
	})
}

func (e *Engine) LoadScenario(doc *spec.ScenarioDoc) error {
	e.docs[doc.Name] = doc
	for _, s := range doc.Interfaces {
		if s.Name == "" {
			return fmt.Errorf("interface %v has no name", s)
		}
		log.Infof("registering interface %s\n", s.Name)
	}
	return nil
}

func (e *Engine) HasInterface(ifaceId string) bool {
	for _, d := range e.docs {
		if d.GetInterface(ifaceId) != nil {
			return true
		}
	}
	return false
}

func (e *Engine) GetInterface(ifaceId string) *spec.InterfaceEntry {
	for _, d := range e.docs {
		if s := d.GetInterface(ifaceId); s != nil {
			return s
		}
	}
	return nil
}

// InvokeOperation invokes a operation of the interface.
func (e *Engine) InvokeOperation(symbol string, name string, args map[string]any) (any, error) {
	log.Infof("%s/%s invoke\n", symbol, name)
	s := e.GetInterface(symbol)
	if s == nil {
		return nil, fmt.Errorf("interface %s not found", symbol)
	}
	m := s.GetOperation(name)
	if m == nil {
		return nil, fmt.Errorf("operation %s not found", name)
	}
	result, err := e.eval.EvalActions(symbol, m.Actions, s.Properties)
	if err != nil {
		return nil, err
	}
	log.Infof("%s/%s result %v\n", symbol, name, result)
	return result, nil
}

// SetProperties sets the properties of the interface.
func (e *Engine) SetProperties(symbol string, props map[string]any) error {
	s := e.GetInterface(symbol)
	if s == nil {
		return fmt.Errorf("interface %s not found", symbol)
	}
	for name, value := range props {
		s.Properties[name] = value
	}
	return nil
}

// FetchProperties returns a copy of the properties of the interface.
func (e *Engine) GetProperties(symbol string) (map[string]any, error) {
	s := e.GetInterface(symbol)
	if s == nil {
		return nil, fmt.Errorf("interface %s not found", symbol)
	}
	return s.Properties, nil
}

func (e *Engine) HasSequence(sequencerId string) bool {
	for _, d := range e.docs {
		if d.GetSequence(sequencerId) != nil {
			return true
		}
	}
	return false
}

func (e *Engine) PlaySequence(sequencerId string) {
	for _, d := range e.docs {
		if s := d.GetSequence(sequencerId); s != nil {
			log.Printf("playing sequencer %s", sequencerId)
		}
	}
}
