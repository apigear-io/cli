package gen

import (
	"objectapi/pkg/gen/filtercpp"
	"text/template"
)

func PopulateFuncMap() template.FuncMap {
	fm := make(template.FuncMap)
	filtercpp.PopulateFuncMap(fm)
	return fm
}
