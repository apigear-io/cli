package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"

)


func CheckIsSimpleType(schema *model.Schema) (bool, error) {
	if schema == nil {
		return false, fmt.Errorf("CheckIsSimpleType schema is nil")
	}

	var result bool
	switch schema.KindType {
	case model.TypeString:
		result = false
	case model.TypeInt:
		result = true
	case model.TypeInt32:
		result = true
	case model.TypeInt64:
		result = true
	case model.TypeFloat:
		result = true
	case model.TypeFloat32:
		result = true
	case model.TypeFloat64:
		result = true
	case model.TypeBool:
		result = true
	case model.TypeEnum:
		result = true
	case model.TypeStruct:
		result = false
	case model.TypeInterface:
		result = false
	default:
		return false, fmt.Errorf("unknown schema kind type: %s", schema.KindType)
	}
	if schema.IsArray {
		result = false
	}
	return result, nil
}

func ueIsStdSimpleType(node *model.TypedNode) (bool, error) {
	if node == nil {
		return false, fmt.Errorf("isStdSimpleType node is nil")
	}
	return CheckIsSimpleType(&node.Schema)
}
