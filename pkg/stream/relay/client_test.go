package relay_test

import (
	"testing"
	"time"

	"github.com/apigear-io/cli/pkg/stream/relay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// MockClient is a test implementation of Client interface
type MockClient struct {
	name  string
	url   string
	state relay.State
}

func (m *MockClient) Name() string                               { return m.name }
func (m *MockClient) URL() string                                { return m.url }
func (m *MockClient) State() relay.State                       { return m.state }
func (m *MockClient) Start() error                               { return nil }
func (m *MockClient) Stop() error                                { return nil }
func (m *MockClient) Connect() error                             { return nil }
func (m *MockClient) Disconnect()                                {}
func (m *MockClient) SendRaw(messageType int, data []byte) error { return nil }

// TestNewClientRegistry verifies ClientRegistry creation
func TestNewClientRegistry(t *testing.T) {
	registry := relay.NewClientRegistry()
	require.NotNil(t, registry)

	assert.Equal(t, 0, registry.Size())
	assert.Empty(t, registry.Names())
	assert.Empty(t, registry.List())
}

// TestClientRegistry_Add tests adding clients
func TestClientRegistry_Add(t *testing.T) {
	registry := relay.NewClientRegistry()

	client := &MockClient{name: "test-client", url: "ws://localhost:8080/ws"}

	err := registry.Add(client)
	require.NoError(t, err)

	assert.Equal(t, 1, registry.Size())
	assert.True(t, registry.Has("test-client"))
}

// TestClientRegistry_Add_Duplicate tests duplicate client handling
func TestClientRegistry_Add_Duplicate(t *testing.T) {
	registry := relay.NewClientRegistry()

	client1 := &MockClient{name: "same-name", url: "ws://localhost:8080/ws"}
	client2 := &MockClient{name: "same-name", url: "ws://localhost:9090/ws"}

	err := registry.Add(client1)
	require.NoError(t, err)

	err = registry.Add(client2)
	assert.ErrorIs(t, err, relay.ErrClientAlreadyExists)
	assert.Equal(t, 1, registry.Size())
}

// TestClientRegistry_Get tests retrieving clients
func TestClientRegistry_Get(t *testing.T) {
	registry := relay.NewClientRegistry()

	client := &MockClient{name: "test-client", url: "ws://localhost:8080/ws"}
	err := registry.Add(client)
	require.NoError(t, err)

	retrieved, err := registry.Get("test-client")
	require.NoError(t, err)
	assert.Equal(t, client, retrieved)
}

// TestClientRegistry_Get_NotFound tests getting non-existent client
func TestClientRegistry_Get_NotFound(t *testing.T) {
	registry := relay.NewClientRegistry()

	client, err := registry.Get("non-existent")
	assert.ErrorIs(t, err, relay.ErrClientNotFound)
	assert.Nil(t, client)
}

// TestClientRegistry_Remove tests removing clients
func TestClientRegistry_Remove(t *testing.T) {
	registry := relay.NewClientRegistry()

	client := &MockClient{name: "test-client", url: "ws://localhost:8080/ws"}
	err := registry.Add(client)
	require.NoError(t, err)

	err = registry.Remove("test-client")
	require.NoError(t, err)

	assert.Equal(t, 0, registry.Size())
	assert.False(t, registry.Has("test-client"))
}

// TestClientRegistry_Remove_NotFound tests removing non-existent client
func TestClientRegistry_Remove_NotFound(t *testing.T) {
	registry := relay.NewClientRegistry()

	err := registry.Remove("non-existent")
	assert.ErrorIs(t, err, relay.ErrClientNotFound)
}

// TestClientRegistry_List tests listing all clients
func TestClientRegistry_List(t *testing.T) {
	registry := relay.NewClientRegistry()

	client1 := &MockClient{name: "client1", url: "ws://localhost:8080/ws"}
	client2 := &MockClient{name: "client2", url: "ws://localhost:9090/ws"}

	registry.Add(client1)
	registry.Add(client2)

	clients := registry.List()
	assert.Len(t, clients, 2)
	assert.Contains(t, clients, client1)
	assert.Contains(t, clients, client2)
}

// TestClientRegistry_Names tests getting client names
func TestClientRegistry_Names(t *testing.T) {
	registry := relay.NewClientRegistry()

	client1 := &MockClient{name: "client1", url: "ws://localhost:8080/ws"}
	client2 := &MockClient{name: "client2", url: "ws://localhost:9090/ws"}

	registry.Add(client1)
	registry.Add(client2)

	names := registry.Names()
	assert.Len(t, names, 2)
	assert.Contains(t, names, "client1")
	assert.Contains(t, names, "client2")
}

// TestClientRegistry_Has tests checking client existence
func TestClientRegistry_Has(t *testing.T) {
	registry := relay.NewClientRegistry()

	client := &MockClient{name: "test-client", url: "ws://localhost:8080/ws"}
	registry.Add(client)

	assert.True(t, registry.Has("test-client"))
	assert.False(t, registry.Has("non-existent"))
}

// TestClientRegistry_Clear tests clearing all clients
func TestClientRegistry_Clear(t *testing.T) {
	registry := relay.NewClientRegistry()

	client1 := &MockClient{name: "client1", url: "ws://localhost:8080/ws"}
	client2 := &MockClient{name: "client2", url: "ws://localhost:9090/ws"}

	registry.Add(client1)
	registry.Add(client2)
	assert.Equal(t, 2, registry.Size())

	registry.Clear()

	assert.Equal(t, 0, registry.Size())
	assert.Empty(t, registry.List())
}

// TestClientRegistry_StopAll tests stopping all clients
func TestClientRegistry_StopAll(t *testing.T) {
	registry := relay.NewClientRegistry()

	client1 := &MockClient{name: "client1", url: "ws://localhost:8080/ws"}
	client2 := &MockClient{name: "client2", url: "ws://localhost:9090/ws"}

	registry.Add(client1)
	registry.Add(client2)

	err := registry.StopAll()
	require.NoError(t, err)

	// Registry should be cleared
	assert.Equal(t, 0, registry.Size())
}

// TestNewEventHub verifies EventHub creation
func TestNewEventHub(t *testing.T) {
	hub := relay.NewEventHub[string](100)
	require.NotNil(t, hub)

	messages := hub.GetMessageBuffer()
	assert.Empty(t, messages)

	statuses := hub.GetAllStatuses()
	assert.Empty(t, statuses)
}

// TestEventHub_UpdateStatus tests updating client status
func TestEventHub_UpdateStatus(t *testing.T) {
	hub := relay.NewEventHub[string](100)

	status := relay.Status{
		Name:  "client1",
		URL:   "ws://localhost:8080/ws",
		State: relay.StateConnected,
	}

	hub.UpdateStatus(status)

	retrieved := hub.GetStatus("client1")
	require.NotNil(t, retrieved)
	assert.Equal(t, "client1", retrieved.Name)
	assert.Equal(t, relay.StateConnected, retrieved.State)
}

// TestEventHub_GetStatus_NotFound tests getting non-existent status
func TestEventHub_GetStatus_NotFound(t *testing.T) {
	hub := relay.NewEventHub[string](100)

	status := hub.GetStatus("non-existent")
	assert.Nil(t, status)
}

// TestEventHub_GetAllStatuses tests getting all statuses
func TestEventHub_GetAllStatuses(t *testing.T) {
	hub := relay.NewEventHub[string](100)

	hub.UpdateStatus(relay.Status{Name: "client1", State: relay.StateConnected})
	hub.UpdateStatus(relay.Status{Name: "client2", State: relay.StateDisconnected})

	statuses := hub.GetAllStatuses()
	assert.Len(t, statuses, 2)
}

// TestEventHub_RemoveStatus tests removing a status
func TestEventHub_RemoveStatus(t *testing.T) {
	hub := relay.NewEventHub[string](100)

	hub.UpdateStatus(relay.Status{Name: "client1", State: relay.StateConnected})
	assert.NotNil(t, hub.GetStatus("client1"))

	hub.RemoveStatus("client1")
	assert.Nil(t, hub.GetStatus("client1"))
}

// TestEventHub_SubscribeStatus tests status subscriptions
func TestEventHub_SubscribeStatus(t *testing.T) {
	hub := relay.NewEventHub[string](100)

	// Subscribe
	ch := hub.SubscribeStatus()
	defer hub.UnsubscribeStatus(ch)

	// Update status
	status := relay.Status{Name: "client1", State: relay.StateConnected}
	hub.UpdateStatus(status)

	// Receive update
	select {
	case received := <-ch:
		assert.Equal(t, "client1", received.Name)
		assert.Equal(t, relay.StateConnected, received.State)
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Did not receive status update")
	}
}

// TestEventHub_UnsubscribeStatus tests unsubscribing from status
func TestEventHub_UnsubscribeStatus(t *testing.T) {
	hub := relay.NewEventHub[string](100)

	ch := hub.SubscribeStatus()
	hub.UnsubscribeStatus(ch)

	// Channel should be closed
	select {
	case _, ok := <-ch:
		assert.False(t, ok, "Channel should be closed")
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Channel not closed after unsubscribe")
	}
}

// TestEventHub_PublishMessage tests publishing messages
func TestEventHub_PublishMessage(t *testing.T) {
	hub := relay.NewEventHub[string](100)

	// Publish
	hub.PublishMessage("test message")

	// Check buffer
	messages := hub.GetMessageBuffer()
	assert.Len(t, messages, 1)
	assert.Equal(t, "test message", messages[0])
}

// TestEventHub_SubscribeMessages tests message subscriptions
func TestEventHub_SubscribeMessages(t *testing.T) {
	hub := relay.NewEventHub[string](100)

	// Subscribe
	ch := hub.SubscribeMessages()
	defer hub.UnsubscribeMessages(ch)

	// Publish
	hub.PublishMessage("test message")

	// Receive
	select {
	case msg := <-ch:
		assert.Equal(t, "test message", msg)
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Did not receive message")
	}
}

// TestEventHub_UnsubscribeMessages tests unsubscribing from messages
func TestEventHub_UnsubscribeMessages(t *testing.T) {
	hub := relay.NewEventHub[string](100)

	ch := hub.SubscribeMessages()
	hub.UnsubscribeMessages(ch)

	// Channel should be closed
	select {
	case _, ok := <-ch:
		assert.False(t, ok, "Channel should be closed")
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Channel not closed after unsubscribe")
	}
}

// TestEventHub_ClearMessageBuffer tests clearing message buffer
func TestEventHub_ClearMessageBuffer(t *testing.T) {
	hub := relay.NewEventHub[string](100)

	hub.PublishMessage("msg1")
	hub.PublishMessage("msg2")
	assert.Len(t, hub.GetMessageBuffer(), 2)

	hub.ClearMessageBuffer()

	assert.Empty(t, hub.GetMessageBuffer())
}

// TestState_Constants verifies state constants
func TestState_Constants(t *testing.T) {
	assert.Equal(t, relay.State("disconnected"), relay.StateDisconnected)
	assert.Equal(t, relay.State("connecting"), relay.StateConnecting)
	assert.Equal(t, relay.State("connected"), relay.StateConnected)
	assert.Equal(t, relay.State("retrying"), relay.StateRetrying)
}

// TestStatus_Fields verifies Status struct fields
func TestStatus_Fields(t *testing.T) {
	now := time.Now().Unix()

	status := relay.Status{
		Name:        "test-client",
		URL:         "ws://localhost:8080/ws",
		State:       relay.StateConnected,
		RetryCount:  5,
		LastError:   "connection failed",
		ConnectedAt: &now,
	}

	assert.Equal(t, "test-client", status.Name)
	assert.Equal(t, "ws://localhost:8080/ws", status.URL)
	assert.Equal(t, relay.StateConnected, status.State)
	assert.Equal(t, 5, status.RetryCount)
	assert.Equal(t, "connection failed", status.LastError)
	assert.Equal(t, now, *status.ConnectedAt)
}
