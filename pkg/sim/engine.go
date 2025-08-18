package sim

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/apigear-io/objectlink-core-go/olink/remote"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/dop251/goja_nodejs/require"
)

func createSourceLoader() require.SourceLoader {
	return func(filename string) ([]byte, error) {
		log.Info().Str("filename", filename).Msg("Loading module")
		return os.ReadFile(filename)
	}
}

func createPathResolver(workDir string) require.PathResolver {
	return func(base, path string) string {
		log.Info().Str("base", base).Str("path", path).Msg("Resolving path")

		// If path doesn't have an extension, try adding .js
		if filepath.Ext(path) == "" {
			path = path + ".js"
		}

		// If path is absolute, return as-is
		if filepath.IsAbs(path) {
			return path
		}

		// For relative paths, resolve relative to workDir (which is the script directory)
		resolved := filepath.Join(workDir, path)
		log.Info().Str("resolved", resolved).Msg("Resolved to workDir")
		return resolved
	}
}

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
	registry  *require.Registry
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

	registry := require.NewRegistry(
		require.WithLoader(createSourceLoader()),
		require.WithPathResolver(createPathResolver(opts.WorkDir)),
		require.WithGlobalFolders(opts.WorkDir),
	)
	e := &Engine{
		loop:      eventloop.NewEventLoop(eventloop.WithRegistry(registry)),
		workDir:   opts.WorkDir,
		server:    opts.Server,
		connector: opts.Connector,
		registry:  registry,
	}
	e.world = NewWorld(e)
	e.loop.Start()
	
	// Initial setup - wait for initialization to complete before returning
	// This ensures e.rt is set and the engine is fully ready
	done := make(chan bool)
	e.loop.RunOnLoop(func(rt *goja.Runtime) {
		e.rt = rt  // Set the runtime once during initialization
		rt.SetFieldNameMapper(goja.UncapFieldNameMapper())
		e.world.register(rt)
		registry.Enable(rt)
		done <- true
	})
	<-done  // Wait for initialization to complete
	
	return e
}

func (e *Engine) SetOlinkServer(server IOlinkServer) {
	e.rw.Lock()
	defer e.rw.Unlock()
	e.server = server
}

func (e *Engine) RunScript(name string, content string) {
	e.RunOnLoop(func(rt *goja.Runtime) {
		log.Info().Str("name", name).Str("workDir", e.workDir).Msg("Run script")

		value, err := rt.RunScript(name, content)
		if err != nil {
			log.Error().Err(err).Msg("Failed to run script")
		}
		log.Info().Interface("value", value).Msg("Script result")
	})
}

func (e *Engine) RunFunction(name string, args ...any) {
	e.RunOnLoop(func(rt *goja.Runtime) {
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

func (e *Engine) RunOnLoop(fn func(rt *goja.Runtime)) {
	// No lock needed here - eventloop.RunOnLoop is thread-safe
	// and queues the function to run on the event loop thread
	e.loop.RunOnLoop(func(rt *goja.Runtime) {
		// e.rt is already set during initialization in NewEngine
		// and remains constant throughout the engine's lifetime
		fn(rt)
	})
}

func (e *Engine) Runtime() *goja.Runtime {
	e.rw.RLock()
	defer e.rw.RUnlock()
	return e.rt
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
