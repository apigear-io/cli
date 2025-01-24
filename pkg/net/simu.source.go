package net

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/sim/model"
	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
)

type Method func(args core.Args) (core.Any, error)

type SimuSourceOptions struct {
	ObjectId     string
	Registry     *remote.Registry
	provider     model.SimulationProvider
	SimulationId string
}

// SimSource is a generic object-link source for simulation.
type SimSource struct {
	id       string
	registry *remote.Registry
	provider model.SimulationProvider
	simuId   string
	node     *remote.Node
	unsub    func()
}

var _ remote.IObjectSource = (*SimSource)(nil)

// NewSimuSource creates a new SimuSource.
func NewSimuSource(o SimuSourceOptions) (*SimSource, error) {
	log.Debug().Msgf("new simu source: %s", o.ObjectId)
	if o.provider == nil && o.SimulationId == "" {
		log.Warn().Msg("no simulation provider")
		return nil, fmt.Errorf("no simulation provider")
	}
	if o.SimulationId == "" {
		o.SimulationId = "demo"
	}
	if o.Registry == nil {
		log.Warn().Msg("no registry")
		return nil, fmt.Errorf("no registry")
	}
	if o.ObjectId == "" {
		log.Warn().Msg("no object id")
		return nil, fmt.Errorf("no object id")
	}
	s := &SimSource{
		id:       o.ObjectId,
		registry: o.Registry,
		provider: o.provider,
		simuId:   o.SimulationId,
	}
	s.unsub = o.provider.GetSimulationAPI(o.SimulationId).OnEvent(s.OnEvent)
	o.provider.OnSimulationChanged(func(id string) {
		if id == o.SimulationId {
			if s.unsub != nil {
				s.unsub()
			}
			s.unsub = s.getSimulation().OnEvent(s.OnEvent)
		}
	})
	return s, nil
}

// getSimulation returns the simulation for the source.
func (s *SimSource) getSimulation() model.SimulationAPI {
	return s.provider.GetSimulationAPI(s.simuId)
}

// ObjectId returns the id of the source.
func (s *SimSource) ObjectId() string {
	return s.id
}

// Invoke invokes a method of the source.
func (s *SimSource) Invoke(name string, args core.Args) (core.Any, error) {
	log.Debug().Msgf("invoke: %s %v", name, args)
	if strings.HasPrefix(name, "$signal.") {
		signal := strings.TrimPrefix(name, "$signal.")
		s.NotifySignal(signal, args)
		return nil, nil
	}
	if name == "$get" {
		return s.CollectProperties()
	}
	return s.getSimulation().InvokeOperation(s.id, name, args)
}

// SetProperty sets a property of the source.
func (s *SimSource) SetProperty(name string, value core.Any) error {
	log.Debug().Msgf("set property: %s %v", name, value)
	return s.getSimulation().SetProperties(s.id, core.KWArgs{name: value})
}

// CollectProperties collects all properties of the source.
func (s *SimSource) CollectProperties() (core.KWArgs, error) {
	log.Debug().Msgf("collect properties: %v", s.id)
	return s.getSimulation().GetProperties(s.id)
}

// BroadcastMessage broadcasts a message to all remote nodes.
func (s *SimSource) BroadcastMessage(msg core.Message) {
	if s.node != nil {
		s.node.SendMessage(msg)
	}
}

// NotifyPropertyChanged notifies all listeners that a property is changed.
func (s *SimSource) NotifyPropertyChanged(name string, value core.Any) {
	log.Debug().Msgf("notify property change: %s %v", name, value)
	propertyId := core.MakeSymbolId(s.id, name)
	msg := core.MakePropertyChangeMessage(propertyId, value)
	s.BroadcastMessage(msg)
}

// NotifySignal notifies all listeners that a signal is emitted.
func (s *SimSource) NotifySignal(name string, args core.Args) {
	log.Debug().Msgf("notify signal: %s %v", name, args)
	signalId := core.MakeSymbolId(s.id, name)
	msg := core.MakeSignalMessage(signalId, args)
	s.BroadcastMessage(msg)
}

// Linked is called when the source is linked to a remote node.
func (s *SimSource) Linked(objectId string, node *remote.Node) error {
	log.Debug().Msgf("linked: %v", objectId)
	s.node = node
	return nil
}

func (s *SimSource) OnEvent(evt *model.SimEvent) {
	log.Debug().Msgf("event: %s", evt.Type)
	switch evt.Type {
	case model.EventActorSignal:
		s.registry.NotifySignal(evt.Actor, evt.Member, evt.Data.([]any))
	case model.EventActorChanged:
		s.registry.NotifyPropertyChange(evt.Actor, evt.Data.(map[string]any))
	}
}
