package filtercpp

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

// ToDefaultString returns the default value for a type
func ToDefaultString(prefix string, schema *model.Schema) (string, error) {
	text := ""
	switch schema.KindType {
	case model.TypeVoid:
		text = "void"
	case model.TypeString:
		text = "std::string()"
	case model.TypeInt, model.TypeInt32:
		text = "0"
	case model.TypeInt64:
		text = "0LL"
	case model.TypeFloat, model.TypeFloat32:
		text = "0.0f"
	case model.TypeFloat64:
		text = "0.0"
	case model.TypeBool:
		text = "false"
	case model.TypeExtern:
		xe := parseCppExtern(schema)
		if xe.NameSpace != "" {
			prefix = fmt.Sprintf("%s::", xe.NameSpace)
		}
		text = fmt.Sprintf("%s%s()", prefix, xe.Name)
	case model.TypeEnum:
		e := schema.LookupEnum(schema.Import, schema.Type)
		NameSpace := ""
		if schema.Import != "" {
			NameSpace = fmt.Sprintf("%s::%s::", common.CamelTitleCase(schema.System().Name), common.CamelTitleCase(schema.Import))
		}
		if e != nil {
			text = fmt.Sprintf("%s%sEnum::%s", NameSpace, e.Name, e.Members[0].Name)
		}
	case model.TypeStruct:
		s := schema.LookupStruct(schema.Import, schema.Type)
		NameSpace := ""
		if schema.Import != "" {
			NameSpace = fmt.Sprintf("%s::%s::", common.CamelTitleCase(schema.System().Name), common.CamelTitleCase(schema.Import))
		}
		if s != nil {
			text = fmt.Sprintf("%s%s()", NameSpace, s.Name)
		}
	case model.TypeInterface:
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
func cppDefault(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("cppDefault node is nil")
	}
	return ToDefaultString(prefix, &node.Schema)
}
