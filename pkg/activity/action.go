package activity

import (
	"fmt"
	"time"
)

type Verb string

const (
	VCreate Verb = "create"
	VUpdate Verb = "update"
	VDelete Verb = "delete"
	VStart  Verb = "start"
	VStop   Verb = "stop"
	VWrite  Verb = "write"
	VRead   Verb = "read"
	VOpen   Verb = "open"
	VClose  Verb = "close"
)

type Action struct {
	Actor  string    `json:"actor"`
	Verb   Verb      `json:"verb"`
	Object string    `json:"object"`
	Target string    `json:"target"`
	Time   time.Time `json:"time"`
}

func (a Action) String() string {
	if a.Target == "" {
		if a.Object == "" {
			return fmt.Sprintf("%s %s", a.Actor, a.Verb)
		}
		return fmt.Sprintf("%s %s %s", a.Actor, a.Verb, a.Object)
	}
	return fmt.Sprintf("%s %s %s on %s", a.Actor, a.Verb, a.Object, a.Target)
}

func ReportAction(actor string, verb Verb, object string, target string) {
	action := &Action{
		Actor:  actor,
		Verb:   verb,
		Object: object,
		Target: target,
		Time:   time.Now(),
	}
	SendAction(action)
}
