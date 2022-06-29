package actions

import "apigear/pkg/log"

type Context map[string]any

type HandlerFunc func(Action, Context) error

// { "cmd": [ "arg0", "arg1" ] }
// Action is a command to be executed by the simulator.
type Action map[string][]any

type Engine struct {
	Handlers map[string]HandlerFunc
}

func (e *Engine) Call(action Action, ctx Context) error {
	log.Debugf("engine.call: %v", action)
	for cmd, args := range action {
		log.Debugf("engine.call: %s %v", cmd, args)
		handler := e.Handlers[cmd]
		if handler != nil {
			return handler(action, ctx)
		}
	}
	return nil
}

func (e *Engine) RegisterHandler(name string, handler HandlerFunc) {
	log.Debugf("actions.registerHandler: %s", name)
	e.Handlers[name] = handler
}

func (e *Engine) SetActionHandler(a Action, ctx Context) {
	log.Debugf("actions.setActionHandler: %v", a)
}
