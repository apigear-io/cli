package filters

import (
	"objectapi/pkg/gen/filters/filtercpp"
	"objectapi/pkg/gen/filters/filtergo"
	"objectapi/pkg/log"
	"text/template"
)

func PopulateFuncMap() template.FuncMap {
	log.Info("populate func map")
	fm := make(template.FuncMap)
	fm["snakeCase"] = SnakeCase
	fm["camelCase"] = CamelCase
	fm["delimitedCase"] = DotCase
	fm["lowerCamelCase"] = LowerCamelCase
	fm["kebabCase"] = KebabCase

	filtercpp.PopulateFuncMap(fm)
	filtergo.PopulateFuncMap(fm)
	return fm
}
