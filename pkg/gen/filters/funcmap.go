package filters

import (
	"apigear/pkg/gen/filters/filtercpp"
	"apigear/pkg/gen/filters/filtergo"
	"apigear/pkg/log"
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
	fm["pathCase"] = PathCase
	fm["lowerCase"] = LowerCase
	fm["upperCase"] = UpperCase
	fm["upperFirst"] = UpperFirst
	fm["lowerFirst"] = LowerFirst
	fm["firstChar"] = FirstChar
	fm["firstCharLower"] = FirstCharLower
	fm["firstCharUpper"] = FirstCharUpper

	filtercpp.PopulateFuncMap(fm)
	filtergo.PopulateFuncMap(fm)
	return fm
}
