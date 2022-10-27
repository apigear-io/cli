package tests

import (
	"log"
	"os"
	"testing"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/stretchr/testify/assert"
)

func TestGenerateCmd(t *testing.T) {
	setup(t)
	output := execute(t, "generate")
	assert.Contains(t, output, "Usage:")
}

func TestGenerateSolutionCmd(t *testing.T) {
	setup(t)
	cwd, err := os.Getwd()
	helper.ListDir(".")
	assert.NoError(t, err)
	log.Printf("cwd: %s", cwd)
	output := execute(t, "generate solution ./apigear/test.solution.yaml")
	assert.Contains(t, output, "generated 1 files")
}
