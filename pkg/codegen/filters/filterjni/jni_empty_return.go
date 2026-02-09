package filterjni

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
)

func jniEmptyReturnString(schema *objmodel.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToType schema is nil")
	}

	var text string
	switch schema.KindType {
	case objmodel.TypeVoid:
		text = ""
	case objmodel.TypeString:
		text = "nullptr"
	case objmodel.TypeInt:
		text = "0"
	case objmodel.TypeInt32:
		text = "0"
	case objmodel.TypeInt64:
		text = "0"
	case objmodel.TypeFloat:
		text = "0"
	case objmodel.TypeFloat32:
		text = "0"
	case objmodel.TypeFloat64:
		text = "0"
	case objmodel.TypeBool:
		text = "false"
	case objmodel.TypeEnum:
		text = "nullptr"
	case objmodel.TypeStruct:
		text = "nullptr"
	case objmodel.TypeInterface:
		text = "nullptr"
	case objmodel.TypeExtern:
		text = "nullptr"
	default:
		return "xxx", fmt.Errorf("ToEnvNameType unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		text = "nullptr"
	}
	return text, nil
}

func jniEmptyReturn(node *objmodel.TypedNode) (string, error) {
	return jniEmptyReturnString(&node.Schema)
}
