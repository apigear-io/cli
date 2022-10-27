package tests

import (
	"os"
	"strings"
	"testing"

	"github.com/apigear-io/cli/pkg/cmd"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/stretchr/testify/assert"
)

func execute(t *testing.T, args string) string {
	t.Helper()
	var b strings.Builder
	root := cmd.NewRootCommand()
	root.SetOut(&b)
	root.SetErr(&b)
	root.SetArgs(strings.Split(args, " "))
	log.OnReportBytes(func(s string) {
		b.WriteString(s)
	})
	err := root.Execute()
	assert.NoError(t, err)
	return b.String()
}

// setup sets up the test environment
// copies testdata to a temporary directory
// changes the current working directory to the temporary directory
// restores the original working directory and removes the temporary directory
func setup(t *testing.T) string {
	t.Helper()
	origDir, err := os.Getwd()
	assert.NoError(t, err)
	tmpDir := t.TempDir()
	err = helper.CopyDir("./testdata", tmpDir)
	assert.NoError(t, err)
	err = os.Chdir(tmpDir)
	assert.NoError(t, err)
	t.Cleanup(func() {
		err = os.Chdir(origDir)
		assert.NoError(t, err)
	})
	return tmpDir
}
