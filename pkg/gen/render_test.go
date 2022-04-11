package gen

import (
	"objectapi/pkg/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderString(t *testing.T) {
	var e = NewRenderer("testdata/templates")
	s, err := e.RenderString("{{.System.Name}}", DataMap{"System": model.NewSystem("test")})
	assert.NoError(t, err)
	assert.Equal(t, "test", s)

}
