package filterqt

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToReturnString(prefix string, schema *model.Schema) (string, error) {
	text := ""
	switch schema.Type {
	case "void":
		text = "void"
	case "string":
		text = "QString"
	case "int":
		text = "int"
	case "int32":
		text = "qint32"
	case "int64":
		text = "qint64"
	case "float":
		text = "qreal"
	case "float32":
		text = "float"
	case "float64":
		text = "double"
	case "bool":
		text = "bool"
	default:
		if schema.Module == nil {
			return "xxx", fmt.Errorf("schema.Module is nil")
		}
		e := schema.Module.LookupEnum(schema.Import, schema.Type)
		if e != nil {
			text = fmt.Sprintf("%s%s::%sEnum", prefix, e.Name, e.Name)
		}
		s := schema.Module.LookupStruct(schema.Import, schema.Type)
		if s != nil {
			text = fmt.Sprintf("%s%s", prefix, s.Name)
		}
		i := schema.Module.LookupInterface(schema.Import, schema.Type)
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
