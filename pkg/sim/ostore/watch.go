package ostore

import "github.com/google/uuid"

type EventType int

func (e EventType) String() string {
	switch e {
	case EventTypeNone:
		return "none"
	case EventTypeCreate:
		return "create"
	case EventTypeUpdate:
		return "update"
	case EventTypeDelete:
		return "delete"
	default:
		return "unknown"
	}
}

const (
	EventTypeNone EventType = iota
	EventTypeCreate
	EventTypeUpdate
	EventTypeDelete
)

type StoreEvent struct {
	// The type of event
	Type EventType
	// The id of the object
	Id string
	// The properties of the object
	Value map[string]any
}

func (e StoreEvent) IsNone() bool {
	return e.Type == EventTypeNone
}

type StoreNotifyFunc func(event StoreEvent)
type StoreUnWatch func()

type StoreWatchObject struct {
	id string
	fn StoreNotifyFunc
}

type StoreWatcher struct {
	objects []StoreWatchObject
}

func (w *StoreWatcher) OnEvent(fn StoreNotifyFunc) StoreUnWatch {
	id := uuid.New().String()
	w.objects = append(w.objects, StoreWatchObject{id: id, fn: fn})
	return func() {
		w.UnWatch(id)
	}
}

func (w *StoreWatcher) UnWatch(id string) {
	for i, obj := range w.objects {
		if obj.id == id {
			w.objects = append(w.objects[:i], w.objects[i+1:]...)
		}
	}
}

func (w *StoreWatcher) notify(event StoreEvent) {
	for _, obj := range w.objects {
		obj.fn(event)
	}
}
