package olnk

import (
	"log"

	"github.com/apigear-io/cli/pkg/sim"
	score "github.com/apigear-io/cli/pkg/sim/core"

	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
)

type Method func(args core.Args) (core.Any, error)

// SimuSource is a generic source for simulation.
type SimuSource struct {
	Id       string
	Registry *remote.Registry
	Simu     *sim.Simulation
}

var _ remote.IObjectSource = (*SimuSource)(nil)

type SimuSourceOptions struct {
	ObjectId string
	Simu     *sim.Simulation
	Registry *remote.Registry
}

// NewSimuSource creates a new SimuSource.
func NewSimuSource(o SimuSourceOptions) *SimuSource {
	s := &SimuSource{
		Id:       o.ObjectId,
		Registry: o.Registry,
		Simu:     o.Simu,
	}
	s.Simu.OnEvent(func(event *score.SimuEvent) {
		s.HandleEvent(event)
	})
	return s
}

func (s *SimuSource) HandleEvent(event *score.SimuEvent) {
	log.Printf("event %v\n", event)
	switch event.Type {
	case score.EventSimuStart:
		log.Println("simulation started")
	case score.EventSimuStop:
		log.Println("simulation stopped")
	case score.EventCall:
		log.Printf("call")
	case score.EventReply:
		log.Printf("reply")
	case score.EventSignal:
		log.Printf("signal")
	case score.EventPropertySet:
		log.Printf("property set")
	case score.EventPropertyChanged:
		log.Printf("property changed")
	case score.EventError:
		log.Printf("error")
	}
}

// ObjectId returns the id of the source.
func (s *SimuSource) ObjectId() string {
	return s.Id
}

// Invoke invokes a method of the source.
func (s *SimuSource) Invoke(name string, args core.Args) (core.Any, error) {
	kwargs := core.AsKwArgs(args)
	return s.Simu.InvokeOperation(s.Id, name, kwargs)
}

// SetProperty sets a property of the source.
func (s *SimuSource) SetProperty(name string, value core.Any) error {
	s.Simu.SetProperties(s.Id, core.KWArgs{name: value})
	return nil
}

// CollectProperties collects all properties of the source.
func (s *SimuSource) CollectProperties() (core.KWArgs, error) {
	return s.Simu.GetProperties(s.Id)
}

// BroadcastMessage broadcasts a message to all remote nodes.
func (s *SimuSource) BroadcastMessage(msg core.Message) {
	for _, node := range s.Registry.GetRemoteNodes(s.Id) {
		node.SendMessage(msg)
	}
}

// NotifyPropertyChanged notifies all listeners that a property is changed.
func (s *SimuSource) NotifyPropertyChanged(name string, value core.Any) {
	propertyId := core.MakeSymbolId(s.Id, name)
	msg := core.MakePropertyChangeMessage(propertyId, value)
	s.BroadcastMessage(msg)
}

// NotifySignal notifies all listeners that a signal is emitted.
func (s *SimuSource) NotifySignal(name string, args core.Args) {
	signalId := core.MakeSymbolId(s.Id, name)
	msg := core.MakeSignalMessage(signalId, args)
	s.BroadcastMessage(msg)
}

// Linked is called when the source is linked to a remote node.
func (s *SimuSource) Linked(objectId string, node *remote.Node) error {
	log.Printf("source %s is linked to node %s", objectId, node.Id())
	return nil
}
