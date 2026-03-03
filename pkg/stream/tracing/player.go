package tracing

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// MessageSender is an interface for sending messages to a proxy or stream.
type MessageSender interface {
	SendMessage(messageType int, data []byte) error
}

// Player plays back trace files by sending messages at the original timing.
type Player struct {
	filename string
	entries  []TraceEntry
	speed    float64 // Playback speed multiplier (1.0 = real-time, 2.0 = 2x speed)
	loop     bool    // Whether to loop playback
	filter   FilterOptions

	ctx    context.Context
	cancel context.CancelFunc

	// State
	mu           sync.RWMutex
	position     int
	state        PlayerState
	entriesPlayed int64
	startTime    time.Time

	// Callbacks
	onMessage  func(entry TraceEntry, index int)
	onComplete func()
	onError    func(err error)
}

// PlayerState represents the current state of the player.
type PlayerState string

const (
	PlayerStateStopped PlayerState = "stopped"
	PlayerStatePlaying PlayerState = "playing"
	PlayerStatePaused  PlayerState = "paused"
)

// PlayerOptions configures trace playback.
type PlayerOptions struct {
	Speed      float64       // Playback speed (default: 1.0)
	Loop       bool          // Loop playback
	Filter     FilterOptions // Filter which entries to play
	OnMessage  func(entry TraceEntry, index int)
	OnComplete func()
	OnError    func(err error)
}

// NewPlayer creates a new trace player.
func NewPlayer(filename string, options PlayerOptions) (*Player, error) {
	// Read trace file
	entries, err := ReadTraceFileFiltered(filename, options.Filter)
	if err != nil {
		return nil, fmt.Errorf("failed to read trace file: %w", err)
	}

	if len(entries) == 0 {
		return nil, fmt.Errorf("no entries found in trace file")
	}

	speed := options.Speed
	if speed <= 0 {
		speed = 1.0
	}

	ctx, cancel := context.WithCancel(context.Background())

	p := &Player{
		filename:   filename,
		entries:    entries,
		speed:      speed,
		loop:       options.Loop,
		filter:     options.Filter,
		ctx:        ctx,
		cancel:     cancel,
		state:      PlayerStateStopped,
		onMessage:  options.OnMessage,
		onComplete: options.OnComplete,
		onError:    options.OnError,
	}

	return p, nil
}

// Play starts playback.
func (p *Player) Play() error {
	p.mu.Lock()
	if p.state == PlayerStatePlaying {
		p.mu.Unlock()
		return fmt.Errorf("already playing")
	}
	p.state = PlayerStatePlaying
	p.startTime = time.Now()
	p.mu.Unlock()

	go p.playbackLoop()
	return nil
}

// Pause pauses playback.
func (p *Player) Pause() {
	p.mu.Lock()
	if p.state == PlayerStatePlaying {
		p.state = PlayerStatePaused
	}
	p.mu.Unlock()
}

// Resume resumes playback after pause.
func (p *Player) Resume() {
	p.mu.Lock()
	if p.state == PlayerStatePaused {
		p.state = PlayerStatePlaying
	}
	p.mu.Unlock()
}

// Stop stops playback and resets position.
func (p *Player) Stop() {
	p.cancel()
	p.mu.Lock()
	p.state = PlayerStateStopped
	p.position = 0
	p.mu.Unlock()
}

// Seek moves to a specific position.
func (p *Player) Seek(position int) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if position < 0 || position >= len(p.entries) {
		return fmt.Errorf("position out of range: %d", position)
	}

	p.position = position
	return nil
}

// GetState returns the current player state.
func (p *Player) GetState() PlayerState {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.state
}

// GetPosition returns the current playback position.
func (p *Player) GetPosition() int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.position
}

// GetProgress returns playback progress (0.0 to 1.0).
func (p *Player) GetProgress() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if len(p.entries) == 0 {
		return 0
	}
	return float64(p.position) / float64(len(p.entries))
}

// GetEntriesPlayed returns the total number of entries played.
func (p *Player) GetEntriesPlayed() int64 {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.entriesPlayed
}

// playbackLoop is the main playback loop.
func (p *Player) playbackLoop() {
	defer func() {
		p.mu.Lock()
		p.state = PlayerStateStopped
		p.mu.Unlock()

		if p.onComplete != nil {
			p.onComplete()
		}
	}()

	for {
		// Check if stopped
		select {
		case <-p.ctx.Done():
			return
		default:
		}

		// Check if paused
		p.mu.RLock()
		paused := p.state == PlayerStatePaused
		p.mu.RUnlock()

		if paused {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		// Get current position
		p.mu.RLock()
		position := p.position
		p.mu.RUnlock()

		// Check if reached end
		if position >= len(p.entries) {
			if p.loop {
				p.mu.Lock()
				p.position = 0
				p.mu.Unlock()
				continue
			}
			return
		}

		// Get current entry
		entry := p.entries[position]

		// Calculate delay based on timestamp difference
		if position > 0 {
			prevEntry := p.entries[position-1]
			delayMs := float64(entry.Timestamp-prevEntry.Timestamp) / p.speed

			if delayMs > 0 {
				delay := time.Duration(delayMs) * time.Millisecond
				select {
				case <-p.ctx.Done():
					return
				case <-time.After(delay):
				}
			}
		}

		// Call onMessage callback
		if p.onMessage != nil {
			p.onMessage(entry, position)
		}

		// Update position and stats
		p.mu.Lock()
		p.position++
		p.entriesPlayed++
		p.mu.Unlock()
	}
}

// PlayToSender plays the trace file by sending messages to a MessageSender.
// This is useful for replaying traces to a proxy or WebSocket connection.
func PlayToSender(filename string, sender MessageSender, options PlayerOptions) error {
	// Override onMessage to send to the sender
	options.OnMessage = func(entry TraceEntry, index int) {
		// Parse message as array to determine type
		var msgArray []interface{}
		if err := json.Unmarshal(entry.Message, &msgArray); err != nil {
			if options.OnError != nil {
				options.OnError(fmt.Errorf("failed to parse message: %w", err))
			}
			return
		}

		// Send message (type 1 = text message in WebSocket)
		if err := sender.SendMessage(1, entry.Message); err != nil {
			if options.OnError != nil {
				options.OnError(fmt.Errorf("failed to send message: %w", err))
			}
		}
	}

	player, err := NewPlayer(filename, options)
	if err != nil {
		return err
	}

	return player.Play()
}
