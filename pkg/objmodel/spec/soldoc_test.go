package spec

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var docTargets = SolutionDoc{
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

func TestAggregateDependencies(t *testing.T) {
	t.Run("returns empty when no targets", func(t *testing.T) {
		doc := &SolutionDoc{
			Name:    "test",
			Targets: []*SolutionTarget{},
		}

		deps := doc.AggregateDependencies()
		require.NotNil(t, deps)
		require.Empty(t, deps)
	})

	t.Run("aggregates dependencies from single target", func(t *testing.T) {
		doc := &SolutionDoc{
			Name: "test",
			Targets: []*SolutionTarget{
				{
					Name:     "target1",
					computed: true,
					dependencies: []string{
						"dep1.yaml",
						"dep2.yaml",
					},
				},
			},
		}

		deps := doc.AggregateDependencies()
		require.Len(t, deps, 2)
		require.Contains(t, deps, "dep1.yaml")
		require.Contains(t, deps, "dep2.yaml")
	})

	t.Run("aggregates dependencies from multiple targets", func(t *testing.T) {
		doc := &SolutionDoc{
			Name: "test",
			Targets: []*SolutionTarget{
				{
					Name:     "target1",
					computed: true,
					dependencies: []string{
						"dep1.yaml",
						"dep2.yaml",
					},
				},
				{
					Name:     "target2",
					computed: true,
					dependencies: []string{
						"dep3.yaml",
						"dep4.yaml",
					},
				},
			},
		}

		deps := doc.AggregateDependencies()
		require.Len(t, deps, 4)
		require.Contains(t, deps, "dep1.yaml")
		require.Contains(t, deps, "dep2.yaml")
		require.Contains(t, deps, "dep3.yaml")
		require.Contains(t, deps, "dep4.yaml")
	})

	t.Run("handles targets with no dependencies", func(t *testing.T) {
		doc := &SolutionDoc{
			Name: "test",
			Targets: []*SolutionTarget{
				{
					Name:     "target1",
					computed: true,
					dependencies: []string{
						"dep1.yaml",
					},
				},
				{
					Name:         "target2",
					computed:     true,
					dependencies: []string{},
				},
			},
		}

		deps := doc.AggregateDependencies()
		require.Len(t, deps, 1)
		require.Contains(t, deps, "dep1.yaml")
	})
}
