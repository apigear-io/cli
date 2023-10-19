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

func TestRulesDoc(t *testing.T) {
	err := rulesDoc.Validate()
	assert.NoError(t, err)
}
