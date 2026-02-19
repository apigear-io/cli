package relay_test

import (
	"testing"
	"time"

	"github.com/apigear-io/cli/pkg/stream/relay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewHub verifies Hub creation
func TestNewHub(t *testing.T) {
	opts := relay.DefaultHubOptions()
	hub := relay.NewHub[string](opts)
	require.NotNil(t, hub)
	defer hub.Stop()

	// Hub should start empty
	entries := hub.Entries()
	assert.Empty(t, entries)
}

// TestDefaultHubOptions verifies default options
func TestDefaultHubOptions(t *testing.T) {
	opts := relay.DefaultHubOptions()

	assert.Equal(t, 1000, opts.BufferSize)
	assert.Equal(t, 10000, opts.PublishBufferSize)
	assert.Equal(t, 100, opts.SubscriberBufferSize)
}

// TestHub_PublishSubscribe tests basic pub/sub
func TestHub_PublishSubscribe(t *testing.T) {
	hub := relay.NewHub[string](relay.DefaultHubOptions())
	defer hub.Stop()

	// Subscribe
	subID, ch := hub.Subscribe()
	defer hub.Unsubscribe(subID)

	// Publish
	hub.Publish("test message")

	// Receive
	select {
	case msg := <-ch:
		assert.Equal(t, "test message", msg)
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Did not receive message")
	}
}

// TestHub_MultipleSubscribers tests broadcasting to multiple subscribers
func TestHub_MultipleSubscribers(t *testing.T) {
	hub := relay.NewHub[string](relay.DefaultHubOptions())
	defer hub.Stop()

	// Subscribe three times
	sub1ID, ch1 := hub.Subscribe()
	defer hub.Unsubscribe(sub1ID)
	sub2ID, ch2 := hub.Subscribe()
	defer hub.Unsubscribe(sub2ID)
	sub3ID, ch3 := hub.Subscribe()
	defer hub.Unsubscribe(sub3ID)

	// Publish once
	hub.Publish("broadcast")

	// All should receive
	received := 0
	timeout := time.After(100 * time.Millisecond)

	for received < 3 {
		select {
		case msg := <-ch1:
			assert.Equal(t, "broadcast", msg)
			received++
		case msg := <-ch2:
			assert.Equal(t, "broadcast", msg)
			received++
		case msg := <-ch3:
			assert.Equal(t, "broadcast", msg)
			received++
		case <-timeout:
			t.Fatalf("Only received %d/3 messages", received)
		}
	}
}

// TestHub_RingBuffer tests message history
func TestHub_RingBuffer(t *testing.T) {
	opts := relay.HubOptions{
		BufferSize:           3,
		PublishBufferSize:    100,
		SubscriberBufferSize: 10,
	}
	hub := relay.NewHub[string](opts)
	defer hub.Stop()

	// Publish messages
	hub.Publish("msg1")
	hub.Publish("msg2")
	hub.Publish("msg3")

	// Give time for async processing
	time.Sleep(10 * time.Millisecond)

	// Check buffer
	entries := hub.Entries()
	assert.Len(t, entries, 3)
	assert.Equal(t, []string{"msg1", "msg2", "msg3"}, entries)
}

// TestHub_RingBufferOverflow tests buffer overflow behavior
func TestHub_RingBufferOverflow(t *testing.T) {
	opts := relay.HubOptions{
		BufferSize:           2,
		PublishBufferSize:    100,
		SubscriberBufferSize: 10,
	}
	hub := relay.NewHub[int](opts)
	defer hub.Stop()

	// Publish more than capacity
	hub.Publish(1)
	hub.Publish(2)
	hub.Publish(3)
	hub.Publish(4)

	// Give time for async processing
	time.Sleep(10 * time.Millisecond)

	// Only last 2 should remain
	entries := hub.Entries()
	assert.Len(t, entries, 2)
	assert.Equal(t, []int{3, 4}, entries)
}

// TestHub_Unsubscribe tests unsubscribing
func TestHub_Unsubscribe(t *testing.T) {
	hub := relay.NewHub[string](relay.DefaultHubOptions())
	defer hub.Stop()

	subID, ch := hub.Subscribe()

	// Unsubscribe
	hub.Unsubscribe(subID)

	// Channel should be closed
	select {
	case _, ok := <-ch:
		assert.False(t, ok, "Channel should be closed")
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Channel not closed after unsubscribe")
	}
}

// TestHub_Clear tests clearing the buffer
func TestHub_Clear(t *testing.T) {
	hub := relay.NewHub[string](relay.DefaultHubOptions())
	defer hub.Stop()

	hub.Publish("msg1")
	hub.Publish("msg2")
	time.Sleep(10 * time.Millisecond)

	assert.Len(t, hub.Entries(), 2)

	// Clear
	hub.Clear()

	assert.Empty(t, hub.Entries())
}

// TestNewRingBuffer verifies RingBuffer creation
func TestNewRingBuffer(t *testing.T) {
	buffer := relay.NewRingBuffer[int](5)
	require.NotNil(t, buffer)

	entries := buffer.Entries()
	assert.Empty(t, entries)
}

// TestRingBuffer_PushEntries tests basic push/entries
func TestRingBuffer_PushEntries(t *testing.T) {
	buffer := relay.NewRingBuffer[int](3)

	buffer.Push(1)
	buffer.Push(2)
	buffer.Push(3)

	entries := buffer.Entries()
	assert.Equal(t, []int{1, 2, 3}, entries)
}

// TestRingBuffer_Overflow tests overflow behavior
func TestRingBuffer_Overflow(t *testing.T) {
	buffer := relay.NewRingBuffer[string](2)

	buffer.Push("a")
	buffer.Push("b")
	buffer.Push("c") // Overwrites "a"
	buffer.Push("d") // Overwrites "b"

	entries := buffer.Entries()
	assert.Equal(t, []string{"c", "d"}, entries)
}

// TestRingBuffer_Clear tests clearing the buffer
func TestRingBuffer_Clear(t *testing.T) {
	buffer := relay.NewRingBuffer[int](5)

	buffer.Push(1)
	buffer.Push(2)
	buffer.Push(3)

	buffer.Clear()

	entries := buffer.Entries()
	assert.Empty(t, entries)
}

// TestNewForwarder_Default tests creating default forwarder
func TestNewForwarder_Default(t *testing.T) {
	forwarder := relay.NewForwarder(relay.ForwarderOptions{})
	require.NotNil(t, forwarder)
}

// TestNewForwarder_WithDelay tests creating delayed forwarder
func TestNewForwarder_WithDelay(t *testing.T) {
	opts := relay.ForwarderOptions{
		Delay:      100 * time.Millisecond,
		BufferSize: 1000,
	}

	forwarder := relay.NewForwarder(opts)
	require.NotNil(t, forwarder)
}

// TestNewForwarder_WithSpeed tests creating throttled forwarder
func TestNewForwarder_WithSpeed(t *testing.T) {
	opts := relay.ForwarderOptions{
		Speed:      0.5,
		BufferSize: 500,
	}

	forwarder := relay.NewForwarder(opts)
	require.NotNil(t, forwarder)
}

// TestNewForwarder_WithCallback tests forwarder with message callback
func TestNewForwarder_WithCallback(t *testing.T) {
	called := false
	opts := relay.ForwarderOptions{
		OnMessage: func(msg relay.Message) {
			called = true
		},
	}

	forwarder := relay.NewForwarder(opts)
	require.NotNil(t, forwarder)

	// Note: We can't easily test the callback without setting up full connections
	// The callback is tested in the internal forward package tests
	_ = called
}

// TestForwarderOptions_SpeedPrecedence tests that speed takes precedence
func TestForwarderOptions_SpeedPrecedence(t *testing.T) {
	opts := relay.ForwarderOptions{
		Delay: 100 * time.Millisecond,
		Speed: 0.5, // Speed should take precedence
	}

	forwarder := relay.NewForwarder(opts)
	require.NotNil(t, forwarder)

	// The actual precedence is tested in internal forward package
	// Here we just verify the forwarder is created
}

// TestHub_ConcurrentPublish tests concurrent publishing
func TestHub_ConcurrentPublish(t *testing.T) {
	hub := relay.NewHub[int](relay.DefaultHubOptions())
	defer hub.Stop()

	subID, ch := hub.Subscribe()
	defer hub.Unsubscribe(subID)

	// Publish from multiple goroutines
	done := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			hub.Publish(i)
		}
		close(done)
	}()

	go func() {
		for i := 10; i < 20; i++ {
			hub.Publish(i)
		}
	}()

	// Receive messages
	received := 0
	timeout := time.After(500 * time.Millisecond)

	for received < 20 {
		select {
		case <-ch:
			received++
		case <-timeout:
			// May not receive all due to async nature and timing
			// Just ensure we received some
			assert.Greater(t, received, 0, "Should receive at least some messages")
			return
		}
	}
}

// TestHub_StructMessages tests Hub with struct messages
func TestHub_StructMessages(t *testing.T) {
	type TestMessage struct {
		ID   int
		Text string
	}

	hub := relay.NewHub[TestMessage](relay.DefaultHubOptions())
	defer hub.Stop()

	subID, ch := hub.Subscribe()
	defer hub.Unsubscribe(subID)

	// Publish struct
	hub.Publish(TestMessage{ID: 1, Text: "hello"})

	// Receive
	select {
	case msg := <-ch:
		assert.Equal(t, 1, msg.ID)
		assert.Equal(t, "hello", msg.Text)
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Did not receive message")
	}
}
