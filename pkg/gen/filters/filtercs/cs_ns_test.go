package filtercs

import (
	"testing"

	"github.com/apigear-io/cli/pkg/model"

	"github.com/stretchr/testify/assert"
)

func TestNS(t *testing.T) {
	t.Parallel()
	table := []struct {
		in  string
		out string
	}{
		{"a", "A"},
		{"a.b", "A.B"},
		{"a.b.c", "A.B.C"},
		{"demo.x", "Demo.X"},
	}
	for _, tt := range table {
		t.Run(tt.in, func(t *testing.T) {
			m := model.NewModule(tt.in, "1.0")
			r, err := csNs(m)
			assert.NoError(t, err)
			assert.Equal(t, tt.out, r)
		})
	}
}

func TestNSOpen(t *testing.T) {
	t.Parallel()
	table := []struct {
		in  string
		out string
	}{
		{"a", "namespace A\n{"},
		{"a.b", "namespace A.B\n{"},
		{"a.b.c", "namespace A.B.C\n{"},
	}
	for _, tt := range table {
		t.Run(tt.in, func(t *testing.T) {
			m := model.NewModule(tt.in, "1.0")
			r, err := csNsOpen(m)
			assert.NoError(t, err)
			assert.Equal(t, tt.out, r)
		})
	}
}

func TestNSClose(t *testing.T) {
	t.Parallel()
	table := []struct {
		in  string
		out string
	}{
		{"a", "} // namespace A"},
		{"a.b", "} // namespace A.B"},
		{"a.b.c", "} // namespace A.B.C"},
	}
	for _, tt := range table {
		t.Run(tt.in, func(t *testing.T) {
			m := model.NewModule(tt.in, "1.0")
			r, err := csNsClose(m)
			assert.NoError(t, err)
			assert.Equal(t, tt.out, r)
		})
	}
}
