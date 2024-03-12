package common

import (
	"text/template"

	"github.com/apigear-io/cli/pkg/helper"
)

func PopulateFuncMap(fm template.FuncMap) {
	fm["snake"] = SnakeCaseLower
	fm["Snake"] = SnakeTitleCase
	fm["SNAKE"] = SnakeUpperCase
	fm["camel"] = CamelLowerCase
	fm["Camel"] = CamelTitleCase
	fm["CAMEL"] = CamelUpperCase
	fm["space"] = SpaceLowerCase
	fm["Space"] = SpaceTitleCase
	fm["SPACE"] = SpaceUpperCase
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
	fm["split"] = Split
	fm["splitLast"] = SplitLast
	fm["splitFirst"] = SplitFirst
	fm["trim"] = Trim
	fm["trimPrefix"] = TrimPrefix
	fm["trimSuffix"] = TrimSuffix
	fm["replace"] = Replace
	fm["contains"] = Contains
	fm["indexOf"] = IndexOf
	fm["int2word"] = IntToWordLower
	fm["Int2Word"] = IntToWordTitle
	fm["INT2WORD"] = IntToWordUpper
	fm["plural"] = Pluralize
	fm["abbreviate"] = helper.Abbreviate
	fm["nl"] = NewLine
	fm["toJson"] = ToJson
}
