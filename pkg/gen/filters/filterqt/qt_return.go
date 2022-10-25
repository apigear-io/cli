package filterqt

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToReturnString(prefix string, schema *model.Schema) (string, error) {
	t := schema.Type
	text := ""
	switch t {
	case "void":
		text = "void"
	case "string":
		text = "QString"
	case "int":
		text = "int"
	case "float":
		text = "qreal"
	case "bool":
		text = "bool"
	default:
		if schema.Module == nil {
			return "xxx", fmt.Errorf("schema.Module is nil")
		}
		e := schema.Module.LookupEnum(t)
		if e != nil {
			text = fmt.Sprintf("%s%s::%sEnum", prefix, e.Name, e.Name)
		}
		s := schema.Module.LookupStruct(t)
		if s != nil {
			text = fmt.Sprintf("%s%s", prefix, s.Name)
		}
		i := schema.Module.LookupInterface(t)
		if i != nil {
			text = fmt.Sprintf("%s%s*", prefix, i.Name)
		}
	}
	if schema.IsArray {
		text = fmt.Sprintf("QList<%s>", text)
	}
	return text, nil
}

// cast value to TypedNode and deduct the cpp return type
func qtReturn(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("cppReturn node is nil")
	}
	return ToReturnString(prefix, &node.Schema)
}
