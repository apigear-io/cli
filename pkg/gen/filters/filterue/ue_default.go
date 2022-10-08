package filterue

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/iancoleman/strcase"
)

func ToDefaultString(prefix string, schema *model.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToDefaultString schema is nil")
	}
	moduleId := strcase.ToCamel(schema.Module.Name)
	var text string
	if schema.IsArray {
		switch schema.KindType {
		case model.TypeString:
			text = "TArray<FString>()"
		case model.TypeInt:
			text = "TArray<int32>()"
		case model.TypeFloat:
			text = "TArray<float>()"
		case model.TypeBool:
			text = "TArray<bool>()"
		case model.TypeEnum:
			text = fmt.Sprintf("TArray<%sE%s%s>()", prefix, moduleId, schema.Type)
		case model.TypeStruct:
			text = fmt.Sprintf("TArray<%sF%s%s>()", prefix, moduleId, schema.Type)
		case model.TypeInterface:
			text = fmt.Sprintf("TArray<%sF%s%s>()", prefix, moduleId, schema.Type)
		default:
			return "", fmt.Errorf("unknown schema kind type: %s", schema.KindType)
		}
	} else {
		switch schema.KindType {
		case model.TypeString:
			text = "FString()"
		case model.TypeInt:
			text = "0"
		case model.TypeFloat:
			text = "0.0f"
		case model.TypeBool:
			text = "false"
		case model.TypeEnum:
			symbol := schema.GetEnum()
			member := symbol.Members[0]
			// upper case first letter
			// TODO: EnumValues: using camel-cases for enum values: strcase.ToCamel(member.Name)
			text = fmt.Sprintf("%sE%s%s::%s", prefix, moduleId, symbol.Name, strings.ToUpper(member.Name))
		case model.TypeStruct:
			symbol := schema.GetStruct()
			text = fmt.Sprintf("%sF%s%s()", prefix, moduleId, symbol.Name)
		case model.TypeInterface:
			symbol := schema.GetInterface()
			text = fmt.Sprintf("%sF%s%s()", prefix, moduleId, symbol.Name)
		case model.TypeNull:
			text = ""
		default:
			return "", fmt.Errorf("unknown schema kind type: %s", schema.KindType)
		}
	}
	return text, nil
}

func ueDefault(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "", fmt.Errorf("goDefault node is nil")
	}
	return ToDefaultString(prefix, &node.Schema)
}