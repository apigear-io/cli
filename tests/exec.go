//go:build integration

package tests

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/apigear-io/cli/pkg/cmd"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/stretchr/testify/assert"
)

func execute(args string) string {
	buf := new(bytes.Buffer)
	root := cmd.NewRootCommand()
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(strings.Split(args, " "))
	root.Execute()
	return buf.String()
}

// setup sets up the test environment
// copies testdata to a temporary directory
// changes the current working directory to the temporary directory
// restores the original working directory and removes the temporary directory
func setup(t *testing.T) {
	origDir, err := os.Getwd()
	assert.NoError(t, err)
	tmpDir := t.TempDir()
	log.Printf("tmp dir: %s", tmpDir)
	err = helper.CopyDir("./testdata", tmpDir)
	assert.NoError(t, err)
	err = os.Chdir(tmpDir)
	assert.NoError(t, err)
	t.Cleanup(func() {
		err = os.Chdir(origDir)
		assert.NoError(t, err)
	})
}
