package filterpy

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

// ToTestValueString returns the test value string for a given schema.
// We intentionally ignore arrays in order to return the test value of the inner type.
func ToTestValueString(prefix string, schema *model.Schema) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("pyTestValue schema is nil")
	}
	if schema.Module == nil {
		return "xxx", fmt.Errorf("pyTestValue schema module is nil")
	}
	var text string
	switch schema.KindType {
	case model.TypeString:
		text = "\"xyz\""
	case model.TypeInt, model.TypeInt32, model.TypeInt64:
		text = "1"
	case model.TypeFloat, model.TypeFloat32, model.TypeFloat64:
		text = "1.1"
	case model.TypeBool:
		text = "True"
	case model.TypeVoid:
		return ToDefaultString(schema, prefix)
	case model.TypeEnum:
		e_local := schema.LookupEnum("", schema.Type)
		e_imported := schema.LookupEnum(schema.Import, schema.Type)
		if e_local == nil && e_imported == nil {
			return "xxx", fmt.Errorf("pyTestValue enum not found: %s", schema.Dump())
		}
		// if enum is local it is found both as e_local and e_imported
		name := common.CamelTitleCase(e_imported.Name)
		member := common.SnakeUpperCase(e_imported.Members[0].Name)
		if len(e_imported.Members) > 1 {
			member = common.SnakeUpperCase(e_imported.Members[1].Name)
		}
		if e_local == nil {
			prefix = fmt.Sprintf("%s.api.", e_imported.Module.Name)
		}
		text = fmt.Sprintf("%s%s.%s", prefix, name, member)
	case model.TypeStruct:
		s_local := schema.LookupStruct("", schema.Type)
		s_imported := schema.LookupStruct(schema.Import, schema.Type)
		if s_local == nil && s_imported == nil {
			return "xxx", fmt.Errorf("pyTestValue struct not found: %s", schema.Dump())
		}
		// if struct is local it is found both as s_local and s_imported
		ident := common.CamelTitleCase(s_imported.Name)
		if s_local == nil {
			prefix = fmt.Sprintf("%s.api.", s_imported.Module.Name)
		}
		text = fmt.Sprintf("%s%s()", prefix, ident)
	case model.TypeExtern:
		xe := parsePyExtern(schema)
		if xe.Default != "" {
			text = xe.Default
		} else {
			py_module := ""
			if xe.Import != "" {
				py_module = fmt.Sprintf("%s.", xe.Import)
			}
			text = fmt.Sprintf("%s%s()", py_module, xe.Name)
		}
	case model.TypeInterface:
		i_local := schema.LookupInterface("", schema.Type)
		i_imported := schema.LookupInterface(schema.Import, schema.Type)
		if i_local == nil && i_imported == nil {
			return "xxx", fmt.Errorf("pyTestValue interface not found: %s", schema.Dump())
		}
		// if interface is local it is found both as s_local and s_imported
		ident := common.CamelTitleCase(i_imported.Name)
		if i_local == nil {
			prefix = fmt.Sprintf("%s.api.", i_imported.Module.Name)
		}
		text = fmt.Sprintf("%s%s()", prefix, ident)
	default:
		return "xxx", fmt.Errorf("pyTestValue unknown schema %s", schema.Dump())
	}
	return text, nil
}

func pyTestValue(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("pyTestValue node is nil")
	}
	return ToTestValueString(prefix, &node.Schema)
}
