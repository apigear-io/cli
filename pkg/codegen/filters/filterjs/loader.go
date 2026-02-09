package filterjs

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
