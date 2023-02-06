package actions

import (
	"fmt"
	"sync"

	"github.com/apigear-io/cli/pkg/helper"
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

func (e *eval) EvalActions(ifaceId string, actions []spec.ActionEntry) (any, error) {
	var result any
	for _, action := range actions {
		r, err := e.EvalAction(ifaceId, action)
		if err != nil {
			return nil, err
		}
		if r != nil {
			result = r
		}
	}
	return result, nil
}

func (e *eval) EvalAction(ifaceId string, action spec.ActionEntry) (any, error) {
	var result any
	for k := range action {
		if h, ok := e.actions[k]; ok {
			v, err := h(ifaceId, action[k])
			if err != nil {
				return nil, fmt.Errorf("error in action %s: %v", k, err)
			}
			if v != nil {
				result = v
			}
		} else {
			return nil, fmt.Errorf("unknown action %s", k)
		}
	}
	return result, nil
}

func (e *eval) EvalActionString(ifaceId string, data []byte) (any, error) {
	var action spec.ActionEntry
	if err := yaml.Unmarshal(data, &action); err != nil {
		return nil, err
	}
	return e.EvalAction(ifaceId, action)
}

func (e *eval) register(name string, handler ActionHandler) {
	e.actions[name] = handler
}

// actionSet sets properties of the interface and notifies the change.
func (e *eval) actionSet(ifaceId string, kwargs map[string]any) (any, error) {
	log.Debug().Fields(kwargs).Msg("$set")
	e.mu.Lock()
	defer e.mu.Unlock()
	e.store.Set(ifaceId, kwargs)
	return nil, nil
}

// actionSet sets properties of the interface and notifies the change.
func (e *eval) actionUpdate(ifaceId string, kwargs map[string]any) (any, error) {
	log.Debug().Msgf("$update: %v", kwargs)
	e.mu.Lock()
	defer e.mu.Unlock()
	e.store.Update(ifaceId, kwargs)
	return nil, nil
}

// actionReturn returns the result of the action.
func (e *eval) actionReturn(ifaceId string, kwargs map[string]any) (any, error) {
	log.Debug().Msgf("$return: %v", kwargs)
	e.mu.Lock()
	defer e.mu.Unlock()
	return kwargs, nil
}

// actionSignal sends a signal to the interface.
func (e *eval) actionSignal(ifaceId string, kwargs map[string]any) (any, error) {
	log.Debug().Msgf("$signal: %s", kwargs)
	for name := range kwargs {
		sigArgs := helper.ToSlice(kwargs[name])
		e.EmitSignal(ifaceId, name, sigArgs)
	}
	return nil, nil
}

// actionChange sends a change to the interface.
func (e *eval) actionChange(ifaceId string, kwargs map[string]any) (any, error) {
	log.Debug().Fields(kwargs).Msg("$change")
	for k := range kwargs {
		e.EmitPropertyChanged(ifaceId, k, kwargs[k])
	}
	return nil, nil
}
