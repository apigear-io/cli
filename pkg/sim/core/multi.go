package core

import (
	"fmt"
)

// engine implements IEngine interface
var _ IEngine = (*MultiEngine)(nil)

type MultiEngine struct {
	entries []IEngine
	Notifier
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
	engine.OnChange(func(ifaceId string, name string, value any) {
		e.EmitOnChange(ifaceId, name, value)
	})
	engine.OnSignal(func(ifaceId string, name string, args KWArgs) {
		e.EmitOnSignal(ifaceId, name, args)
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
func (e *MultiEngine) InvokeOperation(ifaceId string, name string, args KWArgs) (any, error) {
	for _, entry := range e.entries {
		if entry.HasInterface(ifaceId) {
			return entry.InvokeOperation(ifaceId, name, args)
		}
	}
	return nil, fmt.Errorf("operation %s/%s not found", ifaceId, name)
}

// SetProperties sets the properties of the interface.
func (e *MultiEngine) SetProperties(ifaceId string, props KWArgs) error {
	for _, entry := range e.entries {
		if entry.HasInterface(ifaceId) {
			return entry.SetProperties(ifaceId, props)
		}
	}
	return fmt.Errorf("interface %s not found", ifaceId)
}

// FetchProperties fetches the properties of the interface.
func (e *MultiEngine) GetProperties(ifaceId string) (KWArgs, error) {
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

func (e *MultiEngine) PlaySequence(sequenceId string) {
	for _, entry := range e.entries {
		if entry.HasSequence(sequenceId) {
			entry.PlaySequence(sequenceId)
		}
	}
}
