package filtercpp

import (
	"text/template"
)

func GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"cpp_ns_open":  nsOpen,
		"cpp_ns_close": nsClose,
		"cpp_return":   cppReturn,
		"cpp_default":  cppDefault,
		"cpp_param":    cppParam,
	}
}
