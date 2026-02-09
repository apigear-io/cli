package filtergo

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/apimodel"
	"github.com/ettle/strcase"
)

func ToDefaultString(schema *apimodel.Schema, prefix string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToDefaultString schema is nil")
	}
	if schema.IsImported() {
		prefix = fmt.Sprintf("%s.", schema.ShortImportName())
	}
	var text string
	if schema.IsArray {
		switch schema.KindType {
		case apimodel.TypeString:
			text = "[]string{}"
		case apimodel.TypeBytes:
			text = "[][]byte{}"
		case apimodel.TypeInt:
			text = "[]int32{}"
		case apimodel.TypeInt32:
			text = "[]int32{}"
		case apimodel.TypeInt64:
			text = "[]int64{}"
		case apimodel.TypeFloat:
			text = "[]float32{}"
		case apimodel.TypeFloat32:
			text = "[]float32{}"
		case apimodel.TypeFloat64:
			text = "[]float64{}"
		case apimodel.TypeBool:
			text = "[]bool{}"
		case apimodel.TypeAny:
			text = "[]any{}"
		case apimodel.TypeExtern:
			xe := parseGoExtern(schema)
			if xe.Import != "" {
				prefix = fmt.Sprintf("%s.", xe.Import)
			}
			text = fmt.Sprintf("[]%s%s{}", prefix, xe.Name)
		case apimodel.TypeEnum:
			text = fmt.Sprintf("[]%s%s{}", prefix, schema.Type)
		case apimodel.TypeStruct:
			text = fmt.Sprintf("[]%s%s{}", prefix, schema.Type)
		case apimodel.TypeInterface:
			text = fmt.Sprintf("[]%s%s{}", prefix, schema.Type)
		default:
			return "xxx", fmt.Errorf("goDefault: unknown schema %s", schema.Dump())
		}
	} else {
		switch schema.KindType {
		case apimodel.TypeString:
			text = "\"\""
		case apimodel.TypeBytes:
			text = "[]byte{}"
		case apimodel.TypeInt:
			text = "int32(0)"
		case apimodel.TypeInt32:
			text = "int32(0)"
		case apimodel.TypeInt64:
			text = "int64(0)"
		case apimodel.TypeFloat:
			text = "float32(0.0)"
		case apimodel.TypeFloat32:
			text = "float32(0.0)"
		case apimodel.TypeFloat64:
			text = "float64(0.0)"
		case apimodel.TypeBool:
			text = "false"
		case apimodel.TypeAny:
			text = "nil"
		case apimodel.TypeExtern:
			xe := parseGoExtern(schema)
			if xe.Import != "" {
				prefix = fmt.Sprintf("%s.", xe.Import)
			}
			text = fmt.Sprintf("%s%s{}", prefix, xe.Name)
		case apimodel.TypeEnum:
			symbol := schema.GetEnum()
			member := symbol.Members[0]
			// upper case first letter

			text = fmt.Sprintf("%s%s%s", prefix, symbol.Name, strcase.ToPascal(member.Name))
		case apimodel.TypeStruct:
			symbol := schema.GetStruct()
			text = fmt.Sprintf("%s%s{}", prefix, symbol.Name)
		case apimodel.TypeInterface:
			text = "nil"
		case apimodel.TypeVoid:
			text = ""
		default:
			return "xxx", fmt.Errorf("goDefault: unknown schema %s", schema.Dump())
		}
	}
	return text, nil
}

func goDefault(prefix string, node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("goDefault node is nil")
	}
	return ToDefaultString(&node.Schema, prefix)
}
