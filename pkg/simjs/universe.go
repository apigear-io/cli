package simjs

import (
	"fmt"

	"github.com/apigear-io/objectlink-core-go/log"
)

type Universe struct {
	worlds map[string]*World
	pub    SimuPublisher
}

func NewUniverse(pub SimuPublisher) *Universe {
	u := &Universe{
		worlds: make(map[string]*World),
		pub:    pub,
	}
	return u
}

func (u *Universe) Worlds() map[string]*World {
	return u.worlds
}

func (u *Universe) Publish(subject string, msg *SimuMessage) error {
	return u.pub.Publish(subject, msg)
}

func (u *Universe) AddWorld(id string) (*World, error) {
	return u.GetWorld(id), nil
}

func (u *Universe) addWorld(world *World) error {
	id := world.ID
	_, ok := u.worlds[id]
	if ok {
		return fmt.Errorf("world %s already exists", id)
	}
	u.worlds[id] = world
	return nil
}

func (u *Universe) GetWorld(id string) *World {
	w, ok := u.worlds[id]
	if ok {
		return w
	}
	w, err := NewWorld(id, u)
	if err != nil {
		return nil
	}
	u.addWorld(w)
	return w
}

func (u *Universe) GetActor(worldID, actorID string) *Actor {
	w := u.GetWorld(worldID)
	if w == nil {
		return nil
	}
	return w.GetActor(actorID)
}

func (u *Universe) WorldChanged(worldID string, value string) {
	msg := &SimuMessage{
		Event:   EWorldChanged,
		WorldID: worldID,
		Value:   value,
	}
	u.Publish(msg.WorldSubject(), msg)
}

func (u *Universe) WorldRemove(worldID string) {
	w := u.GetWorld(worldID)
	if w != nil {
		w.StopScript()
	}
	w.WorldChanged("removed")
	delete(u.worlds, worldID)
}

func (u *Universe) WorldStart(worldID string) {
	w := u.GetWorld(worldID)
	if w != nil {
		w.StartScript()
	}
}

func (u *Universe) WorldStop(worldID string) {
	w := u.GetWorld(worldID)
	if w != nil {
		w.StopScript()
	}
}

func (u *Universe) WorldRun(worldID string, script string, content string) {
	w := u.GetWorld(worldID)
	if w != nil {
		w.RunScript(script, content)
	}
}

func (u *Universe) UniverseChanged(value any) {
	msg := &SimuMessage{
		Event: EUnivChanged,
		Value: value,
	}
	u.Publish(msg.UniverseSubject(), msg)
}

func (u *Universe) HandleMessage(msg *SimuMessage) *SimuMessage {
	switch msg.Event {
	case EUnivWorldRemove:
		log.Info().Msgf("remove world %s", msg.WorldID)
		u.WorldRemove(msg.WorldID)
	case EUnivWorldAdd:
		log.Info().Msgf("add world %s", msg.WorldID)
		u.AddWorld(msg.WorldID)
	case EUnivChanged:
		// ignore
	default:
		log.Info().Msgf("unknown event %s", msg.Event)
	}
	return nil
}
