package tests

import (
	"fmt"
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
	fmt.Printf("output: %s\n", output)
	assert.Contains(t, output, "generated 1 files")
}

// test usage of generate expert command
func TestGenerateExpertUsageCmd(t *testing.T) {
	setup(t)
	output := execute(t, "generate")
	assert.Contains(t, output, "Usage:")

}

// test generate expert command with input, output and template flags
func TestGenerateExpertCmd(t *testing.T) {
	setup(t)
	output := execute(t, "generate expert -i apigear -o out -t tpl")
	assert.Contains(t, output, "generated 1 files")
}
