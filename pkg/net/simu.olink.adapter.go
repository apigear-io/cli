package net

import (
	"github.com/apigear-io/objectlink-core-go/log"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
)

// Adapter is a object-link adapter for simulation.
type Adapter struct {
	simu     ISimulation
	registry *remote.Registry
}

func NewAdapter(simu ISimulation, r *remote.Registry) *Adapter {
	return &Adapter{
		simu:     simu,
		registry: r,
	}
}

func (a *Adapter) Registry() *remote.Registry {
	return a.registry
}

func (a *Adapter) CreateSource(objectId string) *SimSource {
	s := NewSimuSource(SimuSourceOptions{
		Simu:     a.simu,
		ObjectId: objectId,
		Registry: a.registry,
	})
	err := a.registry.AddObjectSource(s)
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
