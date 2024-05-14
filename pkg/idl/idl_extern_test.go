package idl

import (
	"testing"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/stretchr/testify/assert"
)

func loadExternIdl(t *testing.T) *model.System {
	t.Helper()
	sys1 := model.NewSystem("sys1")
	o := NewParser(sys1)
	err := o.ParseFile("./testdata/extern.idl")
	assert.NoError(t, err)
	err = sys1.Validate()
	assert.NoError(t, err)
	return sys1
}

func loadExternYaml(t *testing.T) *model.System {
	t.Helper()
	sys1 := model.NewSystem("sys1")
	dp := model.NewDataParser(sys1)
	err := dp.ParseFile("./testdata/extern.module.yaml")
	assert.NoError(t, err)
	err = sys1.Validate()
	assert.NoError(t, err)
	return sys1
}
func TestExternYamlDeclaration(t *testing.T) {
	sys := loadExternYaml(t)
	tt := []struct {
		module string
		name   string
	}{
		{"demo", "XType1"},
		{"demo", "XType2"},
		{"demo", "XType3"},
	}

	for _, tc := range tt {
		xe := sys.LookupExtern(tc.module, tc.name)
		assert.NotNil(t, xe)
		assert.Equal(t, tc.name, xe.Name)
	}
}

func TestExternIdlDeclaration(t *testing.T) {
	sys := loadExternIdl(t)
	tt := []struct {
		module string
		name   string
	}{
		{"demo", "XType1"},
		{"demo", "XType2"},
		{"demo", "XType3"},
	}

	for _, tc := range tt {
		xe := sys.LookupExtern(tc.module, tc.name)
		assert.NotNil(t, xe)
		assert.Equal(t, tc.name, xe.Name)
	}
}
