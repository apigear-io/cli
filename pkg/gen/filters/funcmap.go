package filters

import (
	"text/template"

	"github.com/apigear-io/cli/pkg/gen/filters/filtercpp"
	"github.com/apigear-io/cli/pkg/gen/filters/filtergo"
	"github.com/apigear-io/cli/pkg/gen/filters/filterpy"
	"github.com/apigear-io/cli/pkg/gen/filters/filterts"
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
	fm["path"] = PathCaseLower
	fm["Path"] = PathCase
	fm["PATH"] = PathCaseUpper
	fm["lower"] = LowerCase
	fm["upper"] = UpperCase
	fm["upper1"] = UpperFirst
	fm["lower1"] = LowerFirst
	fm["first"] = FirstCharLower
	fm["First"] = FirstChar
	fm["FIRST"] = FirstCharUpper
	fm["join"] = Join
	fm["trimPrefix"] = TrimPrefix
	fm["trimSuffix"] = TrimSuffix
	fm["replace"] = Replace

	filtercpp.PopulateFuncMap(fm)
	filtergo.PopulateFuncMap(fm)
	filterts.PopulateFuncMap(fm)
	filterpy.PopulateFuncMap(fm)
	return fm
}
