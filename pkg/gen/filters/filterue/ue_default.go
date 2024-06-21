package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/model"
	"github.com/ettle/strcase"
)

func ToDefaultString(prefix string, schema *model.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToDefaultString schema is nil")
	}
	moduleId := strcase.ToPascal(schema.Module.Name)
	if schema.Import != "" {
		moduleId = strcase.ToPascal(schema.Import)
	}
	var text string
	if schema.IsArray {
		switch schema.KindType {
		case model.TypeString:
			text = "TArray<FString>()"
		case model.TypeInt:
			text = "TArray<int32>()"
		case model.TypeInt32:
			text = "TArray<int32>()"
		case model.TypeInt64:
			text = "TArray<int64>()"
		case model.TypeFloat:
			text = "TArray<float>()"
		case model.TypeFloat32:
			text = "TArray<float>()"
		case model.TypeFloat64:
			text = "TArray<double>()"
		case model.TypeBool:
			text = "TArray<bool>()"
		case model.TypeEnum:
			text = fmt.Sprintf("TArray<%sE%s%s>()", prefix, moduleId, schema.Type)
		case model.TypeStruct:
			text = fmt.Sprintf("TArray<%sF%s%s>()", prefix, moduleId, schema.Type)
		case model.TypeInterface:
			text = fmt.Sprintf("TArray<%sF%s%s>()", prefix, moduleId, schema.Type)
		default:
			return "xxx", fmt.Errorf("ueDefault unknown schema %s", schema.Dump())
		}
	} else {
		switch schema.KindType {
		case model.TypeString:
			text = "FString()"
		case model.TypeInt, model.TypeInt32:
			text = "0"
		case model.TypeInt64:
			text = "0LL"
		case model.TypeFloat, model.TypeFloat32:
			text = "0.0f"
		case model.TypeFloat64:
			text = "0.0"
		case model.TypeBool:
			text = "false"
		case model.TypeVoid:
			return "xxx", fmt.Errorf("void type not allowed as default value")
		case model.TypeEnum:
			symbol := schema.GetEnum()
			member := symbol.Members[0]
			typename := fmt.Sprintf("%s%s", moduleId, symbol.Name)
			abbreviation := helper.Abbreviate(typename)
			// upper case first letter
			// TODO: EnumValues: using camel-cases for enum values: strcase.ToCamel(member.Name)
			text = fmt.Sprintf("%sE%s::%s_%s", prefix, typename, abbreviation, strcase.ToCase(member.Name, strcase.UpperCase, '\x00'))
		case model.TypeStruct:
			symbol := schema.GetStruct()
			text = fmt.Sprintf("%sF%s%s()", prefix, moduleId, symbol.Name)
		case model.TypeExtern:
			symbol := schema.GetExtern()
			text = fmt.Sprintf("%s()", ueExtern(symbol).Name)
		case model.TypeInterface:
			symbol := schema.GetInterface()
			text = fmt.Sprintf("%sF%s%s()", prefix, moduleId, symbol.Name)
		default:
			return "xxx", fmt.Errorf("ueDefault unknown schema %s", schema.Dump())
		}
	}
	return text, nil
}

func ueDefault(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ueDefault node is nil")
	}
	return ToDefaultString(prefix, &node.Schema)
}
