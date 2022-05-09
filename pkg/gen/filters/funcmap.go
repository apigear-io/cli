package filters

import (
	"objectapi/pkg/gen/filters/filtercpp"
	"objectapi/pkg/gen/filters/filtergo"
	"objectapi/pkg/log"
	"text/template"

	"github.com/iancoleman/strcase"
)

func PopulateFuncMap() template.FuncMap {
	log.Info("populate func map")
	fm := make(template.FuncMap)
	fm["to_snake_case"] = strcase.ToSnake
	fm["to_camel_case"] = strcase.ToCamel
	fm["to_kebab_case"] = strcase.ToKebab
	fm["to_delimited_case"] = strcase.ToDelimited
	fm["to_lower_camel_case"] = strcase.ToLowerCamel

	filtercpp.PopulateFuncMap(fm)
	filtergo.PopulateFuncMap(fm)
	return fm
}
