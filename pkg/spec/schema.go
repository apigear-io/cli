package spec

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

//go:embed data/apigear.module.schema.json
var ApigearModuleSchema []byte

//go:embed data/apigear.solution.schema.json
var ApigearSolutionSchema []byte

//go:embed data/apigear.scenario.schema.json
var ApigearScenarioSchema []byte

//go:embed data/apigear.rules.schema.json
var ApigearRulesSchema []byte

type DocumentType string

const (
	DocumentTypeModule   DocumentType = "module"
	DocumentTypeSolution DocumentType = "solution"
	DocumentTypeScenario DocumentType = "scenario"
	DocumentTypeRules    DocumentType = "rules"
	DocumentTypeUnknown  DocumentType = "unknown"
)

func LintDocumentFromString(t DocumentType, data []byte) (*gojsonschema.Result, error) {
	var schema []byte
	var err error
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
		return nil, fmt.Errorf("unsupported document type: %s", t)
	}
	// load document from json
	schemaLoader := gojsonschema.NewBytesLoader(schema)
	_, err = schemaLoader.LoadJSON()
	if err != nil {
		return nil, fmt.Errorf("error loading schema: %w", err)
	}
	// parse the yaml document to go structure
	// docJson, err := helper.YamlToJson(data)
	// if err != nil {
	// 	return nil, fmt.Errorf("error parsing yaml document: %w", err)
	// }
	v := make(map[string]interface{})
	err = yaml.Unmarshal(data, &v)
	if err != nil {
		return nil, fmt.Errorf("error reading document: %w", err)
	}
	docJson, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("error marshalling document: %w", err)
	}
	// load the go structure to json
	docLoader := gojsonschema.NewBytesLoader(docJson)

	// validate the json document
	_, err = docLoader.LoadJSON()
	if err != nil {
		return nil, fmt.Errorf("error loading document: %w", err)
	}
	result, err := gojsonschema.Validate(schemaLoader, docLoader)
	if err != nil {
		return nil, fmt.Errorf("failed to validate document: %w", err)
	}
	return result, nil
}

func GetDocumentType(fn string) (DocumentType, error) {
	t, err := GetDocumentTypeAsString(fn)
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

func GetDocumentTypeAsString(fn string) (string, error) {
	words := strings.Split(fn, ".")
	if len(words) < 2 {
		return "", fmt.Errorf("invalid filename: %s", fn)
	}
	return words[len(words)-2], nil
}
