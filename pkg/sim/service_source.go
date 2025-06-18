package sim

import (
	"sync"

	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
)

type OLinkSource struct {
	mu      sync.RWMutex
	service *ObjectService
	node    *remote.Node
}

func NewOLinkSource(service *ObjectService) *OLinkSource {
	log.Debug().Str("objectId", service.objectId).Msg("new olink source")
	return &OLinkSource{
		service: service,
	}
}

var _ remote.IObjectSource = (*OLinkSource)(nil)

func (s *OLinkSource) ObjectId() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.service.ObjectId()
}

func (s *OLinkSource) Invoke(methodId string, args core.Args) (core.Any, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	jsValue, err := s.service.CallMethod(methodId, args...)
	if err != nil {
		return nil, err
	}
	return jsValue.Export(), nil
}
func (s *OLinkSource) SetProperty(propertyId string, value core.Any) error {
	log.Debug().Str("propertyId", propertyId).Msg("source set property")
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.service.SetProperty(propertyId, value)
	return nil

}
func (s *OLinkSource) Linked(objectId string, node *remote.Node) error {
	log.Debug().Str("objectId", objectId).Msg("source linked")
	s.mu.Lock()
	defer s.mu.Unlock()
	s.node = node
	return nil
}

func (s *OLinkSource) CollectProperties() (core.KWArgs, error) {
	log.Debug().Msg("source collect properties")
	s.mu.RLock()
	defer s.mu.RUnlock()
	return core.KWArgs(s.service.GetProperties()), nil
}

func (s *OLinkSource) Close() {
}

func (s *OLinkSource) NotifyPropertyChanged(name string, value core.Any) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	log.Debug().Str("name", name).Msg("source notify property changed")
	if s.node == nil {
		log.Debug().Msg("source node is nil")
		return
	}
	symbol := core.MakeSymbolId(s.service.objectId, name)
	s.node.NotifyPropertyChange(symbol, value)
}

func (s *OLinkSource) NotifySignal(name string, args core.Args) {
	log.Debug().Str("name", name).Msg("source notify signal")
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.node == nil {
		log.Debug().Msg("source node is nil")
		return
	}
	symbol := core.MakeSymbolId(s.service.objectId, name)
	s.node.NotifySignal(symbol, args)
}
