package model

type SimulationAPI interface {
	OnEvent(func(evt *SimEvent)) func()
	InvokeOperation(objectId string, method string, args []any) (any, error)
	SetProperties(objectId string, props map[string]any) error
	GetProperties(objectId string) (map[string]any, error)
}

type SimulationProvider interface {
	OnSimulationChanged(func(id string))
	GetSimulationAPI(id string) SimulationAPI
}
