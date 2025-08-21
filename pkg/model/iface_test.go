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
	assert.Equal(t, "1.0.0", string(module.Version))
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

const duplicatesYAML = `schema: apigear.module/1.0
name: demo
version: "0.0.1"
interfaces:
  - name: Hello
  - name: Hello`

func TestInterfaceNameDuplicates(t *testing.T) {
	var module Module
	err := helper.ReadYamlFromString(duplicatesYAML, &module)
	assert.NoError(t, err)
	err = module.Validate()
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "demo: duplicate name Hello")
}

const duplicates2YAML = `schema: apigear.module/1.0
name: demo
version: "0.0.1"
interfaces:
  - name: Hello
structs:
  - name: Hello`

func TestStructNameDuplicates(t *testing.T) {
	var module Module
	err := helper.ReadYamlFromString(duplicates2YAML, &module)
	assert.NoError(t, err)
	err = module.Validate()
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "demo: duplicate name Hello")
}

const duplicates3YAML = `schema: apigear.module/1.0
name: demo
version: "0.0.1"
interfaces:
  - name: Hello
enums:
  - name: Hello
`

func TestEnumNameDuplicates(t *testing.T) {
	var module Module
	err := helper.ReadYamlFromString(duplicates3YAML, &module)
	assert.NoError(t, err)
	err = module.Validate()
	assert.Error(t, err)
	assert.Equal(t, "demo: duplicate name Hello", err.Error())
}

func TestExtends(t *testing.T) {
	var module Module
	err := helper.ReadDocument("./testdata/module.yaml", &module)
	assert.NoError(t, err)
	err = module.Validate()
	assert.NoError(t, err)
	iface6 := module.LookupInterface("Module01", "Interface06")
	iface0 := module.LookupInterface("Module01", "Interface01")
	assert.True(t, iface6.HasExtends())
	assert.Equal(t, "Interface01", iface6.Extends.Name)
	assert.Equal(t, "", iface6.Extends.Import)
	assert.Equal(t, iface0, iface6.Extends.Reference)
}
