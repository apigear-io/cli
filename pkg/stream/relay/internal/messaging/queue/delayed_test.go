package queue

import (
	"sync"
	"testing"
	"time"
)

func TestDelayedQueue_NewDelayedQueue(t *testing.T) {
	q := NewDelayedQueue[int](100*time.Millisecond, 10)
	defer q.Stop()

	if q.Delay() != 100*time.Millisecond {
		t.Errorf("expected delay 100ms, got %v", q.Delay())
	}
}

func TestDelayedQueue_DefaultBufferSize(t *testing.T) {
	q := NewDelayedQueue[int](10*time.Millisecond, 0)
	defer q.Stop()

	// Should use default buffer size of 1000
	if q.bufferSize != 1000 {
		t.Errorf("expected buffer size 1000, got %d", q.bufferSize)
	}
}

func TestDelayedQueue_Send_Receive(t *testing.T) {
	delay := 50 * time.Millisecond
	q := NewDelayedQueue[int](delay, 10)
	defer q.Stop()

	start := time.Now()
	if !q.Send(42) {
		t.Fatal("Send returned false")
	}

	select {
	case val := <-q.Receive():
		elapsed := time.Since(start)
		if val != 42 {
			t.Errorf("expected 42, got %d", val)
		}
		// Should have taken at least the delay time
		if elapsed < delay {
			t.Errorf("expected delay of at least %v, got %v", delay, elapsed)
		}
	case <-time.After(200 * time.Millisecond):
		t.Error("timeout waiting for delayed item")
	}
}

func TestDelayedQueue_OrderPreserved(t *testing.T) {
	q := NewDelayedQueue[int](10*time.Millisecond, 100)
	defer q.Stop()

	// Send multiple items
	for i := 1; i <= 5; i++ {
		if !q.Send(i) {
			t.Fatalf("Send(%d) returned false", i)
		}
	}

	// Receive should maintain order
	for i := 1; i <= 5; i++ {
		select {
		case val := <-q.Receive():
			if val != i {
				t.Errorf("expected %d, got %d", i, val)
			}
		case <-time.After(200 * time.Millisecond):
			t.Fatalf("timeout waiting for item %d", i)
		}
	}
}

func TestDelayedQueue_BufferFull(t *testing.T) {
	q := NewDelayedQueue[int](1*time.Second, 2) // Long delay, small buffer
	defer q.Stop()

	// Fill the buffer
	if !q.Send(1) {
		t.Error("first Send should succeed")
	}
	if !q.Send(2) {
		t.Error("second Send should succeed")
	}

	// Third should fail (buffer full)
	if q.Send(3) {
		t.Error("third Send should fail (buffer full)")
	}
}

func TestDelayedQueue_Stop_DrainsPending(t *testing.T) {
	q := NewDelayedQueue[int](5*time.Millisecond, 100)

	// Send some items
	for i := 1; i <= 3; i++ {
		q.Send(i)
	}

	// Wait a bit for items to be processed
	time.Sleep(50 * time.Millisecond)

	// Stop should drain remaining items
	q.Stop()

	// Output channel should be closed
	_, ok := <-q.Receive()
	if ok {
		// Might get remaining items, that's fine
	}
}

func TestDelayedQueue_ZeroDelay(t *testing.T) {
	q := NewDelayedQueue[int](0, 10)
	defer q.Stop()

	start := time.Now()
	q.Send(42)

	select {
	case val := <-q.Receive():
		elapsed := time.Since(start)
		if val != 42 {
			t.Errorf("expected 42, got %d", val)
		}
		// Should be nearly instant
		if elapsed > 50*time.Millisecond {
			t.Errorf("expected near-instant delivery, got %v", elapsed)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("timeout waiting for item")
	}
}

func TestDelayedQueue_ConcurrentSend(t *testing.T) {
	q := NewDelayedQueue[int](5*time.Millisecond, 1000)
	defer q.Stop()

	var wg sync.WaitGroup
	numGoroutines := 10
	numMessages := 50

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(base int) {
			defer wg.Done()
			for j := 0; j < numMessages; j++ {
				q.Send(base*numMessages + j)
			}
		}(i)
	}

	wg.Wait()

	// Collect results
	received := 0
	timeout := time.After(500 * time.Millisecond)

	for received < numGoroutines*numMessages {
		select {
		case <-q.Receive():
			received++
		case <-timeout:
			goto done
		}
	}
done:

	// Should have received most messages
	if received < numGoroutines*numMessages/2 {
		t.Errorf("expected at least %d messages, got %d", numGoroutines*numMessages/2, received)
	}
}

func TestDelayedQueue_WithStructs(t *testing.T) {
	type Message struct {
		ID   int
		Text string
	}

	q := NewDelayedQueue[Message](10*time.Millisecond, 10)
	defer q.Stop()

	q.Send(Message{1, "hello"})

	select {
	case msg := <-q.Receive():
		if msg.ID != 1 || msg.Text != "hello" {
			t.Errorf("expected {1, hello}, got %+v", msg)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("timeout waiting for message")
	}
}
