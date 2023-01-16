package gen

import (
	"testing"

	"github.com/apigear-io/cli/pkg/spec"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestRulesFeatures(t *testing.T) {
	rules := readRules(t, "testdata/fts/rules.yaml")
	fts := rules.ComputeFeatures([]string{"f1"})
	fs := spec.FeatureRulesToStrings(fts)
	assert.Equal(t, fs, []string{"f1"})

	fts = rules.ComputeFeatures([]string{"f2"})
	fs = spec.FeatureRulesToStrings(fts)
	assert.Equal(t, fs, []string{"f1", "f2"})

	fts = rules.ComputeFeatures([]string{"f3"})
	fs = spec.FeatureRulesToStrings(fts)
	assert.Equal(t, fs, []string{"f1", "f2", "f3"})
}

func TestGeneratorRulesRequireF1(t *testing.T) {
	_, o := createMockGenerator(t, "testdata/fts", []string{"f1"})
	assert.Len(t, o.Writes, 1)
	var fts map[string]interface{}
	err := yaml.Unmarshal([]byte(o.Writes["testdata/output/f1.yml"]), &fts)
	assert.NoError(t, err)
	assert.Equal(t, fts, map[string]interface{}{"f1": true})

}

func TestGeneratorRulesRequireF2(t *testing.T) {
	_, o := createMockGenerator(t, "testdata/fts", []string{"f2"})
	assert.Len(t, o.Writes, 2)
	var fts map[string]interface{}
	err := yaml.Unmarshal([]byte(o.Writes["testdata/output/f1.yml"]), &fts)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"f1": true, "f2": true}, fts)
	err = yaml.Unmarshal([]byte(o.Writes["testdata/output/f2.yml"]), &fts)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"f1": true, "f2": true}, fts)
}

func TestGeneratorRulesRequireF3(t *testing.T) {
	_, o := createMockGenerator(t, "testdata/fts", []string{"f3"})
	assert.Len(t, o.Writes, 3)
	var fts map[string]interface{}
	err := yaml.Unmarshal([]byte(o.Writes["testdata/output/f1.yml"]), &fts)
	assert.NoError(t, err)
	assert.Equal(t, fts, map[string]interface{}{"f1": true, "f2": true, "f3": true})
	err = yaml.Unmarshal([]byte(o.Writes["testdata/output/f2.yml"]), &fts)
	assert.NoError(t, err)
	assert.Equal(t, fts, map[string]interface{}{"f1": true, "f2": true, "f3": true})
	err = yaml.Unmarshal([]byte(o.Writes["testdata/output/f3.yml"]), &fts)
	assert.NoError(t, err)
	assert.Equal(t, fts, map[string]interface{}{"f1": true, "f2": true, "f3": true})
}

func TestGeneratorRulesRequireAll(t *testing.T) {
	_, o := createMockGenerator(t, "testdata/fts", []string{})
	assert.Len(t, o.Writes, 3)
	var fts map[string]interface{}
	err := yaml.Unmarshal([]byte(o.Writes["testdata/output/f1.yml"]), &fts)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"f1": true, "f2": true, "f3": true}, fts)
}
