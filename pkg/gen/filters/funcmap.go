package filters

import (
	"text/template"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/gen/filters/filtercpp"
	"github.com/apigear-io/cli/pkg/gen/filters/filtergo"
	"github.com/apigear-io/cli/pkg/gen/filters/filterjava"
	"github.com/apigear-io/cli/pkg/gen/filters/filterjs"
	"github.com/apigear-io/cli/pkg/gen/filters/filterpy"
	"github.com/apigear-io/cli/pkg/gen/filters/filterqt"
	"github.com/apigear-io/cli/pkg/gen/filters/filterrs"
	"github.com/apigear-io/cli/pkg/gen/filters/filterts"
	"github.com/apigear-io/cli/pkg/gen/filters/filterue"
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

	return fm
}
