package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/apimodel"
)

func CheckIsSimpleType(schema *apimodel.Schema) (bool, error) {
	if schema == nil {
		return false, fmt.Errorf("CheckIsSimpleType schema is nil")
	}

	var result bool
	switch schema.KindType {
	case apimodel.TypeString:
		result = false
	case apimodel.TypeInt:
		result = true
	case apimodel.TypeInt32:
		result = true
	case apimodel.TypeInt64:
		result = true
	case apimodel.TypeFloat:
		result = true
	case apimodel.TypeFloat32:
		result = true
	case apimodel.TypeFloat64:
		result = true
	case apimodel.TypeBool:
		result = true
	case apimodel.TypeEnum:
		result = true
	case apimodel.TypeStruct:
		result = false
	case apimodel.TypeExtern:
		result = false
	case apimodel.TypeInterface:
		result = false
	default:
		return false, fmt.Errorf("unknown schema kind type: %s", schema.KindType)
	}
	if schema.IsArray {
		result = false
	}
	return result, nil
}

func ueIsStdSimpleType(node *apimodel.TypedNode) (bool, error) {
	if node == nil {
		return false, fmt.Errorf("isStdSimpleType node is nil")
	}
	return CheckIsSimpleType(&node.Schema)
}
