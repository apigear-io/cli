package filterqt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQtNamespace(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		in    string
		result string
	}{
		{"", ""},
		{"namespace", "namespace"},
		{"NameSpace", "name_space"},
		{"Name_Space", "name_space"},
		{"NAMESPACE", "namespace"},
		{"Name Space", "name_space"},
	}
	
	for _, testLine := range tests {
		t.Run(testLine.in, func(t *testing.T) {
			out := qtNamespace(testLine.in)
			assert.Equal(t, testLine.result, out)
		})
	}
}