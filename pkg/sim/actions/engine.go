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
	docs []*spec.ScenarioDoc
	core.EventNotifier
	players []*Player
}

func NewEngine() *Engine {
	e := &Engine{
		eval:    NewEval(),
		docs:    make([]*spec.ScenarioDoc, 0),
		players: make([]*Player, 0),
	}
	e.init()
	return e
}

func (e *Engine) init() {
	e.eval.OnEvent(func(evt *core.APIEvent) {
		e.EmitEvent(evt)
	})
}

func (e *Engine) LoadScenario(source string, doc *spec.ScenarioDoc) error {
	doc.Source = source
	e.docs = append(e.docs, doc)
	for _, iface := range doc.Interfaces {
		if iface.Name == "" {
			return fmt.Errorf("interface %v has no name", iface)
		}
		log.Info().Msgf("registering interface %s", iface.Name)
	}
	for _, seq := range doc.Sequences {
		log.Info().Msgf("registering sequence %s", seq.Name)
		if seq.Interface == "" {
			return fmt.Errorf("sequence %v has no interface", seq)
		}
		iface := e.GetInterface(seq.Interface)
		if iface == nil {
			return fmt.Errorf("interface %s not found", seq.Interface)
		}
		p := NewPlayer(iface, seq)
		go func() {
			for frame := range p.FramesC {
				iface := frame.Interface
				if iface != nil {
					_, err := e.eval.EvalAction(iface.Name, frame.Action, iface.Properties)
					if err != nil {
						log.Error().Msgf("eval action %s: %v", frame.Action, err)
					}
				}
			}
			log.Info().Msgf("sequence %s stopped", seq.Name)
		}()
		e.players = append(e.players, p)
	}
	return nil
}

func (a *Engine) UnloadScenario(source string) error {
	for i, d := range a.docs {
		if d.Source == source {
			a.docs = append(a.docs[:i], a.docs[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("scenario %s not found", source)
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
	log.Info().Msgf("%s/%s invoke", symbol, name)
	e.EmitCall(symbol, name, args)

	iface := e.GetInterface(symbol)
	if iface == nil {
		return nil, fmt.Errorf("interface %s not found", symbol)
	}
	op := iface.GetOperation(name)
	if op == nil {
		return nil, fmt.Errorf("operation %s not found", name)
	}
	result, err := e.eval.EvalActions(symbol, op.Actions, iface.Properties)
	if err != nil {
		return nil, err
	}
	log.Info().Msgf("%s/%s result %v", symbol, name, result)
	return result, nil
}

// SetProperties sets the properties of the interface.
func (e *Engine) SetProperties(symbol string, props map[string]any) error {
	e.EmitPropertySet(symbol, props)
	iface := e.GetInterface(symbol)
	if iface == nil {
		return fmt.Errorf("interface %s not found", symbol)
	}
	for name, value := range props {
		iface.Properties[name] = value
	}
	return nil
}

// FetchProperties returns a copy of the properties of the interface.
func (e *Engine) GetProperties(symbol string) (map[string]any, error) {
	iface := e.GetInterface(symbol)
	if iface == nil {
		return nil, fmt.Errorf("interface %s not found", symbol)
	}
	return iface.Properties, nil
}

func (e *Engine) HasSequence(name string) bool {
	for _, d := range e.docs {
		if d.GetSequence(name) != nil {
			return true
		}
	}
	return false
}

func (e *Engine) PlayAllSequences() error {
	log.Info().Msgf("actions engine play all sequences")
	for _, p := range e.players {
		err := p.Play()
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Engine) StopAllSequences() {
	log.Info().Msgf("actions engine stop all sequences")
	for _, p := range e.players {
		err := p.Stop()
		if err != nil {
			log.Warn().Msgf("stop sequence %s: %v", p.SequenceName(), err)
		}
	}
}

func (e *Engine) PlaySequence(name string) error {
	for _, p := range e.players {
		if p.SequenceName() == name {
			return p.Play()
		}
	}
	return fmt.Errorf("sequence %s not found", name)
}

func (e *Engine) StopSequence(name string) {
	for _, p := range e.players {
		if p.SequenceName() == name {
			err := p.Stop()
			if err != nil {
				log.Warn().Msgf("stop sequence %s: %v", name, err)
			}
		}
	}
	log.Warn().Msgf("sequence %s not found", name)
}
