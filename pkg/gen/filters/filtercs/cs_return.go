package filtercs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

// ToReturnString deducts the C# type for a schema. Arrays map to the
// generic System.Collections.Generic.List<T>.
func ToReturnString(prefix string, schema *model.Schema) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToReturnString schema is nil")
	}
	text := ""
	switch schema.KindType {
	case model.TypeVoid:
		text = "void"
	case model.TypeBool:
		text = "bool"
	case model.TypeInt, model.TypeInt32:
		text = "int"
	case model.TypeInt64:
		text = "long"
	case model.TypeFloat, model.TypeFloat32:
		text = "float"
	case model.TypeFloat64:
		text = "double"
	case model.TypeString:
		text = "string"
	case model.TypeBytes:
		text = "byte[]"
	case model.TypeAny:
		text = "object"
	case model.TypeExtern:
		xe := parseCsExtern(schema)
		if xe.Namespace != "" {
			prefix = fmt.Sprintf("%s.", xe.Namespace)
		} else {
			// Externs carry their own fully-qualified name; do not prefix them.
			prefix = ""
		}
		text = fmt.Sprintf("%s%s", prefix, xe.Name)
	case model.TypeEnum:
		e := schema.LookupEnum(schema.Import, schema.Type)
		if schema.Import != "" {
			prefix = fmt.Sprintf("%s.%s.", common.CamelTitleCase(schema.System().Name), common.CamelTitleCase(schema.Import))
		}
		if e != nil {
			text = fmt.Sprintf("%s%s", prefix, common.CamelTitleCase(e.Name))
		}
	case model.TypeStruct:
		s := schema.LookupStruct(schema.Import, schema.Type)
		if schema.Import != "" {
			prefix = fmt.Sprintf("%s.%s.", common.CamelTitleCase(schema.System().Name), common.CamelTitleCase(schema.Import))
		}
		if s != nil {
			text = fmt.Sprintf("%s%s", prefix, common.CamelTitleCase(s.Name))
		}
	case model.TypeInterface:
		i := schema.LookupInterface(schema.Import, schema.Type)
		if schema.Import != "" {
			prefix = fmt.Sprintf("%s.%s.", common.CamelTitleCase(schema.System().Name), common.CamelTitleCase(schema.Import))
		}
		if i != nil {
			text = fmt.Sprintf("%sI%s", prefix, common.CamelTitleCase(i.Name))
		}
	default:
		return "xxx", fmt.Errorf("csReturn unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		return fmt.Sprintf("List<%s>", text), nil
	}
	return text, nil
}

// csReturn casts the value to a TypedNode and deducts the C# return type.
func csReturn(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("csReturn node is nil")
	}
	return ToReturnString(prefix, &node.Schema)
}
