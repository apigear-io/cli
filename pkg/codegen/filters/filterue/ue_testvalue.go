package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/codegen/filters/common"
	"github.com/apigear-io/cli/pkg/foundation"
	"github.com/apigear-io/cli/pkg/objmodel"
	"github.com/ettle/strcase"
)

// ToTestValueString returns the test value string for a given schema.
// We intentionally ignore arrays in order to return the test value of the inner type.
func ToTestValueString(prefix string, schema *objmodel.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToDefaultString schema is nil")
	}
	moduleId := strcase.ToPascal(schema.Module.Name)
	if schema.Import != "" {
		moduleId = strcase.ToPascal(schema.Import)
	}
	var text string
	switch schema.KindType {
	case objmodel.TypeString:
		text = "FString(\"xyz\")"
	case objmodel.TypeInt, objmodel.TypeInt32:
		text = "1"
	case objmodel.TypeInt64:
		text = "1LL"
	case objmodel.TypeFloat, objmodel.TypeFloat32:
		text = "1.0f"
	case objmodel.TypeFloat64:
		text = "1.0"
	case objmodel.TypeBool:
		text = "true"
	case objmodel.TypeVoid:
		return ToDefaultString(prefix, schema)
	case objmodel.TypeEnum:
		symbol := schema.GetEnum()
		member := symbol.Members[0]
		if len(symbol.Members) > 1 {
			member = symbol.Members[1]
		}
		typename := fmt.Sprintf("%s%s", moduleId, symbol.Name)
		abbreviation := foundation.Abbreviate(typename)
		// upper case first letter
		// TODO: EnumValues: using camel-cases for enum values: strcase.ToCamel(member.Name)
		text = fmt.Sprintf("%sE%s::%s_%s", prefix, typename, abbreviation, common.CamelTitleCase(member.Name))
	case objmodel.TypeStruct:
		symbol := schema.GetStruct()
		text = fmt.Sprintf("%sF%s%s()", prefix, moduleId, symbol.Name)
	case objmodel.TypeExtern:
		xe := parseUeExtern(schema)
		if xe.Default != "" {
			text = xe.Default
		} else {
			if xe.NameSpace != "" {
				prefix = fmt.Sprintf("%s::", xe.NameSpace)
			}
			text = fmt.Sprintf("%s%s()", prefix, xe.Name)
		}
	case objmodel.TypeInterface:
		symbol := schema.GetInterface()
		text = fmt.Sprintf("TScriptInterface<%sI%s%sInterface>()", prefix, moduleId, symbol.Name)
	default:
		return "xxx", fmt.Errorf("ueDefault unknown schema %s", schema.Dump())
	}
	return text, nil
}

func ueTestValue(prefix string, node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ueDefault node is nil")
	}
	return ToTestValueString(prefix, &node.Schema)
}
