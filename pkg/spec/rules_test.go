package spec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var rulesDoc = RulesDoc{
	Name: "rules",
	Features: []*FeatureRule{
		{
			Name: "feature1",
			Scopes: []*ScopeRule{
				{
					Match:  ScopeInterface,
					Prefix: "interface",
					Documents: []DocumentRule{
						{
							Source: "interface.yaml",
							Target: "interface.md",
						},
					},
				},
			},
		},
	},
}

var rulesDoc2 = RulesDoc{
	Name: "rules",
	Features: []*FeatureRule{
		{
			Name: "feature1",
			Requires: []string{
				"feature2",
			},
		},
		{
			Name: "feature2",
			Requires: []string{
				"feature3",
				"feature1",
			},
		},
		{
			Name: "feature3",
		},
	},
}

func TestRulesDoc(t *testing.T) {
	err := rulesDoc.Validate()
	assert.NoError(t, err)
}

func TestRulesCircularDependency(t *testing.T) {
	err := rulesDoc2.Validate()
	assert.NoError(t, err)
	for _, f := range rulesDoc2.Features {
		assert.False(t, f.Skip)
	}
}

var rulesDoc3 = RulesDoc{
	Name: "rules",
	Features: []*FeatureRule{
		{
			Name: "feature1",
			Requires: []string{
				"feature2",
			},
		},
		{
			Name: "feature2",
			Requires: []string{
				"feature1",
			},
		},
		{
			Name: "feature3",
		},
	},
}

func TestSkippedDependencies(t *testing.T) {
	doc := rulesDoc3
	err := doc.Validate()
	assert.NoError(t, err)
	err = doc.ComputeFeatures([]string{"feature1"})
	assert.NoError(t, err)
	assert.False(t, doc.Features[0].Skip)
	assert.False(t, doc.Features[1].Skip)
	assert.True(t, doc.Features[2].Skip)
}
