package filterqt

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
	api_next_system := model.NewSystem("api_next_system")
	parser := model.NewDataParser(api_next_system)
	err := parser.ParseFile("../testdata/test.module.yaml")
	assert.NoError(t, err)
	err = api_next_system.Validate()
	assert.NoError(t, err)

	err = parser.ParseFile("../testdata/extern_types.module.yaml")
	assert.NoError(t, err)
	err = api_next_system.Validate()
	assert.NoError(t, err)

	err = parser.ParseFile("../testdata/test_apigear_next.module.yaml")
	assert.NoError(t, err)
	err = api_next_system.Validate()
	assert.NoError(t, err)

	return []*model.System{api_next_system}
}
