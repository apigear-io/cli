package filterpy

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToReturnString(schema *model.Schema, prefix string) (string, error) {
	text := ""
	switch schema.KindType {
	case model.TypeString:
		text = "str"
	case model.TypeInt:
		text = "int"
	case model.TypeFloat:
		text = "float"
	case model.TypeBool:
		text = "bool"
	case model.TypeEnum:
		e := schema.Module.LookupEnum(schema.Type)
		if e == nil {
			return "xxx", fmt.Errorf("enum %s not found", schema.Type)
		}
		text = fmt.Sprintf("%s%s", prefix, e.Name)
	case model.TypeStruct:
		s := schema.Module.LookupStruct(schema.Type)
		if s == nil {
			return "xxx", fmt.Errorf("ToReturnString struct %s not found", schema.Type)
		}
		text = fmt.Sprintf("%s%s", prefix, s.Name)
	case model.TypeInterface:
		i := schema.Module.LookupInterface(schema.Type)
		if i == nil {
			return "xxx", fmt.Errorf("ToReturnString interface %s not found", schema.Type)
		}
		text = fmt.Sprintf("%s%s", prefix, i.Name)
	case model.TypeVoid:
		text = "None"
	default:
		return "xxx", fmt.Errorf("unknown schema kind type: %s", schema.KindType)
	}
	if schema.IsArray {
		text = fmt.Sprintf("list[%s]", text)
	}
	return text, nil
}

// cast value to TypedNode and deduct the cpp return type
func pyReturn(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("called with nil node")
	}
	return ToReturnString(&node.Schema, prefix)
}
