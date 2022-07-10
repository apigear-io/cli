package filters

import (
	"apigear/pkg/gen/filters/filtercpp"
	"apigear/pkg/gen/filters/filtergo"
	"apigear/pkg/log"
	"text/template"
)

func PopulateFuncMap() template.FuncMap {
	log.Debug("populate func map")
	fm := make(template.FuncMap)
	fm["snake"] = SnakeCaseLower
	fm["Snake"] = SnakeCase
	fm["SNAKE"] = SnakeCaseUpper
	fm["camel"] = CamelCaseLower
	fm["Camel"] = CamelCase
	fm["dot"] = DotCaseLower
	fm["Dot"] = DotCase
	fm["DOT"] = DotCaseUpper
	fm["kebap"] = KebabCaseLower
	fm["Kebab"] = KebabCase
	fm["KEBAP"] = KebabCaseUpper
	fm["pathCase"] = PathCase
	fm["lower"] = LowerCase
	fm["upper"] = UpperCase
	fm["upper1"] = UpperFirst
	fm["lower1"] = LowerFirst
	fm["first"] = FirstCharLower
	fm["First"] = FirstChar
	fm["FIRST"] = FirstCharUpper

	filtercpp.PopulateFuncMap(fm)
	filtergo.PopulateFuncMap(fm)
	return fm
}
