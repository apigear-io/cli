package actions

import (
	"testing"

	"github.com/apigear-io/cli/pkg/sim/core"

	"github.com/stretchr/testify/assert"
)

func LoadTest(t *testing.T) *Engine {
	e := NewEngine()
	assert.NotNil(t, e)
	doc, err := ReadScenario("testdata/test1.scenario.yaml")
	assert.NoError(t, err)
	e.LoadScenario(doc)
	return e
}

func TestInvokeOperation(t *testing.T) {
	e := LoadTest(t)
	var table = []struct {
		symbol    string
		operation string
		args      core.KWArgs
		result    core.KWArgs
	}{
		{"demo.Counter", "increment", nil, core.KWArgs{"count": 1}},
		{"demo.Counter", "decrement", nil, core.KWArgs{"count": 0}},
	}
	for _, row := range table {
		t.Run(row.symbol+"."+row.operation, func(t *testing.T) {
			assert.True(t, e.HasInterface(row.symbol))
			result, err := e.InvokeOperation(row.symbol, row.operation, row.args)
			assert.NoError(t, err)
			assert.Nil(t, result)
			props, err := e.GetProperties(row.symbol)
			assert.NoError(t, err)
			assert.Equal(t, row.result, props)
		})
	}
}

func TestEngine_HasInterface(t *testing.T) {
	e := NewEngine()
	assert.False(t, e.HasInterface("Interface1"))
}
