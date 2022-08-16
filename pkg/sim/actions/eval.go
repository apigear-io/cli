package actions

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/sim/core"
)

type ActionHandler func(symbol string, args core.KWArgs, ctx core.KWArgs) (any, error)

type eval struct {
	actions map[string]ActionHandler
	core.Notifier
}

func NewEval() *eval {
	e := &eval{
		actions: map[string]ActionHandler{},
	}
	e.register("$set", e.actionSet)
	e.register("$return", e.actionReturn)
	e.register("$signal", e.actionSignal)
	e.register("$change", e.actionChange)
	return e
}

func (e *eval) EvalActions(symbol string, actions []ActionEntry, ctx core.KWArgs) (any, error) {
	var result any
	for _, action := range actions {
		v, err := e.EvalAction(symbol, action, ctx)
		if err != nil {
			return nil, err
		}
		if v != nil {
			result = v
		}
	}
	return result, nil
}

func (e *eval) EvalAction(symbol string, action ActionEntry, ctx core.KWArgs) (any, error) {
	var result any
	for k := range action {
		if h, ok := e.actions[k]; ok {
			v, err := h(symbol, action[k], ctx)
			if err != nil {
				log.Printf("error: %v", err)
			}
			if v != nil {
				result = v
			}
		} else {
			log.Printf("action %s not found", k)
		}
	}
	return result, nil
}

func (e *eval) register(name string, handler ActionHandler) {
	e.actions[name] = handler
}

func (e *eval) actionSet(symbol string, args core.KWArgs, ctx core.KWArgs) (any, error) {
	log.Debugf("actionSet: %v\n", args)
	for k := range args {
		ctx[k] = args[k]
	}
	return nil, nil
}

// actionReturn sets a _return value for the action.
func (e *eval) actionReturn(symbol string, args core.KWArgs, ctx core.KWArgs) (any, error) {
	log.Debugf("actionReturn: %v\n", args)
	return args, nil
}

// actionSignal sends a signal to the interface.
func (e *eval) actionSignal(symbol string, args core.KWArgs, ctx core.KWArgs) (any, error) {
	log.Debugf("actionSignal: %s\n", args)
	for k := range args {
		sigArgs, ok := args[k].(core.KWArgs)
		if !ok {
			return nil, fmt.Errorf("signal %s has no args", k)
		}
		e.EmitOnSignal(symbol, k, sigArgs)
	}
	return nil, nil
}

// actionChange sends a change to the interface.
func (e *eval) actionChange(symbol string, args core.KWArgs, ctx core.KWArgs) (any, error) {
	log.Debugf("actionChange: %v\n", args)
	for k := range args {
		e.EmitOnChange(symbol, k, args[k])
	}
	return nil, nil
}
