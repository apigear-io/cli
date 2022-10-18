package filtergo

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/model"
	"github.com/ettle/strcase"
)

func ToDefaultString(schema *model.Schema, prefix string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToDefaultString schema is nil")
	}
	var text string
	if schema.IsArray {
		switch schema.KindType {
		case model.TypeString:
			text = "[]string{}"
		case model.TypeInt:
			text = "[]int64{}"
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
			return "xxx", fmt.Errorf("unknown schema kind type: %s", schema.KindType)
		}
	} else {
		switch schema.KindType {
		case model.TypeString:
			text = "\"\""
		case model.TypeInt:
			text = "int64(0)"
		case model.TypeFloat:
			text = "float64(0.0)"
		case model.TypeBool:
			text = "false"
		case model.TypeEnum:
			symbol := schema.GetEnum()
			member := symbol.Members[0]
			// upper case first letter

			text = fmt.Sprintf("%s%s%s", prefix, symbol.Name, strcase.ToPascal(member.Name))
		case model.TypeStruct:
			symbol := schema.GetStruct()
			text = fmt.Sprintf("%s%s{}", prefix, symbol.Name)
		case model.TypeInterface:
			text = "nil"
		case model.TypeVoid:
			text = ""
		default:
			return "xxx", fmt.Errorf("unknown schema kind type: %s", schema.KindType)
		}
	}
	return text, nil
}

func goDefault(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		log.Warn().Msg("goDefault called with nil node")
		return "xxx", fmt.Errorf("goDefault node is nil")
	}
	return ToDefaultString(&node.Schema, prefix)
}
