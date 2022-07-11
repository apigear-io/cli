package tests

import (
	"apigear/cmd"
	"flag"
	"testing"

	cmdtest "github.com/google/go-cmdtest"
	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update test files with results")

var entries = []string{
	"api/demo.module.yaml",
	"api/demo.solution.yaml",
	"tpl/rules.yaml",
	"tpl/templates/api.json.tmpl",
}

func setup(root string) error {
	return copyTestData(TEST_DATA, root, entries...)
}

func runCmdTest(t *testing.T, source string) {
	ts, err := cmdtest.Read(source)
	ts.Setup = setup
	ts.KeepRootDirs = true
	assert.NoError(t, err)
	ts.Commands["apigear"] = cmdtest.InProcessProgram("apigear", cmd.Run)
	ts.Commands["exists"] = exists
	ts.Run(t, *update)
}

func TestGen(t *testing.T) {
	runCmdTest(t, "gen")
}

func TestPrj(t *testing.T) {
	runCmdTest(t, "prj")
}

func TestRoot(t *testing.T) {
	runCmdTest(t, "root")
}
