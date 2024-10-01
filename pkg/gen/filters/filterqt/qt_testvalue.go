package filterqt

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
		text = "QString(\"xyz\")"
	case model.TypeInt, model.TypeInt32:
		text = "1"
	case model.TypeInt64:
		text = "1LL"
	case model.TypeFloat, model.TypeFloat32:
		text = "1.1f"
	case model.TypeFloat64:
		text = "1.1"
	case model.TypeBool:
		text = "true"
	case model.TypeVoid:
		return ToDefaultString(prefix, schema)
	case model.TypeEnum:
		e_local := schema.LookupEnum("", schema.Type)
		e_imported := schema.LookupEnum(schema.Import, schema.Type)
		if e_local == nil && e_imported == nil {
			return "xxx", fmt.Errorf("qtTestValue enum not found: %s", schema.Dump())
		}
		// if enum is local it is found both as e_local and e_imported
		name := e_imported.Name
		member := common.UpperFirst(e_imported.Members[0].Name)
		if len(e_imported.Members) > 1 {
			member = common.UpperFirst(e_imported.Members[1].Name)
		}
		if e_local == nil {
			prefix = fmt.Sprintf("%s::", qtNamespace(e_imported.Module.Name))
		}
		text = fmt.Sprintf("%s%s::%s", prefix, name, member)
	// all types return deafualt value, but cannot be passed to deafult filter
	// due to variants with array. Here we want to return default element, not deafult empty array.
	case model.TypeStruct:
		s_local := schema.LookupStruct("", schema.Type)
		s_imported := schema.LookupStruct(schema.Import, schema.Type)
		if s_local == nil && s_imported == nil {
			return "xxx", fmt.Errorf("qtTestValue struct not found: %s", schema.Dump())
		}
		// if struct is local it is found both as s_local and s_imported
		name := s_imported.Name
		if s_local == nil {
			prefix = fmt.Sprintf("%s::", qtNamespace(s_imported.Module.Name))
		}
		text = fmt.Sprintf("%s%s()", prefix, name)
	case model.TypeExtern:
		xe := parseQtExtern(schema)
		if xe.Default != "" {
			text = xe.Default
		} else {
			namespace_prefix := ""
			if xe.NameSpace != "" {
				namespace_prefix = fmt.Sprintf("%s::", xe.NameSpace)
			}
			text = fmt.Sprintf("%s%s()", namespace_prefix, xe.Name)
		}
	case model.TypeInterface:
		i_local := schema.LookupInterface("", schema.Type)
		i_imported := schema.LookupInterface(schema.Import, schema.Type)
		if i_local == nil && i_imported == nil {
			return "xxx", fmt.Errorf("qtTestValue interface not found: %s", schema.Dump())
		}
		// if interface is local it is found both as s_local and s_imported
		name := i_imported.Name
		if i_local == nil {
			prefix = fmt.Sprintf("%s::", qtNamespace(i_imported.Module.Name))
		}
		text = fmt.Sprintf("%s%s()", prefix, name)
	default:
		return "xxx", fmt.Errorf("pyTestValue unknown schema %s", schema.Dump())
	}
	return text, nil
}

func qtTestValue(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("qtTestValue node is nil")
	}
	return ToTestValueString(prefix, &node.Schema)
}
