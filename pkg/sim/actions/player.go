package actions

import (
	"context"
	"sync"
	"time"

	"github.com/apigear-io/cli/pkg/spec"
)

type PlayFrame struct {
	Action    spec.ActionEntry
	Interface *spec.InterfaceEntry
}

// Player is a player for one sequence
// It plays the sequence and sends the actions to the stream
// The stream is closed when the sequence is finished
// Actions are evaluated in the context of an interface
type Player struct {
	sync.RWMutex
	iface   *spec.InterfaceEntry
	seq     *spec.SequenceEntry
	ctx     context.Context
	cancel  context.CancelFunc
	StepC   chan *spec.ActionListEntry
	FramesC chan PlayFrame
}

func NewPlayer(iface *spec.InterfaceEntry, seq *spec.SequenceEntry) *Player {
	p := &Player{
		iface:   iface,
		seq:     seq,
		StepC:   make(chan *spec.ActionListEntry),
		FramesC: make(chan PlayFrame),
	}
	return p
}

func (p *Player) SequenceName() string {
	return p.seq.Name
}

func (p *Player) Play(ctx context.Context) error {
	log.Debug().Msgf("play sequence %s", p.seq.Name)
	ctx, cancel := context.WithCancel(ctx)
	if p.cancel != nil {
		p.cancel()
	}
	p.cancel = cancel
	go p.loopPump(ctx, p.seq)
	go p.framePump(ctx, p.seq.Interval)
	return nil
}

func (p *Player) loopPump(ctx context.Context, seq *spec.SequenceEntry) {
	loops := seq.Loops
	if loops == 0 {
		loops = 1
	}
	for i := 0; i < loops; i++ {
		for _, step := range seq.Steps {
			select {
			case <-ctx.Done():
				return
			default:
				p.StepC <- step
			}
		}
	}
	close(p.StepC)
}

func (p *Player) framePump(ctx context.Context, interval int) {
	for {
		select {
		case s := <-p.StepC:
			if s == nil {
				return
			}
			for _, action := range s.Actions {
				// every frame the player sends the action to the stream
				// the stream is closed when the sequence is finished
				// or the player is stopped
				select {
				case <-ctx.Done():
					return
				default:
					p.FramesC <- PlayFrame{
						Action:    action,
						Interface: p.iface,
					}
				}
			}
			time.Sleep(time.Duration(interval) * time.Millisecond)
		case <-ctx.Done():
			log.Debug().Msgf("frame pump %s is done", p.seq.Name)
			return
		}
	}
}

func (p *Player) Stop() error {
	p.cancel()
	return nil
}
