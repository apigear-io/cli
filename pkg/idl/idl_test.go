package idl

import (
	"testing"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/stretchr/testify/assert"
)

func loadIdl(name string, files []string) (*model.System, error) {
	system := model.NewSystem(name)
	for _, file := range files {
		parser := NewParser(system)
		err := parser.ParseFile(file)
		if err != nil {
			return nil, err
		}
	}
	return system, nil
}

func loadIdlFromString(name string, content string) (*model.System, error) {
	system := model.NewSystem(name)
	parser := NewParser(system)
	err := parser.ParseString(content)
	if err != nil {
		return nil, err
	}
	return system, nil
}

func TestSimpleIdl(t *testing.T) {
	s, err := loadIdl("simple", []string{"./testdata/simple.idl"})
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, "simple", s.Name)
	assert.Equal(t, 1, len(s.Modules))
	assert.Equal(t, "tb.simple", s.Modules[0].Name)
}
