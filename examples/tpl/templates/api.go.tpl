package api


{{- range .Module.Enums }}

type {{.Name}} int
{{- $enum := .Name }}
{{- range .Members }}
const (
    {{.Name}} {{$enum}} = {{.Value}}
)
{{- end }}
{{- end }}

{{- range .Module.Structs }}

type {{.Name}} struct {
{{- range .Fields }}
    {{Camel .Name}} {{goType .}}
{{- end }}
}
{{- end }}

{{- range .Module.Interfaces }}

type {{Camel .Name }} interface {
{{- range .Properties }}
    {{Camel .Name}}() {{goType "" . }}
    Set{{Camel .Name}}({{goParam "" .}})    
{{- end }}
{{- range .Operations }}
    {{Camel .Name}}({{goParams "" .Params}}) {{goReturn "" .Return }}
{{- end }}
{{- range .Signals }}
    On{{Camel .Name}}(func ({{goParams "" .Params}}))
{{- end }}
}
{{- end }}