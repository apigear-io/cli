// Package forward provides WebSocket message forwarding strategies.
package forward

import (
	"sync"
	"time"

	"github.com/apigear-io/cli/pkg/stream/relay/internal/core"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

// Message represents a WebSocket message to be forwarded.
type Message struct {
	Type int
	Data []byte
}

// MessageHandler is called for each message before forwarding.
// It allows stats collection, logging, etc.
type MessageHandler func(msg Message)

// Options configures a forwarder.
type Options struct {
	OnMessage  MessageHandler // Called for each message received
	BufferSize int            // Max queued messages (default: 1000)
	Delay      time.Duration  // Fixed delay (for DelayedForwarder)
	Speed      float64        // Speed factor 0-1 (for ThrottledForwarder)
}

// Forwarder defines the interface for message forwarding strategies.
type Forwarder interface {
	// Forward reads from src and writes to dst until an error or closure.
	Forward(src, dst core.Connection) error
}

// NewForwarder creates the appropriate forwarder based on options.
func NewForwarder(opts Options) Forwarder {
	if opts.BufferSize == 0 {
		opts.BufferSize = 1000
	}

	// Speed throttling takes precedence
	if opts.Speed > 0 && opts.Speed < 1.0 {
		return &ThrottledForwarder{opts: opts}
	}

	// Fixed delay
	if opts.Delay > 0 {
		return &DelayedForwarder{opts: opts}
	}

	// Direct forwarding (no delay)
	return &DirectForwarder{opts: opts}
}

// DirectForwarder forwards messages without any delay.
type DirectForwarder struct {
	opts Options
}

// Forward reads messages from src and writes to dst immediately.
func (f *DirectForwarder) Forward(src, dst core.Connection) error {
	for {
		messageType, message, err := src.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseAbnormalClosure) {
				log.Error().Err(err).Msg("Read error")
			}
			return err
		}

		if f.opts.OnMessage != nil {
			f.opts.OnMessage(Message{Type: messageType, Data: message})
		}

		if err := dst.WriteMessage(messageType, message); err != nil {
			log.Error().Err(err).Msg("Write error")
			return err
		}
	}
}

// delayedMessage holds a message to be sent after a delay.
type delayedMessage struct {
	messageType int
	data        []byte
	sendAt      time.Time
}

// DelayedForwarder forwards messages with a fixed delay.
type DelayedForwarder struct {
	opts Options
}

// Forward reads messages and queues them to be sent after a delay.
func (f *DelayedForwarder) Forward(src, dst core.Connection) error {
	queue := make(chan delayedMessage, f.opts.BufferSize)
	senderDone := make(chan struct{})
	var senderErr error

	// Sender goroutine: sends messages at their scheduled time
	go func() {
		defer close(senderDone)
		for {
			select {
			case msg, ok := <-queue:
				if !ok {
					return // Queue closed, reader finished
				}
				// Wait until it's time to send (interruptible)
				waitTime := time.Until(msg.sendAt)
				if waitTime > 0 {
					select {
					case <-time.After(waitTime):
					case <-dst.Done():
						return // core.Connection closed
					}
				}

				if err := dst.WriteMessage(msg.messageType, msg.data); err != nil {
					log.Error().Err(err).Msg("Write error")
					senderErr = err
					return
				}
			case <-dst.Done():
				return // core.Connection closed
			}
		}
	}()

	// Reader: reads messages and queues them with delay
	for {
		select {
		case <-senderDone:
			return senderErr
		default:
		}

		messageType, message, err := src.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseAbnormalClosure) {
				log.Error().Err(err).Msg("Read error")
			}
			close(queue)
			<-senderDone
			return err
		}

		if f.opts.OnMessage != nil {
			f.opts.OnMessage(Message{Type: messageType, Data: message})
		}

		select {
		case queue <- delayedMessage{
			messageType: messageType,
			data:        message,
			sendAt:      time.Now().Add(f.opts.Delay),
		}:
		case <-senderDone:
			return senderErr
		}
	}
}

// throttledMessage holds a message with its scheduled send time.
type throttledMessage struct {
	messageType int
	data        []byte
	sendAt      time.Time
}

// ThrottledForwarder forwards messages with speed throttling.
// Speed < 1.0 slows down traffic by stretching the gaps between messages.
type ThrottledForwarder struct {
	opts Options
}

// Forward reads messages and queues them with scaled timing.
func (f *ThrottledForwarder) Forward(src, dst core.Connection) error {
	queue := make(chan throttledMessage, f.opts.BufferSize)
	senderDone := make(chan struct{})
	var senderErr error

	// Track timing for scaling gaps
	var lastSendTime time.Time
	sendTimeMu := sync.Mutex{}

	// Sender goroutine: sends messages at their scheduled time
	go func() {
		defer close(senderDone)
		for {
			select {
			case msg, ok := <-queue:
				if !ok {
					return
				}
				// Wait until it's time to send
				waitTime := time.Until(msg.sendAt)
				if waitTime > 0 {
					select {
					case <-time.After(waitTime):
					case <-dst.Done():
						return
					}
				}

				if err := dst.WriteMessage(msg.messageType, msg.data); err != nil {
					log.Error().Err(err).Msg("Write error")
					senderErr = err
					return
				}

				sendTimeMu.Lock()
				lastSendTime = time.Now()
				sendTimeMu.Unlock()
			case <-dst.Done():
				return
			}
		}
	}()

	var lastRecvTime time.Time
	firstMessage := true

	for {
		select {
		case <-senderDone:
			return senderErr
		default:
		}

		messageType, message, err := src.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseAbnormalClosure) {
				log.Error().Err(err).Msg("Read error")
			}
			close(queue)
			<-senderDone
			return err
		}

		now := time.Now()

		if f.opts.OnMessage != nil {
			f.opts.OnMessage(Message{Type: messageType, Data: message})
		}

		// Calculate when to send this message
		var sendAt time.Time
		if firstMessage {
			sendAt = now
			firstMessage = false
			sendTimeMu.Lock()
			lastSendTime = now
			sendTimeMu.Unlock()
		} else {
			gap := now.Sub(lastRecvTime)
			scaledGap := time.Duration(float64(gap) / f.opts.Speed)

			sendTimeMu.Lock()
			sendAt = lastSendTime.Add(scaledGap)
			sendTimeMu.Unlock()
		}
		lastRecvTime = now

		// Check buffer capacity
		if len(queue) >= f.opts.BufferSize {
			log.Warn().
				Int("bufferSize", f.opts.BufferSize).
				Msg("Buffer full, dropping message")
			continue
		}

		select {
		case queue <- throttledMessage{
			messageType: messageType,
			data:        message,
			sendAt:      sendAt,
		}:
		case <-senderDone:
			return senderErr
		}
	}
}
