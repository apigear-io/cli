package mock

import (
	"context"

	"github.com/apigear-io/cli/pkg/sim/core"
	"github.com/apigear-io/cli/pkg/spec"
)

type MockEvent struct {
	Command string
	Name    string
	Symbol  string
	KWArgs  map[string]any
	Args    []any
	Source  string
	Doc     *spec.ScenarioDoc
}

type MockEngine struct {
	core.INotifier
	Events []MockEvent
}

func NewMockEngine() *MockEngine {
	return &MockEngine{
		Events: []MockEvent{},
	}
}

var _ core.IEngine = (*MockEngine)(nil)

func (e *MockEngine) HasInterface(id string) bool {
	return false
}

func (e *MockEngine) InvokeOperation(symbol string, name string, args []any) (any, error) {
	e.Events = append(e.Events, MockEvent{
		Command: "invoke",
		Name:    name,
		Symbol:  symbol,
		Args:    args,
	})
	return nil, nil
}
func (e *MockEngine) SetProperties(symbol string, params map[string]any) error {
	e.Events = append(e.Events, MockEvent{
		Command: "set",
		Symbol:  symbol,
		KWArgs:  params,
	})
	return nil
}
func (e *MockEngine) GetProperties(symbol string) (map[string]any, error) {
	e.Events = append(e.Events, MockEvent{
		Command: "get",
		Symbol:  symbol,
	})
	return nil, nil
}
func (e *MockEngine) HasSequence(name string) bool {
	e.Events = append(e.Events, MockEvent{
		Command: "has",
		Name:    name,
	})
	return false
}
func (e *MockEngine) PlaySequence(ctx context.Context, name string) error {
	e.Events = append(e.Events, MockEvent{
		Command: "play",
		Name:    name,
	})
	return nil
}
func (e *MockEngine) StopSequence(name string) error {
	e.Events = append(e.Events, MockEvent{
		Command: "stop",
		Name:    name,
	})
	return nil
}
func (e *MockEngine) PlayAllSequences(ctx context.Context) error {
	e.Events = append(e.Events, MockEvent{
		Command: "play_all",
	})
	return nil
}
func (e *MockEngine) StopAllSequences() {
	e.Events = append(e.Events, MockEvent{
		Command: "stop_all",
	})
}

func (e *MockEngine) ClearEvents() {
	e.Events = []MockEvent{}
}

func (e *MockEngine) LastEvent() MockEvent {
	return e.Events[len(e.Events)-1]
}
