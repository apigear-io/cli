package actions

import "github.com/apigear-io/cli/pkg/sim"

type HandlerFunc func(sim.ActionEntry, ActionContext) error

type Engine struct {
	Handlers map[string]HandlerFunc
}

func (e *Engine) Call(action sim.ActionEntry, ctx ActionContext) error {
	log.Debugf("engine.call: %v", action)
	handler := e.Handlers[action.Name]
	if handler == nil {
		return nil
	}
	return handler(action, ctx)
}

func (e *Engine) RegisterHandler(name string, handler HandlerFunc) {
	log.Debugf("actions.registerHandler: %s", name)
	e.Handlers[name] = handler
}

func SetHandler(action sim.ActionEntry, ctx ActionContext) error {
	log.Debugf("actions.setHandler: %v", action)
	key, value := action.Params()
	ctx.Set(key.(string), value)

	return nil
}
