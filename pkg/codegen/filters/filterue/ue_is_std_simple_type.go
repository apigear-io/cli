package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
)

func CheckIsSimpleType(schema *objmodel.Schema) (bool, error) {
	if schema == nil {
		return false, fmt.Errorf("CheckIsSimpleType schema is nil")
	}

	var result bool
	switch schema.KindType {
	case objmodel.TypeString:
		result = false
	case objmodel.TypeInt:
		result = true
	case objmodel.TypeInt32:
		result = true
	case objmodel.TypeInt64:
		result = true
	case objmodel.TypeFloat:
		result = true
	case objmodel.TypeFloat32:
		result = true
	case objmodel.TypeFloat64:
		result = true
	case objmodel.TypeBool:
		result = true
	case objmodel.TypeEnum:
		result = true
	case objmodel.TypeStruct:
		result = false
	case objmodel.TypeExtern:
		result = false
	case objmodel.TypeInterface:
		result = false
	default:
		return false, fmt.Errorf("unknown schema kind type: %s", schema.KindType)
	}
	if schema.IsArray {
		result = false
	}
	return result, nil
}

func ueIsStdSimpleType(node *objmodel.TypedNode) (bool, error) {
	if node == nil {
		return false, fmt.Errorf("isStdSimpleType node is nil")
	}
	return CheckIsSimpleType(&node.Schema)
}
