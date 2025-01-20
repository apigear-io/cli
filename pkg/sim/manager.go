package sim

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/sim/api"
	"github.com/apigear-io/cli/pkg/sim/js"
	"github.com/apigear-io/cli/pkg/sim/model"
	"github.com/apigear-io/cli/pkg/tools"
	"github.com/apigear-io/objectlink-core-go/log"
)

var (
	defaultManager *Manager
)

func GetManager() *Manager {
	if defaultManager == nil {
		defaultManager = New()
	}
	return defaultManager
}

type Manager struct {
	sims    map[string]*js.Runtime
	hooks   tools.Hook[model.SimEvent]
	service *api.Service
	client  *api.Client
}

func New() *Manager {
	return &Manager{
		sims:  make(map[string]*js.Runtime),
		hooks: tools.Hook[model.SimEvent]{},
	}
}

// CreateService creates a new simulation service
func (m *Manager) CreateService(url string) (*api.Service, error) {
	if m.service == nil {
		service, err := api.NewService(url, m)
		if err != nil {
			return nil, err
		}
		m.service = service
	}
	return m.service, nil
}

func (m *Manager) CreateClient(url string) (*api.Client, error) {
	if m.client == nil {
		client, err := api.NewClient(url)
		if err != nil {
			return nil, err
		}
		m.client = client
	}
	return m.client, nil
}

func (m *Manager) GetObjectAPI(world string) (*api.ObjectAPI, error) {
	if m.client == nil {
		return nil, fmt.Errorf("client not created")
	}
	// TODO: cache object APIs
	return m.client.ObjectAPI(world), nil
}

// CreateSimulation creates a new simulation
func (m *Manager) CreateSimulation(id string) *js.Runtime {
	m.DeleteSimulation(id)
	m.sims[id] = js.NewRuntime(id)
	m.fireEvent(model.EventSimCreated, id)
	return m.sims[id]
}

// DeleteSimulation deletes a simulation
func (m *Manager) DeleteSimulation(id string) {
	delete(m.sims, id)
	m.fireEvent(model.EventSimDeleted, id)
}

// GetSimulation returns a simulation
func (m *Manager) GetSimulation(id string) *js.Runtime {
	if m.sims[id] != nil {
		return m.sims[id]
	}
	return m.CreateSimulation(id)
}

// RunScript runs a script in the simulation
func (m *Manager) RunScript(worldId string, script model.Script) (any, error) {
	si := m.CreateSimulation(worldId)
	return si.RunScript(script)
}

// CreateActor creates a new actor in the simulation
func (m *Manager) CreateActor(worldId string, id string, state map[string]any) (*js.Actor, error) {
	si := m.GetSimulation(worldId)
	_, err := si.GetWorld().CreateActor(id, si.MapToJsObject(state))
	if err != nil {
		return nil, err
	}
	return si.GetWorld().GetActor(id), nil
}

// GetActor returns an actor by name
func (m *Manager) GetActor(worldId string, id string) *js.Actor {
	simu := m.GetSimulation(worldId)
	if simu == nil {
		log.Warn().Str("world", worldId).Msg("world not found")
		return nil
	}
	return simu.GetWorld().GetActor(id)
}

// CallActorMethod calls a method on an actor
func (m *Manager) CallActorMethod(worldId string, actorId string, method string, args ...any) (any, error) {
	si := m.GetSimulation(worldId)
	actor := si.GetWorld().GetActor(actorId)
	if actor == nil {
		return nil, fmt.Errorf("actor %s not found", actorId)
	}
	return actor.CallMethod(method, si.ArgsToJsArgs(args)...)
}

// ListActors returns a list of actors
func (m *Manager) ListActors(worldId string) []string {
	si := m.GetSimulation(worldId)
	if si == nil {
		return nil
	}
	return si.GetWorld().ListActors()
}

// GetSimulationStatus returns the status of a simulation
func (m *Manager) GetSimulationStatus(worldId string) model.WorldStatus {
	si := m.GetSimulation(worldId)
	if si == nil {
		return model.WorldStatus{}
	}
	return model.WorldStatus{
		Name:       si.GetWorld().Id_(),
		IsActive:   si.IsActive,
		ActorCount: si.GetWorld().ActorCount(),
		LastUpdate: si.LastUpdate,
	}
}

// RemoveAll removes all simulations
func (m *Manager) RemoveAll() {
	for k := range m.sims {
		m.DeleteSimulation(k)
	}
}

///////////////////////////////////////////////////////////////////////////////
// Actors
///////////////////////////////////////////////////////////////////////////////

// DeleteActor deletes an actor
func (m *Manager) DeleteActor(worldId string, actorId string) {
	simu := m.GetSimulation(worldId)
	if simu == nil {
		return
	}
	simu.GetWorld().DeleteActor(actorId)
}

///////////////////////////////////////////////////////////////////////////////
// World
///////////////////////////////////////////////////////////////////////////////

// CallWorldMethod calls a method on the world
func (m *Manager) CallWorldMethod(worldId string, method string, args ...any) (any, error) {
	simu := m.GetSimulation(worldId)
	if simu == nil {
		return nil, fmt.Errorf("simulation %s not found", worldId)
	}
	return simu.RunFunction(method, args...)
}

///////////////////////////////////////////////////////////////////////////////
// Properties
///////////////////////////////////////////////////////////////////////////////

// SetActorValue sets the value of an actor
func (m *Manager) SetActorValue(worldId string, actorId string, key string, value any) {
	si := m.GetSimulation(worldId)
	if si == nil {
		return
	}
	actor := si.GetActor(actorId)
	if actor == nil {
		return
	}
	actor.SetProperty(key, si.ToJsValue(value))
}

// GetActorValue returns the value of an actor
func (m *Manager) GetActorValue(worldId string, actorId string, key string) any {
	actor := m.GetActor(worldId, actorId)
	if actor == nil {
		return nil
	}
	return actor.GetProperty(key).Export()
}

// GetActorState returns the state of an actor
func (m *Manager) GetActorState(worldId string, actorId string) map[string]any {
	si := m.GetSimulation(worldId)
	if si == nil {
		log.Warn().Str("world", worldId).Msg("world not found")
		return nil
	}
	state, err := si.GetProperties(actorId)
	if err != nil {
		log.Warn().Str("world", worldId).Str("actor", actorId).Msg("actor not found")
		return nil
	}
	return state
}

// SetActorState sets the state of an actor
func (m *Manager) SetActorState(worldId string, actorId string, state map[string]any) {
	si := m.GetSimulation(worldId)
	if si == nil {
		return
	}
	err := si.SetProperties(actorId, state)
	if err != nil {
		log.Warn().Str("world", worldId).Str("actor", actorId).Msg("actor not found")
	}
}

///////////////////////////////////////////////////////////////////////////////
// Events
///////////////////////////////////////////////////////////////////////////////

func (m *Manager) fireEvent(event model.EventType, worldId string) {
	err := m.hooks.Fire(&model.SimEvent{
		Type:  event,
		World: worldId,
	})
	if err != nil {
		log.Error().Err(err).Msg("error firing event")
	}
}
