package model

import (
	"objectapi/pkg/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModuleYaml(t *testing.T) {
	var module Module
	err := helper.ReadDocument("./test/module.yaml", &module)
	assert.NoError(t, err)
	assert.Equal(t, "Module001", module.Name)
	assert.Equal(t, "1.0", module.Version)
}

func TestModuleJson(t *testing.T) {
	var module Module
	err := helper.ReadDocument("./test/module.json", &module)
	assert.NoError(t, err)
	assert.Equal(t, "Module001", module.Name)
	assert.Equal(t, "1.0", module.Version)
}
