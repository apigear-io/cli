package filterts

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
)

func ToReturnString(schema *objmodel.Schema, prefix string) (string, error) {
	text := ""
	switch schema.KindType {
	case objmodel.TypeString:
		text = "string"
	case objmodel.TypeInt, objmodel.TypeInt32, objmodel.TypeInt64:
		text = "number"
	case objmodel.TypeFloat, objmodel.TypeFloat32, objmodel.TypeFloat64:
		text = "number"
	case objmodel.TypeBool:
		text = "boolean"
	case objmodel.TypeEnum:
		e := schema.LookupEnum(schema.Import, schema.Type)
		if e == nil {
			return "xxx", fmt.Errorf("tsReturn enum not found: %s", schema.Dump())
		}
		text = fmt.Sprintf("%s%s", prefix, e.Name)
	case objmodel.TypeStruct:
		s := schema.LookupStruct(schema.Import, schema.Type)
		if s == nil {
			return "xxx", fmt.Errorf("tsReturn struct not found: %s", schema.Dump())
		}
		text = fmt.Sprintf("%s%s", prefix, s.Name)
	case objmodel.TypeInterface:
		i := schema.LookupInterface(schema.Import, schema.Type)
		if i == nil {
			return "xxx", fmt.Errorf("tsReturn interface not found: %s", schema.Dump())
		}
		text = fmt.Sprintf("%s%s", prefix, i.Name)
	case objmodel.TypeVoid:
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
func tsReturn(prefix string, node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("tsReturn called with nil node")
	}
	return ToReturnString(&node.Schema, prefix)
}
