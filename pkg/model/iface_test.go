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

func TestMethods(t *testing.T) {
	var module Module
	err := helper.ReadDocument("./testdata/module.yaml", &module)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(module.Interfaces[1].Methods))
	iface1 := module.Interfaces[1]
	method0 := iface1.Methods[0]
	assert.Equal(t, "method01", method0.Name)
	assert.Equal(t, 1, len(method0.Inputs))
	assert.Equal(t, "input01", method0.Inputs[0].Name)
	assert.Equal(t, "bool", method0.Inputs[0].Schema.Type)
	assert.Equal(t, "bool", method0.Output.Schema.Type)

}
