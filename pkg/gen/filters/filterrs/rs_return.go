package filterrs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToReturnString(prefixComplexType string, schema *model.Schema) (string, error) {
	text := ""
	switch schema.Type {
	case "void":
		text = "()"
	case "string":
		text = "String"
	case "int":
		text = "i32"
	case "int32":
		text = "i32"
	case "int64":
		text = "i64"
	case "float":
		text = "f32"
	case "float32":
		text = "f32"
	case "float64":
		text = "f64"
	case "bool":
		text = "bool"
	default:
		if schema.Module == nil {
			return "xxx", fmt.Errorf("schema.Module is nil")
		}
		e := schema.LookupEnum(schema.Import, schema.Type)
		if e != nil {
			text = fmt.Sprintf("%s%sEnum", prefixComplexType, e.Name)
		}
		s := schema.LookupStruct(schema.Import, schema.Type)
		if s != nil {
			text = fmt.Sprintf("%s%s", prefixComplexType, s.Name)
		}
		i := schema.LookupInterface(schema.Import, schema.Type)
		if i != nil {
			text = fmt.Sprintf("&%s%s", prefixComplexType, i.Name)
		}
	}
	if schema.IsArray {
		return fmt.Sprintf("Vec<%s>", text), nil
	}
	return text, nil
}

// cast value to TypedNode and deduct the rs return type
func rsReturn(prefixComplexType string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("rsReturn node is nil")
	}
	return ToReturnString(prefixComplexType, &node.Schema)
}
