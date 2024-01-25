package spec

import (
	"fmt"
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

func TestCheckEngines(t *testing.T) {
	doc := RulesDoc{
		Name: "rules",
		Engines: Engines{
			Cli: ">= 1.0.0",
		},
	}
	type row struct {
		constraint string
		version    string
		check      bool
		errorsLen  int
	}
	tests := []row{
		{constraint: "1.0.0", version: "1.0.0", check: true, errorsLen: 0},
		{constraint: "1.0.0", version: "1.0.1", check: false, errorsLen: 1},
		{constraint: ">=1.0.0", version: "1.1.0", check: true, errorsLen: 0},
		{constraint: ">=1.0.0", version: "0.1.0", check: false, errorsLen: 1},
		{constraint: ">=1.0.0", version: "1.0.0", check: true, errorsLen: 0},
		{constraint: "<=1.0.0", version: "0.9.0", check: true, errorsLen: 0},
		{constraint: "<=1.0.0", version: "1.0.0", check: true, errorsLen: 0},
		{constraint: "<=1.0.0", version: "1.1.0", check: false, errorsLen: 1},
		{constraint: "", version: "1.0.0", check: true, errorsLen: 0},
		{constraint: ">=1.0.0", version: "", check: true, errorsLen: 0},
		{constraint: ">=1.0.0", version: "1.0.0-beta", check: false, errorsLen: 1},
		{constraint: ">=1.0.0", version: "(devel)", check: true, errorsLen: 0},
		{constraint: ">=1.0.0", version: "v0.9", check: false, errorsLen: 1},
		{constraint: ">=1.0.0", version: "v1.0", check: true, errorsLen: 0},
		{constraint: ">=1.0.0", version: "v1.1", check: true, errorsLen: 0},
	}

	for _, test := range tests {
		doc = RulesDoc{
			Name: "rules",
			Engines: Engines{
				Cli: test.constraint,
			},
		}
		check, errs := doc.CheckEngines(test.version)
		label := fmt.Sprintf("%s vs. version %s", test.constraint, test.version)
		assert.Len(t, errs, test.errorsLen, label)
		assert.Equal(t, test.check, check, label)
	}
}
