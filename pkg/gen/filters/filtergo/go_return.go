package filtergo

import (
	"fmt"
	"objectapi/pkg/log"
	"objectapi/pkg/model"
	"reflect"
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
	default:
		log.Fatalf("unknown schema kind type: %s", schema.KindType)
	}
	if schema.IsArray {
		text = fmt.Sprintf("[]%s", text)
	}
	return text
}

// cast value to TypedNode and deduct the cpp return type
func goReturn(node reflect.Value) (reflect.Value, error) {
	if node.IsNil() {
		log.Debug("goReturn called with nil node")
		return reflect.ValueOf(""), nil
	}
	p, ok := node.Interface().(model.ITypeProvider)
	if !ok {
		return reflect.ValueOf(""), fmt.Errorf("%s can not convert to type node", node.Type())
	}
	t := ToReturnString(p.GetSchema())
	return reflect.ValueOf(t), nil
}
