package filtercpp

import (
	"testing"

	"github.com/apigear-io/cli/pkg/idl"
	"github.com/apigear-io/cli/pkg/model"
	"github.com/stretchr/testify/assert"
)

func loadTestSystems(t *testing.T) []*model.System {
	t.Helper()
	sys1 := model.NewSystem("sys1")
	p := idl.NewParser(sys1)
	err := p.ParseFile("../testdata/test.idl")
	assert.NoError(t, err)
	err = sys1.Validate()
	assert.NoError(t, err)

	sys2 := model.NewSystem("sys2")
	dp := model.NewDataParser(sys2)
	err = dp.ParseFile("../testdata/test.module.yaml")
	assert.NoError(t, err)
	err = sys2.Validate()
	assert.NoError(t, err)
	return []*model.System{sys1}
}

func loadExternSystems(t *testing.T) []*model.System {
	t.Helper()
	sys1 := model.NewSystem("sys1")
	p := idl.NewParser(sys1)
	err := p.ParseFile("../testdata/extern.idl")
	assert.NoError(t, err)

	err = p.ParseFile("../testdata/extern2.idl")
	assert.NoError(t, err)
	err = sys1.Validate()
	assert.NoError(t, err)

	parser := model.NewDataParser(sys1)
	err = parser.ParseFile("../testdata/test.module.yaml")
	assert.NoError(t, err)
	err = sys1.Validate()
	assert.NoError(t, err)

	err = parser.ParseFile("../testdata/extern_types.module.yaml")
	assert.NoError(t, err)
	err = sys1.Validate()
	assert.NoError(t, err)

	err = parser.ParseFile("../testdata/test_apigear_next.module.yaml")
	assert.NoError(t, err)
	err = sys1.Validate()
	assert.NoError(t, err)

	return []*model.System{sys1}
}
