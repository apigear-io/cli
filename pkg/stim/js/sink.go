package js

import (
	"encoding/json"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/dop251/goja"
)

type jsSink struct {
	vm       *goja.Runtime
	node     *client.Node
	objectId string
	props    core.KWArgs
}

var _ client.IObjectSink = (*jsSink)(nil)

func NewJsSink(vm *goja.Runtime, node *client.Node, objectId string) *jsSink {
	s := &jsSink{
		vm:       vm,
		node:     node,
		objectId: objectId,
		props:    make(core.KWArgs),
	}
	return s
}

func (s *jsSink) ObjectId() string {
	return s.objectId
}

func (s *jsSink) OnSignal(signalId string, args core.Args) {
	object, member := core.SymbolIdToParts(signalId)
	log.Info().Msgf("%s <- signal %s(%v)", object, member, args)
}

func (s *jsSink) OnPropertyChange(propertyId string, value core.Any) {
	object, member := core.SymbolIdToParts(propertyId)
	log.Info().Msgf("%s <- property changed %s = %v", object, member, value)
}

func (s *jsSink) OnInit(objectId string, props core.KWArgs, node *client.Node) {
	data, err := json.MarshalIndent(props, "", "  ")
	if err != nil {
		log.Warn().Err(err).Msg("failed to marshal value")
		return
	}
	log.Info().Msgf("%s <- init\n", objectId)
	log.Info().Msg(string(data))
	if objectId != s.objectId {
		return
	}
	s.props = props
	s.node = node
}

func (s *jsSink) OnRelease() {
	log.Info().Msgf("%s <- release", s.objectId)
	s.node = nil
}
