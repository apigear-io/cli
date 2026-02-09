package filterjava

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/codegen/filters/common"
	"github.com/apigear-io/cli/pkg/apimodel"
)

func ToElementTypeString(prefix string, schema *apimodel.Schema) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToReturnString schema is nil")
	}
	var text string
	switch schema.KindType {
	case apimodel.TypeString:
		text = "String"
	case apimodel.TypeInt:
		text = "int"
	case apimodel.TypeInt32:
		text = "int"
	case apimodel.TypeInt64:
		text = "long"
	case apimodel.TypeFloat:
		text = "float"
	case apimodel.TypeFloat32:
		text = "float"
	case apimodel.TypeFloat64:
		text = "double"
	case apimodel.TypeBool:
		text = "boolean"
	case apimodel.TypeEnum:
		symbol := schema.GetEnum()
		text = fmt.Sprintf("%s%s", prefix, symbol.Name)
		e_local := schema.LookupEnum("", schema.Type)
		e_imported := schema.LookupEnum(schema.Import, schema.Type)
		if e_local == nil && e_imported == nil {
			return "xxx", fmt.Errorf("javaAsyncReturn enum not found: %s", schema.Dump())
		}
		// if enum is local it is found both as e_local and e_imported
		name := common.CamelTitleCase(e_imported.Name)
		if e_local == nil {
			prefix = fmt.Sprintf("%s.%s_api.", common.CamelLowerCase(e_imported.Module.Name), common.CamelLowerCase(e_imported.Module.Name))
		}
		text = fmt.Sprintf("%s%s", prefix, name)
	case apimodel.TypeStruct:
		s_local := schema.LookupStruct("", schema.Type)
		s_imported := schema.LookupStruct(schema.Import, schema.Type)
		if s_local == nil && s_imported == nil {
			return "xxx", fmt.Errorf("javaAsyncReturn struct not found: %s", schema.Dump())
		}
		// if struct is local it is found both as s_local and s_imported
		if s_local == nil {
			prefix = fmt.Sprintf("%s.%s_api.", common.CamelLowerCase(s_imported.Module.Name), common.CamelLowerCase(s_imported.Module.Name))
		}
		text = fmt.Sprintf("%s%s", prefix, common.CamelTitleCase(s_imported.Name))
	case apimodel.TypeExtern:
		xe := parseJavaExtern(schema)
		text = fmt.Sprintf("new %s()", xe.Name)
		var java_module string
		java_module = ""
		if xe.Package != "" {
			java_module = fmt.Sprintf("%s.", xe.Package)
		}
		text = fmt.Sprintf("%s%s", java_module, xe.Name)
	case apimodel.TypeInterface:
		i_local := schema.LookupInterface("", schema.Type)
		i_imported := schema.LookupInterface(schema.Import, schema.Type)
		if i_local == nil && i_imported == nil {
			return "xxx", fmt.Errorf("javaAsyncReturn interface not found: %s", schema.Dump())
		}
		// if interface is local it is found both as s_local and s_imported
		if i_local == nil {
			prefix = fmt.Sprintf("%s.%s_api.", common.CamelLowerCase(i_imported.Module.Name), common.CamelLowerCase(i_imported.Module.Name))
		}
		text = fmt.Sprintf("%sI%s", prefix, common.CamelTitleCase(i_imported.Name))
	default:
		return "xxx", fmt.Errorf("javaReturn unknown schema %s", schema.Dump())
	}
	return text, nil
}

func javaElementType(prefix string, node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("javaReturn node is nil")
	}
	return ToElementTypeString(prefix, &node.Schema)
}
