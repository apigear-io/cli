package sim

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeSymbol(t *testing.T) {
	assert.Equal(t, "module.interface/resource", MakeSymbol("module.interface", "resource").String())
	assert.Equal(t, "module.interface", MakeSymbol("module.interface", "").String())
}

func TestSplitSymbol(t *testing.T) {
	iface, resource := MakeSymbol("module.interface", "resource").Split()
	assert.Equal(t, "module.interface", iface)
	assert.Equal(t, "resource", resource)
	iface, resource = MakeSymbol("module.interface", "").Split()
	assert.Equal(t, "module.interface", iface)
	assert.Equal(t, "", resource)
}

func TestGetResource(t *testing.T) {
	assert.Equal(t, "resource", MakeSymbol("module.interface", "resource").Resource())
}

func TestGetPath(t *testing.T) {
	assert.Equal(t, "module.interface", MakeSymbol("module.interface", "resource").Path())
}
