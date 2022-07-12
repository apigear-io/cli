package filtergo

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/model"
)

func ToDefaultString(schema *model.Schema, prefix string) string {
	if schema == nil {
		log.Warn("ToDefaultString called with nil schema")
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
			text = fmt.Sprintf("[]%s%s{}", prefix, schema.Type)
		case model.TypeStruct:
			text = fmt.Sprintf("[]%s%s{}", prefix, schema.Type)
		case model.TypeInterface:
			text = fmt.Sprintf("[]*%s%s{}", prefix, schema.Type)
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
			text = fmt.Sprintf("%s%s%s", prefix, sym.Name, member.Name)
		case model.TypeStruct:
			sym := schema.GetStruct()
			text = fmt.Sprintf("%s%s{}", prefix, sym.Name)
		case model.TypeInterface:
			text = "nil"
		case model.TypeNull:
			text = ""
		default:
			log.Fatalf("unknown schema kind type: %s", schema.KindType)
		}
	}
	return text
}

func goDefault(node *model.TypedNode, prefix string) string {
	if node == nil {
		log.Warn("goDefault called with nil node")
		return ""
	}
	return ToDefaultString(&node.Schema, prefix)
}
