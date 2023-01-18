package actions

import (
	"fmt"
	"sync"

	"github.com/apigear-io/cli/pkg/sim/core"
	"github.com/apigear-io/cli/pkg/sim/ostore"
	"github.com/apigear-io/cli/pkg/spec"
	"gopkg.in/yaml.v3"
)

type ActionHandler func(symbol string, args map[string]any) (any, error)

type eval struct {
	store   ostore.IObjectStore
	actions map[string]ActionHandler
	core.EventNotifier
	mu sync.Mutex
}

func NewEval(store ostore.IObjectStore) *eval {
	e := &eval{
		store:   store,
		actions: map[string]ActionHandler{},
	}
	e.register("$set", e.actionSet)
	e.register("$update", e.actionUpdate)
	e.register("$return", e.actionReturn)
	e.register("$signal", e.actionSignal)
	e.register("$change", e.actionChange)
	return e
}

func (e *eval) EvalActions(symbol string, actions []spec.ActionEntry) (any, error) {
	var result any
	for _, action := range actions {
		r, err := e.EvalAction(symbol, action)
		if err != nil {
			return nil, err
		}
		if r != nil {
			result = r
		}
	}
	return result, nil
}

func (e *eval) EvalAction(symbol string, action spec.ActionEntry) (any, error) {
	var result any
	for k := range action {
		if h, ok := e.actions[k]; ok {
			v, err := h(symbol, action[k])
			if err != nil {
				log.Info().Msgf("error: %v", err)
			}
			if v != nil {
				result = v
			}
		} else {
			log.Info().Msgf("action %s not found", k)
		}
	}
	return result, nil
}

func (e *eval) EvalActionString(symbol string, data []byte) (any, error) {
	var action spec.ActionEntry
	if err := yaml.Unmarshal(data, &action); err != nil {
		log.Info().Msgf("error: %v", err)
	}
	return e.EvalAction(symbol, action)
}

func (e *eval) register(name string, handler ActionHandler) {
	e.actions[name] = handler
}

// actionSet sets properties of the interface and notifies the change.
func (e *eval) actionSet(symbol string, args map[string]any) (any, error) {
	log.Debug().Msgf("actionSet: %v", args)
	e.mu.Lock()
	defer e.mu.Unlock()
	e.store.Set(symbol, args)
	return nil, nil
}

// actionSet sets properties of the interface and notifies the change.
func (e *eval) actionUpdate(symbol string, args map[string]any) (any, error) {
	log.Debug().Msgf("actionUpdate: %v", args)
	e.mu.Lock()
	defer e.mu.Unlock()
	e.store.Update(symbol, args)
	return nil, nil
}

// actionReturn returns the result of the action.
func (e *eval) actionReturn(symbol string, args map[string]any) (any, error) {
	log.Debug().Msgf("actionReturn: %v", args)
	e.mu.Lock()
	defer e.mu.Unlock()
	return args, nil
}

// actionSignal sends a signal to the interface.
func (e *eval) actionSignal(symbol string, args map[string]any) (any, error) {
	log.Debug().Msgf("actionSignal: %s", args)
	for k := range args {
		sigArgs, ok := args[k].(map[string]any)
		if !ok {
			return nil, fmt.Errorf("signal %s has no args", k)
		}
		e.EmitSignal(symbol, k, sigArgs)
	}
	return nil, nil
}

// actionChange sends a change to the interface.
func (e *eval) actionChange(symbol string, args map[string]any) (any, error) {
	log.Debug().Msgf("actionChange: %v", args)
	for k := range args {
		e.EmitPropertyChanged(symbol, k, args[k])
	}
	return nil, nil
}
