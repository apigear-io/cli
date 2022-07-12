package filtercpp

import (
	"testing"

	"github.com/apigear-io/cli/pkg/idl"
	"github.com/apigear-io/cli/pkg/model"

	"github.com/stretchr/testify/assert"
)

func loadSystem(t *testing.T) *model.System {
	p := idl.NewParser(model.NewSystem("test"))
	err := p.ParseFile("../testdata/test.idl")
	assert.NoError(t, err)
	return p.System
}
