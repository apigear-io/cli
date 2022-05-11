package filtergo

import (
	"text/template"
)

// PopulateFuncMap fills the given FuncMap with the functions from this package.
func PopulateFuncMap(fm template.FuncMap) {
	fm["goReturn"] = goReturn
	fm["goDefault"] = goDefault
	fm["goParam"] = goParam
	fm["goParams"] = goParams
}
