package conf

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmdGet(t *testing.T) {
	var cmd = getCmd
	buf := bytes.NewBufferString("")
	cmd.SetOut(buf)
	cmd.Execute()
	out, err := io.ReadAll(buf)
	assert.NoError(t, err)
	assert.Contains(t, string(out), "confGet called")
}
