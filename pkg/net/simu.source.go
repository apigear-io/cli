package net

import (
	"strings"

	"github.com/apigear-io/objectlink-core-go/log"
	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
)

type Method func(args core.Args) (core.Any, error)

// SimSource is a generic object-link source for simulation.
type SimSource struct {
	id       string
	registry *remote.Registry
	simu     ISimulation
}

var _ remote.IObjectSource = (*SimSource)(nil)

type SimuSourceOptions struct {
	ObjectId string
	Registry *remote.Registry
	Simu     ISimulation
}

// NewSimuSource creates a new SimuSource.
func NewSimuSource(o SimuSourceOptions) *SimSource {
	s := &SimSource{
		id:       o.ObjectId,
		registry: o.Registry,
		simu:     o.Simu,
	}
	return s
}

// ObjectId returns the id of the source.
func (s *SimSource) ObjectId() string {
	return s.id
}

// Invoke invokes a method of the source.
func (s *SimSource) Invoke(name string, args core.Args) (core.Any, error) {
	if strings.HasPrefix(name, "$signal.") {
		signal := strings.TrimPrefix(name, "$signal.")
		s.NotifySignal(signal, args)
		return nil, nil
	}
	if name == "$get" {
		return s.CollectProperties()
	}
	return s.simu.InvokeOperation(s.id, name, args)
}

// SetProperty sets a property of the source.
func (s *SimSource) SetProperty(name string, value core.Any) error {
	return s.simu.SetProperties(s.id, core.KWArgs{name: value})
}

// CollectProperties collects all properties of the source.
func (s *SimSource) CollectProperties() (core.KWArgs, error) {
	return s.simu.GetProperties(s.id)
}

// BroadcastMessage broadcasts a message to all remote nodes.
func (s *SimSource) BroadcastMessage(msg core.Message) {
	for _, node := range s.registry.GetRemoteNodes(s.id) {
		node.SendMessage(msg)
	}
}

// NotifyPropertyChanged notifies all listeners that a property is changed.
func (s *SimSource) NotifyPropertyChanged(name string, value core.Any) {
	propertyId := core.MakeSymbolId(s.id, name)
	msg := core.MakePropertyChangeMessage(propertyId, value)
	s.BroadcastMessage(msg)
}

// NotifySignal notifies all listeners that a signal is emitted.
func (s *SimSource) NotifySignal(name string, args core.Args) {
	signalId := core.MakeSymbolId(s.id, name)
	msg := core.MakeSignalMessage(signalId, args)
	s.BroadcastMessage(msg)
}

// Linked is called when the source is linked to a remote node.
func (s *SimSource) Linked(objectId string, node *remote.Node) error {
	log.Info().Msgf("linked: %v", objectId)
	return nil
}
