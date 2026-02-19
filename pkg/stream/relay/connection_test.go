package relay_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/apigear-io/cli/pkg/stream/relay"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var upgrader = websocket.Upgrader{}

// TestNewConnection verifies Connection creation
func TestNewConnection(t *testing.T) {
	server, clientConn := setupTestConnection(t)
	defer server.Close()
	defer clientConn.Close()

	conn := relay.NewConnection(clientConn, "test-id")
	require.NotNil(t, conn)

	assert.Equal(t, "test-id", conn.ID())
}

// TestConnection_WriteRead tests basic read/write operations
func TestConnection_WriteRead(t *testing.T) {
	server, clientWS, serverWS := setupTestConnectionPair(t)
	defer server.Close()

	clientConn := relay.NewConnection(clientWS, "client")

	// Write message
	testData := []byte("hello")
	err := clientConn.WriteMessage(websocket.TextMessage, testData)
	require.NoError(t, err)

	// Read on server side
	msgType, data, err := serverWS.ReadMessage()
	require.NoError(t, err)
	assert.Equal(t, websocket.TextMessage, msgType)
	assert.Equal(t, testData, data)
}

// TestConnection_Done verifies Done channel behavior
func TestConnection_Done(t *testing.T) {
	server, clientWS := setupTestConnection(t)
	defer server.Close()

	conn := relay.NewConnection(clientWS, "client")

	// Done should not be closed initially
	select {
	case <-conn.Done():
		t.Fatal("Done channel closed before Close()")
	case <-time.After(10 * time.Millisecond):
		// Expected
	}

	// Close connection
	err := conn.Close()
	require.NoError(t, err)

	// Done should be closed now
	select {
	case <-conn.Done():
		// Expected
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Done channel not closed after Close()")
	}
}

// TestConnection_CloseIdempotent verifies Close can be called multiple times
func TestConnection_CloseIdempotent(t *testing.T) {
	server, clientWS := setupTestConnection(t)
	defer server.Close()

	conn := relay.NewConnection(clientWS, "client")

	// Close multiple times
	err1 := conn.Close()
	err2 := conn.Close()
	err3 := conn.Close()

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NoError(t, err3)
}

// TestNewConnectionPool verifies ConnectionPool creation
func TestNewConnectionPool(t *testing.T) {
	pool := relay.NewConnectionPool()
	require.NotNil(t, pool)

	assert.Equal(t, 0, pool.Size())
	assert.Empty(t, pool.List())
}

// TestConnectionPool_Add tests adding connections
func TestConnectionPool_Add(t *testing.T) {
	pool := relay.NewConnectionPool()
	server, clientWS := setupTestConnection(t)
	defer server.Close()

	conn := relay.NewConnection(clientWS, "client")

	// Add with auto-generated ID
	id := pool.Add(conn)
	assert.NotEmpty(t, id)
	assert.Equal(t, 1, pool.Size())

	// Verify we can retrieve it
	retrieved, err := pool.Get(id)
	require.NoError(t, err)
	assert.Equal(t, conn, retrieved)
}

// TestConnectionPool_AddWithID tests adding with custom ID
func TestConnectionPool_AddWithID(t *testing.T) {
	pool := relay.NewConnectionPool()
	server, clientWS := setupTestConnection(t)
	defer server.Close()

	conn := relay.NewConnection(clientWS, "client")

	// Add with custom ID
	err := pool.AddWithID("custom-id", conn)
	require.NoError(t, err)
	assert.Equal(t, 1, pool.Size())

	// Verify we can retrieve it
	retrieved, err := pool.Get("custom-id")
	require.NoError(t, err)
	assert.Equal(t, conn, retrieved)
}

// TestConnectionPool_AddWithID_Duplicate tests duplicate ID handling
func TestConnectionPool_AddWithID_Duplicate(t *testing.T) {
	pool := relay.NewConnectionPool()
	server1, clientWS1 := setupTestConnection(t)
	defer server1.Close()
	server2, clientWS2 := setupTestConnection(t)
	defer server2.Close()

	conn1 := relay.NewConnection(clientWS1, "client1")
	conn2 := relay.NewConnection(clientWS2, "client2")

	// Add first connection
	err := pool.AddWithID("same-id", conn1)
	require.NoError(t, err)

	// Try to add second with same ID
	err = pool.AddWithID("same-id", conn2)
	assert.ErrorIs(t, err, relay.ErrDuplicateConnection)
	assert.Equal(t, 1, pool.Size())
}

// TestConnectionPool_Get_NotFound tests getting non-existent connection
func TestConnectionPool_Get_NotFound(t *testing.T) {
	pool := relay.NewConnectionPool()

	conn, err := pool.Get("non-existent")
	assert.ErrorIs(t, err, relay.ErrConnectionNotFound)
	assert.Nil(t, conn)
}

// TestConnectionPool_Remove tests removing connections
func TestConnectionPool_Remove(t *testing.T) {
	pool := relay.NewConnectionPool()
	server, clientWS := setupTestConnection(t)
	defer server.Close()

	conn := relay.NewConnection(clientWS, "client")
	id := pool.Add(conn)
	assert.Equal(t, 1, pool.Size())

	// Remove it
	err := pool.Remove(id)
	require.NoError(t, err)
	assert.Equal(t, 0, pool.Size())

	// Verify it's gone
	_, err = pool.Get(id)
	assert.ErrorIs(t, err, relay.ErrConnectionNotFound)
}

// TestConnectionPool_Remove_NotFound tests removing non-existent connection
func TestConnectionPool_Remove_NotFound(t *testing.T) {
	pool := relay.NewConnectionPool()

	err := pool.Remove("non-existent")
	assert.ErrorIs(t, err, relay.ErrConnectionNotFound)
}

// TestConnectionPool_List tests listing connections
func TestConnectionPool_List(t *testing.T) {
	pool := relay.NewConnectionPool()

	// Empty list
	assert.Empty(t, pool.List())

	// Add connections
	server1, clientWS1 := setupTestConnection(t)
	defer server1.Close()
	server2, clientWS2 := setupTestConnection(t)
	defer server2.Close()

	conn1 := relay.NewConnection(clientWS1, "client1")
	conn2 := relay.NewConnection(clientWS2, "client2")

	id1 := pool.Add(conn1)
	id2 := pool.Add(conn2)

	// List should contain both IDs
	ids := pool.List()
	assert.Len(t, ids, 2)
	assert.Contains(t, ids, id1)
	assert.Contains(t, ids, id2)
}

// TestConnectionPool_Close tests closing all connections
func TestConnectionPool_Close(t *testing.T) {
	pool := relay.NewConnectionPool()

	// Add connections
	server1, clientWS1 := setupTestConnection(t)
	defer server1.Close()
	server2, clientWS2 := setupTestConnection(t)
	defer server2.Close()

	conn1 := relay.NewConnection(clientWS1, "client1")
	conn2 := relay.NewConnection(clientWS2, "client2")

	pool.Add(conn1)
	pool.Add(conn2)
	assert.Equal(t, 2, pool.Size())

	// Close pool
	err := pool.Close()
	require.NoError(t, err)

	// Pool should be empty
	assert.Equal(t, 0, pool.Size())

	// Further operations should fail
	_, err = pool.Get("any-id")
	assert.ErrorIs(t, err, relay.ErrPoolClosed)
}

// TestConnectionPool_Close_Idempotent verifies Close can be called multiple times
func TestConnectionPool_Close_Idempotent(t *testing.T) {
	pool := relay.NewConnectionPool()

	err1 := pool.Close()
	err2 := pool.Close()

	assert.NoError(t, err1)
	assert.ErrorIs(t, err2, relay.ErrPoolClosed)
}

// Helper function to create a test WebSocket connection pair (client only)
func setupTestConnection(t *testing.T) (*httptest.Server, *websocket.Conn) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		require.NoError(t, err)
		defer conn.Close()

		// Keep connection open
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				return
			}
		}
	}))

	// Connect client
	url := "ws" + strings.TrimPrefix(server.URL, "http")
	clientConn, _, err := websocket.DefaultDialer.Dial(url, nil)
	require.NoError(t, err)

	return server, clientConn
}

// Helper function to create a test WebSocket connection pair (both sides)
func setupTestConnectionPair(t *testing.T) (*httptest.Server, *websocket.Conn, *websocket.Conn) {
	serverConnChan := make(chan *websocket.Conn, 1)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		require.NoError(t, err)
		serverConnChan <- conn

		// Block until request context is done (server closes)
		<-r.Context().Done()
	}))

	// Connect client
	url := "ws" + strings.TrimPrefix(server.URL, "http")
	clientConn, _, err := websocket.DefaultDialer.Dial(url, nil)
	require.NoError(t, err)

	// Wait for server connection
	var serverConn *websocket.Conn
	select {
	case serverConn = <-serverConnChan:
	case <-time.After(1 * time.Second):
		t.Fatal("Timeout waiting for server connection")
	}

	return server, clientConn, serverConn
}
