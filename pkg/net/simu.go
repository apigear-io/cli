package net

import (
	"context"

	"github.com/apigear-io/cli/pkg/net/olnk"
	"github.com/apigear-io/cli/pkg/sim"
	score "github.com/apigear-io/cli/pkg/sim/core"
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
	s.OnEvent(func(evt *score.SimuEvent) {
		switch evt.Type {
		case score.EventSignal:
			registry.NotifySignal(evt.Symbol, evt.Name, evt.Args)
		case score.EventPropertyChanged:
			registry.NotifyPropertyChange(evt.Symbol, evt.KWArgs)
		}
	})

	return ws.NewHub(ctx, registry)
}
