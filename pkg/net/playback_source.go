package net

import (
	"fmt"
	"sync"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
)

// PlaybackSourceFactory manages per-object playback sources and dispatches
// recorded ObjectLink messages to the registered nodes.
type PlaybackSourceFactory struct {
	mu      sync.RWMutex
	sources map[string]*PlaybackSource
}

func NewPlaybackSourceFactory() *PlaybackSourceFactory {
	return &PlaybackSourceFactory{sources: make(map[string]*PlaybackSource)}
}

// SourceFactoryFunc returns a remote.SourceFactory compatible function so the
// registry can lazily create playback sources when a node links to an object.
func (f *PlaybackSourceFactory) SourceFactoryFunc() remote.SourceFactory {
	return func(objectID string) remote.IObjectSource {
		return f.getOrCreate(objectID)
	}
}

// Dispatch routes a decoded ObjectLink message to the appropriate playback source.
func (f *PlaybackSourceFactory) Dispatch(msg core.Message) {
	objectID := resolveObjectID(msg)
	if objectID == "" {
		log.Warn().Msg("playback: unable to resolve object id from message")
		return
	}
	src := f.getOrCreate(objectID)
	src.HandleMessage(msg)
}

func (f *PlaybackSourceFactory) getOrCreate(objectID string) *PlaybackSource {
	f.mu.Lock()
	defer f.mu.Unlock()
	if src, ok := f.sources[objectID]; ok {
		return src
	}
	src := NewPlaybackSource(objectID)
	f.sources[objectID] = src
	return src
}

// resolveObjectID extracts the object identifier from a generic ObjectLink message.
func resolveObjectID(msg core.Message) string {
	switch msg.Type() {
	case core.MsgLink, core.MsgInit, core.MsgUnlink:
		return core.AsString(msg[1])
	case core.MsgSetProperty, core.MsgPropertyChange:
		propertyID := core.AsString(msg[1])
		objectID, _ := core.SymbolIdToParts(propertyID)
		return objectID
	case core.MsgInvoke:
		_, methodID, _ := msg.AsInvoke()
		objectID, _ := core.SymbolIdToParts(methodID)
		return objectID
	case core.MsgInvokeReply:
		_, methodID, _ := msg.AsInvokeReply()
		objectID, _ := core.SymbolIdToParts(methodID)
		return objectID
	case core.MsgSignal:
		signalID, _ := msg.AsSignal()
		objectID, _ := core.SymbolIdToParts(signalID)
		return objectID
	default:
		return ""
	}
}

// PlaybackSource implements remote.IObjectSource and replays messages to linked nodes.
type PlaybackSource struct {
	objectID string
	mu       sync.RWMutex
	nodes    map[*remote.Node]struct{}
	props    core.KWArgs
	initMsg  core.Message
}

func NewPlaybackSource(objectID string) *PlaybackSource {
	return &PlaybackSource{
		objectID: objectID,
		nodes:    make(map[*remote.Node]struct{}),
		props:    core.KWArgs{},
	}
}

func (s *PlaybackSource) ObjectId() string {
	return s.objectID
}

func (s *PlaybackSource) Invoke(methodId string, args core.Args) (core.Any, error) {
	return nil, fmt.Errorf("playback source %s: invoke not supported", s.objectID)
}

func (s *PlaybackSource) SetProperty(propertyId string, value core.Any) error {
	return fmt.Errorf("playback source %s: set property not supported", s.objectID)
}

func (s *PlaybackSource) Linked(objectId string, node *remote.Node) error {
	s.mu.Lock()
	s.nodes[node] = struct{}{}
	init := s.initMsg
	s.mu.Unlock()

	if init != nil {
		node.SendMessage(init)
	}
	return nil
}

func (s *PlaybackSource) CollectProperties() (core.KWArgs, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return cloneKWArgs(s.props), nil
}

// HandleMessage updates internal state and broadcasts the message to linked nodes.
func (s *PlaybackSource) HandleMessage(msg core.Message) {
	s.mu.Lock()
	s.updateStateLocked(msg)
	nodes := make([]*remote.Node, 0, len(s.nodes))
	for node := range s.nodes {
		nodes = append(nodes, node)
	}
	outgoing := cloneMessage(msg)
	s.mu.Unlock()

	for _, node := range nodes {
		node.SendMessage(outgoing)
	}
}

func (s *PlaybackSource) updateStateLocked(msg core.Message) {
	switch msg.Type() {
	case core.MsgInit:
		_, props := msg.AsInit()
		s.props = cloneKWArgs(props)
		s.initMsg = core.MakeInitMessage(s.objectID, cloneKWArgs(props))
	case core.MsgPropertyChange:
		propertyID, value := msg.AsPropertyChange()
		objectID, name := core.SymbolIdToParts(propertyID)
		if objectID == s.objectID && name != "" {
			if s.props == nil {
				s.props = core.KWArgs{}
			}
			s.props[name] = value
		}
	case core.MsgSetProperty:
		propertyID, value := msg.AsSetProperty()
		objectID, name := core.SymbolIdToParts(propertyID)
		if objectID == s.objectID && name != "" {
			if s.props == nil {
				s.props = core.KWArgs{}
			}
			s.props[name] = value
		}
	}
}

func cloneKWArgs(in core.KWArgs) core.KWArgs {
	if in == nil {
		return nil
	}
	out := make(core.KWArgs, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

func cloneMessage(msg core.Message) core.Message {
	if msg == nil {
		return nil
	}
	copy := make(core.Message, len(msg))
	for i, v := range msg {
		copy[i] = v
	}
	return copy
}
