package filterjava

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

// ToTestValueString returns the test value string for a given schema.
// We intentionally ignore arrays in order to return the test value of the inner type.
func ToTestValueString(prefix string, schema *model.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("javaTestValue schema is nil")
	}
	var text string
	switch schema.KindType {
	case model.TypeString:
		text = "String(\"xyz\")"
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
		text = ""
	case model.TypeEnum:
		symbol := schema.GetEnum()
		member := symbol.Members[0]
		if len(symbol.Members) > 1 {
			member = symbol.Members[1]
		}
		menberName := common.UpperFirst(member.Name)
		text = fmt.Sprintf("%s%s.%s", prefix, symbol.Name, menberName)
	case model.TypeStruct:
		symbol := schema.GetStruct()
		text = fmt.Sprintf("new %s%s()", prefix, symbol.Name)
	case model.TypeExtern:
		text = "TODO EXTERN"
	case model.TypeInterface:
		symbol := schema.GetInterface()
		text = fmt.Sprintf("%s%s()", prefix, symbol.Name)
	default:
		return "xxx", fmt.Errorf("javaTestValue unknown schema %s", schema.Dump())
	}
	return text, nil
}

func javaTestValue(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("javaTestValue node is nil")
	}
	return ToTestValueString(prefix, &node.Schema)
}
