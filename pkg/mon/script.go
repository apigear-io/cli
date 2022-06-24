package mon

import (
	"apigear/pkg/log"
	"time"

	"github.com/dop251/goja"
	"github.com/google/uuid"
)

type EventScript struct {
	vm      *goja.Runtime
	emitter chan *Event
}

func NewEventScript(emitter chan *Event) *EventScript {
	vm := goja.New()
	s := &EventScript{vm: vm}
	s.init()
	return s
}

func (s *EventScript) RunScript(script string) {
	_, err := s.vm.RunString(script)
	if err != nil {
		panic(err)
	}
	fn, ok := goja.AssertFunction(s.vm.Get("main"))
	if !ok {
		panic("not a function")
	}
	_, err = fn(goja.Undefined())
	if err != nil {
		panic(err)
	}
}

func (s *EventScript) init() {
	s.vm.Set("call", s.jsCall)
	s.vm.Set("signal", s.jsSignal)
	s.vm.Set("set", s.jsSet)
	s.vm.Set("sleep", s.jsSleep)
}

func (s *EventScript) jsCall(symbol string, data Payload) {
	log.Debugf("call: %s %v", symbol, data)
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
	log.Debugf("signal: %s %v", symbol, data)
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
	log.Debugf("get: %s", symbol)
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
	log.Debugf("sleep: %d", duration)
}
