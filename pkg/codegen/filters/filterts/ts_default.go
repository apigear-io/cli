package filterts

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/apimodel"
)

// ToDefaultString returns the default value for a type
func ToDefaultString(schema *apimodel.Schema, prefix string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("tsDefault called with nil schema")
	}
	if schema.Module == nil {
		return "xxx", fmt.Errorf("tsDefault called with nil schema module")
	}
	var text string
	if schema.IsArray {
		text = "[]"
	} else {
		switch schema.KindType {
		case apimodel.TypeString:
			text = "\"\""
		case apimodel.TypeInt, apimodel.TypeInt32, apimodel.TypeInt64:
			text = "0"
		case apimodel.TypeFloat, apimodel.TypeFloat32, apimodel.TypeFloat64:
			text = "0.0"
		case apimodel.TypeBool:
			text = "false"
		case apimodel.TypeEnum:
			e := schema.LookupEnum(schema.Import, schema.Type)
			if e == nil {
				return "xxx", fmt.Errorf("tsDefault enum not found: %s", schema.Dump())
			}
			text = fmt.Sprintf("%s%s.%s", prefix, e.Name, e.Members[0].Name)
		case apimodel.TypeStruct:
			s := schema.LookupStruct(schema.Import, schema.Type)
			if s == nil {
				return "xxx", fmt.Errorf("tsDefault struct not found: %s", schema.Dump())
			}
			text = "{}"
		case apimodel.TypeInterface:
			i := schema.LookupInterface(schema.Import, schema.Type)
			if i == nil {
				return "xxx", fmt.Errorf("tsDefault interface not found: %s", schema.Dump())
			}
			text = "null"
		case apimodel.TypeVoid:
			text = "void"
		default:
			return "xxx", fmt.Errorf("tsDefault unknown schema %s", schema.Dump())
		}
	}
	return text, nil
}

// cppDefault returns the default value for a type
func tsDefault(prefix string, node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("tsDefault called with nil node")
	}
	return ToDefaultString(&node.Schema, prefix)
}
