package core

// IEngine is the interface for the simulation engine.
type IEngine interface {
	INotifier
	HasInterface(symbol string) bool
	InvokeOperation(symbol string, name string, params map[string]any) (any, error)
	SetProperties(symbol string, params map[string]any) error
	GetProperties(symbol string) (map[string]any, error)
	HasSequence(name string) bool
	PlaySequence(name string) error
	StopSequence(name string)
	PlayAllSequences() error
	StopAllSequences()
}
