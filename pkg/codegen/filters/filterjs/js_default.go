package filterjs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
)

// ToDefaultString returns the default value for a type
func ToDefaultString(schema *objmodel.Schema, prefix string) (string, error) {
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
		case objmodel.TypeString:
			text = "\"\""
		case objmodel.TypeInt, objmodel.TypeInt32, objmodel.TypeInt64:
			text = "0"
		case objmodel.TypeFloat, objmodel.TypeFloat32, objmodel.TypeFloat64:
			text = "0.0"
		case objmodel.TypeBool:
			text = "false"
		case objmodel.TypeEnum:
			e := schema.LookupEnum(schema.Import, schema.Type)
			if e == nil {
				return "xxx", fmt.Errorf("jsDefault: enum not found: %s", schema.Dump())
			}
			text = fmt.Sprintf("%s.%s", e.Name, e.Members[0].Name)
		case objmodel.TypeStruct:
			s := schema.LookupStruct(schema.Import, schema.Type)
			if s == nil {
				return "xxx", fmt.Errorf("jsDefault: struct not found: %s", schema.Dump())
			}
			text = fmt.Sprintf("new %s%s()", prefix, s.Name)
		case objmodel.TypeInterface:
			i := schema.LookupInterface(schema.Import, schema.Type)
			if i == nil {
				return "xxx", fmt.Errorf("jsDefault: interface not found: %s", schema.Dump())
			}
			text = "null"
		case objmodel.TypeVoid:
			text = "void"
		default:
			return "xxx", fmt.Errorf("jsDefault unknown schema %s", schema.Dump())
		}
	}
	return text, nil
}

// cppDefault returns the default value for a type
func jsDefault(prefix string, node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("jsDefault called with nil node")
	}
	return ToDefaultString(&node.Schema, prefix)
}
