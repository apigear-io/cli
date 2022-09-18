package filtergo

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/iancoleman/strcase"
)

func ToDefaultString(schema *model.Schema, prefix string) (string, error) {
	if schema == nil {
		log.Warn("ToDefaultString called with nil schema")
		return "", fmt.Errorf("ToDefaultString schema is nil")
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
			return "", fmt.Errorf("unknown schema kind type: %s", schema.KindType)
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

			text = fmt.Sprintf("%s%s%s", prefix, symbol.Name, strcase.ToCamel(member.Name))
		case model.TypeStruct:
			symbol := schema.GetStruct()
			text = fmt.Sprintf("%s%s{}", prefix, symbol.Name)
		case model.TypeInterface:
			text = "nil"
		case model.TypeNull:
			text = ""
		default:
			return "", fmt.Errorf("unknown schema kind type: %s", schema.KindType)
		}
	}
	return text, nil
}

func goDefault(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		log.Warn("goDefault called with nil node")
		return "", fmt.Errorf("goDefault node is nil")
	}
	return ToDefaultString(&node.Schema, prefix)
}
