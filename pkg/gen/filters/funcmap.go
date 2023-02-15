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
	fm["Snake"] = SnakeTitleCase
	fm["SNAKE"] = SnakeUpperCase
	fm["camel"] = CamelLowerCase
	fm["Camel"] = CamelTitleCase
	fm["CAMEL"] = CamelUpperCase
	fm["dot"] = DotLowerCase
	fm["Dot"] = DotTitleCase
	fm["DOT"] = DotUpperCase
	fm["kebap"] = KebabLowerCase
	fm["Kebab"] = KebabTitleCase
	fm["KEBAP"] = KebabUpperCase
	fm["path"] = PathLowerCase
	fm["Path"] = PathTitleCase
	fm["PATH"] = PathUpperCase
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
	fm["int2word"] = IntToWordLower
	fm["Int2Word"] = IntToWordTitle
	fm["INT2WORD"] = IntToWordUpper
	fm["plural"] = Pluralize
	fm["abbreviate"] = Abbreviate
	fm["nl"] = NewLine

	filtercpp.PopulateFuncMap(fm)
	filtergo.PopulateFuncMap(fm)
	filterts.PopulateFuncMap(fm)
	filterpy.PopulateFuncMap(fm)
	filterue.PopulateFuncMap(fm)
	filterqt.PopulateFuncMap(fm)

	return fm
}
