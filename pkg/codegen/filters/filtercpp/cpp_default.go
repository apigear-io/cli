package filtercpp

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/codegen/filters/common"
	"github.com/apigear-io/cli/pkg/apimodel"
)

// ToDefaultString returns the default value for a type
func ToDefaultString(prefix string, schema *apimodel.Schema) (string, error) {
	text := ""
	switch schema.KindType {
	case apimodel.TypeVoid:
		text = "void"
	case apimodel.TypeString:
		text = "std::string()"
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
	case apimodel.TypeExtern:
		xe := parseCppExtern(schema)
		if xe.Default != "" {
			text = xe.Default
		} else {
			if xe.NameSpace != "" {
				prefix = fmt.Sprintf("%s::", xe.NameSpace)
			} else {
				prefix = "" // Externs should not be prefixed with any other prefix than given in extern info.
			}
			text = fmt.Sprintf("%s%s()", prefix, xe.Name)
		}
	case apimodel.TypeEnum:
		e := schema.LookupEnum(schema.Import, schema.Type)
		NameSpace := prefix
		if schema.Import != "" {
			NameSpace = fmt.Sprintf("%s::%s::", common.CamelTitleCase(schema.System().Name), common.CamelTitleCase(schema.Import))
		}
		if e != nil {
			text = fmt.Sprintf("%s%sEnum::%s", NameSpace, e.Name, e.Members[0].Name)
		}
	case apimodel.TypeStruct:
		s := schema.LookupStruct(schema.Import, schema.Type)
		NameSpace := prefix
		if schema.Import != "" {
			NameSpace = fmt.Sprintf("%s::%s::", common.CamelTitleCase(schema.System().Name), common.CamelTitleCase(schema.Import))
		}
		if s != nil {
			text = fmt.Sprintf("%s%s()", NameSpace, s.Name)
		}
	case apimodel.TypeInterface:
		i := schema.LookupInterface(schema.Import, schema.Type)
		if i != nil {
			text = "nullptr"
		}
	}
	if schema.IsArray {
		inner := schema.InnerSchema()
		ret, err := ToReturnString(prefix, &inner)
		if err != nil {
			return "xxx", fmt.Errorf("ToDefaultString inner value error: %s", err)
		}
		text = fmt.Sprintf("std::list<%s>()", ret)
	}
	return text, nil
}

// cppDefault returns the default value for a type
func cppDefault(prefix string, node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("cppDefault node is nil")
	}
	return ToDefaultString(prefix, &node.Schema)
}
