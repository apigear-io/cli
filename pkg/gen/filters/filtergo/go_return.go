package filtergo

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

// TODO: need to return error case
func ToReturnString(prefix string, schema *model.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToReturnString schema is nil")
	}
	var text string
	switch schema.KindType {
	case model.TypeString:
		text = "string"
	case model.TypeInt:
		text = "int"
	case model.TypeFloat:
		text = "float64"
	case model.TypeBool:
		text = "bool"
	case model.TypeEnum:
		text = fmt.Sprintf("%s%s", prefix, schema.Type)
	case model.TypeStruct:
		text = fmt.Sprintf("%s%s", prefix, schema.Type)
	case model.TypeInterface:
		text = fmt.Sprintf("*%s%s", prefix, schema.Type)
	case model.TypeNull:
		text = ""
	default:
		return "", fmt.Errorf("unknown schema kind type: %s", schema.KindType)
	}
	if schema.IsArray {
		text = fmt.Sprintf("[]%s", text)
	}
	return text, nil
}

func goReturn(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "", fmt.Errorf("goReturn node is nil")
	}
	return ToReturnString(prefix, &node.Schema)
}
