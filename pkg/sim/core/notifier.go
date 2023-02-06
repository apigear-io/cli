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

func (n *EventNotifier) EmitCall(ifaceId string, opName string, args []any) {
	n.EmitEvent(&SimuEvent{
		Type:   EventCall,
		Symbol: ifaceId,
		Name:   opName,
		Args:   args,
	})
}

func (n *EventNotifier) EmitReply(ifaceId string, opName string, value any) {
	n.EmitEvent(&SimuEvent{
		Type:   EventReply,
		Symbol: ifaceId,
		Name:   opName,
		KWArgs: map[string]any{"value": value},
	})
}

func (n *EventNotifier) EmitSignal(ifaceId string, signName string, args []any) {
	n.EmitEvent(&SimuEvent{
		Type:   EventSignal,
		Symbol: ifaceId,
		Name:   signName,
		Args:   args,
	})
}

func (n *EventNotifier) EmitPropertySet(ifaceId string, kwargs map[string]any) {
	n.EmitEvent(&SimuEvent{
		Type:   EventPropertySet,
		Symbol: ifaceId,
		KWArgs: kwargs,
	})
}

func (n *EventNotifier) EmitPropertyChanged(ifaceId string, propName string, value any) {
	n.EmitEvent(&SimuEvent{
		Type:   EventPropertyChanged,
		Symbol: ifaceId,
		KWArgs: map[string]any{propName: value},
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
