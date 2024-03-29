package filterqt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParam(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "bool propBool"},
		{"test", "Test1", "propInt", "int propInt"},
		{"test", "Test1", "propInt32", "qint32 propInt32"},
		{"test", "Test1", "propInt64", "qint64 propInt64"},
		{"test", "Test1", "propFloat", "qreal propFloat"},
		{"test", "Test1", "propFloat32", "float propFloat32"},
		{"test", "Test1", "propFloat64", "double propFloat64"},
		{"test", "Test1", "propString", "const QString& propString"},
		{"test", "Test1", "propBoolArray", "const QList<bool>& propBoolArray"},
		{"test", "Test1", "propIntArray", "const QList<int>& propIntArray"},
		{"test", "Test1", "propInt32Array", "const QList<qint32>& propInt32Array"},
		{"test", "Test1", "propInt64Array", "const QList<qint64>& propInt64Array"},
		{"test", "Test1", "propFloatArray", "const QList<qreal>& propFloatArray"},
		{"test", "Test1", "propFloat32Array", "const QList<float>& propFloat32Array"},
		{"test", "Test1", "propFloat64Array", "const QList<double>& propFloat64Array"},
		{"test", "Test1", "propStringArray", "const QList<QString>& propStringArray"},
		{"test", "Test1", "prop_Bool", "bool prop_Bool"},
		{"test", "Test1", "prop_bool", "bool prop_bool"},
		{"test", "Test1", "prop_1", "bool prop_1"},
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
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Enum1::Enum1Enum propEnum"},
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

func TestParamWithPrefix_prefixIsIgnoredForBuiltInTypes(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "bool propBool"},
		{"test", "Test1", "propInt", "int propInt"},
		{"test", "Test1", "propInt32", "qint32 propInt32"},
		{"test", "Test1", "propInt64", "qint64 propInt64"},
		{"test", "Test1", "propFloat", "qreal propFloat"},
		{"test", "Test1", "propFloat32", "float propFloat32"},
		{"test", "Test1", "propFloat64", "double propFloat64"},
		{"test", "Test1", "propString", "const QString& propString"},
		{"test", "Test1", "propBoolArray", "const QList<bool>& propBoolArray"},
		{"test", "Test1", "propIntArray", "const QList<int>& propIntArray"},
		{"test", "Test1", "propInt32Array", "const QList<qint32>& propInt32Array"},
		{"test", "Test1", "propInt64Array", "const QList<qint64>& propInt64Array"},
		{"test", "Test1", "propFloatArray", "const QList<qreal>& propFloatArray"},
		{"test", "Test1", "propFloat32Array", "const QList<float>& propFloat32Array"},
		{"test", "Test1", "propFloat64Array", "const QList<double>& propFloat64Array"},
		{"test", "Test1", "propStringArray", "const QList<QString>& propStringArray"},
	}
	prefix := "my_prefix::"
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := qtParam(prefix, prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamSymbolsWithPrefix(t *testing.T) {
	t.Parallel()
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "my_prefix::Enum1::Enum1Enum propEnum"},
		{"test", "Test2", "propStruct", "const my_prefix::Struct1& propStruct"},
		{"test", "Test2", "propInterface", "my_prefix::Interface1 *propInterface"},
		{"test", "Test2", "propEnumArray", "const QList<my_prefix::Enum1::Enum1Enum>& propEnumArray"},
		{"test", "Test2", "propStructArray", "const QList<my_prefix::Struct1>& propStructArray"},
		{"test", "Test2", "propInterfaceArray", "const QList<my_prefix::Interface1*>& propInterfaceArray"},
	}
	prefix := "my_prefix::"
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := qtParam(prefix, prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
