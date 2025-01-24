package net

import (
	"context"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/sim/model"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
	"github.com/apigear-io/objectlink-core-go/olink/ws"
)

func NewSimuWSServer(ctx context.Context, provider model.SimulationProvider, simuId string) *ws.Hub {
	registry := remote.NewRegistry()
	registry.SetSourceFactory(func(objectId string) remote.IObjectSource {
		log.Info().Msgf("new simu source: %s", objectId)
		sf, err := NewSimuSource(SimuSourceOptions{
			ObjectId:     objectId,
			Registry:     registry,
			provider:     provider,
			SimulationId: simuId,
		})
		if err != nil {
			log.Error().Err(err).Msg("failed to create simu source")
			return nil
		}
		return sf
	})
	return ws.NewHub(ctx, registry)
}
