package filtergo

import (
	"fmt"
	"objectapi/pkg/log"
	"objectapi/pkg/model"
)

func ToReturnString(schema *model.Schema) string {
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
		text = schema.Type
	case model.TypeStruct:
		text = schema.Type
	case model.TypeInterface:
		text = fmt.Sprintf("*%s", schema.Type)
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
func goReturn(p *model.TypedNode) string {
	return ToReturnString(&p.Schema)
}
