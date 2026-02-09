package filters

import (
	"text/template"

	"github.com/apigear-io/cli/pkg/codegen/filters/common"
	"github.com/apigear-io/cli/pkg/codegen/filters/filtercpp"
	"github.com/apigear-io/cli/pkg/codegen/filters/filtergo"
	"github.com/apigear-io/cli/pkg/codegen/filters/filterjava"
	"github.com/apigear-io/cli/pkg/codegen/filters/filterjni"
	"github.com/apigear-io/cli/pkg/codegen/filters/filterjs"
	"github.com/apigear-io/cli/pkg/codegen/filters/filterpy"
	"github.com/apigear-io/cli/pkg/codegen/filters/filterqt"
	"github.com/apigear-io/cli/pkg/codegen/filters/filterrs"
	"github.com/apigear-io/cli/pkg/codegen/filters/filterts"
	"github.com/apigear-io/cli/pkg/codegen/filters/filterue"
)

func PopulateFuncMap() template.FuncMap {
	fm := make(template.FuncMap)

	common.PopulateFuncMap(fm)
	filtercpp.PopulateFuncMap(fm)
	filtergo.PopulateFuncMap(fm)
	filterts.PopulateFuncMap(fm)
	filterpy.PopulateFuncMap(fm)
	filterue.PopulateFuncMap(fm)
	filterqt.PopulateFuncMap(fm)
	filterjs.PopulateFuncMap(fm)
	filterrs.PopulateFuncMap(fm)
	filterjava.PopulateFuncMap(fm)
	filterjni.PopulateFuncMap(fm)

	return fm
}
