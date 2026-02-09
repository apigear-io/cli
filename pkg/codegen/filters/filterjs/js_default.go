package filterjs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/apimodel"
)

// ToDefaultString returns the default value for a type
func ToDefaultString(schema *apimodel.Schema, prefix string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToDefaultString schema is nil")
	}
	if schema.Module == nil {
		return "xxx", fmt.Errorf("ToDefaultString schema module is nil")
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
				return "xxx", fmt.Errorf("jsDefault: enum not found: %s", schema.Dump())
			}
			text = fmt.Sprintf("%s.%s", e.Name, e.Members[0].Name)
		case apimodel.TypeStruct:
			s := schema.LookupStruct(schema.Import, schema.Type)
			if s == nil {
				return "xxx", fmt.Errorf("jsDefault: struct not found: %s", schema.Dump())
			}
			text = fmt.Sprintf("new %s%s()", prefix, s.Name)
		case apimodel.TypeInterface:
			i := schema.LookupInterface(schema.Import, schema.Type)
			if i == nil {
				return "xxx", fmt.Errorf("jsDefault: interface not found: %s", schema.Dump())
			}
			text = "null"
		case apimodel.TypeVoid:
			text = "void"
		default:
			return "xxx", fmt.Errorf("jsDefault unknown schema %s", schema.Dump())
		}
	}
	return text, nil
}

// cppDefault returns the default value for a type
func jsDefault(prefix string, node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("jsDefault called with nil node")
	}
	return ToDefaultString(&node.Schema, prefix)
}
