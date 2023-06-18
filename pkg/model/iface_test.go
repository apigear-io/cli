package model

import (
	"testing"

	"github.com/apigear-io/cli/pkg/helper"

	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	var module Module
	err := helper.ReadDocument("./testdata/module.yaml", &module)
	assert.NoError(t, err)
	assert.Equal(t, "Module01", module.Name)
	assert.Equal(t, "1.0", string(module.Version))
	assert.Equal(t, 5, len(module.Interfaces))
	iface0 := module.Interfaces[0]
	assert.Equal(t, "Interface01", iface0.Name)

}

func TestProperties(t *testing.T) {
	var module Module
	err := helper.ReadDocument("./testdata/module.yaml", &module)
	assert.NoError(t, err)
	iface0 := module.Interfaces[0]
	assert.Equal(t, 1, len(iface0.Properties))
	prop0 := iface0.Properties[0]
	assert.Equal(t, "prop01", prop0.Name)
	assert.Equal(t, "bool", prop0.Schema.Type)
}

func TestReadonlyProperties(t *testing.T) {
	var module Module
	err := helper.ReadDocument("./testdata/module.yaml", &module)
	assert.NoError(t, err)
	iface0 := module.Interfaces[4]
	assert.Equal(t, 3, len(iface0.Properties))
	p1 := iface0.LookupProperty("prop01")
	assert.NotNil(t, p1)
	assert.False(t, p1.IsReadOnly)
	p2 := iface0.LookupProperty("prop02")
	assert.NotNil(t, p2)
	assert.True(t, p2.IsReadOnly)
	p3 := iface0.LookupProperty("prop03")
	assert.NotNil(t, p3)
	assert.False(t, p3.IsReadOnly)

}

func TestOperations(t *testing.T) {
	var module Module
	err := helper.ReadDocument("./testdata/module.yaml", &module)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(module.Interfaces[1].Operations))
	iface1 := module.Interfaces[1]
	op0 := iface1.Operations[0]
	assert.Equal(t, "operation01", op0.Name)
	assert.Equal(t, 1, len(op0.Params))
	assert.Equal(t, "param01", op0.Params[0].Name)
	assert.Equal(t, "bool", op0.Params[0].Schema.Type)
	assert.Equal(t, "bool", op0.Return.Schema.Type)

}
