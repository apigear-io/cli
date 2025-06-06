package sim

import (
	"sync"

	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/rs/zerolog/log"
)

type ObjectClientSink struct {
	mu     sync.Mutex
	client *ObjectClient
	node   *client.Node
}

func NewObjectClientSink(client *ObjectClient) *ObjectClientSink {
	return &ObjectClientSink{client: client}
}

func (s *ObjectClientSink) ObjectId() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.client.object
}

func (s *ObjectClientSink) HandleSignal(signalId string, args core.Args) {
	log.Debug().Interface("args", args).Msg("ObjectClientSink.HandleSignal")
	s.mu.Lock()
	defer s.mu.Unlock()

	s.client.emitLocalSignal(signalId, args)
}

func (s *ObjectClientSink) HandlePropertyChange(propertyId string, value core.Any) {
	log.Debug().Interface("value", value).Msg("ObjectClientSink.HandlePropertyChange")
	s.mu.Lock()
	defer s.mu.Unlock()
	s.client.setLocalProperty(propertyId, value)
}

func (s *ObjectClientSink) HandleInit(objectId string, props core.KWArgs, node *client.Node) {
	log.Debug().Interface("props", props).Msg("ObjectClientSink.HandleInit")
	s.mu.Lock()
	defer s.mu.Unlock()
	s.node = node
	if s.client.object != objectId {
		log.Error().Msgf("ObjectClientSink.HandleInit: objectId mismatch: %s != %s", s.client.object, objectId)
		return
	}
	s.client.setLocalProperties(props)
}

func (s *ObjectClientSink) HandleRelease() {
	log.Info().Msg("ObjectClientSink.HandleRelease")
	s.mu.Lock()
	defer s.mu.Unlock()
	s.node = nil
}
