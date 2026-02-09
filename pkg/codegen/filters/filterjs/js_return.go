package filterjs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/apimodel"
)

func ToReturnString(schema *apimodel.Schema, prefix string) (string, error) {
	text := ""
	switch schema.KindType {
	case apimodel.TypeString:
		text = ""
	case apimodel.TypeInt, apimodel.TypeInt32, apimodel.TypeInt64:
		text = ""
	case apimodel.TypeFloat, apimodel.TypeFloat32, apimodel.TypeFloat64:
		text = ""
	case apimodel.TypeBool:
		text = ""
	case apimodel.TypeEnum:
		e := schema.LookupEnum(schema.Import, schema.Type)
		if e == nil {
			return "xxx", fmt.Errorf("jsReturn enum not found: %s", schema.Dump())
		}
		text = ""
	case apimodel.TypeStruct:
		s := schema.LookupStruct(schema.Import, schema.Type)
		if s == nil {
			return "xxx", fmt.Errorf("jsReturn struct not found: %s", schema.Dump())
		}
		text = ""
	case apimodel.TypeInterface:
		i := schema.LookupInterface(schema.Import, schema.Type)
		if i == nil {
			return "xxx", fmt.Errorf("jsReturn interface not found: %s", schema.Dump())
		}
		text = ""
	case apimodel.TypeVoid:
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
func jsReturn(prefix string, node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("jsReturn called with nil node")
	}
	return ToReturnString(&node.Schema, prefix)
}
