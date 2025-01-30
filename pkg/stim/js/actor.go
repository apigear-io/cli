package js

import (
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/dop251/goja"
)

// sink is a helper sink to implement IObjectSink for the js actor
type sink struct {
	actor *jsActor
}

var _ client.IObjectSink = (*sink)(nil)

func (s *sink) ObjectId() string {
	return s.actor.objectId
}

func (s *sink) OnSignal(signalId string, args core.Args) {
	log.Info().Msgf("<- signal %s(%v)", signalId, args)
	jsArgs := make([]goja.Value, len(args))
	for i, arg := range args {
		jsArgs[i] = s.actor.node.vm.ToValue(arg)
	}
	s.actor.signals.Emit(signalId, jsArgs...)
}

func (s *sink) OnPropertyChange(propertyId string, value core.Any) {
	log.Info().Msgf("<- property %s = %v", propertyId, value)
	jsValue := s.actor.node.vm.ToValue(value)
	s.actor.changes.Emit(propertyId, jsValue)
	s.actor.properties[propertyId] = jsValue
}

func (s *sink) OnInit(objectId string, props core.KWArgs, node *client.Node) {
	log.Info().Msgf("-> init %s %v", objectId, props)
	for k, v := range props {
		s.actor.properties[k] = s.actor.node.vm.ToValue(v)
	}
}

func (s *sink) OnRelease() {
	log.Info().Msgf("-> release %s", s.ObjectId())
}

type jsActor struct {
	objectId   string
	node       *jsNode
	properties map[string]goja.Value
	signals    *Emitter
	changes    *Emitter
	sink       *sink
}

func NewJsActor(objectId string, node *jsNode) *jsActor {
	a := &jsActor{
		objectId:   objectId,
		node:       node,
		signals:    NewEmitter(node.vm),
		changes:    NewEmitter(node.vm),
		properties: map[string]goja.Value{},
	}
	a.sink = &sink{actor: a}
	// register sink
	log.Info().Msgf("-> register sink for %s", a.objectId)
	node.node.Registry().AddObjectSink(a.sink)
	return a
}

func (a *jsActor) ObjectId() string {
	return a.objectId
}

// SetProperty sets the value of a remote property
func (a *jsActor) SetProperty(property string, value any) error {
	return a.node.SetProperty(a.objectId, property, value)
}

// GetProperty returns the value of a localproperty
func (a *jsActor) GetProperty(property string) any {
	return a.properties[property]
}

// HasProperty checks if the actor has a value
func (a *jsActor) HasProperty(property string) bool {
	_, ok := a.properties[property]
	return ok
}

// OnPropertyChange registers a callback that is called when the property changes
func (a *jsActor) OnPropertyChange(property string, fn goja.Callable) {
	a.changes.On(property, fn)
}

func (a *jsActor) Invoke(method string, args core.Args) (any, error) {
	return a.node.Invoke(a.objectId, method, args)
}

func (a *jsActor) EmitSignal(signal string, args core.Args) error {
	return a.node.Signal(a.objectId, signal, args)
}

func (a *jsActor) OnSignal(signal string, fn goja.Callable) {
	a.signals.On(signal, fn)
}

func (a *jsActor) Unlink() error {
	return a.node.Unlink(a.objectId)
}

// SetProperties sets the properties of the actor
func (a *jsActor) SetProperties(properties core.KWArgs) error {
	for k, v := range properties {
		a.node.SetProperty(a.objectId, k, v)
	}
	return nil
}
