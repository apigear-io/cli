package filterqt

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToParamString(prefix string, schema *model.Schema, name string) (string, error) {
	if schema.IsArray {
		inner := schema.InnerSchema()
		ret, err := ToReturnString(prefix, &inner)
		if err != nil {
			return "xxx", fmt.Errorf("qtParam inner value error: %s", err)
		}
		return fmt.Sprintf("const QList<%s>& %s", ret, name), nil
	}
	switch schema.Type {
	case "string":
		return fmt.Sprintf("const QString& %s", name), nil
	case "int":
		return fmt.Sprintf("int %s", name), nil
	case "int32":
		return fmt.Sprintf("qint32 %s", name), nil
	case "int64":
		return fmt.Sprintf("qint64 %s", name), nil
	case "float":
		return fmt.Sprintf("qreal %s", name), nil
	case "float32":
		return fmt.Sprintf("float %s", name), nil
	case "float64":
		return fmt.Sprintf("double %s", name), nil
	case "bool":
		return fmt.Sprintf("bool %s", name), nil
	}
	ex := schema.LookupExtern(schema.Import, schema.Type)
	if ex != nil {
		exQt := qtExtern(schema.GetExtern())
		namespace :=""
		if exQt.NameSpace != "" {
			namespace = fmt.Sprintf("%s::", exQt.NameSpace)
		}
		return fmt.Sprintf("const %s%s& %s", namespace, exQt.Name, name), nil
	}
	e := schema.LookupEnum("", schema.Type)
	if e != nil {
		return fmt.Sprintf("%s%s::%sEnum %s", prefix, e.Name, e.Name, name), nil
	}
	e_imported := schema.LookupEnum(schema.Import, schema.Type)
	if e_imported != nil {
		return fmt.Sprintf("%s::%s::%sEnum %s", qtNamespace(e_imported.Module.Name), e_imported.Name, e_imported.Name, name), nil
	}
	s := schema.LookupStruct("", schema.Type)
	if s != nil {
		return fmt.Sprintf("const %s%s& %s", prefix, s.Name, name), nil
	}
	s_imported := schema.LookupStruct(schema.Import, schema.Type)
	if s_imported != nil {
		return fmt.Sprintf("const %s::%s& %s", qtNamespace(s_imported.Module.Name), s_imported.Name, name), nil
	}
	i := schema.LookupInterface("", schema.Type)
	if i != nil {
		return fmt.Sprintf("%s%s *%s", prefix, i.Name, name), nil
	}
	i_imported := schema.LookupInterface(schema.Import, schema.Type)
	if i_imported != nil {
		return fmt.Sprintf("%s::%s *%s", qtNamespace(i_imported.Module.Name), i_imported.Name, name), nil
	}
	return "xxx", fmt.Errorf("qtParam unknown schema %s", schema.Dump())
}

func qtParam(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("qtParam node is nil")
	}
	return ToParamString(prefix, &node.Schema, node.Name)
}
