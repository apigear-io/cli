package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/codegen/filters/common"
	"github.com/apigear-io/cli/pkg/foundation"
	"github.com/apigear-io/cli/pkg/apimodel"
	"github.com/ettle/strcase"
)

func ToDefaultString(prefix string, schema *apimodel.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToDefaultString schema is nil")
	}
	moduleId := strcase.ToPascal(schema.Module.Name)
	if schema.Import != "" {
		moduleId = strcase.ToPascal(schema.Import)
	}
	var text string
	switch schema.KindType {
	case apimodel.TypeString:
		text = "FString()"
	case apimodel.TypeInt, apimodel.TypeInt32:
		text = "0"
	case apimodel.TypeInt64:
		text = "0LL"
	case apimodel.TypeFloat, apimodel.TypeFloat32:
		text = "0.0f"
	case apimodel.TypeFloat64:
		text = "0.0"
	case apimodel.TypeBool:
		text = "false"
	case apimodel.TypeVoid:
		return "xxx", fmt.Errorf("void type not allowed as default value")
	case apimodel.TypeEnum:
		symbol := schema.GetEnum()
		member := symbol.Members[0]
		typename := fmt.Sprintf("%s%s", moduleId, symbol.Name)
		abbreviation := foundation.Abbreviate(typename)
		// upper case first letter
		// TODO: EnumValues: using camel-cases for enum values: strcase.ToCamel(member.Name)
		text = fmt.Sprintf("%sE%s::%s_%s", prefix, typename, abbreviation, common.CamelTitleCase(member.Name))
	case apimodel.TypeStruct:
		symbol := schema.GetStruct()
		text = fmt.Sprintf("%sF%s%s()", prefix, moduleId, symbol.Name)
	case apimodel.TypeExtern:
		xe := parseUeExtern(schema)
		if xe.Default != "" {
			text = xe.Default
		} else {
			if xe.NameSpace != "" {
				prefix = fmt.Sprintf("%s::", xe.NameSpace)
			}
			text = fmt.Sprintf("%s%s()", prefix, xe.Name)
		}
	case apimodel.TypeInterface:
		symbol := schema.GetInterface()
		text = fmt.Sprintf("TScriptInterface<%sI%s%sInterface>()", prefix, moduleId, symbol.Name)
	default:
		return "xxx", fmt.Errorf("ueDefault unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		inner := schema.InnerSchema()
		ret, err := ToReturnString(prefix, &inner)
		if err != nil {
			return "xxx", fmt.Errorf("ToDefaultString inner value error: %s", err)
		}
		text = fmt.Sprintf("TArray<%s>()", ret)
	}
	return text, nil
}

func ueDefault(prefix string, node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ueDefault node is nil")
	}
	return ToDefaultString(prefix, &node.Schema)
}
