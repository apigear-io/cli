package net

import (
	"context"

	"github.com/apigear-io/cli/pkg/sim/model"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
	"github.com/apigear-io/objectlink-core-go/olink/ws"
)

func NewSimuWSServer(ctx context.Context, provider SimulationProviderFunc) *ws.Hub {
	simu := provider("")
	registry := remote.NewRegistry()
	registry.SetSourceFactory(func(objectId string) remote.IObjectSource {
		return NewSimuSource(SimuSourceOptions{
			ObjectId: objectId,
			Registry: registry,
			Simu:     simu,
		})
	})
	simu.OnEvent(func(evt *model.SimEvent) {
		switch evt.Type {
		case model.EventActorSignal:
			registry.NotifySignal(evt.Actor, evt.Member, evt.Data.([]any))
		case model.EventActorChanged:
			registry.NotifyPropertyChange(evt.Actor, evt.Data.(map[string]any))
		}
	})

	return ws.NewHub(ctx, registry)
}
