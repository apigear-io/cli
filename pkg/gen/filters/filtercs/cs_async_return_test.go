package filtercs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsyncReturn(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opVoid", "Task"},
		{"test", "Test3", "opBool", "Task<bool>"},
		{"test", "Test3", "opInt", "Task<int>"},
		{"test", "Test3", "opInt32", "Task<int>"},
		{"test", "Test3", "opInt64", "Task<long>"},
		{"test", "Test3", "opFloat", "Task<float>"},
		{"test", "Test3", "opFloat64", "Task<double>"},
		{"test", "Test3", "opString", "Task<string>"},
		{"test", "Test3", "opBoolArray", "Task<List<bool>>"},
		{"test", "Test3", "opIntArray", "Task<List<int>>"},
		{"test", "Test3", "opStringArray", "Task<List<string>>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := csAsyncReturn("", op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestAsyncReturnSymbols(t *testing.T) {
	t.Parallel()
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test4", "opEnum", "Task<Enum1>"},
		{"test", "Test4", "opStruct", "Task<Struct1>"},
		{"test", "Test4", "opInterface", "Task<IInterface1>"},
		{"test", "Test4", "opEnumArray", "Task<List<Enum1>>"},
		{"test", "Test4", "opStructArray", "Task<List<Struct1>>"},
		{"test", "Test4", "opInterfaceArray", "Task<List<IInterface1>>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := csAsyncReturn("", op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
