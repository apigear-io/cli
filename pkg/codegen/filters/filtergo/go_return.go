package filtergo

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
)

// TODO: need to return error case
func ToReturnString(prefix string, schema *objmodel.Schema) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToReturnString schema is nil")
	}
	if schema.IsImported() {
		prefix = fmt.Sprintf("%s.", schema.ShortImportName())
	}
	var text string
	switch schema.KindType {
	case objmodel.TypeString:
		text = "string"
	case objmodel.TypeBytes:
		text = "[]byte"
	case objmodel.TypeInt:
		text = "int32"
	case objmodel.TypeInt32:
		text = "int32"
	case objmodel.TypeInt64:
		text = "int64"
	case objmodel.TypeFloat:
		text = "float32"
	case objmodel.TypeFloat32:
		text = "float32"
	case objmodel.TypeFloat64:
		text = "float64"
	case objmodel.TypeBool:
		text = "bool"
	case objmodel.TypeAny:
		text = "any"
	case objmodel.TypeExtern:
		x := schema.LookupExtern(schema.Import, schema.Type)
		if x == nil {
			return "xxx", fmt.Errorf("goReturn extern not found: %s", schema.Dump())
		}
		xe := parseGoExtern(schema)
		if xe.Import != "" {
			prefix = fmt.Sprintf("%s.", xe.Import)
		}
		text = fmt.Sprintf("%s%s", prefix, xe.Name)
	case objmodel.TypeEnum:
		text = fmt.Sprintf("%s%s", prefix, schema.Type)
	case objmodel.TypeStruct:
		text = fmt.Sprintf("%s%s", prefix, schema.Type)
	case objmodel.TypeInterface:
		text = fmt.Sprintf("%s%s", prefix, schema.Type)
	case objmodel.TypeVoid:
		text = ""
	default:
		return "xxx", fmt.Errorf("goReturn: unknown schema: %s", schema.Dump())
	}
	if schema.IsArray {
		text = fmt.Sprintf("[]%s", text)
	}
	return text, nil
}

func goReturn(prefix string, node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("goReturn node is nil")
	}
	return ToReturnString(prefix, &node.Schema)
}
