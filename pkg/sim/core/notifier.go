package core

type OnChangeFunc func(symbol string, name string, value any)
type OnSignalFunc func(symbol string, name string, args map[string]any)

type Notifier struct {
	onChange OnChangeFunc
	onSignal OnSignalFunc
}

func (n *Notifier) OnChange(f OnChangeFunc) {
	n.onChange = f
}

func (n *Notifier) OnSignal(f OnSignalFunc) {
	n.onSignal = f
}

func (n *Notifier) EmitOnChange(symbol string, name string, value any) {
	if n.onChange != nil {
		n.onChange(symbol, name, value)
	}
}

func (n *Notifier) EmitOnSignal(symbol string, name string, args map[string]any) {
	if n.onSignal != nil {
		n.onSignal(symbol, name, args)
	}
}
