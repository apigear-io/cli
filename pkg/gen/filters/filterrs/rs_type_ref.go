package filterrs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToTypeRefString(prefix string, schema *model.Schema) (string, error) {
	if schema.IsArray {
		inner := schema.InnerSchema()
		ret, err := ToReturnString(prefix, &inner)
		if err != nil {
			return "xxx", err
		}
		return fmt.Sprintf("&Vec<%s>", ret), nil
	}
	text := ""
	switch schema.Type {
	case "void":
		text = "()"
	case "string":
		text = "&String"
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
		e := schema.LookupEnum(schema.Import, schema.Type)
		if e != nil {
			text = fmt.Sprintf("%s%sEnum", prefix, e.Name)
		}
		s := schema.LookupStruct(schema.Import, schema.Type)
		if s != nil {
			text = fmt.Sprintf("&%s%s", prefix, s.Name)
		}
		i := schema.LookupInterface(schema.Import, schema.Type)
		if i != nil {
			text = fmt.Sprintf("&%s%s", prefix, i.Name)
		}
	}
	return text, nil
}

// cast value to TypedNode and deduct the rs return type
func rsTypeRef(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("rsTypeRef node is nil")
	}
	return ToTypeRefString(prefix, &node.Schema)
}
