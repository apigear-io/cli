package js

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/dop251/goja"
)

type jsNode struct {
	vm     *goja.Runtime
	node   *client.Node
	actors map[string]*jsActor
}

func NewJsNode(vm *goja.Runtime, node *client.Node) *jsNode {
	s := &jsNode{
		vm:     vm,
		node:   node,
		actors: map[string]*jsActor{},
	}
	return s
}

func (n *jsNode) GetProperty(objectId string) (any, error) {
	log.Info().Msgf("-> get %s", objectId)
	if n.node == nil {
		return nil, fmt.Errorf("not connected")
	}
	methodId := core.MakeSymbolId(objectId, "$get")
	log.Info().Msgf("-> invoke %s(%v)", methodId, core.Args{})
	return n.node.InvokeRemoteSync(methodId, core.Args{})
}

func (n *jsNode) SetProperty(objectId string, property string, value interface{}) error {
	log.Info().Msgf("-> set %s.%s = %v", objectId, property, value)
	if n.node == nil {
		return fmt.Errorf("not connected")
	}
	log.Info().Msgf("-> set %s/%s = %v", objectId, property, value)
	propertyId := core.MakeSymbolId(objectId, property)
	n.node.SetRemoteProperty(propertyId, value)
	return nil
}

// init is called when the node is initialized
func (n *jsNode) Link(objectId string) (*jsActor, error) {
	log.Info().Msgf("-> link %s", objectId)
	if n.actors[objectId] != nil {
		log.Info().Msgf("already linked %s", objectId)
		return n.actors[objectId], nil
	}
	if n.node == nil {
		log.Warn().Msg("no node")
		return nil, fmt.Errorf("not connected")
	}
	if n.node.Registry() == nil {
		log.Warn().Msg("no registry")
		return nil, fmt.Errorf("not registered")
	}
	registry := n.node.Registry()
	if registry.ObjectSink(objectId) == nil {
		log.Info().Msgf("register sink for %s", objectId)
		sink := &jsSink{
			objectId: objectId,
		}
		err := registry.AddObjectSink(sink)
		if err != nil {
			log.Warn().Err(err).Msg("failed to add sink")
			return nil, err
		}
	} else {
		log.Info().Msgf("sink for %s already registered", objectId)
	}
	log.Info().Msgf("-> link %s", objectId)
	n.node.LinkRemoteNode(objectId)
	actor := NewJsActor(objectId, n)
	n.actors[objectId] = actor
	return actor, nil
}

// Unlink is called when the node is disconnected
func (n *jsNode) Unlink(objectId string) error {
	log.Info().Msgf("-> unlink %s", objectId)
	if n.node == nil {
		return fmt.Errorf("not connected")
	}
	if n.actors[objectId] != nil {
		delete(n.actors, objectId)
	}
	log.Info().Msgf("-> unlink %s", objectId)
	n.node.UnlinkRemoteNode(objectId)
	registry := n.node.Registry()
	if registry != nil {
		registry.RemoveObjectSink(objectId)
	}
	return nil
}

func (n *jsNode) Invoke(objectId string, method string, args core.Args) (any, error) {
	log.Info().Msgf("-> invoke %s.%s(%v)", objectId, method, args)
	if n.node == nil {
		return nil, fmt.Errorf("not connected")
	}
	methodId := core.MakeSymbolId(objectId, method)
	return n.node.InvokeRemoteSync(methodId, args)
}

func (n *jsNode) Signal(objectId string, signal string, args core.Args) error {
	log.Info().Msgf("-> signal %s.%s(%v)", objectId, signal, args)
	if n.node == nil {
		return fmt.Errorf("not connected")
	}
	methodId := core.MakeSymbolId(objectId, fmt.Sprintf("$signal.%s", signal))
	n.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {})
	return nil
}
