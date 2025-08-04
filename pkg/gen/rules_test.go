package gen

import (
	"testing"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/assert"
)

// TestRulesFeatures tests that the rules can compute the features correctly
// See testdata/fts/rules.yaml for the dependency graph
func TestRulesFeatures(t *testing.T) {
	t.Parallel()
	rules := readRules(t, "testdata/fts/rules.yaml")
	err := rules.ComputeFeatures([]string{"f1"})
	assert.NoError(t, err)
	f := rules.FeatureNamesMap()
	assert.Equal(t, f, map[string]bool{"f1": true, "f2": false, "f3": false})

	err = rules.ComputeFeatures([]string{"f2"})
	assert.NoError(t, err)
	f = rules.FeatureNamesMap()
	assert.Equal(t, f, map[string]bool{"f1": true, "f2": true, "f3": false})

	err = rules.ComputeFeatures([]string{"f3"})
	assert.NoError(t, err)
	f = rules.FeatureNamesMap()
	assert.Equal(t, f, map[string]bool{"f1": true, "f2": true, "f3": true})
}

// TestGeneratorRulesRequireF1 tests that the generator will generate the featurs based on dependency f1
// See testdata/fts/rules.yaml for the dependency graph
func TestGeneratorRulesRequireF1(t *testing.T) {
	t.Parallel()
	_, o := createMockGenerator(t, "testdata/fts", []string{"f1"})
	assert.Len(t, o.Writes, 1)
	var fts map[string]interface{}
	target := helper.Join("testdata", "output", "f1.yml")
	err := yaml.Unmarshal([]byte(o.Writes[target]), &fts)
	assert.NoError(t, err)
	assert.Equal(t, fts, map[string]interface{}{"f1": true, "f2": false, "f3": false})
}

// TestGeneratorRulesRequireF2 tests that the generator will generate the featurs based on dependency f2
// See testdata/fts/rules.yaml for the dependency graph
func TestGeneratorRulesRequireF2(t *testing.T) {
	t.Parallel()
	_, o := createMockGenerator(t, "testdata/fts", []string{"f2"})
	assert.Len(t, o.Writes, 2)
	var fts map[string]interface{}
	target := helper.Join("testdata", "output", "f1.yml")
	err := yaml.Unmarshal([]byte(o.Writes[target]), &fts)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"f1": true, "f2": true, "f3": false}, fts)
	target = helper.Join("testdata", "output", "f2.yml")
	err = yaml.Unmarshal([]byte(o.Writes[target]), &fts)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"f1": true, "f2": true, "f3": false}, fts)
}

// TestGeneratorRulesRequireF3 tests that the generator will generate the featurs based on dependency f3
// See testdata/fts/rules.yaml for the dependency graph
func TestGeneratorRulesRequireF3(t *testing.T) {
	t.Parallel()
	_, o := createMockGenerator(t, "testdata/fts", []string{"f3"})
	assert.Len(t, o.Writes, 3)
	var fts map[string]interface{}
	target := helper.Join("testdata", "output", "f1.yml")
	err := yaml.Unmarshal([]byte(o.Writes[target]), &fts)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"f1": true, "f2": true, "f3": true}, fts)
	target = helper.Join("testdata", "output", "f2.yml")
	err = yaml.Unmarshal([]byte(o.Writes[target]), &fts)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"f1": true, "f2": true, "f3": true}, fts)
	target = helper.Join("testdata", "output", "f3.yml")
	err = yaml.Unmarshal([]byte(o.Writes[target]), &fts)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"f1": true, "f2": true, "f3": true}, fts)
}

// TestGeneratorRulesRequireAll tests that the generator will generate the when all featurs are required
func TestGeneratorRulesRequireAll(t *testing.T) {
	t.Parallel()
	_, o := createMockGenerator(t, "testdata/fts", []string{})
	assert.Len(t, o.Writes, 3)
	var fts map[string]interface{}
	target := helper.Join("testdata", "output", "f1.yml")
	err := yaml.Unmarshal([]byte(o.Writes[target]), &fts)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"f1": true, "f2": true, "f3": true}, fts)
}

// TestErrorForMissingFeature tests that the generator will generate error when the feature is missing
func TestErrorForMissingFeature(t *testing.T) {
	t.Parallel()
	rules := readRules(t, "testdata/fts/rules.yaml")
	err := rules.ComputeFeatures([]string{"f4"})
	assert.Error(t, err)
}
