package actions

import (
	"context"
	"fmt"

	"github.com/apigear-io/cli/pkg/sim/core"
	"github.com/apigear-io/cli/pkg/sim/ostore"
	"github.com/apigear-io/cli/pkg/spec"
)

// engine implements core.IEngine interface
var _ core.IEngine = (*Engine)(nil)

type Engine struct {
	store ostore.IObjectStore
	eval  *eval
	docs  []*spec.ScenarioDoc
	core.EventNotifier
	players []*Player
}

func NewEngine(store ostore.IObjectStore) *Engine {
	eval := NewEval(store)
	e := &Engine{
		store:   store,
		eval:    eval,
		docs:    make([]*spec.ScenarioDoc, 0),
		players: make([]*Player, 0),
	}
	eval.OnEvent(func(evt *core.SimuEvent) {
		e.EmitEvent(evt)
	})
	return e
}

func (e *Engine) LoadScenario(source string, doc *spec.ScenarioDoc) error {
	doc.Source = source
	for _, doc := range e.docs {
		if doc.Source == source {
			err := e.UnloadScenario(source)
			if err != nil {
				log.Error().Err(err).Str("source", source).Msg("unload scenario")
			}
		}
	}
	e.docs = append(e.docs, doc)
	for _, iface := range doc.Interfaces {
		if iface.Name == "" {
			return fmt.Errorf("interface %v has no name", iface)
		}
		log.Debug().Msgf("registering interface %s", iface.Name)
	}
	for _, seq := range doc.Sequences {
		log.Debug().Msgf("registering sequence %s", seq.Name)
		if seq.Interface == "" {
			return fmt.Errorf("sequence %v has no interface", seq)
		}
		iface := e.GetInterface(seq.Interface)
		if iface == nil {
			return fmt.Errorf("interface %s not found", seq.Interface)
		}
		p := NewPlayer(iface, seq)
		go func(framesC chan PlayFrame) {
			for frame := range framesC {
				iface := frame.Interface
				if iface != nil {
					_, err := e.eval.EvalAction(iface.Name, frame.Action)
					if err != nil {
						log.Error().Msgf("eval action %s: %v", frame.Action, err)
					}
				}
			}
			log.Debug().Msgf("sequence %s stopped", seq.Name)
		}(p.FramesC)
		e.players = append(e.players, p)
	}
	return nil
}

func (a *Engine) UnloadScenario(source string) error {
	// make sure all players are stopped
	a.StopAllSequences()
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
func (e *Engine) InvokeOperation(ifaceId string, opName string, args []any) (any, error) {
	log.Debug().Msgf("%s/%s invoke", ifaceId, opName)
	e.EmitCall(ifaceId, opName, args)

	iface := e.GetInterface(ifaceId)
	if iface == nil {
		return nil, fmt.Errorf("interface %s not found", ifaceId)
	}
	op := iface.GetOperation(opName)
	if op == nil {
		return nil, fmt.Errorf("operation %s not found", opName)
	}
	result, err := e.eval.EvalActions(ifaceId, op.Actions)
	if err != nil {
		e.EmitCallError(ifaceId, opName, err)
		return nil, err
	}
	if result != nil {
		e.EmitReply(ifaceId, opName, result)
	}
	log.Debug().Msgf("%s/%s result %v", ifaceId, opName, result)
	return result, nil
}

// SetProperties sets the properties of the interface.
func (e *Engine) SetProperties(symbol string, props map[string]any) error {
	e.store.Set(symbol, props)
	return nil
}

// FetchProperties returns a copy of the properties of the interface.
func (e *Engine) GetProperties(symbol string) (map[string]any, error) {
	return e.store.Get(symbol), nil
}

func (e *Engine) HasSequence(name string) bool {
	for _, d := range e.docs {
		if d.GetSequence(name) != nil {
			return true
		}
	}
	return false
}

func (e *Engine) PlayAllSequences(ctx context.Context) error {
	log.Debug().Msgf("actions engine play all sequences")
	for _, p := range e.players {
		err := p.Play(ctx)
		if err != nil {
			log.Error().Err(err).Msgf("play sequence %s", p.SequenceName())
		}
	}
	return nil
}

func (e *Engine) StopAllSequences() {
	log.Debug().Msgf("actions engine stop all sequences")
	for _, p := range e.players {
		err := p.Stop()
		if err != nil {
			log.Warn().Msgf("stop sequence %s: %v", p.SequenceName(), err)
		}
	}
}

func (e *Engine) PlaySequence(ctx context.Context, name string) error {
	for _, p := range e.players {
		if p.SequenceName() == name {
			err := p.Play(ctx)
			if err != nil {
				log.Warn().Msgf("play sequence %s: %v", name, err)
			}
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
