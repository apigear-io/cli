package filtercs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

// ToDefaultString returns the C# default value for a schema. Arrays default to
// an empty List<T>; strings to string.Empty; byte arrays to Array.Empty<byte>().
func ToDefaultString(prefix string, schema *model.Schema) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToDefaultString schema is nil")
	}
	if schema.IsArray {
		inner := schema.InnerSchema()
		ret, err := ToReturnString(prefix, &inner)
		if err != nil {
			return "xxx", fmt.Errorf("csDefault inner value error: %s", err)
		}
		return fmt.Sprintf("new List<%s>()", ret), nil
	}
	text := ""
	switch schema.KindType {
	case model.TypeVoid:
		text = "void"
	case model.TypeBool:
		text = "false"
	case model.TypeInt, model.TypeInt32:
		text = "0"
	case model.TypeInt64:
		text = "0L"
	case model.TypeFloat, model.TypeFloat32:
		text = "0.0f"
	case model.TypeFloat64:
		text = "0.0"
	case model.TypeString:
		text = "string.Empty"
	case model.TypeBytes:
		text = "System.Array.Empty<byte>()"
	case model.TypeAny:
		text = "null"
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
			return "xxx", fmt.Errorf("csDefault enum not found: %s", schema.Dump())
		}
		if schema.Import != "" {
			prefix = fmt.Sprintf("%s.%s.", common.CamelTitleCase(schema.System().Name), common.CamelTitleCase(schema.Import))
		}
		text = fmt.Sprintf("%s%s.%s", prefix, common.CamelTitleCase(e.Name), common.CamelTitleCase(e.Members[0].Name))
	case model.TypeStruct:
		s := schema.LookupStruct(schema.Import, schema.Type)
		if s == nil {
			return "xxx", fmt.Errorf("csDefault struct not found: %s", schema.Dump())
		}
		if schema.Import != "" {
			prefix = fmt.Sprintf("%s.%s.", common.CamelTitleCase(schema.System().Name), common.CamelTitleCase(schema.Import))
		}
		text = fmt.Sprintf("new %s%s()", prefix, common.CamelTitleCase(s.Name))
	case model.TypeInterface:
		text = "null"
	default:
		return "xxx", fmt.Errorf("csDefault unknown schema %s", schema.Dump())
	}
	return text, nil
}

func csDefault(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("csDefault node is nil")
	}
	return ToDefaultString(prefix, &node.Schema)
}
