package filterjava

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

	sys2 := apimodel.NewSystem("sys2")
	dp := apimodel.NewDataParser(sys2)
	err = dp.ParseFile("../testdata/test.module.yaml")
	assert.NoError(t, err)
	err = sys2.Validate()
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

func loadExternSystemsYAML(t *testing.T) []*apimodel.System {
	t.Helper()
	api_next_system := apimodel.NewSystem("api_next_system")
	parser := apimodel.NewDataParser(api_next_system)
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

	return []*apimodel.System{api_next_system}
}
