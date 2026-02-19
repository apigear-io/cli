package relay_test

import (
	"sync"
	"testing"
	"time"

	"github.com/apigear-io/cli/pkg/stream/relay"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestConnectionPool_ConcurrentAddRemove tests concurrent pool operations
func TestConnectionPool_ConcurrentAddRemove(t *testing.T) {
	pool := relay.NewConnectionPool()
	const numGoroutines = 10
	const opsPerGoroutine = 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines * 2) // Add and Remove goroutines

	// Concurrent adds
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < opsPerGoroutine; j++ {
				server, clientWS := setupTestConnection(t)
				conn := relay.NewConnection(clientWS, "client")
				pool.Add(conn)
				server.Close()
				clientWS.Close()
			}
		}(i)
	}

	// Concurrent removes
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < opsPerGoroutine; j++ {
				ids := pool.List()
				if len(ids) > 0 {
					pool.Remove(ids[0])
				}
				time.Sleep(time.Microsecond)
			}
		}()
	}

	wg.Wait()
}

// TestConnectionPool_ConcurrentGetList tests concurrent reads
func TestConnectionPool_ConcurrentGetList(t *testing.T) {
	pool := relay.NewConnectionPool()

	// Add some connections
	type serverInfo struct {
		closer func()
		id     string
	}
	servers := make([]serverInfo, 20)
	for i := 0; i < 20; i++ {
		server, clientWS := setupTestConnection(t)
		conn := relay.NewConnection(clientWS, "client")
		id := pool.Add(conn)
		servers[i] = serverInfo{
			closer: func() {
				server.Close()
				clientWS.Close()
			},
			id: id,
		}
	}

	const numGoroutines = 50
	const opsPerGoroutine = 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines * 2)

	// Concurrent Gets
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < opsPerGoroutine; j++ {
				ids := pool.List()
				if len(ids) > 0 {
					pool.Get(ids[0])
				}
			}
		}()
	}

	// Concurrent Lists
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < opsPerGoroutine; j++ {
				pool.List()
				pool.Size()
			}
		}()
	}

	wg.Wait()

	// Cleanup
	for _, s := range servers {
		s.closer()
	}
	pool.Close()
}

// TestConnection_ConcurrentWrites tests concurrent writes to same connection
func TestConnection_ConcurrentWrites(t *testing.T) {
	server, clientWS := setupTestConnection(t)
	defer server.Close()

	conn := relay.NewConnection(clientWS, "client")
	defer conn.Close()

	const numGoroutines = 10
	const writesPerGoroutine = 50

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < writesPerGoroutine; j++ {
				data := []byte("test message")
				conn.WriteMessage(websocket.TextMessage, data)
			}
		}(i)
	}

	wg.Wait()
}

// TestHub_ConcurrentSubscribeUnsubscribe tests concurrent sub/unsub operations
func TestHub_ConcurrentSubscribeUnsubscribe(t *testing.T) {
	opts := relay.DefaultHubOptions()
	opts.BufferSize = 100
	hub := relay.NewHub[string](opts)

	const numGoroutines = 20
	const opsPerGoroutine = 50

	var wg sync.WaitGroup
	wg.Add(numGoroutines * 2)

	// Concurrent subscribes and unsubscribes
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < opsPerGoroutine; j++ {
				subID, _ := hub.Subscribe()
				time.Sleep(time.Microsecond)
				hub.Unsubscribe(subID)
			}
		}()
	}

	// Concurrent publishes
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < opsPerGoroutine; j++ {
				hub.Publish("message")
				time.Sleep(time.Microsecond)
			}
		}(i)
	}

	wg.Wait()
	hub.Stop()
}

// TestHub_ConcurrentPublishAndRead tests publishing while reading
func TestHub_ConcurrentPublishAndRead(t *testing.T) {
	opts := relay.DefaultHubOptions()
	opts.BufferSize = 1000
	hub := relay.NewHub[int](opts)

	const numPublishers = 10
	const numSubscribers = 10
	const messagesPerPublisher = 100

	var wg sync.WaitGroup
	wg.Add(numPublishers + numSubscribers)

	// Start subscribers
	receivedCounts := make([]int, numSubscribers)
	for i := 0; i < numSubscribers; i++ {
		go func(subIdx int) {
			defer wg.Done()
			subID, ch := hub.Subscribe()
			defer hub.Unsubscribe(subID)

			timeout := time.After(2 * time.Second)
			for {
				select {
				case _, ok := <-ch:
					if !ok {
						return
					}
					receivedCounts[subIdx]++
				case <-timeout:
					return
				}
			}
		}(i)
	}

	// Give subscribers time to subscribe
	time.Sleep(10 * time.Millisecond)

	// Start publishers
	for i := 0; i < numPublishers; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < messagesPerPublisher; j++ {
				hub.Publish(id*1000 + j)
			}
		}(i)
	}

	wg.Wait()
	hub.Stop()

	// Verify messages were received
	for i, count := range receivedCounts {
		assert.Greater(t, count, 0, "Subscriber %d should receive messages", i)
	}
}

// TestClientRegistry_ConcurrentOperations tests concurrent registry operations
func TestClientRegistry_ConcurrentOperations(t *testing.T) {
	registry := relay.NewClientRegistry()

	const numGoroutines = 10
	const opsPerGoroutine = 50

	var wg sync.WaitGroup
	wg.Add(numGoroutines * 3)

	// Concurrent adds
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < opsPerGoroutine; j++ {
				client := &MockClient{
					name: "client-" + string(rune(id*opsPerGoroutine+j)),
					url:  "ws://localhost:8080/ws",
				}
				registry.Add(client)
			}
		}(i)
	}

	// Concurrent gets
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < opsPerGoroutine; j++ {
				names := registry.Names()
				if len(names) > 0 {
					registry.Get(names[0])
				}
				time.Sleep(time.Microsecond)
			}
		}()
	}

	// Concurrent removes
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < opsPerGoroutine; j++ {
				names := registry.Names()
				if len(names) > 0 {
					registry.Remove(names[0])
				}
				time.Sleep(time.Microsecond)
			}
		}()
	}

	wg.Wait()
}

// TestEventHub_ConcurrentStatusUpdates tests concurrent status operations
func TestEventHub_ConcurrentStatusUpdates(t *testing.T) {
	hub := relay.NewEventHub[string](100)

	const numGoroutines = 20
	const opsPerGoroutine = 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines * 2)

	// Concurrent status updates
	for i := 0; i < numGoroutines; i++ {
		go func(clientID int) {
			defer wg.Done()
			for j := 0; j < opsPerGoroutine; j++ {
				status := relay.Status{
					Name:  "client-" + string(rune(clientID)),
					URL:   "ws://localhost:8080/ws",
					State: relay.StateConnected,
				}
				hub.UpdateStatus(status)
			}
		}(i)
	}

	// Concurrent status reads
	for i := 0; i < numGoroutines; i++ {
		go func(clientID int) {
			defer wg.Done()
			for j := 0; j < opsPerGoroutine; j++ {
				hub.GetStatus("client-" + string(rune(clientID)))
				hub.GetAllStatuses()
			}
		}(i)
	}

	wg.Wait()
}

// TestEventHub_ConcurrentMessagePublish tests concurrent message operations
func TestEventHub_ConcurrentMessagePublish(t *testing.T) {
	hub := relay.NewEventHub[int](1000)

	const numPublishers = 10
	const numSubscribers = 5
	const messagesPerPublisher = 100

	var wg sync.WaitGroup
	wg.Add(numPublishers + numSubscribers)

	// Start subscribers
	for i := 0; i < numSubscribers; i++ {
		go func(id int) {
			defer wg.Done()
			ch := hub.SubscribeMessages()
			defer hub.UnsubscribeMessages(ch)

			timeout := time.After(2 * time.Second)
			count := 0
			for {
				select {
				case _, ok := <-ch:
					if !ok {
						return
					}
					count++
				case <-timeout:
					assert.Greater(t, count, 0, "Subscriber %d should receive messages", id)
					return
				}
			}
		}(i)
	}

	// Give subscribers time to subscribe
	time.Sleep(10 * time.Millisecond)

	// Start publishers
	for i := 0; i < numPublishers; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < messagesPerPublisher; j++ {
				hub.PublishMessage(id*1000 + j)
			}
		}(i)
	}

	wg.Wait()
}

// TestRingBuffer_ConcurrentPush tests concurrent ring buffer operations
func TestRingBuffer_ConcurrentPush(t *testing.T) {
	buffer := relay.NewRingBuffer[int](100)

	const numGoroutines = 20
	const pushesPerGoroutine = 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines * 2)

	// Concurrent pushes
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < pushesPerGoroutine; j++ {
				buffer.Push(id*1000 + j)
			}
		}(i)
	}

	// Concurrent reads
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < pushesPerGoroutine; j++ {
				entries := buffer.Entries()
				_ = len(entries) // Just verify we can read size
			}
		}()
	}

	wg.Wait()

	// Verify buffer has entries
	entries := buffer.Entries()
	require.NotEmpty(t, entries, "Buffer should have entries")
	assert.LessOrEqual(t, len(entries), 100, "Buffer should not exceed capacity")
}

// TestConnectionPool_StressTest performs heavy concurrent load
func TestConnectionPool_StressTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	pool := relay.NewConnectionPool()
	defer pool.Close()

	const numGoroutines = 10  // Reduced from 50
	const duration = 500 * time.Millisecond

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	stop := make(chan struct{})
	time.AfterFunc(duration, func() {
		close(stop)
	})

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stop:
					return
				default:
					server, clientWS := setupTestConnection(t)
					conn := relay.NewConnection(clientWS, "client")
					id := pool.Add(conn)

					// Random operations
					pool.Get(id)
					pool.List()
					pool.Size()
					pool.Remove(id)

					server.Close()
					clientWS.Close()

					// Brief pause to avoid exhausting ephemeral ports
					time.Sleep(time.Millisecond)
				}
			}
		}()
	}

	wg.Wait()
}
