package actions

import (
	"testing"

	"github.com/apigear-io/cli/pkg/spec"
	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	e := NewEval()
	assert.NotNil(t, e)
	var table = []struct {
		symbol string
		action string
		args   map[string]any
		props  map[string]any
		result map[string]any
	}{
		{"demo.Counter", "$set", map[string]any{"count": 1}, map[string]any{"count": 1}, nil},
		{"demo.Counter", "$return", map[string]any{"value": 2}, map[string]any{"count": 0}, map[string]any{"value": 2}},
	}
	for _, row := range table {
		t.Run(row.symbol+"."+row.action, func(t *testing.T) {
			ctx := map[string]any{"count": 0}
			result, err := e.EvalAction(row.symbol, spec.ActionEntry{row.action: row.args}, ctx)
			assert.NoError(t, err)
			assert.Equal(t, row.props, ctx)
			assert.Equal(t, row.result, result)
		})
	}
}

func TestActionSet(t *testing.T) {
	e := NewEval()
	assert.NotNil(t, e)
	ctx := map[string]any{"count": 0}
	args := map[string]any{"count": 1}
	symbol := "demo.Counter"
	result, err := e.EvalAction(symbol, spec.ActionEntry{"$set": args}, ctx)
	assert.NoError(t, err)
	assert.Equal(t, 1, ctx["count"])
	assert.Nil(t, result)
	args = map[string]any{"count": 2}
	result, err = e.EvalAction(symbol, spec.ActionEntry{"$set": args}, ctx)
	assert.NoError(t, err)
	assert.Equal(t, 2, ctx["count"])
	assert.Nil(t, result)
}

func TestActionReturn(t *testing.T) {
	e := NewEval()
	assert.NotNil(t, e)
	ctx := map[string]any{"count": 0}
	args := map[string]any{"value": 2}
	symbol := "demo.Counter"
	result, err := e.EvalAction(symbol, spec.ActionEntry{"$return": args}, ctx)
	assert.NoError(t, err)
	assert.Equal(t, 0, ctx["count"])
	assert.Equal(t, map[string]any{"value": 2}, result)
}

func TestActionSignal(t *testing.T) {
	e := NewEval()
	var sigName string
	var sigArgs map[string]any
	e.OnSignal(func(symbol string, name string, args map[string]any) {
		sigName = name
		sigArgs = args
	})

	assert.NotNil(t, e)
	ctx := map[string]any{"count": 0}
	args := map[string]any{"shutdown": map[string]any{"timeout": 1}}
	symbol := "demo.Counter"
	// $signal: { shutdown: { timeout: 1 } }
	result, err := e.EvalAction(symbol, spec.ActionEntry{"$signal": args}, ctx)
	assert.NoError(t, err)
	assert.Equal(t, 0, ctx["count"])
	assert.Nil(t, result)
	assert.Equal(t, "shutdown", sigName)
	assert.Equal(t, map[string]any{"timeout": 1}, sigArgs)

	// $signal: { shutdown: { timeout: 2 } }
	action := []byte("$signal: { shutdown2: { timeout: 2 } }")
	result, err = e.EvalActionString(symbol, action, ctx)
	assert.NoError(t, err)
	assert.Equal(t, 0, ctx["count"])
	assert.Nil(t, result)
	assert.Equal(t, "shutdown2", sigName)
	assert.Equal(t, map[string]any{"timeout": 2}, sigArgs)
}
