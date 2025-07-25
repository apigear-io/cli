package spec

import (
	_ "embed"
	"fmt"
)

//go:embed schema/apigear.module.schema.yaml
var ApigearModuleYamlSchema []byte

//go:embed schema/apigear.solution.schema.yaml
var ApigearSolutionYamlSchema []byte

//go:embed schema/apigear.scenario.schema.yaml
var ApigearScenarioYamlSchema []byte

//go:embed schema/apigear.rules.schema.yaml
var ApigearRulesYamlSchema []byte

type SchemaFormat string

const (
	SchemaFormatJson SchemaFormat = "json"
	SchemaFormatYaml SchemaFormat = "yaml"
)

func ShowSchemaFile(t DocumentType, f SchemaFormat) (*string, error) {
	var schema []byte
	switch t {
	case DocumentTypeModule:
		switch f {
		case SchemaFormatJson:
			schema = ApigearModuleSchema
		case SchemaFormatYaml:
			schema = ApigearModuleYamlSchema
		default:
			return nil, fmt.Errorf("unsupported schema format: %s", f)
		}
	case DocumentTypeSolution:
		switch f {
		case SchemaFormatJson:
			schema = ApigearSolutionSchema
		case SchemaFormatYaml:
			schema = ApigearSolutionYamlSchema
		default:
			return nil, fmt.Errorf("unsupported schema format: %s", f)
		}
	case DocumentTypeScenario:
		switch f {
		case SchemaFormatJson:
			schema = ApigearScenarioSchema
		case SchemaFormatYaml:
			schema = ApigearScenarioYamlSchema
		default:
			return nil, fmt.Errorf("unsupported schema format: %s", f)
		}
	case DocumentTypeRules:
		switch f {
		case SchemaFormatJson:
			schema = ApigearRulesSchema
		case SchemaFormatYaml:
			schema = ApigearRulesYamlSchema
		default:
			return nil, fmt.Errorf("unsupported schema format: %s", f)
		}
	default:
		return nil, fmt.Errorf("unsupported document type: %s", t)
	}
	result := string(schema)
	return &result, nil
}
