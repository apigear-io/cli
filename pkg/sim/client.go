package sim

import (
	"sync"

	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/core"
)

type ObjectClient struct {
	mu           sync.RWMutex
	object       string
	state        map[string]any
	stateEmitter *Emitter[any]
	signals      *Emitter[[]any]
	sink         *ObjectClientSink
	channel      *Channel
}

func NewObjectClient(channel *Channel, objectId string) *ObjectClient {
	c := &ObjectClient{
		object:       objectId,
		state:        make(map[string]any),
		stateEmitter: NewEmitter[any](),
		signals:      NewEmitter[[]any](),
		channel:      channel,
	}
	c.sink = NewObjectClientSink(c)
	c.connector().RegisterSink(channel.url, c.sink)
	return c
}

func (c *ObjectClient) connector() IOlinkConnector {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.channel.connector()
}

func (c *ObjectClient) node() *client.Node {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.channel.node()
}

func (c *ObjectClient) Close() {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.connector().UnregisterSink(c.channel.url, c.sink)
}

func (o *ObjectClient) ObjectId() string {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.object
}

// setLocalProperties
func (o *ObjectClient) setLocalProperties(properties map[string]any) {
	o.mu.Lock()
	defer o.mu.Unlock()
	for name, value := range properties {
		o.state[name] = value
		o.stateEmitter.Emit(name, value)
	}
}

// setLocalProperty
func (o *ObjectClient) setLocalProperty(name string, value any) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.state[name] = value
	o.stateEmitter.Emit(name, value)
}

// SetProperty
func (o *ObjectClient) SetProperty(name string, value any) {
	o.mu.RLock()
	defer o.mu.RUnlock()
	node := o.channel.node()
	if node == nil {
		log.Error().Msg("ObjectClient.SetProperty: node is nil")
		return
	}
	symbol := core.MakeSymbolId(o.object, name)
	node.SetRemoteProperty(symbol, value)
}

func (o *ObjectClient) GetProperty(name string) any {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.state[name]
}

func (o *ObjectClient) OnProperty(name string, fn func(value any)) {
	o.mu.RLock()
	defer o.mu.RUnlock()
	o.stateEmitter.Add(name, fn)
}

func (o *ObjectClient) CallMethod(method string, args ...any) any {
	o.mu.RLock()
	node := o.node()
	o.mu.RUnlock()
	if node == nil {
		log.Error().Msg("ObjectClient.CallMethod: node is nil")
		return nil
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	var reply any
	symbol := core.MakeSymbolId(o.object, method)
	node.InvokeRemote(symbol, core.Args(args), func(arg client.InvokeReplyArg) {
		log.Debug().Interface("arg", arg).Msg("ObjectClient.CallMethod: InvokeRemote: arg")
		reply = arg.Value
		wg.Done()
	})
	wg.Wait()
	return reply
}

func (o *ObjectClient) OnSignal(signal string, fn func(args ...any)) {
	o.mu.RLock()
	defer o.mu.RUnlock()
	symbol := core.MakeSymbolId(o.object, signal)
	o.signals.Add(symbol, func(args []any) {
		fn(args...)
	})
}

func (o *ObjectClient) emitLocalSignal(signal string, args ...any) {
	o.mu.RLock()
	defer o.mu.RUnlock()
	o.signals.Emit(signal, args)
}
