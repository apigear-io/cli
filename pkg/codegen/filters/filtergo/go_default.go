package filtergo

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
	"github.com/ettle/strcase"
)

func ToDefaultString(schema *objmodel.Schema, prefix string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToDefaultString schema is nil")
	}
	if schema.IsImported() {
		prefix = fmt.Sprintf("%s.", schema.ShortImportName())
	}
	var text string
	if schema.IsArray {
		switch schema.KindType {
		case objmodel.TypeString:
			text = "[]string{}"
		case objmodel.TypeBytes:
			text = "[][]byte{}"
		case objmodel.TypeInt:
			text = "[]int32{}"
		case objmodel.TypeInt32:
			text = "[]int32{}"
		case objmodel.TypeInt64:
			text = "[]int64{}"
		case objmodel.TypeFloat:
			text = "[]float32{}"
		case objmodel.TypeFloat32:
			text = "[]float32{}"
		case objmodel.TypeFloat64:
			text = "[]float64{}"
		case objmodel.TypeBool:
			text = "[]bool{}"
		case objmodel.TypeAny:
			text = "[]any{}"
		case objmodel.TypeExtern:
			xe := parseGoExtern(schema)
			if xe.Import != "" {
				prefix = fmt.Sprintf("%s.", xe.Import)
			}
			text = fmt.Sprintf("[]%s%s{}", prefix, xe.Name)
		case objmodel.TypeEnum:
			text = fmt.Sprintf("[]%s%s{}", prefix, schema.Type)
		case objmodel.TypeStruct:
			text = fmt.Sprintf("[]%s%s{}", prefix, schema.Type)
		case objmodel.TypeInterface:
			text = fmt.Sprintf("[]%s%s{}", prefix, schema.Type)
		default:
			return "xxx", fmt.Errorf("goDefault: unknown schema %s", schema.Dump())
		}
	} else {
		switch schema.KindType {
		case objmodel.TypeString:
			text = "\"\""
		case objmodel.TypeBytes:
			text = "[]byte{}"
		case objmodel.TypeInt:
			text = "int32(0)"
		case objmodel.TypeInt32:
			text = "int32(0)"
		case objmodel.TypeInt64:
			text = "int64(0)"
		case objmodel.TypeFloat:
			text = "float32(0.0)"
		case objmodel.TypeFloat32:
			text = "float32(0.0)"
		case objmodel.TypeFloat64:
			text = "float64(0.0)"
		case objmodel.TypeBool:
			text = "false"
		case objmodel.TypeAny:
			text = "nil"
		case objmodel.TypeExtern:
			xe := parseGoExtern(schema)
			if xe.Import != "" {
				prefix = fmt.Sprintf("%s.", xe.Import)
			}
			text = fmt.Sprintf("%s%s{}", prefix, xe.Name)
		case objmodel.TypeEnum:
			symbol := schema.GetEnum()
			member := symbol.Members[0]
			// upper case first letter

			text = fmt.Sprintf("%s%s%s", prefix, symbol.Name, strcase.ToPascal(member.Name))
		case objmodel.TypeStruct:
			symbol := schema.GetStruct()
			text = fmt.Sprintf("%s%s{}", prefix, symbol.Name)
		case objmodel.TypeInterface:
			text = "nil"
		case objmodel.TypeVoid:
			text = ""
		default:
			return "xxx", fmt.Errorf("goDefault: unknown schema %s", schema.Dump())
		}
	}
	return text, nil
}

func goDefault(prefix string, node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("goDefault node is nil")
	}
	return ToDefaultString(&node.Schema, prefix)
}
