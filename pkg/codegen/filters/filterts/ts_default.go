package filterts

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
)

// ToDefaultString returns the default value for a type
func ToDefaultString(schema *objmodel.Schema, prefix string) (string, error) {
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
				return "xxx", fmt.Errorf("tsDefault enum not found: %s", schema.Dump())
			}
			text = fmt.Sprintf("%s%s.%s", prefix, e.Name, e.Members[0].Name)
		case objmodel.TypeStruct:
			s := schema.LookupStruct(schema.Import, schema.Type)
			if s == nil {
				return "xxx", fmt.Errorf("tsDefault struct not found: %s", schema.Dump())
			}
			text = "{}"
		case objmodel.TypeInterface:
			i := schema.LookupInterface(schema.Import, schema.Type)
			if i == nil {
				return "xxx", fmt.Errorf("tsDefault interface not found: %s", schema.Dump())
			}
			text = "null"
		case objmodel.TypeVoid:
			text = "void"
		default:
			return "xxx", fmt.Errorf("tsDefault unknown schema %s", schema.Dump())
		}
	}
	return text, nil
}

// cppDefault returns the default value for a type
func tsDefault(prefix string, node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("tsDefault called with nil node")
	}
	return ToDefaultString(&node.Schema, prefix)
}
