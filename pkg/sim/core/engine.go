package core

// IEngine is the interface for the simulation engine.
type IEngine interface {
	HasInterface(symbol string) bool
	InvokeOperation(symbol string, name string, params KWArgs) (any, error)
	SetProperties(symbol string, params KWArgs) error
	GetProperties(symbol string) (KWArgs, error)
	HasSequence(name string) bool
	PlaySequence(name string)
	OnChange(f OnChangeFunc)
	OnSignal(f OnSignalFunc)
}
