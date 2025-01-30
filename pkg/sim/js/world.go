package js

import (
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/sim/model"
	"github.com/apigear-io/cli/pkg/tools"
	"github.com/dop251/goja"
)

const (
	DefaultWorldId = "demo"
)

type World struct {
	id     string
	vm     *goja.Runtime
	actors map[string]*Actor
	hooks  *tools.Hook[model.SimEvent]
}

func NewDemoWorld() *World {
	return NewWorld(DefaultWorldId, goja.New())
}

func NewWorld(id string, vm *goja.Runtime) *World {
	if id == "" {
		log.Warn().Msgf("world id is empty, using default world id=%s", DefaultWorldId)
		id = DefaultWorldId
	}
	log.Debug().Str("id", id).Msg("new world")
	w := &World{
		id:     id,
		vm:     vm,
		actors: map[string]*Actor{},
		hooks:  tools.NewHook[model.SimEvent](),
	}
	return w
}

// Id_ returns the world's ID
func (w *World) Id_() string {
	return w.id
}

func (w *World) CreateActor(id string, state *goja.Object) (*Actor, error) {
	log.Debug().Str("actor", id).Str("world", w.id).Msg("world create actor")
	if id == "" {
		id = "actor"
	}
	actor, ok := w.actors[id]
	if ok {
		log.Debug().Str("actor", id).Msg("actor already exists")
		return actor, nil
	}
	actor, err := NewActor(id, state, w)
	if err != nil {
		return nil, err
	}
	w.actors[id] = actor
	// proxy := actor.toProxy()
	w.fireEvent(model.EventActorCreated, id, "", nil)
	return actor, nil
}

// GetActor returns an actor by name
func (w *World) GetActor(id string) *Actor {
	if id == "" {
		id = "actor"
	}
	log.Debug().Str("actor", id).Msg("world get actor")
	actor, ok := w.actors[id]
	if ok {
		return actor
	}
	actor, err := w.CreateActor(id, w.vm.NewObject())
	if err != nil {
		log.Error().Err(err).Str("actor", id).Msg("actor not found")
		return nil
	}
	w.actors[id] = actor
	return actor
}

// ListActors returns a list of actors in the simulation
func (w *World) ListActors() []string {
	actors := make([]string, 0, len(w.actors))
	for id := range w.actors {
		actors = append(actors, id)
	}
	return actors
}

// ActorCount returns the number of actors in the simulation
func (w *World) ActorCount() int {
	return len(w.actors)
}

// DeleteActor deletes an actor from the simulation
func (w *World) DeleteActor(id string) {
	if id == "" {
		id = "actor"
	}
	delete(w.actors, id)
	w.fireEvent(model.EventActorDeleted, id, "", nil)
}

func (w *World) fireEvent(event model.EventType, actorId string, memberId string, data map[string]any) {
	w.hooks.Fire(&model.SimEvent{
		Type:   event,
		World:  w.id,
		Actor:  actorId,
		Member: memberId,
		Data:   data,
	})
}

// Hooks returns the world's hooks
func (w *World) Hooks() *tools.Hook[model.SimEvent] {
	return w.hooks
}
