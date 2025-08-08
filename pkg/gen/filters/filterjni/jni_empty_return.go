package filterjni

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func jniEmptyReturnString(schema *model.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToType schema is nil")
	}

	var text string
	switch schema.KindType {
	case model.TypeVoid:
		text = ""
	case model.TypeString:
		text = "nullptr"
	case model.TypeInt:
		text = "0"
	case model.TypeInt32:
		text = "0"
	case model.TypeInt64:
		text = "0"
	case model.TypeFloat:
		text = "0"
	case model.TypeFloat32:
		text = "0"
	case model.TypeFloat64:
		text = "0"
	case model.TypeBool:
		text = "false"
	case model.TypeEnum:
		text = "nullptr"
	case model.TypeStruct:
		text = "nullptr"
	case model.TypeInterface:
		text = "nullptr"
	case model.TypeExtern:
		text = "TODO"
	default:
		return "xxx", fmt.Errorf("ToEnvNameType unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		text = "nullptr"
	}
	return text, nil
}

func jniEmptyReturn(node *model.TypedNode) (string, error) {
	return jniEmptyReturnString(&node.Schema)
}
