package filtergo

import (
	"fmt"
	"objectapi/pkg/log"
	"objectapi/pkg/model"
	"reflect"
)

func ToDefaultString(schema *model.Schema) string {
	if schema == nil {
		return ""
	}
	var text string
	if schema.IsArray {
		switch schema.KindType {
		case model.TypeString:
			text = "[]string{}"
		case model.TypeInt:
			text = "[]int{}"
		case model.TypeFloat:
			text = "[]float64{}"
		case model.TypeBool:
			text = "[]bool{}"
		case model.TypeEnum:
			text = fmt.Sprintf("[]%s{}", schema.Type)
		case model.TypeStruct:
			text = fmt.Sprintf("[]%s{}", schema.Type)
		case model.TypeInterface:
			text = fmt.Sprintf("[]*%s{}", schema.Type)
		default:
			log.Fatalf("unknown schema kind type: %s", schema.KindType)
		}
	} else {
		switch schema.KindType {
		case model.TypeString:
			text = "\"\""
		case model.TypeInt:
			text = "0"
		case model.TypeFloat:
			text = "0.0"
		case model.TypeBool:
			text = "false"
		case model.TypeEnum:
			sym := schema.GetEnum()
			member := sym.Members[0]
			text = fmt.Sprintf("%s%s", sym.Name, member.Name)
		case model.TypeStruct:
			sym := schema.GetStruct()
			text = fmt.Sprintf("%s{}", sym.Name)
		case model.TypeInterface:
			text = "nil"
		default:
			log.Fatalf("unknown schema kind type: %s", schema.KindType)
		}
	}
	return text
}

func goDefault(node reflect.Value) (reflect.Value, error) {
	if node.IsNil() {
		log.Debug("goDefault called with nil node")
		return reflect.ValueOf(""), nil
	}
	p, ok := node.Interface().(model.ITypeProvider)
	if !ok {
		log.Debugf("goDefault called with non-typeprovider node: %v", node)
		return reflect.ValueOf(""), nil
	}
	t := ToDefaultString(p.GetSchema())
	return reflect.ValueOf(t), nil
}
