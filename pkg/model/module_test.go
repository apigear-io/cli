package model

import (
	"testing"

	"github.com/apigear-io/cli/pkg/helper"

	"github.com/stretchr/testify/assert"
)

func readSystem(t *testing.T) *System {
	var system System
	var aModule Module
	err := helper.ReadDocument("./testdata/a.module.yaml", &aModule)
	assert.NoError(t, err)
	system.AddModule(&aModule)

	var bModule Module
	err = helper.ReadDocument("./testdata/b.module.yaml", &bModule)
	assert.NoError(t, err)
	system.AddModule(&bModule)
	return &system
}

func TestModuleYaml(t *testing.T) {
	var module Module
	err := helper.ReadDocument("./testdata/module.yaml", &module)
	assert.NoError(t, err)
	assert.Equal(t, "Module01", module.Name)
	assert.Equal(t, "1.0.0", string(module.Version))
}

func TestModuleJson(t *testing.T) {
	var module Module
	err := helper.ReadDocument("./testdata/module.json", &module)
	assert.NoError(t, err)
	assert.Equal(t, "Module01", module.Name)
	assert.Equal(t, "1.0.0", string(module.Version))
}

func TestChecksum(t *testing.T) {
	var module Module
	err := helper.ReadDocument("./testdata/module.yaml", &module)
	assert.NoError(t, err)
	err = module.Validate()
	assert.NoError(t, err)
	module.computeChecksum()
	assert.Equal(t, "aacb40d122fb8a126754d15e1c78e2ad", module.Checksum)
	assert.Equal(t, 32, len(module.Checksum))
}

func TestModuleImport(t *testing.T) {
	system := readSystem(t)
	assert.NotNil(t, system)
	assert.Equal(t, 2, len(system.Modules))
	module := system.LookupModule("b")
	assert.NotNil(t, module)
	assert.Equal(t, 1, len(module.Imports))
	assert.Equal(t, "a", module.Imports[0].Name)

	prop := module.Interfaces[0].Properties[0]
	assert.Equal(t, "value", prop.Name)
	assert.Equal(t, "a.A", prop.Type)

	s := module.LookupStruct("a", "A")
	assert.NotNil(t, s)
	s2 := module.LookupStruct("a", "B")
	assert.Nil(t, s2)

	s3 := system.LookupStruct("a", "A")
	assert.NotNil(t, s3)

	assert.Equal(t, s, s3)

}
