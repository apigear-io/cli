package filterpy

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

func ToReturnString(schema *model.Schema, prefix string) (string, error) {
	text := ""
	switch schema.KindType {
	case model.TypeString:
		text = "str"
	case model.TypeInt, model.TypeInt32, model.TypeInt64:
		text = "int"
	case model.TypeFloat, model.TypeFloat32, model.TypeFloat64:
		text = "float"
	case model.TypeBool:
		text = "bool"
	case model.TypeEnum:
		e := schema.Module.LookupEnum(schema.Type)
		if e == nil {
			return "xxx", fmt.Errorf("enum %s not found", schema.Type)
		}
		ident := common.CamelTitleCase(e.Name)
		text = fmt.Sprintf("%s%s", prefix, ident)
	case model.TypeStruct:
		s := schema.Module.LookupStruct(schema.Type)
		if s == nil {
			return "xxx", fmt.Errorf("ToReturnString struct %s not found", schema.Type)
		}
		ident := common.CamelTitleCase(s.Name)
		text = fmt.Sprintf("%s%s", prefix, ident)
	case model.TypeInterface:
		i := schema.Module.LookupInterface(schema.Type)
		if i == nil {
			return "xxx", fmt.Errorf("ToReturnString interface %s not found", schema.Type)
		}
		ident := common.CamelTitleCase(i.Name)
		text = fmt.Sprintf("%s%s", prefix, ident)
	case model.TypeVoid:
		text = "None"
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
