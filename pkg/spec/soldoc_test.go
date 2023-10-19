package spec

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var docTargets = SolutionDoc{
	Schema:      "apigear/solution",
	Version:     "1.0.0",
	Name:        "solution",
	Description: "a description",
	RootDir:     "testdata",
	Targets: []*SolutionTarget{
		{
			Name:     "target1",
			Template: "./tpl/",
			Output:   "./output/",
		},
		{
			Name:     "target2",
			Template: "./tpl/",
			Output:   "./output/",
		},
	},
}

var docLayers = SolutionDoc{
	Schema:      "apigear/solution",
	Version:     "1.0.0",
	Name:        "solution",
	Description: "a description",
	RootDir:     "testdata",
	Layers: []*SolutionTarget{
		{
			Name:     "layer1",
			Template: "./tpl/",
			Output:   "./output/",
		},
		{
			Name:     "layer2",
			Template: "./tpl/",
			Output:   "./output/",
		},
	},
}

func TestUseTargets(t *testing.T) {
	doc := docTargets
	err := doc.Validate()
	require.NoError(t, err)
	require.Equal(t, 2, len(doc.Targets))
	require.Equal(t, "target1", doc.Targets[0].Name)
	require.Equal(t, "target2", doc.Targets[1].Name)
}

func TestUseLayers(t *testing.T) {
	doc := docLayers
	err := doc.Validate()
	require.NoError(t, err)
	require.Equal(t, 0, len(doc.Layers))
	require.Equal(t, 2, len(doc.Targets))
	require.Equal(t, "layer1", doc.Targets[0].Name)
	require.Equal(t, "layer2", doc.Targets[1].Name)
}
