package filtergo

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/apimodel"
)

// TODO: need to return error case
func ToReturnString(prefix string, schema *apimodel.Schema) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToReturnString schema is nil")
	}
	if schema.IsImported() {
		prefix = fmt.Sprintf("%s.", schema.ShortImportName())
	}
	var text string
	switch schema.KindType {
	case apimodel.TypeString:
		text = "string"
	case apimodel.TypeBytes:
		text = "[]byte"
	case apimodel.TypeInt:
		text = "int32"
	case apimodel.TypeInt32:
		text = "int32"
	case apimodel.TypeInt64:
		text = "int64"
	case apimodel.TypeFloat:
		text = "float32"
	case apimodel.TypeFloat32:
		text = "float32"
	case apimodel.TypeFloat64:
		text = "float64"
	case apimodel.TypeBool:
		text = "bool"
	case apimodel.TypeAny:
		text = "any"
	case apimodel.TypeExtern:
		x := schema.LookupExtern(schema.Import, schema.Type)
		if x == nil {
			return "xxx", fmt.Errorf("goReturn extern not found: %s", schema.Dump())
		}
		xe := parseGoExtern(schema)
		if xe.Import != "" {
			prefix = fmt.Sprintf("%s.", xe.Import)
		}
		text = fmt.Sprintf("%s%s", prefix, xe.Name)
	case apimodel.TypeEnum:
		text = fmt.Sprintf("%s%s", prefix, schema.Type)
	case apimodel.TypeStruct:
		text = fmt.Sprintf("%s%s", prefix, schema.Type)
	case apimodel.TypeInterface:
		text = fmt.Sprintf("%s%s", prefix, schema.Type)
	case apimodel.TypeVoid:
		text = ""
	default:
		return "xxx", fmt.Errorf("goReturn: unknown schema: %s", schema.Dump())
	}
	if schema.IsArray {
		text = fmt.Sprintf("[]%s", text)
	}
	return text, nil
}

func goReturn(prefix string, node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("goReturn node is nil")
	}
	return ToReturnString(prefix, &node.Schema)
}
