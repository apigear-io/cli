package net

import "github.com/apigear-io/cli/pkg/sim/model"

type ISimulation interface {
	OnEvent(func(evt *model.SimEvent))
	InvokeOperation(objectId string, method string, args []any) (any, error)
	SetProperties(objectId string, props map[string]any) error
	GetProperties(objectId string) (map[string]any, error)
}

type SimulationProviderFunc = func(id string) ISimulation
