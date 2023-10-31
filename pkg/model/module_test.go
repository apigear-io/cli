package model

import (
	"testing"

	"github.com/apigear-io/cli/pkg/helper"

	"github.com/stretchr/testify/assert"
)

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
