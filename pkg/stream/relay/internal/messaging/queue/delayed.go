package queue

import (
	"sync"
	"time"
)

// delayedItem wraps an item with its scheduled send time.
type delayedItem[T any] struct {
	item   T
	sendAt time.Time
}

// DelayedQueue is a channel-based queue that delays items by a fixed duration.
// Items are delivered in order, each delayed by the configured amount.
type DelayedQueue[T any] struct {
	delay      time.Duration
	bufferSize int
	inputCh    chan delayedItem[T]
	outputCh   chan T
	done       chan struct{}
	wg         sync.WaitGroup
}

// NewDelayedQueue creates a new delayed queue.
// delay is the duration to delay each item.
// bufferSize is the capacity of the internal buffer.
func NewDelayedQueue[T any](delay time.Duration, bufferSize int) *DelayedQueue[T] {
	if bufferSize <= 0 {
		bufferSize = 1000
	}

	q := &DelayedQueue[T]{
		delay:      delay,
		bufferSize: bufferSize,
		inputCh:    make(chan delayedItem[T], bufferSize),
		outputCh:   make(chan T, bufferSize),
		done:       make(chan struct{}),
	}
	q.start()
	return q
}

// Send queues an item for delayed delivery.
// Returns false if the buffer is full (item is dropped).
func (q *DelayedQueue[T]) Send(item T) bool {
	di := delayedItem[T]{
		item:   item,
		sendAt: time.Now().Add(q.delay),
	}

	select {
	case q.inputCh <- di:
		return true
	default:
		return false // Buffer full
	}
}

// Receive returns the output channel for receiving delayed items.
func (q *DelayedQueue[T]) Receive() <-chan T {
	return q.outputCh
}

// Stop gracefully stops the queue.
// Remaining items in the buffer are delivered before stopping.
func (q *DelayedQueue[T]) Stop() {
	close(q.done)
	q.wg.Wait()
	close(q.outputCh)
}

// Delay returns the configured delay duration.
func (q *DelayedQueue[T]) Delay() time.Duration {
	return q.delay
}

// start begins the background goroutine that processes delayed items.
func (q *DelayedQueue[T]) start() {
	q.wg.Add(1)
	go func() {
		defer q.wg.Done()
		for {
			select {
			case di, ok := <-q.inputCh:
				if !ok {
					return
				}
				q.processItem(di)
			case <-q.done:
				// Drain remaining items
				for {
					select {
					case di := <-q.inputCh:
						q.processItem(di)
					default:
						return
					}
				}
			}
		}
	}()
}

// processItem waits until the scheduled time and delivers the item.
func (q *DelayedQueue[T]) processItem(di delayedItem[T]) {
	waitTime := time.Until(di.sendAt)
	if waitTime > 0 {
		select {
		case <-time.After(waitTime):
		case <-q.done:
			// Still deliver the item even if stopping
		}
	}

	select {
	case q.outputCh <- di.item:
	default:
		// Output buffer full, drop item
	}
}
