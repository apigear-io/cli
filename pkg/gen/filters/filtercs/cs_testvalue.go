package filtercs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

// ToTestValueString returns a non-default test value for a given schema.
// We intentionally ignore arrays in order to return the test value of the
// inner type (it can be placed into a collection by the template).
func ToTestValueString(prefix string, schema *model.Schema) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("csTestValue schema is nil")
	}
	var text string
	switch schema.KindType {
	case model.TypeString:
		text = "\"xyz\""
	case model.TypeInt, model.TypeInt32:
		text = "1"
	case model.TypeInt64:
		text = "1L"
	case model.TypeFloat, model.TypeFloat32:
		text = "1.0f"
	case model.TypeFloat64:
		text = "1.0"
	case model.TypeBool:
		text = "true"
	case model.TypeBytes:
		text = "new byte[] { 1 }"
	case model.TypeAny:
		text = "1"
	case model.TypeVoid:
		text = ""
	case model.TypeExtern:
		xe := parseCsExtern(schema)
		if xe.Default != "" {
			text = xe.Default
		} else {
			ns := ""
			if xe.Namespace != "" {
				ns = fmt.Sprintf("%s.", xe.Namespace)
			}
			text = fmt.Sprintf("new %s%s()", ns, xe.Name)
		}
	case model.TypeEnum:
		e := schema.LookupEnum(schema.Import, schema.Type)
		if e == nil {
			return "xxx", fmt.Errorf("csTestValue enum not found: %s", schema.Dump())
		}
		member := e.Members[0].Name
		if len(e.Members) > 1 {
			member = e.Members[1].Name
		}
		if schema.Import != "" {
			prefix = fmt.Sprintf("%s.%s.", common.CamelTitleCase(schema.System().Name), common.CamelTitleCase(schema.Import))
		}
		text = fmt.Sprintf("%s%s.%s", prefix, common.CamelTitleCase(e.Name), common.CamelTitleCase(member))
	case model.TypeStruct:
		s := schema.LookupStruct(schema.Import, schema.Type)
		if s == nil {
			return "xxx", fmt.Errorf("csTestValue struct not found: %s", schema.Dump())
		}
		if schema.Import != "" {
			prefix = fmt.Sprintf("%s.%s.", common.CamelTitleCase(schema.System().Name), common.CamelTitleCase(schema.Import))
		}
		text = fmt.Sprintf("new %s%s()", prefix, common.CamelTitleCase(s.Name))
	case model.TypeInterface:
		i := schema.LookupInterface(schema.Import, schema.Type)
		if i == nil {
			return "xxx", fmt.Errorf("csTestValue interface not found: %s", schema.Dump())
		}
		if schema.Import != "" {
			prefix = fmt.Sprintf("%s.%s.", common.CamelTitleCase(schema.System().Name), common.CamelTitleCase(schema.Import))
		}
		text = fmt.Sprintf("new %s%s()", prefix, common.CamelTitleCase(i.Name))
	default:
		return "xxx", fmt.Errorf("csTestValue unknown schema %s", schema.Dump())
	}
	return text, nil
}

func csTestValue(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("csTestValue node is nil")
	}
	return ToTestValueString(prefix, &node.Schema)
}
