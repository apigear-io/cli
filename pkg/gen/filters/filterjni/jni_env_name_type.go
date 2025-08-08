package filterjni

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToEnvNameType(schema *model.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToType schema is nil")
	}

	var text string
	switch schema.KindType {
	case model.TypeString:
		text = "Object"
	case model.TypeInt:
		text = "Int"
	case model.TypeInt32:
		text = "Int"
	case model.TypeInt64:
		text = "Long"
	case model.TypeFloat:
		text = "Float"
	case model.TypeFloat32:
		text = "Float"
	case model.TypeFloat64:
		text = "Double"
	case model.TypeBool:
		text = "Boolean"
	case model.TypeEnum:
		text = "Object"
	case model.TypeStruct:
		text = "Object"
	case model.TypeExtern:
		text = "Object"
	case model.TypeInterface:
		text = "Object"
	default:
		return "xxx", fmt.Errorf("ToEnvNameType unknown schema %s", schema.Dump())
	}
	return text, nil
}

func jniToEnvNameType(node *model.TypedNode) (string, error) {
	return ToEnvNameType(&node.Schema)
}
