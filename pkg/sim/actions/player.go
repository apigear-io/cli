package actions

import (
	"context"
	"sync"
	"time"

	"github.com/apigear-io/cli/pkg/helper"
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
	id string
	sync.RWMutex
	e      *Engine
	iface  *spec.InterfaceEntry
	seq    *spec.SequenceEntry
	cancel context.CancelFunc
}

// NewPlayer creates a new sequence player
func NewPlayer(e *Engine, iface *spec.InterfaceEntry, seq *spec.SequenceEntry) *Player {
	p := &Player{
		id:    helper.NewUUID(),
		e:     e,
		iface: iface,
		seq:   seq,
	}
	log.Debug().Msgf("%s: NewPlayer", p.id)
	return p
}

// SequenceName returns the name of the sequence
func (p *Player) SequenceName() string {
	return p.seq.Name
}

// Play starts the sequence
// It runs the sequence in a goroutine
func (p *Player) Play(ctx context.Context) error {
	log.Debug().Msgf("%s: Player.play", p.id)
	if p.cancel != nil {
		p.cancel()
	}
	ctx, cancel := context.WithCancel(ctx)
	p.cancel = cancel
	go p.runSequence(ctx, p.seq)
	return nil
}

// runSequence runs the sequence
func (p *Player) runSequence(ctx context.Context, seq *spec.SequenceEntry) {
	log.Debug().Msgf("%s: Player.run", p.id)
	defer func() {
		log.Info().Msgf("%s Player.run finished", p.id)
	}()
	loops := seq.Loops
	if loops == 0 {
		loops = 1
	}
	for i := 0; i < loops; i++ {
		helper.EachTicked(ctx, seq.Steps, func(step *spec.ActionListEntry) {
			for _, action := range step.Actions {
				if p.iface == nil {
					log.Error().Msgf("interface %s not found", seq.Interface)
					return
				}
				_, err := p.e.eval.EvalAction(p.iface.Name, action)
				if err != nil {
					log.Error().Msgf("error evaluating action %#v: %s", action, err)
					return
				}
			}
		}, time.Duration(seq.Interval)*time.Millisecond)
	}
}

// Stop stops the play loop
func (p *Player) Stop() error {
	log.Debug().Msgf("%s: Player.stop", p.id)
	if p.cancel != nil {
		p.cancel()
		p.cancel = nil
	}
	return nil
}
