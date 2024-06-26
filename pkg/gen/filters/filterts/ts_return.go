package filterts

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToReturnString(schema *model.Schema, prefix string) (string, error) {
	text := ""
	switch schema.KindType {
	case model.TypeString:
		text = "string"
	case model.TypeInt, model.TypeInt32, model.TypeInt64:
		text = "number"
	case model.TypeFloat, model.TypeFloat32, model.TypeFloat64:
		text = "number"
	case model.TypeBool:
		text = "boolean"
	case model.TypeEnum:
		e := schema.LookupEnum(schema.Import, schema.Type)
		if e == nil {
			return "xxx", fmt.Errorf("tsReturn enum not found: %s", schema.Dump())
		}
		text = fmt.Sprintf("%s%s", prefix, e.Name)
	case model.TypeStruct:
		s := schema.LookupStruct(schema.Import, schema.Type)
		if s == nil {
			return "xxx", fmt.Errorf("tsReturn struct not found: %s", schema.Dump())
		}
		text = fmt.Sprintf("%s%s", prefix, s.Name)
	case model.TypeInterface:
		i := schema.LookupInterface(schema.Import, schema.Type)
		if i == nil {
			return "xxx", fmt.Errorf("tsReturn interface not found: %s", schema.Dump())
		}
		text = fmt.Sprintf("%s%s", prefix, i.Name)
	case model.TypeVoid:
		text = "void"
	default:
		return "xxx", fmt.Errorf("tsReturn unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		text = fmt.Sprintf("%s[]", text)
	}
	return text, nil
}

// cast value to TypedNode and deduct the cpp return type
func tsReturn(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("tsReturn called with nil node")
	}
	return ToReturnString(&node.Schema, prefix)
}
