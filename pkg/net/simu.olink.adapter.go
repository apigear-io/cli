package net

import (
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/sim/model"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
)

// Adapter is a object-link adapter for simulation.
type Adapter struct {
	provider model.SimulationProvider
	registry *remote.Registry
}

func NewAdapter(provider model.SimulationProvider, r *remote.Registry) *Adapter {
	return &Adapter{
		provider: provider,
		registry: r,
	}
}

func (a *Adapter) Registry() *remote.Registry {
	return a.registry
}

func (a *Adapter) CreateSource(objectId string) *SimSource {
	s, err := NewSimuSource(SimuSourceOptions{
		ObjectId:     objectId,
		Registry:     a.registry,
		provider:     a.provider,
		SimulationId: "demo",
	})
	if err != nil {
		log.Error().Err(err).Str("id", objectId).Msg("failed to create simu source")
		return nil
	}
	err = a.registry.AddObjectSource(s)
	if err != nil {
		log.Error().Err(err).Str("id", objectId).Msg("failed to add object source")
		return nil
	}
	return s
}

// SourceById returns a source by id.
func (a *Adapter) SourceById(objectId string) *SimSource {
	s := a.registry.GetObjectSource(objectId)
	if s == nil {
		return nil
	}
	return s.(*SimSource)
}
