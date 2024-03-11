package filterrs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

// ToDefaultString returns the default value for a type
func ToDefaultString(prefix string, schema *model.Schema) (string, error) {
	text := ""
	switch schema.Type {
	case "void":
		text = "Default::default()"
	case "string":
		text = "Default::default()"
	case "int", "int32":
		text = "Default::default()"
	case "int64":
		text = "Default::default()"
	case "float", "float32":
		text = "Default::default()"
	case "float64":
		text = "Default::default()"
	case "bool":
		text = "Default::default()"
	default:
		if schema.Module == nil {
			return "xxx", fmt.Errorf("schema.Module is nil")
		}
		e := schema.Module.LookupEnum(schema.Import, schema.Type)
		if e != nil {
			text = "Default::default()"
		}
		s := schema.Module.LookupStruct(schema.Import, schema.Type)
		if s != nil {
			text = "Default::default()"
		}
		i := schema.Module.LookupInterface(schema.Import, schema.Type)
		if i != nil {
			text = "Default::default()"
		}
	}
	if schema.IsArray {
		text = "Default::default()"
	}
	return text, nil
}

// rsDefault returns the default value for a type
func rsDefault(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("rsDefault node is nil")
	}
	return ToDefaultString(prefix, &node.Schema)
}
