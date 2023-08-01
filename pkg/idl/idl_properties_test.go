package idl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProperties(t *testing.T) {
	s, err := LoadIdlFromFiles("meta", []string{"./testdata/properties.idl"})
	assert.NoError(t, err)
	iface := s.LookupInterface("demo", "Demo")
	assert.NotNil(t, iface)
	table := []struct {
		name     string
		meta     map[string]interface{}
		readonly bool
	}{
		{"prop01", nil, false},
		{"prop02", nil, true},
		{"prop03", map[string]interface{}{"IsReadOnly": false}, true},
	}
	for _, tr := range table {
		t.Run(tr.name, func(t *testing.T) {
			p := iface.LookupProperty(tr.name)
			assert.NotNil(t, p)
			assert.Equal(t, tr.meta, p.Meta)
			assert.Equal(t, tr.readonly, p.IsReadOnly)
		})
	}

}
