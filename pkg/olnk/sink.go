package olnk

import (
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/core"
)

// client messages supported for feed
// - ["link", "demo.Calc"]
// - ["set", "demo.Calc/total", 20]
// - ["invoke", 1, "demo.Calc/add", [1]]
// - ["unlink", "demo.Calc"]
// server messages not supported for feed
// - ["init", "demo.Calc", { "total": 10 }]
// - ["change", "demo.Calc/total", 20]
// - ["reply", 1, "demo.Calc/add", 21]
// - ["signal", "demo.Calc/clearDone", []]
// - ["error", "init", 0, "init error"]

type ObjectSink struct {
	objectId string
}

func (s *ObjectSink) ObjectId() string {
	return s.objectId
}

func (s *ObjectSink) HandleSignal(signalId string, args core.Args) {
	log.Info().Msgf("<- signal %s(%v)", signalId, args)
}
func (s *ObjectSink) HandlePropertyChange(propertyId string, value core.Any) {
	log.Info().Msgf("<- property %s = %v", propertyId, value)
}
func (s *ObjectSink) HandleInit(objectId string, props core.KWArgs, node *client.Node) {
	s.objectId = objectId
	log.Info().Msgf("<- init %s with %v", objectId, props)
}
func (s *ObjectSink) HandleRelease() {
	log.Info().Msgf("<- release %s", s.objectId)
	s.objectId = ""
}
