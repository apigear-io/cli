package model

type EventType string

const (
	EventActorGet     EventType = "actor.get"
	EventActorSet     EventType = "actor.set"
	EventActorChanged EventType = "actor.changed"
	EventActorSignal  EventType = "actor.signal"
	EventActorCall    EventType = "actor.call"
	EventWorldCall    EventType = "world.call"
	EventActorCreated EventType = "actor.created"
	EventActorDeleted EventType = "actor.deleted"
	EventSimCreated   EventType = "sim.created"
	EventSimDeleted   EventType = "sim.deleted"
)

type SimEvent struct {
	Type   EventType `json:"type"`
	World  string    `json:"world"`
	Actor  string    `json:"actor"`
	Member string    `json:"member"`
	Data   any       `json:"data"`
}

func (e SimEvent) Subject() string {
	return "sim." + string(e.Type)
}
