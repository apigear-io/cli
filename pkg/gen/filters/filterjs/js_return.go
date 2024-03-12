package filterjs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToReturnString(schema *model.Schema, prefix string) (string, error) {
	text := ""
	switch schema.KindType {
	case model.TypeString:
		text = ""
	case model.TypeInt, model.TypeInt32, model.TypeInt64:
		text = ""
	case model.TypeFloat, model.TypeFloat32, model.TypeFloat64:
		text = ""
	case model.TypeBool:
		text = ""
	case model.TypeEnum:
		e := schema.LookupEnum(schema.Import, schema.Type)
		if e == nil {
			return "xxx", fmt.Errorf("jsReturn enum not found: %s", schema.Dump())
		}
		text = ""
	case model.TypeStruct:
		s := schema.LookupStruct(schema.Import, schema.Type)
		if s == nil {
			return "xxx", fmt.Errorf("jsReturn struct not found: %s", schema.Dump())
		}
		text = ""
	case model.TypeInterface:
		i := schema.LookupInterface(schema.Import, schema.Type)
		if i == nil {
			return "xxx", fmt.Errorf("jsReturn interface not found: %s", schema.Dump())
		}
		text = ""
	case model.TypeVoid:
		text = ""
	default:
		return "xxx", fmt.Errorf("jsReturn unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		text = ""
	}
	return text, nil
}

// cast value to TypedNode and deduct the cpp return type
func jsReturn(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("jsReturn called with nil node")
	}
	return ToReturnString(&node.Schema, prefix)
}
