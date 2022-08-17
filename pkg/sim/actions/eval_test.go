package actions

import (
	"testing"

	"github.com/apigear-io/cli/pkg/spec"
	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	e := NewEval()
	assert.NotNil(t, e)
	e.EvalAction("demo.Counter", spec.ActionEntry{"$return": map[string]any{}}, map[string]any{})
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
