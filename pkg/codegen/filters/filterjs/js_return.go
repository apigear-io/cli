package filterjs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
)

func ToReturnString(schema *objmodel.Schema, prefix string) (string, error) {
	text := ""
	switch schema.KindType {
	case objmodel.TypeString:
		text = ""
	case objmodel.TypeInt, objmodel.TypeInt32, objmodel.TypeInt64:
		text = ""
	case objmodel.TypeFloat, objmodel.TypeFloat32, objmodel.TypeFloat64:
		text = ""
	case objmodel.TypeBool:
		text = ""
	case objmodel.TypeEnum:
		e := schema.LookupEnum(schema.Import, schema.Type)
		if e == nil {
			return "xxx", fmt.Errorf("jsReturn enum not found: %s", schema.Dump())
		}
		text = ""
	case objmodel.TypeStruct:
		s := schema.LookupStruct(schema.Import, schema.Type)
		if s == nil {
			return "xxx", fmt.Errorf("jsReturn struct not found: %s", schema.Dump())
		}
		text = ""
	case objmodel.TypeInterface:
		i := schema.LookupInterface(schema.Import, schema.Type)
		if i == nil {
			return "xxx", fmt.Errorf("jsReturn interface not found: %s", schema.Dump())
		}
		text = ""
	case objmodel.TypeVoid:
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
func jsReturn(prefix string, node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("jsReturn called with nil node")
	}
	return ToReturnString(&node.Schema, prefix)
}
