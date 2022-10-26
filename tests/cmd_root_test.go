//go:build integration

package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCmd(t *testing.T) {
	setup(t)
	output := execute("")
	assert.Contains(t, output, "Usage:")
}

func TestRootHelpCmd(t *testing.T) {
	setup(t)
	out := execute("help")
	assert.Contains(t, out, "Usage:")
}
