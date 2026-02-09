package filterjni

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/apimodel"
)

func jniEmptyReturnString(schema *apimodel.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToType schema is nil")
	}

	var text string
	switch schema.KindType {
	case apimodel.TypeVoid:
		text = ""
	case apimodel.TypeString:
		text = "nullptr"
	case apimodel.TypeInt:
		text = "0"
	case apimodel.TypeInt32:
		text = "0"
	case apimodel.TypeInt64:
		text = "0"
	case apimodel.TypeFloat:
		text = "0"
	case apimodel.TypeFloat32:
		text = "0"
	case apimodel.TypeFloat64:
		text = "0"
	case apimodel.TypeBool:
		text = "false"
	case apimodel.TypeEnum:
		text = "nullptr"
	case apimodel.TypeStruct:
		text = "nullptr"
	case apimodel.TypeInterface:
		text = "nullptr"
	case apimodel.TypeExtern:
		text = "nullptr"
	default:
		return "xxx", fmt.Errorf("ToEnvNameType unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		text = "nullptr"
	}
	return text, nil
}

func jniEmptyReturn(node *apimodel.TypedNode) (string, error) {
	return jniEmptyReturnString(&node.Schema)
}
