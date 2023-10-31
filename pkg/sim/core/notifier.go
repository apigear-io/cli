package core

import "time"

// OnEventFunc is the type of the event handler function.
type OnEventFunc func(event *SimuEvent)

// INotifier is the interface for the event notifier.
type INotifier interface {
	OnEvent(f OnEventFunc)
	EmitEvent(e *SimuEvent)
}

// EventNotifier implements INotifier interface.
type EventNotifier struct {
	handlers []OnEventFunc
}

// ensures that EventNotifier implements INotifier interface.
var _ INotifier = (*EventNotifier)(nil)

// OnEvent registers the event handler function.
func (n *EventNotifier) OnEvent(f OnEventFunc) {
	n.handlers = append(n.handlers, f)
}

// EmitEvent emits the event to all registered handlers.
// The event timestamp is set to the current time.
func (n *EventNotifier) EmitEvent(e *SimuEvent) {
	e.Timestamp = time.Now()
	for _, f := range n.handlers {
		f(e)
	}
}

// EmitCall emits the call event.
func (n *EventNotifier) EmitCall(ifaceId string, opName string, args []any) {
	n.EmitEvent(&SimuEvent{
		Type:   EventCall,
		Symbol: ifaceId,
		Name:   opName,
		Args:   args,
	})
}

// EmitReply emits the reply event.
func (n *EventNotifier) EmitReply(ifaceId string, opName string, value any) {
	n.EmitEvent(&SimuEvent{
		Type:   EventReply,
		Symbol: ifaceId,
		Name:   opName,
		KWArgs: map[string]any{"value": value},
	})
}

// EmitSignal emits the signal event.
func (n *EventNotifier) EmitSignal(ifaceId string, signName string, args []any) {
	n.EmitEvent(&SimuEvent{
		Type:   EventSignal,
		Symbol: ifaceId,
		Name:   signName,
		Args:   args,
	})
}

// EmitPropertySet emits the property set event.
func (n *EventNotifier) EmitPropertySet(ifaceId string, kwargs map[string]any) {
	n.EmitEvent(&SimuEvent{
		Type:   EventPropertySet,
		Symbol: ifaceId,
		KWArgs: kwargs,
	})
}

// EmitPropertyChanged emits the property changed event.
func (n *EventNotifier) EmitPropertyChanged(ifaceId string, propName string, value any) {
	n.EmitEvent(&SimuEvent{
		Type:   EventPropertyChanged,
		Symbol: ifaceId,
		KWArgs: map[string]any{propName: value},
	})
}

// EmitSimuStart emits the simulation start event.
func (n *EventNotifier) EmitSimuStart(name string) {
	n.EmitEvent(&SimuEvent{
		Type: EventSimuStart,
		Name: name,
	})
}

// EmitSimuStop emits the simulation stop event.
func (n *EventNotifier) EmitSimuStop(name string) {
	n.EmitEvent(&SimuEvent{
		Type: EventSimuStop,
		Name: name,
	})
}

// EmitCallError emits the call error event.
func (n *EventNotifier) EmitCallError(symbol string, name string, err error) {
	n.EmitEvent(&SimuEvent{
		Type:   EventError,
		Symbol: symbol,
		Name:   name,
		Error:  err.Error(),
	})
}

// EmitError emits the error event.
func (n *EventNotifier) EmitError(err error) {
	n.EmitEvent(&SimuEvent{
		Type:  EventError,
		Error: err.Error(),
	})
}
