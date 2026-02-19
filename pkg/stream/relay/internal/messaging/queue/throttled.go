package queue

import (
	"sync"
	"time"
)

// throttledItem wraps an item with its scheduled send time.
type throttledItem[T any] struct {
	item   T
	sendAt time.Time
}

// ThrottledQueue is a channel-based queue that scales timing gaps between items.
// Speed < 1.0 slows down traffic by stretching gaps between items.
// For example, speed=0.5 means gaps between items are doubled.
type ThrottledQueue[T any] struct {
	speed      float64
	bufferSize int
	inputCh    chan throttledItem[T]
	outputCh   chan T
	done       chan struct{}
	wg         sync.WaitGroup

	// Timing state
	lastRecvTime time.Time
	lastSendTime time.Time
	firstMessage bool
	mu           sync.Mutex

	// Stats
	dropped int64
}

// NewThrottledQueue creates a new throttled queue.
// speed is the throttling factor (0.5 = half speed, 1.0 = normal).
// bufferSize is the capacity of the internal buffer.
func NewThrottledQueue[T any](speed float64, bufferSize int) *ThrottledQueue[T] {
	if speed <= 0 {
		speed = 1.0
	}
	if bufferSize <= 0 {
		bufferSize = 1000
	}

	q := &ThrottledQueue[T]{
		speed:        speed,
		bufferSize:   bufferSize,
		inputCh:      make(chan throttledItem[T], bufferSize),
		outputCh:     make(chan T, bufferSize),
		done:         make(chan struct{}),
		firstMessage: true,
	}
	q.start()
	return q
}

// Send queues an item with scaled timing.
// Returns false if the buffer is full (item is dropped).
func (q *ThrottledQueue[T]) Send(item T) bool {
	now := time.Now()

	q.mu.Lock()
	var sendAt time.Time
	if q.firstMessage {
		// First message: send immediately
		sendAt = now
		q.firstMessage = false
		q.lastSendTime = now
	} else {
		// Calculate gap since last received message
		gap := now.Sub(q.lastRecvTime)
		// Scale the gap by 1/speed (speed=0.5 means gaps are doubled)
		scaledGap := time.Duration(float64(gap) / q.speed)
		sendAt = q.lastSendTime.Add(scaledGap)
	}
	q.lastRecvTime = now
	q.mu.Unlock()

	ti := throttledItem[T]{
		item:   item,
		sendAt: sendAt,
	}

	// Check buffer capacity
	if len(q.inputCh) >= q.bufferSize {
		q.mu.Lock()
		q.dropped++
		q.mu.Unlock()
		return false
	}

	select {
	case q.inputCh <- ti:
		return true
	default:
		q.mu.Lock()
		q.dropped++
		q.mu.Unlock()
		return false
	}
}

// Receive returns the output channel for receiving throttled items.
func (q *ThrottledQueue[T]) Receive() <-chan T {
	return q.outputCh
}

// Stop gracefully stops the queue.
// Remaining items in the buffer are delivered before stopping.
func (q *ThrottledQueue[T]) Stop() {
	close(q.done)
	q.wg.Wait()
	close(q.outputCh)
}

// Speed returns the configured speed factor.
func (q *ThrottledQueue[T]) Speed() float64 {
	return q.speed
}

// Dropped returns the number of items dropped due to buffer overflow.
func (q *ThrottledQueue[T]) Dropped() int64 {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.dropped
}

// start begins the background goroutine that processes throttled items.
func (q *ThrottledQueue[T]) start() {
	q.wg.Add(1)
	go func() {
		defer q.wg.Done()
		for {
			select {
			case ti, ok := <-q.inputCh:
				if !ok {
					return
				}
				q.processItem(ti)
			case <-q.done:
				// Drain remaining items
				for {
					select {
					case ti := <-q.inputCh:
						q.processItem(ti)
					default:
						return
					}
				}
			}
		}
	}()
}

// processItem waits until the scheduled time and delivers the item.
func (q *ThrottledQueue[T]) processItem(ti throttledItem[T]) {
	waitTime := time.Until(ti.sendAt)
	if waitTime > 0 {
		select {
		case <-time.After(waitTime):
		case <-q.done:
			// Still deliver the item even if stopping
		}
	}

	// Update actual send time
	q.mu.Lock()
	q.lastSendTime = time.Now()
	q.mu.Unlock()

	select {
	case q.outputCh <- ti.item:
	default:
		// Output buffer full, drop item
		q.mu.Lock()
		q.dropped++
		q.mu.Unlock()
	}
}
