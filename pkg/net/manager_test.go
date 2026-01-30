package net

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewManager(t *testing.T) {
	t.Run("creates new network manager", func(t *testing.T) {
		manager := NewManager()
		assert.NotNil(t, manager)
	})
}

func TestDefaultOptions(t *testing.T) {
	t.Run("has correct default values", func(t *testing.T) {
		opts := DefaultOptions
		assert.Equal(t, "localhost:5555", opts.HttpAddr)
		assert.False(t, opts.HttpDisabled)
		assert.False(t, opts.MonitorDisabled)
		assert.False(t, opts.ObjectAPIDisabled)
		assert.False(t, opts.Logging)
	})
}

func TestNetworkManagerHttpServer(t *testing.T) {
	t.Run("returns nil when http server not started", func(t *testing.T) {
		manager := NewManager()
		assert.Nil(t, manager.HttpServer())
	})
}

func TestNetworkManagerMonitorEmitter(t *testing.T) {
	t.Run("returns monitor emitter", func(t *testing.T) {
		manager := NewManager()
		emitter := manager.MonitorEmitter()
		assert.NotNil(t, emitter)
	})
}

func TestNetworkManagerGetMonitorAddress(t *testing.T) {
	t.Run("returns error when http server not started", func(t *testing.T) {
		manager := NewManager()
		addr, err := manager.GetMonitorAddress()
		assert.Error(t, err)
		assert.Empty(t, addr)
		assert.Contains(t, err.Error(), "http server not started")
	})
}

func TestNetworkManagerGetSimulationAddress(t *testing.T) {
	t.Run("returns error when http server not started", func(t *testing.T) {
		manager := NewManager()
		addr, err := manager.GetSimulationAddress()
		assert.Error(t, err)
		assert.Empty(t, addr)
		assert.Contains(t, err.Error(), "http server not started")
	})
}

func TestNetworkManagerEnableMonitor(t *testing.T) {
	t.Run("returns error when http server not started", func(t *testing.T) {
		manager := NewManager()
		err := manager.EnableMonitor()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "http server not started")
	})
}

func TestNetworkManagerStopHTTP(t *testing.T) {
	t.Run("handles stop when no http server running", func(t *testing.T) {
		manager := NewManager()
		err := manager.StopHTTP()
		assert.NoError(t, err)
	})
}

func TestNetworkManagerStop(t *testing.T) {
	t.Run("stops manager without errors", func(t *testing.T) {
		manager := NewManager()
		err := manager.Stop()
		assert.NoError(t, err)
	})
}
