package hub

import (
	"sync"
	"testing"
	"time"
)

func TestHub_NewHub_Defaults(t *testing.T) {
	hub := NewHub[int](HubOptions{})
	defer hub.Stop()

	if hub.buffer.Cap() != 1000 {
		t.Errorf("expected buffer capacity 1000, got %d", hub.buffer.Cap())
	}
}

func TestHub_NewHub_CustomOptions(t *testing.T) {
	hub := NewHub[int](HubOptions{
		BufferSize:           50,
		PublishBufferSize:    100,
		SubscriberBufferSize: 10,
	})
	defer hub.Stop()

	if hub.buffer.Cap() != 50 {
		t.Errorf("expected buffer capacity 50, got %d", hub.buffer.Cap())
	}
}

func TestHub_Subscribe_Unsubscribe(t *testing.T) {
	hub := NewHub[int](DefaultHubOptions())
	defer hub.Stop()

	id, ch := hub.Subscribe()

	if hub.SubscriberCount() != 1 {
		t.Errorf("expected 1 subscriber, got %d", hub.SubscriberCount())
	}

	hub.Unsubscribe(id)

	if hub.SubscriberCount() != 0 {
		t.Errorf("expected 0 subscribers, got %d", hub.SubscriberCount())
	}

	// Channel should be closed
	_, ok := <-ch
	if ok {
		t.Error("expected channel to be closed")
	}
}

func TestHub_Publish_Receive(t *testing.T) {
	hub := NewHub[int](DefaultHubOptions())
	defer hub.Stop()

	_, ch := hub.Subscribe()

	hub.Publish(42)

	select {
	case val := <-ch:
		if val != 42 {
			t.Errorf("expected 42, got %d", val)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("timeout waiting for message")
	}
}

func TestHub_Publish_MultipleSubscribers(t *testing.T) {
	hub := NewHub[int](DefaultHubOptions())
	defer hub.Stop()

	_, ch1 := hub.Subscribe()
	_, ch2 := hub.Subscribe()
	_, ch3 := hub.Subscribe()

	hub.Publish(42)

	// All subscribers should receive the message
	for i, ch := range []<-chan int{ch1, ch2, ch3} {
		select {
		case val := <-ch:
			if val != 42 {
				t.Errorf("subscriber %d: expected 42, got %d", i, val)
			}
		case <-time.After(100 * time.Millisecond):
			t.Errorf("subscriber %d: timeout waiting for message", i)
		}
	}
}

func TestHub_Entries(t *testing.T) {
	hub := NewHub[int](HubOptions{BufferSize: 5})
	defer hub.Stop()

	for i := 1; i <= 3; i++ {
		hub.Publish(i)
	}

	// Wait for async processing
	time.Sleep(50 * time.Millisecond)

	entries := hub.Entries()
	if len(entries) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(entries))
	}

	for i, v := range entries {
		if v != i+1 {
			t.Errorf("entry[%d]: expected %d, got %d", i, i+1, v)
		}
	}
}

func TestHub_Entries_Overflow(t *testing.T) {
	hub := NewHub[int](HubOptions{BufferSize: 3})
	defer hub.Stop()

	for i := 1; i <= 5; i++ {
		hub.Publish(i)
	}

	// Wait for async processing
	time.Sleep(50 * time.Millisecond)

	entries := hub.Entries()
	if len(entries) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(entries))
	}

	// Should have 3, 4, 5 (oldest entries overwritten)
	expected := []int{3, 4, 5}
	for i, v := range expected {
		if entries[i] != v {
			t.Errorf("entry[%d]: expected %d, got %d", i, v, entries[i])
		}
	}
}

func TestHub_Clear(t *testing.T) {
	hub := NewHub[int](DefaultHubOptions())
	defer hub.Stop()

	hub.Publish(1)
	hub.Publish(2)
	time.Sleep(50 * time.Millisecond)

	hub.Clear()

	if hub.Len() != 0 {
		t.Errorf("expected 0 items after clear, got %d", hub.Len())
	}
}

func TestHub_Stop_Drains(t *testing.T) {
	hub := NewHub[int](HubOptions{BufferSize: 100})

	// Publish many items
	for i := 0; i < 50; i++ {
		hub.Publish(i)
	}

	// Stop should drain all pending items
	hub.Stop()

	// All items should be in the buffer
	if hub.Len() != 50 {
		t.Errorf("expected 50 items after stop, got %d", hub.Len())
	}
}

func TestHub_SlowSubscriber_DropsMessages(t *testing.T) {
	hub := NewHub[int](HubOptions{
		BufferSize:           100,
		SubscriberBufferSize: 2, // Small buffer
	})
	defer hub.Stop()

	_, ch := hub.Subscribe()

	// Publish more than the subscriber buffer can hold
	for i := 0; i < 10; i++ {
		hub.Publish(i)
	}

	// Wait for processing
	time.Sleep(50 * time.Millisecond)

	// Subscriber should only have received up to buffer size
	received := 0
	for {
		select {
		case <-ch:
			received++
		default:
			goto done
		}
	}
done:

	// Should have received at most the buffer size
	if received > 2 {
		t.Errorf("expected at most 2 messages, got %d", received)
	}
}

func TestHub_ConcurrentPublish(t *testing.T) {
	hub := NewHub[int](HubOptions{BufferSize: 1000})
	defer hub.Stop()

	_, ch := hub.Subscribe()

	var wg sync.WaitGroup
	numGoroutines := 10
	numMessages := 100

	// Concurrent publishers
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(base int) {
			defer wg.Done()
			for j := 0; j < numMessages; j++ {
				hub.Publish(base*numMessages + j)
			}
		}(i)
	}

	wg.Wait()

	// Give time for messages to be processed
	time.Sleep(100 * time.Millisecond)

	// Buffer should have messages
	if hub.Len() != numGoroutines*numMessages {
		t.Errorf("expected %d items, got %d", numGoroutines*numMessages, hub.Len())
	}

	// Drain the subscriber channel
	received := 0
	for {
		select {
		case <-ch:
			received++
		default:
			goto done2
		}
	}
done2:

	// Should have received messages (may not be all due to buffer limits)
	if received == 0 {
		t.Error("expected to receive at least some messages")
	}
}

func TestHub_UnsubscribeTwice(t *testing.T) {
	hub := NewHub[int](DefaultHubOptions())
	defer hub.Stop()

	id, _ := hub.Subscribe()
	hub.Unsubscribe(id)
	hub.Unsubscribe(id) // Should not panic
}

func TestHub_UnsubscribeInvalidID(t *testing.T) {
	hub := NewHub[int](DefaultHubOptions())
	defer hub.Stop()

	hub.Unsubscribe("nonexistent") // Should not panic
}

func TestHub_WithStructs(t *testing.T) {
	type Event struct {
		Type string
		Data int
	}

	hub := NewHub[Event](DefaultHubOptions())
	defer hub.Stop()

	_, ch := hub.Subscribe()

	hub.Publish(Event{"test", 42})

	select {
	case evt := <-ch:
		if evt.Type != "test" || evt.Data != 42 {
			t.Errorf("expected {test, 42}, got %+v", evt)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("timeout waiting for event")
	}
}
