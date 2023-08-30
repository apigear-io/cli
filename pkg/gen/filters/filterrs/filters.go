package filterrs

import (
	"text/template"
)

// PopulateFuncMap fills the given FuncMap with the functions from this package.
func PopulateFuncMap(fm template.FuncMap) {
	fm["rsNs"] = ns
	fm["rsNsOpen"] = nsOpen
	fm["rsNsClose"] = nsClose
	fm["rsReturn"] = rsReturn
	fm["rsDefault"] = rsDefault
	fm["rsParam"] = rsParam
	fm["rsParams"] = rsParams
	fm["rsVar"] = rsVar
	fm["rsVars"] = rsVars
	fm["rsType"] = rsType
	fm["rsTypeRef"] = rsTypeRef
}
