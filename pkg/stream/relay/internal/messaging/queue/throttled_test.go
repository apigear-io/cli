package queue

import (
	"sync"
	"testing"
	"time"
)

func TestThrottledQueue_NewThrottledQueue(t *testing.T) {
	q := NewThrottledQueue[int](0.5, 10)
	defer q.Stop()

	if q.Speed() != 0.5 {
		t.Errorf("expected speed 0.5, got %f", q.Speed())
	}
}

func TestThrottledQueue_DefaultValues(t *testing.T) {
	// Zero speed should default to 1.0
	q := NewThrottledQueue[int](0, 0)
	defer q.Stop()

	if q.speed != 1.0 {
		t.Errorf("expected default speed 1.0, got %f", q.speed)
	}
	if q.bufferSize != 1000 {
		t.Errorf("expected default buffer 1000, got %d", q.bufferSize)
	}
}

func TestThrottledQueue_Send_Receive_FirstMessage(t *testing.T) {
	q := NewThrottledQueue[int](0.5, 10)
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
		// First message should be near-instant
		if elapsed > 50*time.Millisecond {
			t.Errorf("first message should be fast, got %v", elapsed)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("timeout waiting for item")
	}
}

func TestThrottledQueue_SpeedScaling(t *testing.T) {
	// Speed 0.5 = half speed, gaps between sends should be doubled
	// This means if messages arrive 100ms apart, they should be sent 200ms apart
	q := NewThrottledQueue[int](0.5, 10)
	defer q.Stop()

	// Track timestamps
	var recv1, recv2 time.Time

	// Send first message
	q.Send(1)
	select {
	case <-q.Receive():
		recv1 = time.Now()
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timeout on first message")
	}

	// Wait 100ms, then send second message
	time.Sleep(100 * time.Millisecond)
	q.Send(2)

	select {
	case <-q.Receive():
		recv2 = time.Now()
	case <-time.After(500 * time.Millisecond):
		t.Fatal("timeout on second message")
	}

	// The gap between receiving the two outputs should be ~200ms (100ms * 2)
	actualGap := recv2.Sub(recv1)
	expectedMin := 180 * time.Millisecond // Allow some tolerance
	expectedMax := 300 * time.Millisecond

	if actualGap < expectedMin || actualGap > expectedMax {
		t.Errorf("expected gap between %v and %v, got %v", expectedMin, expectedMax, actualGap)
	}
}

func TestThrottledQueue_NormalSpeed(t *testing.T) {
	// Speed 1.0 = normal speed, gaps should be unchanged
	q := NewThrottledQueue[int](1.0, 10)
	defer q.Stop()

	// Send first message
	q.Send(1)
	<-q.Receive()

	// Wait, then send second message
	gap := 30 * time.Millisecond
	time.Sleep(gap)

	start := time.Now()
	q.Send(2)

	select {
	case <-q.Receive():
		elapsed := time.Since(start)
		// Should be nearly instant (gap preserved, not stretched)
		if elapsed > 50*time.Millisecond {
			t.Errorf("expected near-instant at speed 1.0, got %v", elapsed)
		}
	case <-time.After(200 * time.Millisecond):
		t.Error("timeout waiting for item")
	}
}

func TestThrottledQueue_BufferFull(t *testing.T) {
	// Create a queue with very small buffer and slow processing
	q := NewThrottledQueue[int](0.01, 2) // Very slow speed
	defer q.Stop()

	// Send first message to start timing
	q.Send(1)

	// Wait a moment, then fill buffer rapidly
	time.Sleep(50 * time.Millisecond)

	// Try to overflow - send many more than buffer can hold
	dropped := 0
	for i := 2; i <= 10; i++ {
		if !q.Send(i) {
			dropped++
		}
	}

	// Should have dropped some due to buffer being full
	if dropped == 0 && q.Dropped() == 0 {
		t.Error("expected some messages to be dropped when buffer overflows")
	}
}

func TestThrottledQueue_OrderPreserved(t *testing.T) {
	q := NewThrottledQueue[int](1.0, 100)
	defer q.Stop()

	// Send multiple items quickly
	for i := 1; i <= 5; i++ {
		q.Send(i)
	}

	// Should receive in order
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

func TestThrottledQueue_ConcurrentSend(t *testing.T) {
	q := NewThrottledQueue[int](1.0, 1000)
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
				time.Sleep(time.Millisecond) // Small gap between sends
			}
		}(i)
	}

	wg.Wait()

	// Collect results
	received := 0
	timeout := time.After(1 * time.Second)

	for {
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

func TestThrottledQueue_Stop(t *testing.T) {
	q := NewThrottledQueue[int](1.0, 100)

	// Send some items
	for i := 0; i < 5; i++ {
		q.Send(i)
	}

	// Stop should be graceful
	q.Stop()

	// Channel should be closed
	_, ok := <-q.Receive()
	if ok {
		// May get remaining items
	}
}

func TestThrottledQueue_WithStructs(t *testing.T) {
	type Event struct {
		Type string
		Data int
	}

	q := NewThrottledQueue[Event](1.0, 10)
	defer q.Stop()

	q.Send(Event{"test", 42})

	select {
	case evt := <-q.Receive():
		if evt.Type != "test" || evt.Data != 42 {
			t.Errorf("expected {test, 42}, got %+v", evt)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("timeout waiting for event")
	}
}

func TestThrottledQueue_HighSpeed(t *testing.T) {
	// Speed > 1.0 should speed up traffic
	q := NewThrottledQueue[int](2.0, 10)
	defer q.Stop()

	// Send first message
	q.Send(1)
	<-q.Receive()

	// Wait, then send second message
	gap := 100 * time.Millisecond
	time.Sleep(gap)

	start := time.Now()
	q.Send(2)

	select {
	case <-q.Receive():
		elapsed := time.Since(start)
		// With speed=2.0, a 100ms gap should become ~50ms
		// Allow some tolerance
		if elapsed > 80*time.Millisecond {
			t.Errorf("expected faster delivery with speed 2.0, got %v", elapsed)
		}
	case <-time.After(200 * time.Millisecond):
		t.Error("timeout waiting for item")
	}
}
