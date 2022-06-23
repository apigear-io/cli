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

func (s *EventScript) jsCall(symbol string, args ...interface{}) {
	log.Debugf("call: %s %v", symbol, args)
	evt := &Event{
		Id:        uuid.New().String(),
		DeviceId:  "1234",
		Kind:      KindCall,
		Timestamp: time.Now(),
		Source:    "345",
		Symbol:    symbol,
		Params:    args,
	}
	s.emitter <- evt
}

func (s *EventScript) jsSignal(symbol string, args ...interface{}) {
	log.Debugf("signal: %s %v", symbol, args)
	evt := &Event{
		Id:        uuid.New().String(),
		DeviceId:  "1234",
		Kind:      KindSignal,
		Timestamp: time.Now(),
		Source:    "345",
		Symbol:    symbol,
		Params:    args,
	}
	s.emitter <- evt
}

func (s *EventScript) jsSet(symbol string, props map[string]any) {
	log.Debugf("get: %s", symbol)
	evt := &Event{
		Id:        uuid.New().String(),
		DeviceId:  "1234",
		Kind:      KindState,
		Timestamp: time.Now(),
		Source:    "345",
		Symbol:    symbol,
		Props:     props,
	}
	s.emitter <- evt

}

func (s *EventScript) jsSleep(duration int) {
	log.Debugf("sleep: %d", duration)
}
