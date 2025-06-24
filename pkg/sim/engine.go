package sim

import (
	"sync"

	"github.com/apigear-io/objectlink-core-go/olink/remote"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/dop251/goja_nodejs/require"

	_ "embed"
)

//go:embed proxy.js
var proxyJS string

type EngineOptions struct {
	WorkDir   string
	Server    IOlinkServer
	Connector IOlinkConnector
}
type Engine struct {
	rw        sync.RWMutex
	world     *World
	loop      *eventloop.EventLoop
	workDir   string
	server    IOlinkServer
	connector IOlinkConnector
	rt        *goja.Runtime
}

func NewEngine(opts EngineOptions) *Engine {
	log.Info().Msg("NewEngine")
	if opts.WorkDir == "" {
		opts.WorkDir = "."
	}
	if opts.Server == nil {
		opts.Server = NewOlinkServer()
	}
	if opts.Connector == nil {
		opts.Connector = NewOlinkConnector()
	}
	printer := NewLogPrinter(&log)
	require.RegisterCoreModule(console.ModuleName, console.RequireWithPrinter(printer))
	registry := require.NewRegistry(require.WithLoader(require.DefaultSourceLoader), require.WithGlobalFolders(opts.WorkDir))
	e := &Engine{
		loop:      eventloop.NewEventLoop(eventloop.WithRegistry(registry)),
		workDir:   opts.WorkDir,
		server:    opts.Server,
		connector: opts.Connector,
	}
	e.world = NewWorld(e)
	e.loop.Start()
	e.loop.RunOnLoop(func(rt *goja.Runtime) {
		rt.SetFieldNameMapper(goja.UncapFieldNameMapper())
		e.world.register(rt)
		if _, err := rt.RunScript("proxy.js", proxyJS); err != nil {
			log.Error().Err(err).Msg("failed to run proxy.js script")
		}
	})
	return e
}

func (e *Engine) SetOlinkServer(server IOlinkServer) {
	e.rw.Lock()
	defer e.rw.Unlock()
	e.server = server
}

func (e *Engine) RunScript(name string, content string) {
	e.rw.Lock()
	defer e.rw.Unlock()
	e.loop.RunOnLoop(func(rt *goja.Runtime) {
		log.Info().Str("name", name).Msg("Run script")
		value, err := rt.RunScript(name, content)
		if err != nil {
			log.Error().Err(err).Msg("Failed to run script")
		}
		log.Info().Interface("value", value).Msg("Script result")
		e.rt = rt
	})
}

func (e *Engine) RunFunction(name string, args ...any) {
	e.rw.Lock()
	defer e.rw.Unlock()
	e.loop.RunOnLoop(func(rt *goja.Runtime) {
		log.Info().Str("name", name).Msg("Run function")
		fn, ok := goja.AssertFunction(rt.Get(name))
		if !ok {
			log.Error().Str("name", name).Msg("Function not found")
			return
		}
		if fn == nil {
			log.Error().Str("name", name).Msg("Function not found")
			return
		}
		var jsArgs []goja.Value
		for _, arg := range args {
			jsArgs = append(jsArgs, rt.ToValue(arg))
		}
		_, err := fn(goja.Undefined(), jsArgs...)
		if err != nil {
			log.Error().Err(err).Msg("Failed to run function")
		}
	})
}

func (e *Engine) CompileScript(name string, src string) error {
	_, err := goja.Compile(name, src, true)
	if err != nil {
		return err
	}
	return nil
}

func (e *Engine) Close() {
	log.Info().Msg("Stop engine")
	e.rw.Lock()
	defer e.rw.Unlock()
	e.loop.StopNoWait()
	e.loop.Terminate()
}

func (e *Engine) registerSource(source remote.IObjectSource) {
	e.rw.Lock()
	defer e.rw.Unlock()
	if e.server != nil {
		e.server.RegisterSource(source)
	}
}

func (e *Engine) unregisterSource(source remote.IObjectSource) {
	e.rw.Lock()
	defer e.rw.Unlock()
	if e.server != nil {
		e.server.UnregisterSource(source)
	}
}
