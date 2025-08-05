package sim

import (
	"github.com/apigear-io/cli/pkg/net"
)

type ManagerOptions struct {
	Server IOlinkServer
}

type Manager struct {
	engine *Engine
	server IOlinkServer
}

func NewManager(opts ManagerOptions) *Manager {
	m := &Manager{
		engine: nil,
		server: opts.Server,
	}
	return m
}

func (m *Manager) Start(netman *net.NetworkManager) {
	server := NewOlinkServer()
	addr := netman.HttpServer().Address()
	log.Info().Msgf("starting Olink server at ws://%s/ws", addr)
	netman.HttpServer().Router().Handle("/ws", server)
	m.server = server
}

func (m *Manager) Stop() {
	if m.engine != nil {
		m.engine.Close()
	}
}

func (m *Manager) ScriptRun(script Script) string {
	log.Info().Msgf("manager run script %s", script)
	if m.engine != nil {
		m.engine.Close()
	}
	m.engine = NewEngine(EngineOptions{Server: m.server, WorkDir: script.Dir})
	m.engine.RunScript(script.Name, script.Content)
	log.Info().Msgf("manager running script %s", script.Name)
	return script.Name
}

func (m *Manager) ScriptStop(worldId string) error {
	log.Info().Msgf("manager stopping script %s", worldId)
	if m.engine != nil {
		m.engine.Close()
	}
	return nil
}

func (m *Manager) FunctionRun(fn string, args []any) {
	log.Info().Msgf("manager run function %s", fn)
	if m.engine == nil {
		return
	}
	m.engine.RunFunction(fn, args...)
}
