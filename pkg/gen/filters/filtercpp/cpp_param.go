package filtercpp

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToParamString(prefix string, schema *model.Schema, name string) (string, error) {
	if schema.IsArray {
		inner := schema.InnerSchema()
		ret, err := ToReturnString(prefix, &inner)
		if err != nil {
			return "xxx", fmt.Errorf("ToParamString inner value error: %s", err)
		}
		return fmt.Sprintf("const std::list<%s>& %s", ret, name), nil
	}
	switch schema.KindType {
	case model.TypeString:
		return fmt.Sprintf("const std::string& %s", name), nil
	case model.TypeInt:
		return fmt.Sprintf("int %s", name), nil
	case model.TypeInt32:
		return fmt.Sprintf("int32_t %s", name), nil
	case model.TypeInt64:
		return fmt.Sprintf("int64_t %s", name), nil
	case model.TypeFloat:
		return fmt.Sprintf("float %s", name), nil
	case model.TypeFloat32:
		return fmt.Sprintf("float %s", name), nil
	case model.TypeFloat64:
		return fmt.Sprintf("double %s", name), nil
	case model.TypeBool:
		return fmt.Sprintf("bool %s", name), nil
	case model.TypeExtern:
		xe := parseCppExtern(schema)
		if xe.NameSpace != "" {
			prefix = fmt.Sprintf("%s::", xe.NameSpace)
		}
		return fmt.Sprintf("const %s%s& %s", prefix, xe.Name, name), nil
	case model.TypeEnum:
		e := schema.LookupEnum(schema.Import, schema.Type)
		if e != nil {
			return fmt.Sprintf("%sEnum %s", e.Name, name), nil
		}
	case model.TypeStruct:
		s := schema.LookupStruct(schema.Import, schema.Type)
		if s != nil {
			return fmt.Sprintf("const %s& %s", s.Name, name), nil
		}
	case model.TypeInterface:
		i := schema.LookupInterface(schema.Import, schema.Type)
		if i != nil {
			return fmt.Sprintf("%s* %s", i.Name, name), nil
		}
	}
	return "xxx", fmt.Errorf("cppParam: unknown schema %s", schema.Dump())
}

func cppParam(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("cppParam node is nil")
	}
	return ToParamString(prefix, &node.Schema, node.Name)
}
