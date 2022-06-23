package conf

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func ExecuteCmd(t *testing.T, cmd *cobra.Command, args ...string) string {
	t.Helper()
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)
	err := cmd.Execute()
	assert.NoError(t, err)
	return strings.TrimSpace(buf.String())
}

func TestCmdGetAllSettings(t *testing.T) {
	out := ExecuteCmd(t, NewGetCmd())
	assert.Contains(t, out, "All settings")
}

func TestCmdGetOneSetting(t *testing.T) {
	out := ExecuteCmd(t, NewGetCmd(), "foo")
	assert.Contains(t, out, "foo")
}
