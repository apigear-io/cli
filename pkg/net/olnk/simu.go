package olnk

import (
	"github.com/apigear-io/cli/pkg/sim"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
)

// Adapter is a object-link adapter for simulation.
type Adapter struct {
	simu     *sim.Simulation
	registry *remote.Registry
}

func NewAdapter(simu *sim.Simulation, r *remote.Registry) *Adapter {
	return &Adapter{
		simu:     simu,
		registry: r,
	}
}

func (a *Adapter) Registry() *remote.Registry {
	return a.registry
}

func (a *Adapter) CreateSource(objectId string) *SimuSource {
	s := NewSimuSource(SimuSourceOptions{
		Simu:     a.simu,
		ObjectId: objectId,
		Registry: a.registry,
	})
	a.registry.AddObjectSource(s)
	return s
}

// SourceById returns a source by id.
func (a *Adapter) SourceById(objectId string) *SimuSource {
	s := a.registry.GetObjectSource(objectId)
	if s == nil {
		return nil
	}
	return s.(*SimuSource)
}
