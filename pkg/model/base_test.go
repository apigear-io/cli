package model

import (
	"testing"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/stretchr/testify/assert"
)

func TestVoidReturn(t *testing.T) {
	var module Module
	err := helper.ReadDocument("./testdata/module.yaml", &module)
	assert.NoError(t, module.Validate())
	assert.NoError(t, err)
	assert.Equal(t, 5, len(module.Interfaces))
	iface3 := module.Interfaces[3]
	assert.Equal(t, 3, len(iface3.Operations))
	op0 := iface3.Operations[0]
	assert.True(t, op0.Return.IsVoid())
	op1 := iface3.Operations[1]
	assert.True(t, op1.Return.IsVoid())
	op2 := iface3.Operations[2]
	assert.False(t, op2.Return.IsVoid())
}
