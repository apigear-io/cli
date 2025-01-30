package api

import (
	"github.com/apigear-io/cli/pkg/sim/js"
	"github.com/apigear-io/cli/pkg/sim/model"
	"github.com/apigear-io/cli/pkg/tools"
)

type Simulation interface {
}

// TODO: refacrtor api to not expose js internals and move API to model

type SimulationManager interface {
	CallActorMethod(worldId string, actorId string, method string, args ...any) (any, error)
	RunScript(worldId string, script model.Script) (any, error)
	GetActorState(worldId string, actorId string) map[string]any
	GetActorValue(worldId string, actorId string, property string) any
	SetActorValue(worldId string, actorId string, property string, value any)
	ListActors(worldId string) []string
	GetSimulation(worldId string) *js.Runtime
	CreateSimulation(worldId string) *js.Runtime
	DeleteSimulation(worldId string)
	GetSimulationStatus(worldId string) model.WorldStatus
	CreateActor(worldId string, id string, state map[string]any) (*js.Actor, error)
	DeleteActor(worldId string, id string)
	CallWorldMethod(worldId string, method string, args ...any) (any, error)
	WorldHooks(worldId string) *tools.Hook[model.SimEvent]
}

// TODO: refactor to use model.SimulationAPI or merge with it
type IObjectAPI interface {
	OnEvent(func(evt *model.SimEvent)) func()
	InvokeMethod(objectId string, method string, args []any) (any, error)
	SetProperties(objectId string, props map[string]any) error
	GetProperties(objectId string) (map[string]any, error)
	EmitSignal(objectId string, signal string, args []any)
}
