package olnk

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/apigear-io/objectlink-core-go/olink/ws"
)

type Feeder struct {
	registry *client.Registry
	conn     *ws.Connection
	node     *client.Node
}

func NewFeeder() *Feeder {
	registry := client.NewRegistry()
	registry.SetSinkFactory(func(objectId string) client.IObjectSink {
		return &ObjectSink{objectId: objectId}
	})
	node := client.NewNode(registry)
	registry.AttachClientNode(node)

	return &Feeder{
		registry: registry,
		node:     node,
	}
}

func (f *Feeder) Connect(ctx context.Context, addr string) error {
	conn, err := ws.Dial(ctx, addr)
	if err != nil {
		return err
	}
	f.conn = conn
	conn.SetOutput(f.node)
	f.node.SetOutput(conn)
	return nil
}

func (f *Feeder) Close() error {
	if f.conn == nil {
		return nil
	}
	err := f.conn.Close()
	if err != nil {
		log.Error().Err(err).Msg("failed to close connection")
		return err
	}
	f.conn = nil
	return nil
}

func (f *Feeder) Feed(data []byte) error {
	var m core.Message
	err := json.Unmarshal(data, &m)
	if err != nil {
		log.Error().Err(err).Msgf("invalid message: %s", data)
		return err
	}
	s, ok := m[0].(string)
	if !ok {
		log.Error().Msgf("invalid message type, expected string: %v", m)
		return fmt.Errorf("invalid message type, expected string: %v", m)
	}
	m[0] = core.MsgTypeFromString(s)
	switch m[0] {
	case core.MsgLink:
		objectId := m.AsLink()
		f.node.LinkRemoteNode(objectId)
	case core.MsgUnlink:
		objectId := m.AsLink()
		f.node.UnlinkRemoteNode(objectId)
	case core.MsgSetProperty:
		propertyId, value := m.AsSetProperty()
		f.node.SetRemoteProperty(propertyId, value)
	case core.MsgInvoke:
		_, methodId, args := m.AsInvoke()
		f.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {
			log.Info().Msgf("<- reply %s : %v", arg.Identifier, arg.Value)
		})
	default:
		log.Info().Msgf("not supported message type: %v", m)
		return fmt.Errorf("not supported message type: %v", m)
	}
	return nil
}
