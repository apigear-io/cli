package spec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDocumentType(t *testing.T) {
	tests := []struct {
		name string
		want DocumentType
	}{
		{
			name: "demo.idl",
			want: DocumentTypeModule,
		},
		{
			name: "demo.module.yaml",
			want: DocumentTypeModule,
		},
		{
			name: "demo.module.json",
			want: DocumentTypeModule,
		},
		{
			name: "demo.solution.yaml",
			want: DocumentTypeSolution,
		},
		{
			name: "demo.solution.json",
			want: DocumentTypeSolution,
		},
		{
			name: "rules.yaml",
			want: DocumentTypeRules,
		},
		{
			name: "rules.json",
			want: DocumentTypeRules,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDocumentType(tt.name)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
