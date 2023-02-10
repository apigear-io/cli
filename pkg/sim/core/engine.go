package core

import "context"

// IEngine is the interface for the simulation engine.
type IEngine interface {
	INotifier
	HasInterface(symbol string) bool
	InvokeOperation(symbol string, name string, args []any) (any, error)
	SetProperties(symbol string, params map[string]any) error
	GetProperties(symbol string) (map[string]any, error)
	HasSequence(name string) bool
	PlaySequence(ctx context.Context, name string) error
	StopSequence(name string)
	PlayAllSequences(ctx context.Context) error
	StopAllSequences()
}
