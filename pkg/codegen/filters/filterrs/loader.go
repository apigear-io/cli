package filterrs

import (
	"testing"

	"github.com/apigear-io/cli/pkg/objmodel/idl"
	"github.com/apigear-io/cli/pkg/objmodel"
	"github.com/stretchr/testify/assert"
)

func loadTestSystems(t *testing.T) []*objmodel.System {
	t.Helper()
	sys1 := objmodel.NewSystem("sys1")
	p := idl.NewParser(sys1)
	err := p.ParseFile("../testdata/test.idl")
	assert.NoError(t, err)
	err = sys1.Validate()
	assert.NoError(t, err)

	sys2 := objmodel.NewSystem("sys2")
	dp := objmodel.NewDataParser(sys2)
	err = dp.ParseFile("../testdata/test.module.yaml")
	assert.NoError(t, err)
	err = sys2.Validate()
	assert.NoError(t, err)
	return []*objmodel.System{sys1}
}
