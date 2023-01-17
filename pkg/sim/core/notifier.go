package core

import "time"

type OnEventFunc func(event *SimuEvent)

type INotifier interface {
	OnEvent(f OnEventFunc)
	EmitEvent(e *SimuEvent)
}

type EventNotifier struct {
	handlers []OnEventFunc
}

func (n *EventNotifier) OnEvent(f OnEventFunc) {
	n.handlers = append(n.handlers, f)
}

func (n *EventNotifier) EmitEvent(e *SimuEvent) {
	e.Timestamp = time.Now()
	for _, f := range n.handlers {
		f(e)
	}
}

func (n *EventNotifier) EmitCall(symbol string, name string, params map[string]any) {
	n.EmitEvent(&SimuEvent{
		Type:   EventCall,
		Symbol: symbol,
		Name:   name,
		KWArgs: params,
	})
}

func (n *EventNotifier) EmitReply(symbol string, name string, value any) {
	n.EmitEvent(&SimuEvent{
		Type:   EventReply,
		Symbol: symbol,
		Name:   name,
		KWArgs: map[string]any{"value": value},
	})
}

func (n *EventNotifier) EmitSignal(symbol string, name string, args map[string]any) {
	n.EmitEvent(&SimuEvent{
		Type:   EventSignal,
		Symbol: symbol,
		Name:   name,
		KWArgs: args,
	})
}

func (n *EventNotifier) EmitPropertySet(symbol string, kwargs map[string]any) {
	n.EmitEvent(&SimuEvent{
		Type:   EventPropertySet,
		Symbol: symbol,
		KWArgs: kwargs,
	})
}

func (n *EventNotifier) EmitPropertyChanged(symbol string, name string, value any) {
	n.EmitEvent(&SimuEvent{
		Type:   EventPropertyChanged,
		Symbol: symbol,
		KWArgs: map[string]any{name: value},
	})
}

func (n *EventNotifier) EmitSimuStart() {
	n.EmitEvent(&SimuEvent{
		Type: EventSimuStart,
	})
}

func (n *EventNotifier) EmitSimuStop() {
	n.EmitEvent(&SimuEvent{
		Type: EventSimuStop,
	})
}

func (n *EventNotifier) EmitCallError(symbol string, name string, err error) {
	n.EmitEvent(&SimuEvent{
		Type:   EventError,
		Symbol: symbol,
		Name:   name,
		Error:  err.Error(),
	})
}

func (n *EventNotifier) EmitError(err error) {
	n.EmitEvent(&SimuEvent{
		Type:  EventError,
		Error: err.Error(),
	})
}
