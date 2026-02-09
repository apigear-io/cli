package filtergo

import (
	"testing"

	"github.com/apigear-io/cli/pkg/apimodel/idl"
	"github.com/apigear-io/cli/pkg/apimodel"
	"github.com/stretchr/testify/assert"
)

func loadTestSystems(t *testing.T) []*apimodel.System {
	t.Helper()
	sys1 := apimodel.NewSystem("sys1")
	p := idl.NewParser(sys1)
	err := p.ParseFile("../testdata/test.idl")
	assert.NoError(t, err)
	err = sys1.Validate()
	assert.NoError(t, err)

	return []*apimodel.System{sys1}
}

func loadExternSystems(t *testing.T) []*apimodel.System {
	t.Helper()
	sys1 := apimodel.NewSystem("sys1")
	p := idl.NewParser(sys1)
	err := p.ParseFile("../testdata/extern.idl")
	assert.NoError(t, err)

	err = p.ParseFile("../testdata/extern2.idl")
	assert.NoError(t, err)
	err = sys1.Validate()
	assert.NoError(t, err)

	return []*apimodel.System{sys1}
}
