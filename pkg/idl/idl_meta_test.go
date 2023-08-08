package idl

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestSimpleTag(t *testing.T) {
	s, err := LoadIdlFromFiles("meta", []string{"./testdata/meta.idl"})
	assert.NoError(t, err)
	table := []struct {
		ifaceId string
		meta    map[string]interface{}
		desc    string
	}{
		{"SingleLine", map[string]interface{}{"tag1": true}, "first line"},
		{"MultiLine", map[string]interface{}{"tag1": false, "tag2": true, "tag3": true}, "first line\nsecond line\nthird line"},
	}
	for _, tr := range table {
		t.Run(fmt.Sprintf("%s.%s", tr.ifaceId, tr.meta), func(t *testing.T) {
			i := s.LookupInterface("tb.meta", tr.ifaceId)
			assert.NotNil(t, i)
			assert.Equal(t, tr.meta, i.Meta)
			assert.Equal(t, tr.desc, i.Description)
		})
	}
}

func TestPropertyMeta(t *testing.T) {
	s, err := LoadIdlFromFiles("meta", []string{"./testdata/meta.idl"})
	assert.NoError(t, err)
	table := []struct {
		ifaceId string
		propId  string
		meta    map[string]interface{}
		desc    string
	}{
		{"FullMeta", "prop1", map[string]interface{}{"prop1": true}, "prop1"},
	}
	for _, tr := range table {
		t.Run(fmt.Sprintf("%s.%s", tr.ifaceId, tr.propId), func(t *testing.T) {
			i := s.LookupInterface("tb.meta", tr.ifaceId)
			assert.NotNil(t, i)
			p := i.LookupProperty(tr.propId)
			assert.NotNil(t, p)
			assert.Equal(t, tr.meta, p.Meta)
			assert.Equal(t, tr.desc, p.Description)
		})
	}
}

func TestOperationMeta(t *testing.T) {
	s, err := LoadIdlFromFiles("meta", []string{"./testdata/meta.idl"})
	assert.NoError(t, err)
	table := []struct {
		ifaceId string
		opId    string
		meta    map[string]interface{}
		desc    string
	}{
		{"FullMeta", "op1", map[string]interface{}{"op1": true}, "op1"},
	}
	for _, tr := range table {
		t.Run(fmt.Sprintf("%s.%s", tr.ifaceId, tr.opId), func(t *testing.T) {
			i := s.LookupInterface("tb.meta", tr.ifaceId)
			assert.NotNil(t, i)
			o := i.LookupOperation(tr.opId)
			assert.NotNil(t, o)
			assert.Equal(t, tr.meta, o.Meta)
			assert.Equal(t, tr.desc, o.Description)
		})
	}
}

func TestSignalMeta(t *testing.T) {
	s, err := LoadIdlFromFiles("meta", []string{"./testdata/meta.idl"})
	assert.NoError(t, err)
	table := []struct {
		ifaceId string
		sigId   string
		meta    map[string]interface{}
		desc    string
	}{
		{"FullMeta", "sig1", map[string]interface{}{"sig1": true}, "sig1"},
	}
	for _, tr := range table {
		t.Run(fmt.Sprintf("%s.%s", tr.ifaceId, tr.sigId), func(t *testing.T) {
			i := s.LookupInterface("tb.meta", tr.ifaceId)
			assert.NotNil(t, i)
			s := i.LookupSignal(tr.sigId)
			assert.NotNil(t, s)
			assert.Equal(t, tr.meta, s.Meta)
			assert.Equal(t, tr.desc, s.Description)
		})
	}
}

func TestStructMeta(t *testing.T) {
	s, err := LoadIdlFromFiles("meta", []string{"./testdata/meta.idl"})
	assert.NoError(t, err)
	table := []struct {
		structId string
		meta     map[string]interface{}
		desc     string
	}{
		{"MetaStruct", map[string]interface{}{"tag1": true}, "line 1"},
	}
	for _, tr := range table {
		t.Run(tr.structId, func(t *testing.T) {
			st := s.LookupStruct("tb.meta", tr.structId)
			assert.NotNil(t, st)
			assert.Equal(t, tr.meta, st.Meta)
			assert.Equal(t, tr.desc, st.Description)
		})
	}
}

const idl1 = `
@tag0
// line 0
module tb.meta
@tag1
// line 1
interface SingleLine {
	@tag2: true
	// line 2
	prop1: bool
	@tag3: true
	// line 3
	fun1()
}

@tag4: true
// line 4
struct Struct1 {
	@tag5: true
	// line 5
	field1: bool
}

@tag6: true
// line 6
enum Enum1 {
	@tag7: true
	// line 7
	Value1
}
`

const tpl1 = `
@tag0: {{ .Meta.tag0 }}
// {{ .Description }}
{{- range .Interfaces }}
@tag1: {{ .Meta.tag1 }}
// {{ .Description }}
{{- range .Properties }}
@tag2: {{ .Meta.tag2 }}
// {{ .Description }}
{{- end }}
{{- range .Operations }}
@tag3: {{ .Meta.tag3 }}
// {{ .Description }}
{{- end }}
{{- end }}
{{- range .Structs }}
@tag4: {{ .Meta.tag4 }}
// {{ .Description }}
{{- range .Fields }}
@tag5: {{ .Meta.tag5 }}
// {{ .Description }}
{{- end }}
{{- end }}
{{- range .Enums }}
@tag6: {{ .Meta.tag6 }}
// {{ .Description }}
{{- range .Members }}
@tag7: {{ .Meta.tag7 }}
// {{ .Description }}
{{- end }}
{{- end }}
`

const out1 = `
@tag0: true
// line 0
@tag1: true
// line 1
@tag2: true
// line 2
@tag3: true
// line 3
@tag4: true
// line 4
@tag5: true
// line 5
@tag6: true
// line 6
@tag7: true
// line 7
`

func TestTemplate(t *testing.T) {
	s, err := LoadIdlFromString("meta", idl1)
	assert.NoError(t, err)
	tpl, err := template.New("test").Parse(tpl1)
	assert.NoError(t, err)
	var buf bytes.Buffer
	err = tpl.Execute(&buf, s.Modules[0])
	assert.NoError(t, err)
	assert.Equal(t, out1, buf.String())
}
