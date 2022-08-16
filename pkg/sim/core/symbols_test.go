package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeSymbol(t *testing.T) {
	assert.Equal(t, "module.interface/resource", MakeSymbol("module.interface", "resource"))
	assert.Equal(t, "module.interface", MakeSymbol("module.interface", ""))
}

func TestSplitSymbol(t *testing.T) {
	iface, resource := SplitSymbol("module.interface/resource")
	assert.Equal(t, "module.interface", iface)
	assert.Equal(t, "resource", resource)
	iface, resource = SplitSymbol("module.interface")
	assert.Equal(t, "module.interface", iface)
	assert.Equal(t, "", resource)
}
