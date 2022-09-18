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
	assert.Equal(t, "1.0", module.Version)
	assert.Equal(t, 3, len(module.Interfaces))
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
