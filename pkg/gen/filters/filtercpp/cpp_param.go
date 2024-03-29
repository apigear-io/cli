package filtercpp

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToParamString(prefix string, schema *model.Schema, name string) (string, error) {
	t := schema.Type
	if schema.IsArray {
		inner := *schema
		inner.IsArray = false
		ret, err := ToReturnString(prefix, &inner)
		if err != nil {
			return "xxx", fmt.Errorf("ToParamString inner value error: %s", err)
		}
		return fmt.Sprintf("const std::list<%s>& %s", ret, name), nil
	}
	switch t {
	case "string":
		return fmt.Sprintf("const std::string& %s", name), nil
	case "int":
		return fmt.Sprintf("int %s", name), nil
	case "int32":
		return fmt.Sprintf("int32_t %s", name), nil
	case "int64":
		return fmt.Sprintf("int64_t %s", name), nil
	case "float":
		return fmt.Sprintf("float %s", name), nil
	case "float32":
		return fmt.Sprintf("float %s", name), nil
	case "float64":
		return fmt.Sprintf("double %s", name), nil
	case "bool":
		return fmt.Sprintf("bool %s", name), nil
	}
	e := schema.Module.LookupEnum(t)
	if e != nil {
		return fmt.Sprintf("%sEnum %s", e.Name, name), nil
	}
	s := schema.Module.LookupStruct(t)
	if s != nil {
		return fmt.Sprintf("const %s& %s", s.Name, name), nil
	}
	i := schema.Module.LookupInterface(t)
	if i != nil {
		return fmt.Sprintf("%s* %s", i.Name, name), nil
	}
	return "xxx", fmt.Errorf("ToParamString: unknown type %s", t)
}

func cppParam(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("cppParam node is nil")
	}
	return ToParamString(prefix, &node.Schema, node.Name)
}
