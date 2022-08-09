package tests

import (
	"flag"
	"testing"

	"github.com/apigear-io/cli/cmd"

	cmdtest "github.com/google/go-cmdtest"
	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update test files with results")

var entries = []string{
	"config.json",
	"api/demo.module.yaml",
	"api/demo.solution.yaml",
	"tpl/rules.yaml",
	"tpl/templates/api.json.tpl",
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

func TestSDK(t *testing.T) {
	runCmdTest(t, "sdk")
}

func TestProject(t *testing.T) {
	runCmdTest(t, "project")
}

func TestConfig(t *testing.T) {
	runCmdTest(t, "config")
}

func TestRoot(t *testing.T) {
	runCmdTest(t, "root")
}

func TestTools(t *testing.T) {
	runCmdTest(t, "tools")
}

func TestTemplate(t *testing.T) {
	runCmdTest(t, "template")
}
