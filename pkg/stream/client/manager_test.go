package client

import (
	"testing"

	"github.com/apigear-io/cli/pkg/stream/config"
)

func TestNewClient(t *testing.T) {
	client := NewClient("test", "ws://localhost:5560/ws", []string{"demo.Counter"}, true, false)

	if client.name != "test" {
		t.Errorf("expected name test, got %s", client.name)
	}
	if client.url != "ws://localhost:5560/ws" {
		t.Errorf("expected url ws://localhost:5560/ws, got %s", client.url)
	}
	if len(client.interfaces) != 1 || client.interfaces[0] != "demo.Counter" {
		t.Errorf("expected interfaces [demo.Counter], got %v", client.interfaces)
	}
	if !client.autoReconnect {
		t.Error("expected autoReconnect to be true")
	}
	if client.enabled {
		t.Error("expected enabled to be false")
	}
	if client.Status() != StatusDisconnected {
		t.Errorf("expected status disconnected, got %s", client.Status())
	}
}

func TestClientInfo(t *testing.T) {
	client := NewClient("test", "ws://localhost:5560/ws", []string{"demo.Counter"}, true, true)
	info := client.Info()

	if info.Name != "test" {
		t.Errorf("expected name test, got %s", info.Name)
	}
	if info.URL != "ws://localhost:5560/ws" {
		t.Errorf("expected URL ws://localhost:5560/ws, got %s", info.URL)
	}
	if len(info.Interfaces) != 1 || info.Interfaces[0] != "demo.Counter" {
		t.Errorf("expected interfaces [demo.Counter], got %v", info.Interfaces)
	}
	if !info.AutoReconnect {
		t.Error("expected AutoReconnect to be true")
	}
	if !info.Enabled {
		t.Error("expected Enabled to be true")
	}
	if info.Status != StatusDisconnected {
		t.Errorf("expected status disconnected, got %s", info.Status)
	}
}

func TestNewManager(t *testing.T) {
	manager := NewManager()

	if manager == nil {
		t.Fatal("expected manager to be created")
	}

	clients := manager.ListClients()
	if len(clients) != 0 {
		t.Errorf("expected no clients, got %d", len(clients))
	}
}

func TestManagerAddClient(t *testing.T) {
	manager := NewManager()

	cfg := config.ClientConfig{
		URL:           "ws://localhost:5560/ws",
		Interfaces:    []string{"demo.Counter"},
		Enabled:       false,
		AutoReconnect: true,
	}

	err := manager.AddClient("test", cfg)
	if err != nil {
		t.Fatalf("AddClient failed: %v", err)
	}

	// Try to add the same client again
	err = manager.AddClient("test", cfg)
	if err == nil {
		t.Error("expected error when adding duplicate client")
	}

	// Verify client exists
	client, err := manager.GetClient("test")
	if err != nil {
		t.Fatalf("GetClient failed: %v", err)
	}

	if client.name != "test" {
		t.Errorf("expected name test, got %s", client.name)
	}

	// Verify list
	clients := manager.ListClients()
	if len(clients) != 1 {
		t.Errorf("expected 1 client, got %d", len(clients))
	}
	if clients[0].Name != "test" {
		t.Errorf("expected client name test, got %s", clients[0].Name)
	}
}

func TestManagerRemoveClient(t *testing.T) {
	manager := NewManager()

	cfg := config.ClientConfig{
		URL:           "ws://localhost:5560/ws",
		Interfaces:    []string{"demo.Counter"},
		Enabled:       false,
		AutoReconnect: false,
	}

	err := manager.AddClient("test", cfg)
	if err != nil {
		t.Fatalf("AddClient failed: %v", err)
	}

	// Remove client
	err = manager.RemoveClient("test")
	if err != nil {
		t.Fatalf("RemoveClient failed: %v", err)
	}

	// Verify client doesn't exist
	_, err = manager.GetClient("test")
	if err == nil {
		t.Error("expected error when getting removed client")
	}

	// Try to remove non-existent client
	err = manager.RemoveClient("test")
	if err == nil {
		t.Error("expected error when removing non-existent client")
	}
}

func TestManagerLoadFromConfig(t *testing.T) {
	manager := NewManager()

	clients := map[string]config.ClientConfig{
		"client1": {
			URL:           "ws://localhost:5560/ws",
			Interfaces:    []string{"demo.Counter"},
			Enabled:       false,
			AutoReconnect: true,
		},
		"client2": {
			URL:           "ws://localhost:5561/ws",
			Interfaces:    []string{"demo.Calculator"},
			Enabled:       false,
			AutoReconnect: false,
		},
	}

	err := manager.LoadFromConfig(clients)
	if err != nil {
		t.Fatalf("LoadFromConfig failed: %v", err)
	}

	// Verify clients loaded
	clientList := manager.ListClients()
	if len(clientList) != 2 {
		t.Errorf("expected 2 clients, got %d", len(clientList))
	}

	// Verify client1
	client1, err := manager.GetClient("client1")
	if err != nil {
		t.Fatalf("GetClient failed: %v", err)
	}
	if client1.url != "ws://localhost:5560/ws" {
		t.Errorf("expected client1 url ws://localhost:5560/ws, got %s", client1.url)
	}

	// Verify client2
	client2, err := manager.GetClient("client2")
	if err != nil {
		t.Fatalf("GetClient failed: %v", err)
	}
	if client2.url != "ws://localhost:5561/ws" {
		t.Errorf("expected client2 url ws://localhost:5561/ws, got %s", client2.url)
	}
}

func TestManagerClose(t *testing.T) {
	manager := NewManager()

	cfg := config.ClientConfig{
		URL:           "ws://localhost:5560/ws",
		Interfaces:    []string{"demo.Counter"},
		Enabled:       false,
		AutoReconnect: false,
	}

	err := manager.AddClient("test", cfg)
	if err != nil {
		t.Fatalf("AddClient failed: %v", err)
	}

	// Close manager
	err = manager.Close()
	if err != nil {
		t.Fatalf("Close failed: %v", err)
	}

	// Verify all clients removed
	clients := manager.ListClients()
	if len(clients) != 0 {
		t.Errorf("expected no clients after close, got %d", len(clients))
	}
}
