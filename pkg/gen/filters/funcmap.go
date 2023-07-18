package filters

import (
	"text/template"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/gen/filters/filtercpp"
	"github.com/apigear-io/cli/pkg/gen/filters/filtergo"
	"github.com/apigear-io/cli/pkg/gen/filters/filterjs"
	"github.com/apigear-io/cli/pkg/gen/filters/filterpy"
	"github.com/apigear-io/cli/pkg/gen/filters/filterqt"
	"github.com/apigear-io/cli/pkg/gen/filters/filterts"
	"github.com/apigear-io/cli/pkg/gen/filters/filterue"
	"github.com/apigear-io/cli/pkg/helper"
)

func PopulateFuncMap() template.FuncMap {
	fm := make(template.FuncMap)
	fm["snake"] = common.SnakeCaseLower
	fm["Snake"] = common.SnakeTitleCase
	fm["SNAKE"] = common.SnakeUpperCase
	fm["camel"] = common.CamelLowerCase
	fm["Camel"] = common.CamelTitleCase
	fm["CAMEL"] = common.CamelUpperCase
	fm["space"] = common.SpaceLowerCase
	fm["Space"] = common.SpaceTitleCase
	fm["SPACE"] = common.SpaceUpperCase
	fm["dot"] = common.DotLowerCase
	fm["Dot"] = common.DotTitleCase
	fm["DOT"] = common.DotUpperCase
	fm["kebap"] = common.KebabLowerCase
	fm["Kebab"] = common.KebabTitleCase
	fm["KEBAP"] = common.KebabUpperCase
	fm["path"] = common.PathLowerCase
	fm["Path"] = common.PathTitleCase
	fm["PATH"] = common.PathUpperCase
	fm["lower"] = common.LowerCase
	fm["upper"] = common.UpperCase
	fm["upper1"] = common.UpperFirst
	fm["lower1"] = common.LowerFirst
	fm["first"] = common.FirstCharLower
	fm["First"] = common.FirstChar
	fm["FIRST"] = common.FirstCharUpper
	fm["join"] = common.Join
	fm["trimPrefix"] = common.TrimPrefix
	fm["trimSuffix"] = common.TrimSuffix
	fm["replace"] = common.Replace
	fm["int2word"] = common.IntToWordLower
	fm["Int2Word"] = common.IntToWordTitle
	fm["INT2WORD"] = common.IntToWordUpper
	fm["plural"] = common.Pluralize
	fm["abbreviate"] = helper.Abbreviate
	fm["nl"] = common.NewLine

	filtercpp.PopulateFuncMap(fm)
	filtergo.PopulateFuncMap(fm)
	filterts.PopulateFuncMap(fm)
	filterpy.PopulateFuncMap(fm)
	filterue.PopulateFuncMap(fm)
	filterqt.PopulateFuncMap(fm)
	filterjs.PopulateFuncMap(fm)

	return fm
}
