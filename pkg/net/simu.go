package net

import (
	"context"

	"github.com/apigear-io/cli/pkg/net/olnk"
	"github.com/apigear-io/cli/pkg/sim"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
	"github.com/apigear-io/objectlink-core-go/olink/ws"
)

func NewSimuHub(ctx context.Context, s *sim.Simulation) *ws.Hub {
	registry := remote.NewRegistry()
	registry.SetSourceFactory(func(objectId string) remote.IObjectSource {
		return olnk.NewSimuSource(olnk.SimuSourceOptions{
			ObjectId: objectId,
			Registry: registry,
			Simu:     s,
		})
	})
	return ws.NewHub(ctx, registry)
}
