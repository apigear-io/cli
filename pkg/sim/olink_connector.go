package sim

import (
	"context"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/ws"
	"github.com/rs/zerolog/log"
)

var nextChannelId = helper.MakeIdGenerator("c")

type connEntry struct {
	conn     *ws.Connection
	id       string
	url      string
	node     *client.Node
	registry *client.Registry
}

func (e *connEntry) Close() {
	e.conn.Close()
	e.node.Close()
}

type IOlinkConnector interface {
	Connect(url string) error
	Disconnect(url string) error
	RegisterSink(url string, sink client.IObjectSink)
	UnregisterSink(url string, sink client.IObjectSink)
	Node(objectId string) *client.Node
}

type OlinkConnector struct {
	conns map[string]*connEntry
}

var _ IOlinkConnector = (*OlinkConnector)(nil)

func NewOlinkConnector() *OlinkConnector {
	return &OlinkConnector{
		conns: make(map[string]*connEntry),
	}
}

// Connect connects to a given url and returns a connection id.
// The connection id can be used to disconnect from the server using Disconnect.
func (c *OlinkConnector) Connect(url string) error {
	log.Info().Str("url", url).Msg("connect")
	_, ok := c.conns[url]
	if ok {
		log.Info().Str("url", url).Msg("connection already exists")
		return nil
	}
	log.Info().Str("url", url).Msg("create new connection")
	conn, err := ws.Dial(context.Background(), url)
	if err != nil {
		return err
	}
	registry := client.NewRegistry()
	node := client.NewNode(registry)
	node.SetOutput(conn)
	conn.SetOutput(node)

	entry := &connEntry{
		conn:     conn,
		id:       nextChannelId(),
		url:      url,
		node:     node,
		registry: registry,
	}
	c.conns[url] = entry
	return nil
}

// Disconnect closes a connection to the server.
// The connection id is the string returned when connecting to the server using Connect.
// If the connection id is not found, this function does nothing and returns nil.
func (c *OlinkConnector) Disconnect(url string) error {
	log.Info().Str("url", url).Msg("disconnect")
	entry, ok := c.conns[url]
	if !ok {
		return nil
	}
	entry.Close()
	delete(c.conns, url)
	return nil
}

func (c *OlinkConnector) Close() {
	for _, conn := range c.conns {
		conn.Close()
	}
	c.conns = nil
}

func (c *OlinkConnector) RegisterSink(url string, sink client.IObjectSink) {
	log.Info().Str("sink", sink.ObjectId()).Msg("register sink")
	entry, ok := c.conns[url]
	if !ok {
		log.Error().Str("url", url).Msg("connection not found")
		return
	}
	entry.registry.AddObjectSink(sink)
	entry.node.LinkRemoteNode(sink.ObjectId())
}

func (c *OlinkConnector) UnregisterSink(url string, sink client.IObjectSink) {
	log.Info().Str("sink", sink.ObjectId()).Msg("unregister sink")
	entry, ok := c.conns[url]
	if !ok {
		log.Error().Str("url", url).Msg("connection not found")
		return
	}
	entry.registry.RemoveObjectSink(sink.ObjectId())
}

func (c *OlinkConnector) Node(url string) *client.Node {
	entry, ok := c.conns[url]
	if !ok {
		log.Error().Str("url", url).Msg("connection not found")
		return nil
	}
	node := entry.node
	if node == nil {
		log.Error().Str("url", url).Msg("node is nil")
		return nil
	}
	return node
}
