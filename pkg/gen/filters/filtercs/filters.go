package filtercs

import "text/template"

// PopulateFuncMap fills the given FuncMap with the C# filter functions.
func PopulateFuncMap(fm template.FuncMap) {
	fm["csReturn"] = csReturn
	fm["csType"] = csType
	fm["csDefault"] = csDefault
	fm["csParam"] = csParam
	fm["csParams"] = csParams
	fm["csVar"] = csVar
	fm["csVars"] = csVars
	fm["csTestValue"] = csTestValue
	fm["csAsyncReturn"] = csAsyncReturn
	fm["csNs"] = csNs
	fm["csNsOpen"] = csNsOpen
	fm["csNsClose"] = csNsClose
	fm["csExtern"] = csExtern
	fm["csExterns"] = csExterns
}
