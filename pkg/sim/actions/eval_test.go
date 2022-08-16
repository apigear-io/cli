package actions

import (
	"testing"

	"github.com/apigear-io/cli/pkg/sim/core"
	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	e := NewEval()
	assert.NotNil(t, e)
	e.EvalAction("demo.Counter", ActionEntry{"$return": core.KWArgs{}}, core.KWArgs{})
	var table = []struct {
		symbol string
		action string
		args   core.KWArgs
		props  core.KWArgs
		result core.KWArgs
	}{
		{"demo.Counter", "$set", core.KWArgs{"count": 1}, core.KWArgs{"count": 1}, nil},
		{"demo.Counter", "$return", core.KWArgs{"value": 2}, core.KWArgs{"count": 0}, core.KWArgs{"value": 2}},
	}
	for _, row := range table {
		t.Run(row.symbol+"."+row.action, func(t *testing.T) {
			ctx := core.KWArgs{"count": 0}
			result, err := e.EvalAction(row.symbol, ActionEntry{row.action: row.args}, ctx)
			assert.NoError(t, err)
			assert.Equal(t, row.props, ctx)
			assert.Equal(t, row.result, result)
		})
	}
}
