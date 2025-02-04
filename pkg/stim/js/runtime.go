package js

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/stim/model"
	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/ws"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/sasha-s/go-deadlock"
)

type Runtime struct {
	id         string
	loop       *eventloop.EventLoop
	vm         *goja.Runtime
	node       *client.Node
	lock       deadlock.Mutex
	IsActive   bool
	LastUpdate time.Time
	registry   *client.Registry
}

func Must(err error) {
	if err != nil {
		log.Fatal().Msgf("error: %s", err)
	}
}

func NewRuntime(id string) *Runtime {
	if id == "" {
		id = "demo"
	}
	log.Info().Str("id", id).Msg("creating runtime")
	rt := &Runtime{
		id:         id,
		node:       nil,
		registry:   client.NewRegistry(),
		IsActive:   false,
		LastUpdate: time.Now(),
	}
	rt.loop = eventloop.NewEventLoop()
	// initial run wait for result
	rt.loop.Run(func(vm *goja.Runtime) {
		log.Info().Msg("loop run loop initial run")
		vm.SetFieldNameMapper(goja.UncapFieldNameMapper())
		Must(vm.Set("connect", rt.connect))
		Must(vm.Set("disconnect", rt.disconnect))
		rt.vm = vm
	})
	// run in background
	rt.loop.Start()
	return rt
}

func (rt *Runtime) RunScript(script model.Script) (any, error) {
	log.Info().Str("script", script.Name).Msg("running script")
	rt.lock.Lock()
	defer rt.lock.Unlock()
	rt.IsActive = true
	rt.LastUpdate = time.Now()
	rt.loop.RunOnLoop(func(vm *goja.Runtime) {
		_, err := vm.RunScript(script.Name, script.Source)
		if err != nil {
			log.Error().Err(err).Msg("run script")
			return
		}
	})

	return nil, nil
}

func (rt *Runtime) connect(url string) (*jsNode, error) {
	if url == "" {
		url = "ws://localhost:5555/ws"
	}
	ctx := context.Background()
	conn, err := ws.Dial(ctx, url)
	if err != nil {
		log.Error().Err(err).Str("url", url).Msg("dialing ws")
		return nil, err
	}
	node := client.NewNode(rt.registry)
	node.SetOutput(conn)
	conn.SetOutput(node)
	rt.node = node
	log.Info().Str("id", conn.Id()).Str("url", url).Msg("connection connected")
	return NewJsNode(rt.vm, node), nil
}

// disconnect is called when the node is disconnected
func (rt *Runtime) disconnect() error {
	if rt.node == nil {
		return fmt.Errorf("not connected")
	}
	log.Info().Str("id", rt.node.Id()).Msg("connection disconnected")
	rt.node.Close()
	rt.node = nil
	return nil
}

// Interrupt is called when the simulation is interrupted
func (rt *Runtime) Interrupt() {
	log.Info().Str("id", rt.id).Msg("interrupting simulation")
	rt.lock.Lock()
	defer rt.lock.Unlock()
	rt.IsActive = false
	rt.loop.Stop()
	rt.vm.Interrupt(errors.New("interrupted"))
}
