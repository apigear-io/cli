package gen

import (
	"io/ioutil"
	"objectapi/pkg/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

type MockFileWriter struct {
	Writes map[string]string
}

func (m *MockFileWriter) WriteFile(fn string, content string) error {
	m.Writes[fn] = content
	return nil
}

func NewMockFileWriter() *MockFileWriter {
	return &MockFileWriter{
		Writes: make(map[string]string),
	}
}

type MockRenderEngine struct {
	Results []string
}

func (m *MockRenderEngine) RenderString(template string, ctx Context) (string, error) {
	m.Results = append(m.Results, template)
	return template, nil
}

func (m *MockRenderEngine) RenderFile(name string, ctx Context) (string, error) {
	m.Results = append(m.Results, name)
	return name, nil
}

func NewMockRenderEngine() *MockRenderEngine {
	return &MockRenderEngine{
		Results: make([]string, 0),
	}
}

func readRules(t *testing.T, filename string) RulesDoc {
	content, err := ioutil.ReadFile(filename)
	assert.NoError(t, err)
	var file RulesDoc
	err = yaml.Unmarshal(content, &file)
	assert.NoError(t, err)
	return file
}

func createProcessor() *Generator {
	var engine = NewMockRenderEngine()
	var writer = NewMockFileWriter()
	var processor = NewGenerator(engine, writer)
	return processor
}
func TestEmptyRules(t *testing.T) {
	s := model.NewSystem("test")
	processor := createProcessor()
	r := readRules(t, "testdata/empty.rules.yaml")
	processor.Process(r, s)
}

func TestHelloRules(t *testing.T) {
	s := model.NewSystem("test")
	var e = NewMockRenderEngine()
	var w = NewMockFileWriter()
	var p = NewGenerator(e, w)
	r := readRules(t, "testdata/hello.rules.yaml")
	p.Process(r, s)
	assert.Equal(t, "hello.txt", w.Writes["hello.txt"])
	assert.Equal(t, "hello.txt", e.Results[0])
}
