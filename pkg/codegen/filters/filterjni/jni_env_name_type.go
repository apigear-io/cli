package filterjni

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
)

func ToEnvNameType(schema *objmodel.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToType schema is nil")
	}

	var text string
	switch schema.KindType {
	case objmodel.TypeString:
		text = "Object"
	case objmodel.TypeInt:
		text = "Int"
	case objmodel.TypeInt32:
		text = "Int"
	case objmodel.TypeInt64:
		text = "Long"
	case objmodel.TypeFloat:
		text = "Float"
	case objmodel.TypeFloat32:
		text = "Float"
	case objmodel.TypeFloat64:
		text = "Double"
	case objmodel.TypeBool:
		text = "Boolean"
	case objmodel.TypeEnum:
		text = "Object"
	case objmodel.TypeStruct:
		text = "Object"
	case objmodel.TypeExtern:
		text = "Object"
	case objmodel.TypeInterface:
		text = "Object"
	default:
		return "xxx", fmt.Errorf("ToEnvNameType unknown schema %s", schema.Dump())
	}
	return text, nil
}

func jniToEnvNameType(node *objmodel.TypedNode) (string, error) {
	return ToEnvNameType(&node.Schema)
}
