package filtergo

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/apimodel"
)

func ToParamString(prefix string, schema *apimodel.Schema, name string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToParamString schema is nil")
	}
	if schema.IsImported() {
		prefix = fmt.Sprintf("%s.", schema.ShortImportName())
	}
	if schema.IsArray {
		inner := schema.InnerSchema()
		innerValue, err := ToReturnString(prefix, &inner)
		if err != nil {
			return "xxx", fmt.Errorf("ToParamString inner value error: %s", err)
		}
		return fmt.Sprintf("%s []%s", name, innerValue), nil
	}
	switch schema.KindType {
	case apimodel.TypeString:
		return fmt.Sprintf("%s string", name), nil
	case apimodel.TypeBytes:
		return fmt.Sprintf("%s []byte", name), nil
	case apimodel.TypeInt:
		return fmt.Sprintf("%s int32", name), nil
	case apimodel.TypeInt32:
		return fmt.Sprintf("%s int32", name), nil
	case apimodel.TypeInt64:
		return fmt.Sprintf("%s int64", name), nil
	case apimodel.TypeFloat:
		return fmt.Sprintf("%s float32", name), nil
	case apimodel.TypeFloat32:
		return fmt.Sprintf("%s float32", name), nil
	case apimodel.TypeFloat64:
		return fmt.Sprintf("%s float64", name), nil
	case apimodel.TypeBool:
		return fmt.Sprintf("%s bool", name), nil
	case apimodel.TypeAny:
		return fmt.Sprintf("%s any", name), nil
	case apimodel.TypeExtern:
		x := schema.LookupExtern(schema.Import, schema.Type)
		if x == nil {
			return "xxx", fmt.Errorf("goParam extern not found: %s", schema.Dump())
		}
		xe := parseGoExtern(schema)
		prefix = ""
		if xe.Import != "" {
			prefix = fmt.Sprintf("%s.", xe.Import)
		}
		return fmt.Sprintf("%s %s%s", name, prefix, xe.Name), nil
	case apimodel.TypeEnum:
		e := schema.LookupEnum(schema.Import, schema.Type)
		if e == nil {
			return "xxx", fmt.Errorf("goParam enum not found: %s", schema.Dump())
		}
		return fmt.Sprintf("%s %s%s", name, prefix, e.Name), nil
	case apimodel.TypeStruct:
		s := schema.LookupStruct(schema.Import, schema.Type)
		if s == nil {
			return "xxx", fmt.Errorf("goParam struct not found: %s", schema.Dump())
		}
		return fmt.Sprintf("%s %s%s", name, prefix, s.Name), nil
	case apimodel.TypeInterface:
		i := schema.LookupInterface(schema.Import, schema.Type)
		if i == nil {
			return "xxx", fmt.Errorf("goParam interface not found: %s", schema.Dump())
		}
		return fmt.Sprintf("%s %s%s", name, prefix, i.Name), nil
	}
	return "xxx", fmt.Errorf("goParam: unknown schema %s", schema.Dump())
}

func goParam(prefix string, node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("goParam called with nil node")
	}
	return ToParamString(prefix, &node.Schema, node.Name)
}
