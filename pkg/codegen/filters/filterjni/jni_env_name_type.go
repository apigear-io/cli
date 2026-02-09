package filterjni

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/apimodel"
)

func ToEnvNameType(schema *apimodel.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToType schema is nil")
	}

	var text string
	switch schema.KindType {
	case apimodel.TypeString:
		text = "Object"
	case apimodel.TypeInt:
		text = "Int"
	case apimodel.TypeInt32:
		text = "Int"
	case apimodel.TypeInt64:
		text = "Long"
	case apimodel.TypeFloat:
		text = "Float"
	case apimodel.TypeFloat32:
		text = "Float"
	case apimodel.TypeFloat64:
		text = "Double"
	case apimodel.TypeBool:
		text = "Boolean"
	case apimodel.TypeEnum:
		text = "Object"
	case apimodel.TypeStruct:
		text = "Object"
	case apimodel.TypeExtern:
		text = "Object"
	case apimodel.TypeInterface:
		text = "Object"
	default:
		return "xxx", fmt.Errorf("ToEnvNameType unknown schema %s", schema.Dump())
	}
	return text, nil
}

func jniToEnvNameType(node *apimodel.TypedNode) (string, error) {
	return ToEnvNameType(&node.Schema)
}
