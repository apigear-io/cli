package filterts

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

// ToDefaultString returns the default value for a type
func ToDefaultString(schema *model.Schema, prefix string) (string, error) {
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
		case model.TypeString:
			text = "\"\""
		case model.TypeInt:
			text = "0"
		case model.TypeFloat:
			text = "0.0"
		case model.TypeBool:
			text = "false"
		case model.TypeEnum:
			e := schema.Module.LookupEnum(schema.Type)
			if e == nil {
				return "xxx", fmt.Errorf("ToDefaultString enum %s not found", schema.Type)
			}
			text = fmt.Sprintf("%s.%s", e.Name, e.Members[0].Name)
		case model.TypeStruct:
			s := schema.Module.LookupStruct(schema.Type)
			if s == nil {
				return "xxx", fmt.Errorf("ToDefaultString struct %s not found", schema.Type)
			}
			text = fmt.Sprintf("new %s%s()", prefix, s.Name)
		case model.TypeInterface:
			i := schema.Module.LookupInterface(schema.Type)
			if i == nil {
				return "xxx", fmt.Errorf("ToDefaultString interface %s not found", schema.Type)
			}
			text = "null"
		case model.TypeVoid:
			text = "void"
		default:
			return "xxx", fmt.Errorf("unknown schema kind type: %s", schema.KindType)
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
