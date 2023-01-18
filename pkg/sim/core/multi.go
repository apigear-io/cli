package core

import (
	"fmt"
)

// engine implements IEngine interface
var _ IEngine = (*MultiEngine)(nil)

type MultiEngine struct {
	entries []IEngine
	EventNotifier
}

func NewMultiEngine(entries ...IEngine) *MultiEngine {
	e := &MultiEngine{
		entries: entries,
	}
	for _, entry := range entries {
		e.registerNotifier(entry)
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
	for _, entry := range e.entries {
		if entry.HasInterface(ifaceId) {
			return true
		}
	}
	return false
}

// InvokeOperation invokes the operation of the interface.
func (e *MultiEngine) InvokeOperation(ifaceId string, name string, args map[string]any) (any, error) {
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

func (e *MultiEngine) invokeOperation(ifaceId string, name string, args map[string]any) (any, error) {
	for _, entry := range e.entries {
		if entry.HasInterface(ifaceId) {
			return entry.InvokeOperation(ifaceId, name, args)
		}
	}
	return nil, fmt.Errorf("operation %s/%s not found", ifaceId, name)
}

// SetProperties sets the properties of the interface.
func (e *MultiEngine) SetProperties(ifaceId string, props map[string]any) error {
	err := e.setProperties(ifaceId, props)
	e.EmitPropertySet(ifaceId, props)
	if err != nil {
		e.EmitError(err)
	}
	return err
}

func (e *MultiEngine) setProperties(ifaceId string, props map[string]any) error {
	for _, entry := range e.entries {
		if entry.HasInterface(ifaceId) {
			return entry.SetProperties(ifaceId, props)
		}
	}
	return fmt.Errorf("interface %s not found", ifaceId)
}

// FetchProperties fetches the properties of the interface.
func (e *MultiEngine) GetProperties(ifaceId string) (map[string]any, error) {
	for _, entry := range e.entries {
		if entry.HasInterface(ifaceId) {
			return entry.GetProperties(ifaceId)
		}
	}
	return nil, fmt.Errorf("interface %s not found", ifaceId)
}

func (e *MultiEngine) HasSequence(sequencerId string) bool {
	for _, entry := range e.entries {
		if entry.HasSequence(sequencerId) {
			return true
		}
	}
	return false
}

func (e *MultiEngine) PlaySequence(sequenceId string) error {
	for _, entry := range e.entries {
		if entry.HasSequence(sequenceId) {
			return entry.PlaySequence(sequenceId)
		}
	}
	return fmt.Errorf("sequence %s not found", sequenceId)
}

func (e *MultiEngine) StopSequence(sequenceId string) {
	for _, entry := range e.entries {
		if entry.HasSequence(sequenceId) {
			entry.StopSequence(sequenceId)
		}
	}
}

func (e *MultiEngine) PlayAllSequences() error {
	for _, entry := range e.entries {
		if err := entry.PlayAllSequences(); err != nil {
			return err
		}
	}
	return nil
}

func (e *MultiEngine) StopAllSequences() {
	for _, entry := range e.entries {
		entry.StopAllSequences()
	}
}
