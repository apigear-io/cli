package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/model"
	"github.com/ettle/strcase"
)

// ToTestValueString returns the test value string for a given schema.
// We intentionally ignore arrays in order to return the test value of the inner type.
func ToTestValueString(prefix string, schema *model.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToDefaultString schema is nil")
	}
	moduleId := strcase.ToPascal(schema.Module.Name)
	if schema.Import != "" {
		moduleId = strcase.ToPascal(schema.Import)
	}
	var text string
	switch schema.KindType {
	case model.TypeString:
		text = "FString(\"xyz\")"
	case model.TypeInt, model.TypeInt32:
		text = "1"
	case model.TypeInt64:
		text = "1LL"
	case model.TypeFloat, model.TypeFloat32:
		text = "1.0f"
	case model.TypeFloat64:
		text = "1.0"
	case model.TypeBool:
		text = "true"
	case model.TypeVoid:
		return ToDefaultString(prefix, schema)
	case model.TypeEnum:
		symbol := schema.GetEnum()
		member := symbol.Members[0]
		if len(symbol.Members) > 1 {
			member = symbol.Members[1]
		}
		typename := fmt.Sprintf("%s%s", moduleId, symbol.Name)
		abbreviation := helper.Abbreviate(typename)
		// upper case first letter
		// TODO: EnumValues: using camel-cases for enum values: strcase.ToCamel(member.Name)
		text = fmt.Sprintf("%sE%s::%s_%s", prefix, typename, abbreviation, strcase.ToCase(member.Name, strcase.UpperCase, '\x00'))
	case model.TypeStruct:
		symbol := schema.GetStruct()
		text = fmt.Sprintf("%sF%s%s()", prefix, moduleId, symbol.Name)
	case model.TypeExtern:
		xe := parseUeExtern(schema)
		if xe.Default != "" {
			text = xe.Default
		} else {
			if xe.NameSpace != "" {
				prefix = fmt.Sprintf("%s::", xe.NameSpace)
			}
			text = fmt.Sprintf("%s%s()", prefix, xe.Name)
		}
	case model.TypeInterface:
		symbol := schema.GetInterface()
		text = fmt.Sprintf("%sF%s%s()", prefix, moduleId, symbol.Name)
	default:
		return "xxx", fmt.Errorf("ueDefault unknown schema %s", schema.Dump())
	}
	return text, nil
}

func ueTestValue(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ueDefault node is nil")
	}
	return ToTestValueString(prefix, &node.Schema)
}
