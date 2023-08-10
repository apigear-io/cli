package filterrust

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToReturnString(prefix string, schema *model.Schema) (string, error) {
	t := schema.Type
	text := ""
	switch t {
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
		e := schema.Module.LookupEnum(t)
		if e != nil {
			text = fmt.Sprintf("%s%sEnum", prefix, e.Name)
		}
		s := schema.Module.LookupStruct(t)
		if s != nil {
			text = fmt.Sprintf("%s%s", prefix, s.Name)
		}
		i := schema.Module.LookupInterface(t)
		if i != nil {
			text = fmt.Sprintf("&%s%s", prefix, i.Name)
		}
	}
	if schema.IsArray {
		return fmt.Sprintf("Vec<%s>", text), nil
	}
	return text, nil
}

// cast value to TypedNode and deduct the rust return type
func rustReturn(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("rustReturn node is nil")
	}
	return ToReturnString(prefix, &node.Schema)
}
