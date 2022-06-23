package filtercpp

import (
	"apigear/pkg/idl"
	"apigear/pkg/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func loadSystem(t *testing.T) *model.System {
	p := idl.NewParser(model.NewSystem("test"))
	err := p.ParseFile("../testdata/test.idl")
	assert.NoError(t, err)
	return p.System
}
