package filterjava

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToElementTypeString(prefix string, schema *model.Schema) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToReturnString schema is nil")
	}
	var text string
	switch schema.KindType {
	case model.TypeString:
		text = "String"
	case model.TypeInt:
		text = "int"
	case model.TypeInt32:
		text = "int"
	case model.TypeInt64:
		text = "long"
	case model.TypeFloat:
		text = "float"
	case model.TypeFloat32:
		text = "float"
	case model.TypeFloat64:
		text = "double"
	case model.TypeBool:
		text = "boolean"
	case model.TypeEnum:
		symbol := schema.GetEnum()
		text = fmt.Sprintf("%s%s", prefix, symbol.Name)
	case model.TypeStruct:
		symbol := schema.GetStruct()
		text = fmt.Sprintf("%s%s", prefix, symbol.Name)
	case model.TypeExtern:
		xe := parseJavaExtern(schema)
		text = fmt.Sprintf("%s%s", prefix, xe.Name)
	case model.TypeInterface:
		symbol := schema.GetInterface()
		text = fmt.Sprintf("%s%s", prefix, symbol.Name)
	default:
		return "xxx", fmt.Errorf("javaReturn unknown schema %s", schema.Dump())
	}
	return text, nil
}

func javaElementType(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("javaReturn node is nil")
	}
	return ToElementTypeString(prefix, &node.Schema)
}
