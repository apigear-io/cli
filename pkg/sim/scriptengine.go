package sim

import (
	"apigear/pkg/log"

	"github.com/dop251/goja"
)

// ScriptEngine is a runs the different scripts in the context of a simulation.
// The engine produces events which are sent via a channel.
// Typically the simulation listens to all script engines passes them on to the network communication.
type ScriptEngine struct {
	vm *goja.Runtime
}

func NewScriptEngine() *ScriptEngine {
	log.Debugf("new script engine")
	return &ScriptEngine{
		vm: goja.New(),
	}
}

func (s *ScriptEngine) Run(script ScriptEntry) error {
	log.Debugf("run script %s", script.Name)
	value, err := s.vm.RunScript(script.Name, script.Source)
	if err != nil {
		log.Errorf("ScriptEngine.run: %s", err)
		return err
	}
	log.Debugf("script %s returned %v", script.Name, value)
	return nil
}
