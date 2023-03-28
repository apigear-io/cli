package actions

import (
	"context"
	"fmt"

	"github.com/apigear-io/cli/pkg/sim/core"
	"github.com/apigear-io/cli/pkg/sim/ostore"
	"github.com/apigear-io/cli/pkg/spec"
)

type ScenarioEntry struct {
	Source  string
	Doc     *spec.ScenarioDoc
	Players []*Player
}

// engine implements core.IEngine interface
var _ core.IEngine = (*Engine)(nil)

type Engine struct {
	store ostore.IObjectStore
	eval  *eval
	core.EventNotifier
	entries map[string]*ScenarioEntry
}

func NewEngine(store ostore.IObjectStore) *Engine {
	e := &Engine{
		store:   store,
		entries: make(map[string]*ScenarioEntry),
	}
	eval := NewActionsEvaluator(e, store)
	eval.OnEvent(func(evt *core.SimuEvent) {
		e.EmitEvent(evt)
	})
	e.eval = eval
	return e
}

// ScenarioBySource returns the scenario with the given source
func (e *Engine) ScenarioBySource(source string) *ScenarioEntry {
	entry, ok := e.entries[source]
	if !ok {
		return nil
	}
	return entry
}

func (e *Engine) LoadScenario(source string, doc *spec.ScenarioDoc) error {

	entry := e.ScenarioBySource(source)
	if entry != nil {
		// remove existing scenario and stop all players
		err := e.UnloadScenario(source)
		if err != nil {
			return err
		}
	}
	// create new scenario and add it to the engine
	doc.Source = source
	entry = &ScenarioEntry{
		Source:  source,
		Doc:     doc,
		Players: make([]*Player, 0),
	}
	e.entries[source] = entry

	for _, iface := range entry.Doc.Interfaces {
		if iface.Name == "" {
			return fmt.Errorf("interface %v has no name", iface)
		}
		log.Debug().Msgf("registering interface %s", iface.Name)
	}
	for _, seq := range entry.Doc.Sequences {
		log.Debug().Msgf("registering sequence %s", seq.Name)
		if seq.Interface == "" {
			return fmt.Errorf("sequence %v has no interface", seq)
		}
		iface := e.GetInterface(seq.Interface)
		if iface == nil {
			return fmt.Errorf("interface %s not found", seq.Interface)
		}
		p := NewPlayer(e, iface, seq)
		entry.Players = append(entry.Players, p)
	}
	return nil
}

func (a *Engine) UnloadScenario(source string) error {
	// make sure all players are stopped
	a.StopAllSequences()
	delete(a.entries, source)
	return nil
}

func (e *Engine) HasInterface(ifaceId string) bool {
	return e.GetInterface(ifaceId) != nil
}

func (e *Engine) GetInterface(ifaceId string) *spec.InterfaceEntry {
	for _, entry := range e.entries {
		iface := entry.Doc.GetInterface(ifaceId)
		if iface != nil {
			return iface
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

// GetPlayer returns the player with the given name
func (e *Engine) GetPlayer(name string) *Player {
	for _, entry := range e.entries {
		for _, p := range entry.Players {
			if p.SequenceName() == name {
				return p
			}
		}
	}
	return nil
}

func (e *Engine) HasSequence(name string) bool {
	return e.GetPlayer(name) != nil
}

func (e *Engine) PlayAllSequences(ctx context.Context) error {
	log.Debug().Msgf("actions engine play all sequences")
	for _, entry := range e.entries {
		for _, p := range entry.Players {
			e.EmitSimuStart(p.SequenceName())
			err := p.Play(ctx)
			if err != nil {
				log.Error().Err(err).Msgf("play sequence %s", p.SequenceName())
			}
		}
	}
	return nil
}

func (e *Engine) StopAllSequences() {
	log.Debug().Msgf("actions engine stop all sequence players")
	for _, entry := range e.entries {
		for _, p := range entry.Players {
			e.EmitSimuStop(p.SequenceName())
			err := p.Stop()
			if err != nil {
				log.Warn().Msgf("stop sequence player %s: %v", p.SequenceName(), err)
			}
		}
	}
}

func (e *Engine) PlaySequence(ctx context.Context, name string) error {
	p := e.GetPlayer(name)
	if p == nil {
		return fmt.Errorf("sequence player %s not found", name)
	}
	e.EmitSimuStart(name)
	return p.Play(ctx)
}

func (e *Engine) StopSequence(name string) error {
	p := e.GetPlayer(name)
	if p != nil {
		return fmt.Errorf("sequence player %s not found", name)
	}
	e.EmitSimuStop(name)
	return p.Stop()
}
