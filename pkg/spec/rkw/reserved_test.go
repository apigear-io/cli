package rkw

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsKeywordReserved(t *testing.T) {
	for _, keyword := range pyReservedKeywords() {
		langs, ok := IsKeywordReserved(keyword)
		require.True(t, ok, "Expected %s to be reserved", keyword)
		require.Contains(t, langs, PY)

	}
	for _, keyword := range cppReservedKeywords() {
		langs, ok := IsKeywordReserved(keyword)
		require.True(t, ok, "Expected %s to be reserved", keyword)
		require.Contains(t, langs, CPP)
	}
}

func TestIsKeywordReservedInLang(t *testing.T) {
	type test struct {
		kw    string
		ok    bool
		langs []Lang
	}

	tests := []test{
		{kw: "else", ok: true, langs: []Lang{CPP, PY, TS, JS, GO, UE, QT}},
	}

	for _, tc := range tests {
		langs, ok := IsKeywordReserved(tc.kw)
		require.Equal(t, tc.ok, ok)
		require.ElementsMatch(t, tc.langs, langs)
	}
}
