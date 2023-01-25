package olnk

import (
	"github.com/apigear-io/cli/pkg/sim"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
	"github.com/apigear-io/objectlink-core-go/olink/ws"
)

func NewHub(s *sim.Simulation) *ws.Hub {
	registry := remote.NewRegistry()
	registry.SetSourceFactory(func(objectId string) remote.IObjectSource {
		return NewSimuSource(SimuSourceOptions{
			ObjectId: objectId,
			Registry: registry,
			Simu:     s,
		})
	})
	return ws.NewHub(registry)
}
