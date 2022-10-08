package mon

import (
	"fmt"
	"os"
	"time"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
	"github.com/google/uuid"
)

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

type EventScript struct {
	vm      *goja.Runtime
	emitter chan *Event
}

func NewEventScript(emitter chan *Event) *EventScript {
	vm := goja.New()
	new(require.Registry).Enable(vm)
	console.Enable(vm)
	s := &EventScript{vm: vm, emitter: emitter}
	s.init()
	return s
}

func (s *EventScript) RunScriptFromFile(file string) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("failed to read script file: %v", err)
	}
	return s.RunScript(string(content))
}

func (s *EventScript) RunScript(script string) error {
	prog, err := goja.Compile("", script, true)
	defer func() {
		close(s.emitter)
	}()
	if err != nil {
		return fmt.Errorf("compile error: %v", err)
	}
	_, err = s.vm.RunProgram(prog)
	if err != nil {
		return fmt.Errorf("failed to run script: %v", err)
	}
	return nil
}

func (s *EventScript) init() {
	Must(s.vm.Set("call", s.jsCall))
	Must(s.vm.Set("signal", s.jsSignal))
	Must(s.vm.Set("set", s.jsSet))
	Must(s.vm.Set("sleep", s.jsSleep))
}

func (s *EventScript) jsCall(symbol string, data Payload) {
	log.Debug().Msgf("call: %s %v", symbol, data)
	evt := &Event{
		Id:        uuid.New().String(),
		Type:      TypeCall,
		Timestamp: time.Now(),
		Source:    "345",
		Symbol:    symbol,
		Data:      data,
	}
	s.emitter <- evt
}

func (s *EventScript) jsSignal(symbol string, data Payload) {
	log.Debug().Msgf("signal: %s %v", symbol, data)
	evt := &Event{
		Id:        uuid.New().String(),
		Type:      TypeSignal,
		Timestamp: time.Now(),
		Source:    "345",
		Symbol:    symbol,
		Data:      data,
	}
	s.emitter <- evt
}

func (s *EventScript) jsSet(symbol string, data Payload) {
	log.Debug().Msgf("set: %s", symbol)
	evt := &Event{
		Id:        uuid.New().String(),
		Type:      TypeState,
		Timestamp: time.Now(),
		Source:    "345",
		Symbol:    symbol,
		Data:      data,
	}
	s.emitter <- evt
}

func (s *EventScript) jsSleep(duration int) {
	log.Debug().Msgf("sleep: %d", duration)
}
