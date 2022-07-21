package filtergo

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToReturnString(schema *model.Schema, prefix string) string {
	if schema == nil {
		log.Debug("ToReturnString called with nil schema")
		return ""
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
		log.Fatalf("unknown schema kind type: %s", schema.KindType)
	}
	if schema.IsArray {
		text = fmt.Sprintf("[]%s", text)
	}
	return text
}

// cast value to TypedNode and deduct the cpp return type
func goReturn(node *model.TypedNode, prefix string) string {
	if node == nil {
		log.Warnf("goReturn called with nil node")
		return ""
	}
	return ToReturnString(&node.Schema, prefix)
}
