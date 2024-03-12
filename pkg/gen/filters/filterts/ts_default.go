package filterts

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

// ToDefaultString returns the default value for a type
func ToDefaultString(schema *model.Schema, prefix string) (string, error) {
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
		case model.TypeString:
			text = "\"\""
		case model.TypeInt, model.TypeInt32, model.TypeInt64:
			text = "0"
		case model.TypeFloat, model.TypeFloat32, model.TypeFloat64:
			text = "0.0"
		case model.TypeBool:
			text = "false"
		case model.TypeEnum:
			e := schema.LookupEnum(schema.Import, schema.Type)
			if e == nil {
				return "xxx", fmt.Errorf("tsDefault enum not found: %s", schema.Dump())
			}
			text = fmt.Sprintf("%s%s.%s", prefix, e.Name, e.Members[0].Name)
		case model.TypeStruct:
			s := schema.LookupStruct(schema.Import, schema.Type)
			if s == nil {
				return "xxx", fmt.Errorf("tsDefault struct not found: %s", schema.Dump())
			}
			text = "{}"
		case model.TypeInterface:
			i := schema.LookupInterface(schema.Import, schema.Type)
			if i == nil {
				return "xxx", fmt.Errorf("tsDefault interface not found: %s", schema.Dump())
			}
			text = "null"
		case model.TypeVoid:
			text = "void"
		default:
			return "xxx", fmt.Errorf("tsDefault unknown schema %s", schema.Dump())
		}
	}
	return text, nil
}

// cppDefault returns the default value for a type
func tsDefault(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("tsDefault called with nil node")
	}
	return ToDefaultString(&node.Schema, prefix)
}
