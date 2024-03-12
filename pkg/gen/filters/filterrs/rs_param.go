package filterrs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToParamString(prefixVarName string, prefixComplexType string, schema *model.Schema, node *model.TypedNode) (string, error) {
	name, err := ToVarString(prefixVarName, node)
	if err != nil {
		return "xxx", fmt.Errorf("rsParam inner value error: %s", err)
	}
	if schema.IsArray {
		inner := schema.InnerSchema()
		ret, err := ToReturnString(prefixComplexType, &inner)
		if err != nil {
			return "xxx", fmt.Errorf("rsParam inner value error: %s", err)
		}
		return fmt.Sprintf("%s: &[%s]", name, ret), nil
	}
	switch schema.Type {
	case "string":
		return fmt.Sprintf("%s: &str", name), nil
	case "int":
		return fmt.Sprintf("%s: i32", name), nil
	case "int32":
		return fmt.Sprintf("%s: i32", name), nil
	case "int64":
		return fmt.Sprintf("%s: i64", name), nil
	case "float":
		return fmt.Sprintf("%s: f32", name), nil
	case "float32":
		return fmt.Sprintf("%s: f32", name), nil
	case "float64":
		return fmt.Sprintf("%s: f64", name), nil
	case "bool":
		return fmt.Sprintf("%s: bool", name), nil
	}
	e := schema.LookupEnum(schema.Import, schema.Type)
	if e != nil {
		return fmt.Sprintf("%s: %sEnum", name, e.Name), nil
	}
	s := schema.LookupStruct(schema.Import, schema.Type)
	if s != nil {
		return fmt.Sprintf("%s: &%s", name, s.Name), nil
	}
	i := schema.LookupInterface(schema.Import, schema.Type)
	if i != nil {
		return fmt.Sprintf("%s: &%s", name, i.Name), nil
	}
	return "xxx", fmt.Errorf("rsParam unknown schema %s", schema.Dump())
}

func rsParam(prefixVarName string, prefixComplexType string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("rsParam node is nil")
	}
	return ToParamString(prefixVarName, prefixComplexType, &node.Schema, node)
}
