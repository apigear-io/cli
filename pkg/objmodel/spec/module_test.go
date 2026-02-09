package spec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNames(t *testing.T) {
	result, err := CheckFile("./testdata/names.module.yaml")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(result.Errors))
}
