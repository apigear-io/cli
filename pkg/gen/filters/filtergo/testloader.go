package filtergo

import (
	"testing"

	"github.com/apigear-io/cli/pkg/idl"
	"github.com/apigear-io/cli/pkg/model"

	"github.com/stretchr/testify/assert"
)

func loadSystem(t *testing.T) *model.System {
	system := model.NewSystem("test")
	p := idl.NewParser(system)
	err := p.ParseFile("../testdata/test.idl")
	assert.NoError(t, err)
	err = system.ResolveAll()
	assert.NoError(t, err)
	return system
}
