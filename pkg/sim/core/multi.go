package core

import (
	"context"
	"fmt"
)

// engine implements IEngine interface
var _ IEngine = (*MultiEngine)(nil)

type MultiEngine struct {
	engines []IEngine
	EventNotifier
}

func NewMultiEngine(engines ...IEngine) *MultiEngine {
	e := &MultiEngine{
		engines: engines,
	}
	for _, engine := range engines {
		e.registerNotifier(engine)
	}
	return e
}

func (e *MultiEngine) registerNotifier(engine IEngine) {
	engine.OnEvent(func(evt *SimuEvent) {
		e.EmitEvent(evt)
	})
}

// HasInterface returns true if the interface is served by the simulation.
func (e *MultiEngine) HasInterface(ifaceId string) bool {
	for _, engine := range e.engines {
		if engine.HasInterface(ifaceId) {
			return true
		}
	}
	return false
}

// InvokeOperation invokes the operation of the interface.
func (e *MultiEngine) InvokeOperation(ifaceId string, name string, args []any) (any, error) {
	e.EmitCall(ifaceId, name, args)
	result, err := e.invokeOperation(ifaceId, name, args)
	if err != nil {
		e.EmitCallError(ifaceId, name, err)
	}
	if result != nil {
		e.EmitReply(ifaceId, name, result)
	}
	return result, err
}

func (e *MultiEngine) invokeOperation(ifaceId string, name string, args []any) (any, error) {
	for _, engine := range e.engines {
		if engine.HasInterface(ifaceId) {
			return engine.InvokeOperation(ifaceId, name, args)
		}
	}
	return nil, fmt.Errorf("operation %s/%s not found", ifaceId, name)
}

// SetProperties updates the properties of the interface.
func (e *MultiEngine) SetProperties(ifaceId string, props map[string]any) error {
	err := e.setProperties(ifaceId, props)
	e.EmitPropertySet(ifaceId, props)
	if err != nil {
		e.EmitError(err)
	}
	return err
}

func (e *MultiEngine) setProperties(ifaceId string, props map[string]any) error {
	for _, entry := range e.engines {
		if entry.HasInterface(ifaceId) {
			return entry.SetProperties(ifaceId, props)
		}
	}
	return fmt.Errorf("interface %s not found", ifaceId)
}

// FetchProperties fetches the properties of the interface.
func (e *MultiEngine) GetProperties(ifaceId string) (map[string]any, error) {
	for _, engine := range e.engines {
		if engine.HasInterface(ifaceId) {
			return engine.GetProperties(ifaceId)
		}
	}
	return nil, fmt.Errorf("interface %s not found", ifaceId)
}

func (e *MultiEngine) HasSequence(sequencerId string) bool {
	for _, engine := range e.engines {
		if engine.HasSequence(sequencerId) {
			return true
		}
	}
	return false
}

func (e *MultiEngine) PlaySequence(ctx context.Context, sequenceId string) error {
	for _, engine := range e.engines {
		if engine.HasSequence(sequenceId) {
			return engine.PlaySequence(ctx, sequenceId)
		}
	}
	return fmt.Errorf("sequence %s not found", sequenceId)
}

func (e *MultiEngine) StopSequence(sequenceId string) error {
	var lastError error
	for _, engine := range e.engines {
		if engine.HasSequence(sequenceId) {
			lastError = engine.StopSequence(sequenceId)
		}
	}
	return lastError
}

func (e *MultiEngine) PlayAllSequences(ctx context.Context) error {
	var lastError error
	for _, engine := range e.engines {
		lastError = engine.PlayAllSequences(ctx)
	}
	return lastError
}

func (e *MultiEngine) StopAllSequences() {
	for _, engine := range e.engines {
		engine.StopAllSequences()
	}
}
