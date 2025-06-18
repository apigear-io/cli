package sim

import (
	"fmt"

	"github.com/apigear-io/objectlink-core-go/olink/client"
)

type Channel struct {
	engine  *Engine
	clients map[string]*ObjectClient
	url     string
}

func NewChannel(engine *Engine, url string) (*Channel, error) {
	if url == "" {
		url = "ws://localhost:5555/ws"
	}
	c := &Channel{
		engine:  engine,
		clients: make(map[string]*ObjectClient),
		url:     url,
	}
	err := c.Connect()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Channel) connector() IOlinkConnector {
	if c.engine.connector == nil {
		log.Error().Msg("connector is nil")
		return nil
	}
	return c.engine.connector
}

func (c *Channel) node() *client.Node {
	return c.connector().Node(c.url)
}

func (c *Channel) Url() string {
	return c.url
}

func (c *Channel) String() string {
	return fmt.Sprintf("Channel{url=%s}", c.url)
}

func (c *Channel) Connect() error {
	return c.engine.connector.Connect(c.url)
}

func (c *Channel) Disconnect() error {
	c.engine.connector.Disconnect(c.url)
	return nil
}

func (c *Channel) CreateClient(object string) *ObjectClient {
	client, ok := c.clients[object]
	if ok {
		log.Warn().Msgf("client %s already exists", object)
		return client
	}
	client = NewObjectClient(c, object)
	c.clients[object] = client
	return client
}

func (c *Channel) DestroyClient(object string) {
	c.node().UnlinkRemoteNode(object)
	delete(c.clients, object)

}

func (c *Channel) GetClient(object string) *ObjectClient {
	if c.clients[object] == nil {
		log.Warn().Msgf("client %s not found", object)
		return nil
	}
	return c.clients[object]
}
