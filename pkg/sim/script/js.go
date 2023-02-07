package script

import (
	"fmt"

	js "github.com/dop251/goja"
)

func (s *Engine) jsRegisterInterface(obj *js.Object) error {
	jsId := obj.Get("_id")
	if jsId == nil {
		return fmt.Errorf("_id not found")
	}
	id := jsId.String()
	if id == "" {
		return fmt.Errorf("interface id is empty")
	}
	if _, ok := s.interfaces[id]; ok {
		return fmt.Errorf("interface %s already registered", id)
	}
	log.Info().Msgf("registering interface %s", id)
	s.interfaces[id] = obj
	return nil
}

func (s *Engine) jsSignal(interfaceId string, name string, args []any) {
	s.EmitSignal(interfaceId, name, args)
}

func (s *Engine) jsChange(interfaceId string, name string, value any) {
	s.EmitPropertyChanged(interfaceId, name, value)
}

func (s *Engine) jsRegisterSequence(obj *js.Object) error {
	jsId := obj.Get("_id")
	if jsId == nil {
		return fmt.Errorf("_id not found")
	}
	id := jsId.String()
	if id == "" {
		return fmt.Errorf("sequencer id is empty")
	}
	if _, ok := s.sequencers[id]; ok {
		return fmt.Errorf("sequencer %s already registered", id)
	}
	s.sequencers[id] = obj
	return nil
}
