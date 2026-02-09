package filterts

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/apimodel"
)

func ToReturnString(schema *apimodel.Schema, prefix string) (string, error) {
	text := ""
	switch schema.KindType {
	case apimodel.TypeString:
		text = "string"
	case apimodel.TypeInt, apimodel.TypeInt32, apimodel.TypeInt64:
		text = "number"
	case apimodel.TypeFloat, apimodel.TypeFloat32, apimodel.TypeFloat64:
		text = "number"
	case apimodel.TypeBool:
		text = "boolean"
	case apimodel.TypeEnum:
		e := schema.LookupEnum(schema.Import, schema.Type)
		if e == nil {
			return "xxx", fmt.Errorf("tsReturn enum not found: %s", schema.Dump())
		}
		text = fmt.Sprintf("%s%s", prefix, e.Name)
	case apimodel.TypeStruct:
		s := schema.LookupStruct(schema.Import, schema.Type)
		if s == nil {
			return "xxx", fmt.Errorf("tsReturn struct not found: %s", schema.Dump())
		}
		text = fmt.Sprintf("%s%s", prefix, s.Name)
	case apimodel.TypeInterface:
		i := schema.LookupInterface(schema.Import, schema.Type)
		if i == nil {
			return "xxx", fmt.Errorf("tsReturn interface not found: %s", schema.Dump())
		}
		text = fmt.Sprintf("%s%s", prefix, i.Name)
	case apimodel.TypeVoid:
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
func tsReturn(prefix string, node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("tsReturn called with nil node")
	}
	return ToReturnString(&node.Schema, prefix)
}
