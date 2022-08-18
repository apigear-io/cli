package actions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func LoadTest(t *testing.T) *Engine {
	e := NewEngine()
	assert.NotNil(t, e)
	doc, err := ReadScenario("testdata/test1.scenario.yaml")
	assert.NoError(t, err)
	err = e.LoadScenario("123", doc)
	assert.NoError(t, err)
	return e
}

func TestInvokeOperation(t *testing.T) {
	e := LoadTest(t)
	var table = []struct {
		symbol    string
		operation string

		args   map[string]any
		props  map[string]any
		result any
	}{
		{"demo.Counter", "increment", nil, map[string]any{"count": 1}, nil},
		{"demo.Counter", "decrement", nil, map[string]any{"count": 0}, nil},
	}
	for _, row := range table {
		t.Run(row.symbol+"."+row.operation, func(t *testing.T) {
			assert.True(t, e.HasInterface(row.symbol))
			result, err := e.InvokeOperation(row.symbol, row.operation, row.args)
			assert.NoError(t, err)
			assert.Nil(t, result)
			props, err := e.GetProperties(row.symbol)
			assert.NoError(t, err)
			assert.Equal(t, row.props, props)
		})
	}
}

func TestResultOfInvokeOperation(t *testing.T) {
	e := LoadTest(t)
	var table = []struct {
		symbol    string
		operation string
		args      map[string]any
		result    map[string]any
	}{
		{"demo.Counter", "increment", nil, nil},
		{"demo.Counter", "decrement", nil, nil},
		{"demo.Counter", "getCount", nil, map[string]any{"value": 2}},
	}
	for _, row := range table {
		t.Run(row.symbol+"."+row.operation, func(t *testing.T) {
			assert.True(t, e.HasInterface(row.symbol))
			result, err := e.InvokeOperation(row.symbol, row.operation, row.args)
			assert.NoError(t, err)
			assert.Equal(t, row.result, result)
		})
	}
}

func TestEngine_HasInterface(t *testing.T) {
	e := NewEngine()
	assert.False(t, e.HasInterface("Interface1"))
}
