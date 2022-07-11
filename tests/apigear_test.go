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

func TestApiGear(t *testing.T) {
	ts, err := cmdtest.Read("./testdata")
	ts.Setup = setup
	ts.KeepRootDirs = true
	// ts.Setup = setupWrapped(cwd)
	assert.NoError(t, err)
	ts.Commands["apigear"] = cmdtest.InProcessProgram("apigear", cmd.Run)
	ts.Commands["exists"] = exists
	ts.Run(t, *update)
}
