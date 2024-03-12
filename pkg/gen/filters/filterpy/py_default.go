package filterpy

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

// ToDefaultString returns the default value for a type
func ToDefaultString(schema *model.Schema, prefix string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("pyDefault schema is nil")
	}
	if schema.Module == nil {
		return "xxx", fmt.Errorf("pyDefault schema module is nil")
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
			text = "False"
		case model.TypeEnum:
			e := schema.LookupEnum(schema.Import, schema.Type)
			if e == nil {
				return "xxx", fmt.Errorf("pyDefault enum not found: %s", schema.Dump())
			}
			name := common.CamelTitleCase(e.Name)
			member := common.SnakeUpperCase(e.Members[0].Name)
			text = fmt.Sprintf("%s%s.%s", prefix, name, member)
		case model.TypeStruct:
			s := schema.LookupStruct(schema.Import, schema.Type)
			if s == nil {
				return "xxx", fmt.Errorf("pyDefault struct not found: %s", schema.Dump())
			}
			ident := common.CamelTitleCase(s.Name)
			text = fmt.Sprintf("%s%s()", prefix, ident)
		case model.TypeInterface:
			i := schema.LookupInterface(schema.Import, schema.Type)
			if i == nil {
				return "xxx", fmt.Errorf("pyDefault interface not found: %s", schema.Dump())
			}
			text = "None"
		case model.TypeVoid:
			text = "None"
		default:
			return "xxx", fmt.Errorf("pyDefault unknown schema %s", schema.Dump())
		}
	}
	if text == "" {
		return "xxx", fmt.Errorf("pyDefault text is empty: %s", schema.Dump())
	}
	return text, nil
}

// cppDefault returns the default value for a type
func pyDefault(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("pyDefault called with nil node")
	}
	return ToDefaultString(&node.Schema, prefix)
}
