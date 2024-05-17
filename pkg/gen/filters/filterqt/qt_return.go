package filterqt

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToReturnString(prefix string, schema *model.Schema) (string, error) {
	text := ""
	switch schema.Type {
	case "void":
		text = "void"
	case "string":
		text = "QString"
	case "int":
		text = "int"
	case "int32":
		text = "qint32"
	case "int64":
		text = "qint64"
	case "float":
		text = "qreal"
	case "float32":
		text = "float"
	case "float64":
		text = "double"
	case "bool":
		text = "bool"
	default:
		if schema.KindType == model.TypeExtern {
			xe := qtExtern(schema.GetExtern())
			namespace_prefix := ""
			if xe.NameSpace != "" {
				namespace_prefix = fmt.Sprintf("%s::", xe.NameSpace)
			}
			text = fmt.Sprintf("%s%s", namespace_prefix, xe.Name)
		}
		e := schema.LookupEnum("", schema.Type)
		e_imported := schema.LookupEnum(schema.Import, schema.Type)
		if e != nil {
			text = fmt.Sprintf("%s%s::%sEnum", prefix, e.Name, e.Name)
		} else if e_imported != nil {
			text = fmt.Sprintf("%s::%s::%sEnum", qtNamespace(e_imported.Module.Name), e.Name, e.Name)
		}
		s := schema.LookupStruct("", schema.Type)
		s_imported := schema.LookupStruct(schema.Import, schema.Type)
		if s != nil {
			text = fmt.Sprintf("%s%s", prefix, s.Name)
		} else if s_imported != nil {
			text = fmt.Sprintf("%s::%s", qtNamespace(s_imported.Module.Name), s_imported.Name)
		}
		
		i := schema.LookupInterface("", schema.Type)
		i_imported := schema.LookupInterface(schema.Import, schema.Type)
		if i != nil {
			text = fmt.Sprintf("%s%s*", prefix, i.Name)
		} else if i_imported != nil {
			text = fmt.Sprintf("%s::%s*", qtNamespace(i_imported.Module.Name), i_imported.Name)
		}
	}
	if schema.IsArray {
		text = fmt.Sprintf("QList<%s>", text)
	}
	return text, nil
}

// cast value to TypedNode and deduct the cpp return type
func qtReturn(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("qtReturn node is nil")
	}
	return ToReturnString(prefix, &node.Schema)
}
