package actions

import (
	"testing"

	"github.com/apigear-io/cli/pkg/sim/core"
	"github.com/apigear-io/cli/pkg/sim/ostore"
	"github.com/apigear-io/cli/pkg/spec"
	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	store := ostore.NewMemoryStore()
	e := NewEval(store)
	assert.NotNil(t, e)
	var table = []struct {
		symbol string
		action string
		args   map[string]any
		props  map[string]any
		result any
	}{
		{"demo.Counter", "$set", map[string]any{"count": 1}, map[string]any{"count": 1}, nil},
		{"demo.Counter", "$return", map[string]any{"value": 2}, map[string]any{"count": 0}, map[string]any{"value": 2}},
	}
	for _, row := range table {
		t.Run(row.symbol+"."+row.action, func(t *testing.T) {
			store.Set(row.symbol, map[string]any{"count": 0})
			result, err := e.EvalAction(row.symbol, spec.ActionEntry{row.action: row.args})
			assert.NoError(t, err)
			assert.Equal(t, row.props, store.Get(row.symbol))
			assert.Equal(t, row.result, result)
		})
	}
}

func TestActionSet(t *testing.T) {
	store := ostore.NewMemoryStore()
	e := NewEval(store)
	assert.NotNil(t, e)
	store.Set("demo.Counter", map[string]any{"count": 0})
	args := map[string]any{"count": 1}
	symbol := "demo.Counter"
	result, err := e.EvalAction(symbol, spec.ActionEntry{"$set": args})
	assert.NoError(t, err)
	assert.Equal(t, 1, store.Get("demo.Counter")["count"])
	assert.Nil(t, result)
	args = map[string]any{"count": 2}
	result, err = e.EvalAction(symbol, spec.ActionEntry{"$set": args})
	assert.NoError(t, err)
	assert.Equal(t, 2, store.Get("demo.Counter")["count"])
	assert.Nil(t, result)
}

func TestActionReturn(t *testing.T) {
	store := ostore.NewMemoryStore()
	e := NewEval(store)
	assert.NotNil(t, e)
	store.Set("demo.Counter", map[string]any{"count": 0})
	args := map[string]any{"value": 2}
	symbol := "demo.Counter"
	result, err := e.EvalAction(symbol, spec.ActionEntry{"$return": args})
	assert.NoError(t, err)
	assert.Equal(t, 0, store.Get("demo.Counter")["count"])
	assert.Equal(t, map[string]any{"value": 2}, result)
}

func TestActionSignal(t *testing.T) {
	store := ostore.NewMemoryStore()
	e := NewEval(store)
	var sigName string
	var sigArgs []any
	e.OnEvent(func(e *core.SimuEvent) {
		sigName = e.Name
		sigArgs = e.Args
	})

	assert.NotNil(t, e)
	ctx := map[string]any{"count": 0}
	args := []any{1}
	symbol := "demo.Counter"
	// $signal: { shutdown: [ 1 ] }
	result, err := e.EvalAction(symbol, spec.ActionEntry{"$signal": {"shutdown": args}})
	assert.NoError(t, err)
	assert.Equal(t, 0, ctx["count"])
	assert.Nil(t, result)
	assert.Equal(t, "shutdown", sigName)
	assert.Equal(t, []any{1}, sigArgs)

	// $signal: { shutdown: { timeout: 2 } }
	action := []byte("$signal: { shutdown2: [ 2 ] }")
	result, err = e.EvalActionString(symbol, action)
	assert.NoError(t, err)
	assert.Equal(t, 0, ctx["count"])
	assert.Nil(t, result)
	assert.Equal(t, "shutdown2", sigName)
	assert.Equal(t, []any{2}, sigArgs)
}
