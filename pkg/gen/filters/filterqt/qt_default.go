package filterqt

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

// ToDefaultString returns the default value for a type
func ToDefaultString(prefix string, schema *model.Schema) (string, error) {
	text := ""
	switch schema.Type {
	case "void":
		text = "void"
	case "string":
		text = "QString()"
	case "int", "int32":
		text = "0"
	case "int64":
		text = "0LL"
	case "float", "float32":
		text = "0.0f"
	case "float64":
		text = "0.0"
	case "bool":
		text = "false"
	default:
		e := schema.LookupEnum(schema.Import, schema.Type)
		if e != nil {
			text = fmt.Sprintf("%s%s::%s", prefix, e.Name, common.UpperFirst(e.Members[0].Name))
		}
		s := schema.LookupStruct(schema.Import, schema.Type)
		if s != nil {
			text = fmt.Sprintf("%s%s()", prefix, s.Name)
		}
		i := schema.LookupInterface(schema.Import, schema.Type)
		if i != nil {
			text = "nullptr"
		}
	}
	if schema.IsArray {
		inner := model.Schema{
			Import: schema.Import,
			Type:   schema.Type,
			Module: schema.Module,
		}
		ret, err := ToReturnString(prefix, &inner)
		if err != nil {
			return "xxx", fmt.Errorf("qtDefault inner value error: %s", err)
		}
		text = fmt.Sprintf("QList<%s>()", ret)
	}
	return text, nil
}

// qtDefault returns the default value for a type
func qtDefault(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("qtDefault node is nil")
	}
	return ToDefaultString(prefix, &node.Schema)
}
