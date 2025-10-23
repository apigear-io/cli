package sim

import (
	"context"

	"github.com/apigear-io/cli/pkg/net"
)

type ManagerOptions struct {
}

type Manager struct {
	engine *Engine
	netman *net.NetworkManager
	opts   ManagerOptions
}

func NewManager(opts ManagerOptions) *Manager {
	m := &Manager{
		engine: nil,
		opts:   opts,
	}
	return m
}

func (m *Manager) Start(ctx context.Context, netman *net.NetworkManager) error {
	m.netman = netman
	return nil
}

func (m *Manager) OlinkServer() *net.OlinkServer {
	return m.netman.OlinkServer()
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
	m.engine = NewEngine(EngineOptions{Server: m.OlinkServer(), WorkDir: script.Dir})
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
