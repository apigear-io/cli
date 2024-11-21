package simjs

type SimuEvent string

const (
	EUnivWorldAdd         SimuEvent = "univ.add.world"
	EUnivWorldRemove      SimuEvent = "univ.remove.world"
	EUnivChanged          SimuEvent = "univ.changed"
	EWorldStart           SimuEvent = "world.start"
	EWorldStop            SimuEvent = "world.stop"
	EWorldRun             SimuEvent = "world.run"
	EWorldChanged         SimuEvent = "world.changed"
	EActorAdded           SimuEvent = "actor.add"
	EActorRemoved         SimuEvent = "actor.remove"
	EActorPropertySet     SimuEvent = "actor.prop.set"
	EActorPropertyGet     SimuEvent = "actor.prop.get"
	EActorPropertyChanged SimuEvent = "actor.prop.changed"
	EActorStateSet        SimuEvent = "actor.state.set"
	EActorStateGet        SimuEvent = "actor.state.get"
	EActorStateChanged    SimuEvent = "actor.state.changed"
	EActorSignalFire      SimuEvent = "actor.signal.fire"
	EActorSignal          SimuEvent = "actor.signal"
	EActorRPCCall         SimuEvent = "actor.rpc.call"
)

type SimuMessage struct {
	Event   SimuEvent      `json:"event"`
	WorldID string         `json:"world"`
	ActorID string         `json:"actor"`
	Meta    map[string]any `json:"meta"`
	Member  string         `json:"member"`
	Value   any            `json:"value"`
	Args    []any          `json:"args"`
	KWArgs  map[string]any `json:"kwargs"`
}

func (msg *SimuMessage) WorldSubject() string {
	return "simu.world." + msg.WorldID
}

func (msg *SimuMessage) ActorSubject() string {
	return "simu.world." + msg.WorldID + ".actor." + msg.ActorID
}

func (msg *SimuMessage) UniverseSubject() string {
	return "simu.univ"
}
