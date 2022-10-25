package filters

import (
	"text/template"

	"github.com/apigear-io/cli/pkg/gen/filters/filtercpp"
	"github.com/apigear-io/cli/pkg/gen/filters/filtergo"
	"github.com/apigear-io/cli/pkg/gen/filters/filterpy"
	"github.com/apigear-io/cli/pkg/gen/filters/filterqt"
	"github.com/apigear-io/cli/pkg/gen/filters/filterts"
	"github.com/apigear-io/cli/pkg/gen/filters/filterue"
	"github.com/apigear-io/cli/pkg/log"
)

func PopulateFuncMap() template.FuncMap {
	log.Debug().Msg("populate func map")
	fm := make(template.FuncMap)
	fm["snake"] = SnakeCaseLower
	fm["Snake"] = SnakeCase
	fm["SNAKE"] = SnakeCaseUpper
	fm["camel"] = CamelCaseLower
	fm["Camel"] = CamelCase
	fm["CAMEL"] = CamelCaseUpper
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
	fm["int2word"] = IntToWord
	fm["Int2Word"] = func(i int, prefix, postfix string) string {
		return CamelCase(IntToWord(i, prefix, postfix))
	}
	fm["INT2WORD"] = func(i int, prefix, postfix string) string {
		return UpperCase(IntToWord(i, prefix, postfix))
	}
	fm["plural"] = Pluralize
	fm["nl"] = NewLine

	filtercpp.PopulateFuncMap(fm)
	filtergo.PopulateFuncMap(fm)
	filterts.PopulateFuncMap(fm)
	filterpy.PopulateFuncMap(fm)
	filterue.PopulateFuncMap(fm)
	filterqt.PopulateFuncMap(fm)
	return fm
}
