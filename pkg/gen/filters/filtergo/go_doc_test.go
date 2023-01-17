package filtergo

import (
	"testing"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestDoc(t *testing.T) {
	table := []struct {
		in string
		rt string
	}{
		{"", ""},
		{"test", "// test\n"},
		{"test1\ntest2", "// test1\n// test2\n"},
		{"test1\ntest2\n", "// test1\n// test2\n"},
		{"test1\ntest2\n\n", "// test1\n// test2\n"},
	}
	for _, tt := range table {
		t.Run(tt.in, func(t *testing.T) {
			node := &model.NamedNode{
				Name:        "test",
				Description: tt.in,
			}
			r, err := goDoc(node)
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}
