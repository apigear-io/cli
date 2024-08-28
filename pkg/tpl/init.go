package tpl

import (
	"fmt"
	"os"

	"github.com/apigear-io/apigear-by-example/tplcpp"
	"github.com/apigear-io/apigear-by-example/tplgo"
	"github.com/apigear-io/apigear-by-example/tplpy"
	"github.com/apigear-io/apigear-by-example/tplrs"
	"github.com/apigear-io/apigear-by-example/tplts"
	"github.com/apigear-io/apigear-by-example/tplue"
	"github.com/apigear-io/cli/pkg/helper"
)

func InitTemplate(dir string, lang string) error {
	var rules []byte
	var apiTpl []byte
	var apiTplName string
	switch lang {
	case "cpp":
		rules = tplcpp.RulesYaml
		apiTpl = tplcpp.ApiTpl
		apiTplName = tplcpp.ApiTplName
	case "go":
		rules = tplgo.RulesYaml
		apiTpl = tplgo.ApiTpl
		apiTplName = tplgo.ApiTplName
	case "py":
		rules = tplpy.RulesYaml
		apiTpl = tplpy.ApiTpl
		apiTplName = tplpy.ApiTplName
	case "ts":
		rules = tplts.RulesYaml
		apiTpl = tplts.ApiTpl
		apiTplName = tplts.ApiTplName
	case "rs":
		rules = tplrs.RulesYaml
		apiTpl = tplrs.ApiTpl
		apiTplName = tplrs.ApiTplName
	case "ue":
		rules = tplue.RulesYaml
		apiTpl = tplue.ApiTpl
		apiTplName = tplue.ApiTplName
	default:
		return fmt.Errorf("unsupported language: %s", lang)
	}
	log.Info().Msgf("init template %s", dir)
	os.MkdirAll(dir, 0755)
	target := helper.Join(dir, "rules.yaml")
	log.Info().Msgf("write %s", target)
	err := os.WriteFile(target, rules, 0644)
	if err != nil {
		return err
	}
	target = helper.Join(dir, "templates")
	os.MkdirAll(target, 0755)
	target = helper.Join(target, apiTplName)
	log.Info().Msgf("write %s", target)
	err = os.WriteFile(target, apiTpl, 0644)
	if err != nil {
		return err
	}
	return nil
}
