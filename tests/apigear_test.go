package tests

import (
	"apigear/cmd"
	"flag"
	"os"
	"path"
	"testing"

	cmdtest "github.com/google/go-cmdtest"
	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update test files with results")

// func copyFile(src string, dstDir string) error {
// 	buf, err := os.ReadFile(src)
// 	if err != nil {
// 		return err
// 	}
// 	dst := path.Join(dstDir, src)
// 	return os.WriteFile(dst, buf, 0644)
// }

// func setupWrapped(cwd string) func(string) error {
// 	return func(root string) error {
// 		err := copyFile(path.Join(cwd, "api/demo.module.yaml"), root)
// 		if err != nil {
// 			return err
// 		}
// 		err = copyFile(path.Join(cwd, "api/demo.solution.yaml"), root)
// 		if err != nil {
// 			return err
// 		}
// 		err = copyFile(path.Join(cwd, "tpl/rules.yaml"), root)
// 		if err != nil {
// 			return err
// 		}
// 		return copyFile(path.Join(cwd, "tpl/templates/api.json.tmpl"), root)
// 	}
// }

func TestApiGear(t *testing.T) {
	cwd, err := os.Getwd()
	assert.NoError(t, err)
	os.Setenv("DATA", path.Join(cwd, "testdata"))
	// cwd, err := os.Getwd()
	// assert.NoError(t, err)
	ts, err := cmdtest.Read("./testdata")
	ts.KeepRootDirs = true
	// ts.Setup = setupWrapped(cwd)
	assert.NoError(t, err)
	ts.Commands["apigear"] = cmdtest.InProcessProgram("apigear", cmd.Run)
	ts.Run(t, *update)
}
