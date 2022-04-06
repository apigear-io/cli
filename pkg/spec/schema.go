package spec

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"

	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

//go:embed schema/apigear.module.schema.json
var ApigearModuleSchema []byte

//go:embed schema/apigear.solution.schema.json
var ApigearSolutionSchema []byte

//go:embed schema/apigear.scenario.schema.json
var ApigearScenarioSchema []byte

//go:embed schema/apigear.rules.schema.json
var ApigearRulesSchema []byte

type DocumentType string

const (
	DocumentTypeModule   DocumentType = "module"
	DocumentTypeSolution DocumentType = "solution"
	DocumentTypeScenario DocumentType = "scenario"
	DocumentTypeRules    DocumentType = "rules"
	DocumentTypeUnknown  DocumentType = "unknown"
)

func LintJsonDoc(t DocumentType, jsonDoc []byte) (*gojsonschema.Result, error) {
	schemaLoader, err := LoadSchema(t)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("error marshalling document: %w", err)
	}
	// load the go structure to json
	jsonLoader := gojsonschema.NewBytesLoader(jsonDoc)

	// validate the json document
	_, err = jsonLoader.LoadJSON()
	if err != nil {
		return nil, fmt.Errorf("error loading json document: %w", err)
	}
	result, err := gojsonschema.Validate(schemaLoader, jsonLoader)
	if err != nil {
		return nil, fmt.Errorf("failed to validate document: %w", err)
	}
	return result, nil
}

func CheckFile(file string) (*gojsonschema.Result, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	if path.Ext(file) == ".yaml" || path.Ext(file) == ".yml" {
		data, err = YamlToJson(data)
		if err != nil {
			return nil, err
		}
	}
	t, err := GetDocumentType(file)
	if err != nil {
		return nil, err
	}
	return LintJsonDoc(t, data)
}

func YamlToJson(data []byte) ([]byte, error) {
	v := make(map[string]interface{})
	err := yaml.Unmarshal(data, &v)
	if err != nil {
		return nil, fmt.Errorf("error reading document: %w", err)
	}
	return json.MarshalIndent(v, "", "  ")
}

func JsonToYaml(data []byte) ([]byte, error) {
	v := make(map[string]interface{})
	err := json.Unmarshal(data, &v)
	if err != nil {
		return nil, fmt.Errorf("error reading document: %w", err)
	}
	return yaml.Marshal(v)
}

func LoadSchema(t DocumentType) (gojsonschema.JSONLoader, error) {
	var schema []byte
	switch t {
	case DocumentTypeModule:
		schema = ApigearModuleSchema
	case DocumentTypeSolution:
		schema = ApigearSolutionSchema
	case DocumentTypeScenario:
		schema = ApigearScenarioSchema
	case DocumentTypeRules:
		schema = ApigearRulesSchema
	default:
		panic(fmt.Errorf("unsupported document type: %s", t))
	}
	// load document from json
	schemaLoader := gojsonschema.NewBytesLoader(schema)
	_, err := schemaLoader.LoadJSON()
	if err != nil {
		return nil, fmt.Errorf("error loading schema: %w", err)
	}
	return schemaLoader, nil
}

func GetDocumentType(file string) (DocumentType, error) {
	base := filepath.Base(file)
	t, err := DocumentTypeFromFileName(base)
	if err != nil {
		return DocumentTypeUnknown, err
	}
	switch t {
	case "module":
		return DocumentTypeModule, nil
	case "solution":
		return DocumentTypeSolution, nil
	case "scenario":
		return DocumentTypeScenario, nil
	case "rules":
		return DocumentTypeRules, nil
	default:
		return DocumentTypeUnknown, fmt.Errorf("unsupported document type: %s", t)
	}
}

func DocumentTypeFromFileName(fn string) (string, error) {
	if fn == "rules.yaml" {
		return "rules", nil
	}
	words := strings.Split(fn, ".")
	if len(words) < 2 {
		return "", fmt.Errorf("invalid filename: %s", fn)
	}
	return words[len(words)-2], nil
}
