package filtergo

import (
	"text/template"
)

// PopulateFuncMap fills the given FuncMap with the functions from this package.
func PopulateFuncMap(fm template.FuncMap) {
	fm["go_return"] = goReturn
	fm["go_default"] = goDefault
}
