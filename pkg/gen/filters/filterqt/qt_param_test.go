package filterqt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParam(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "bool propBool"},
		{"test", "Test1", "propInt", "int propInt"},
		{"test", "Test1", "propFloat", "double propFloat"},
		{"test", "Test1", "propString", "const QString& propString"},
		{"test", "Test1", "propBoolArray", "const QList<bool>& propBoolArray"},
		{"test", "Test1", "propIntArray", "const QList<int>& propIntArray"},
		{"test", "Test1", "propFloatArray", "const QList<double>& propFloatArray"},
		{"test", "Test1", "propStringArray", "const QList<QString>& propStringArray"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := qtParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamSymbols(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "const Enum1::Enum1Enum propEnum"},
		{"test", "Test2", "propStruct", "const Struct1& propStruct"},
		{"test", "Test2", "propInterface", "Interface1 *propInterface"},
		{"test", "Test2", "propEnumArray", "const QList<Enum1::Enum1Enum>& propEnumArray"},
		{"test", "Test2", "propStructArray", "const QList<Struct1>& propStructArray"},
		{"test", "Test2", "propInterfaceArray", "const QList<Interface1*>& propInterfaceArray"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := qtParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
