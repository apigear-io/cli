package sim

import (
	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
	"github.com/rs/zerolog/log"
)

type NullConnector struct {
}

var _ IOlinkConnector = (*NullConnector)(nil)

func NewNullConnector() *NullConnector {
	return &NullConnector{}
}

func (c *NullConnector) Connect(url string) error {
	log.Info().Str("url", url).Msg("Connect")
	return nil
}

func (c *NullConnector) Disconnect(url string) error {
	log.Info().Str("url", url).Msg("Disconnect")
	return nil
}
func (c *NullConnector) RegisterSink(url string, sink client.IObjectSink) {
	log.Info().Str("sink", sink.ObjectId()).Msg("Register sink")
}
func (c *NullConnector) UnregisterSink(url string, sink client.IObjectSink) {
	log.Info().Str("sink", sink.ObjectId()).Msg("Unregister sink")
}

func (c *NullConnector) Node(url string) *client.Node {
	return nil
}

type NullServer struct {
}

var _ IOlinkServer = (*NullServer)(nil)

func NewNullServer() *NullServer {
	return &NullServer{}
}

func (c *NullServer) RegisterSource(sink remote.IObjectSource) {
	log.Info().Msg("Register source")
}
func (c *NullServer) UnregisterSource(sink remote.IObjectSource) {
	log.Info().Msg("Unregister source")
}
