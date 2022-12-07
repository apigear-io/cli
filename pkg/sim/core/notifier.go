package core

import "time"

type OnEventFunc func(event *APIEvent)

type INotifier interface {
	OnEvent(f OnEventFunc)
	EmitEvent(e *APIEvent)
}

type EventNotifier struct {
	onEvent OnEventFunc
}

func (n *EventNotifier) OnEvent(f OnEventFunc) {
	n.onEvent = f
}

func (n *EventNotifier) EmitEvent(e *APIEvent) {
	if n.onEvent != nil {
		e.Timestamp = time.Now()
		n.onEvent(e)
	}
}

func (n *EventNotifier) EmitCall(symbol string, name string, params map[string]any) {
	n.EmitEvent(&APIEvent{
		Type:   EventCall,
		Symbol: symbol,
		Name:   name,
		KWArgs: params,
	})
}

func (n *EventNotifier) EmitReply(symbol string, name string, value any, err error) {
	n.EmitEvent(&APIEvent{
		Type:   EventReply,
		Symbol: symbol,
		Name:   name,
		KWArgs: map[string]any{"value": value},
		Error:  err,
	})
}

func (n *EventNotifier) EmitSignal(symbol string, name string, args map[string]any) {
	n.EmitEvent(&APIEvent{
		Type:   EventSignal,
		Symbol: symbol,
		Name:   name,
		KWArgs: args,
	})
}

func (n *EventNotifier) EmitPropertySet(symbol string, kwargs map[string]any) {
	n.EmitEvent(&APIEvent{
		Type:   EventPropertySet,
		Symbol: symbol,
		KWArgs: kwargs,
	})
}

func (n *EventNotifier) EmitPropertyChanged(symbol string, name string, value any) {
	n.EmitEvent(&APIEvent{
		Type:   EventPropertyChanged,
		Symbol: symbol,
		KWArgs: map[string]any{name: value},
	})
}

func (n *EventNotifier) EmitSimuStart() {
	n.EmitEvent(&APIEvent{
		Type: EventSimuStart,
	})
}

func (n *EventNotifier) EmitSimuStop() {
	n.EmitEvent(&APIEvent{
		Type: EventSimuStop,
	})
}
