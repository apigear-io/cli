package actions

import (
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
	iface   *spec.InterfaceEntry
	seq     *spec.SequenceEntry
	DoneC   chan bool
	StepC   chan *spec.ActionListEntry
	FramesC chan PlayFrame
}

func NewPlayer(iface *spec.InterfaceEntry, seq *spec.SequenceEntry) *Player {
	p := &Player{
		iface:   iface,
		seq:     seq,
		DoneC:   make(chan bool),
		StepC:   make(chan *spec.ActionListEntry),
		FramesC: make(chan PlayFrame),
	}
	return p
}

func (p *Player) SequenceName() string {
	return p.seq.Name
}

func (p *Player) Play() error {
	log.Info().Msgf("play sequence %s", p.seq.Name)
	go p.loopPump(p.seq)
	go func() {
		err := p.framePump(p.seq.Interval)
		if err != nil {
			log.Error().Msgf("frame pump error: %v", err)
		}
	}()
	return nil
}

func (p *Player) loopPump(seq *spec.SequenceEntry) {
	loops := seq.Loops
	if loops == 0 {
		loops = 1
	}
	for i := 0; i < loops; i++ {
		for _, step := range seq.Steps {
			p.StepC <- step
		}
	}
	p.DoneC <- true
	close(p.StepC)
}

func (p *Player) framePump(interval int) error {
	for {
		select {
		case s := <-p.StepC:
			if s == nil {
				return nil
			}
			for _, action := range s.Actions {
				p.FramesC <- PlayFrame{
					Action:    action,
					Interface: p.iface,
				}
			}
			time.Sleep(time.Duration(interval) * time.Millisecond)
		case <-p.DoneC:
			log.Info().Msgf("frame pump %s is done", p.seq.Name)
			close(p.FramesC)
			close(p.DoneC)
			return nil
		}
	}
}

func (p *Player) Stop() error {
	close(p.DoneC)
	return nil
}
