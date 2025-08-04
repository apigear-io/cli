module {{.Module.Name}} {{.Module.Version}}
{{- range .Module.Enums }}

enum {{Camel .Name}} {
    {{- range .Members }}
    {{.Name}} = {{.Value}}
    {{- end}}
}
{{- end }}

{{- range .Module.Structs }}

struct {{Camel .Name}} {
    {{- range .Fields -}}
    {{.Name}}: {{.Type}}
    {{- end}}
}
{{- end }}

{{- range .Module.Interfaces }}

interface {{Camel .Name}} {
    {{- range .Properties }}
    {{.Name}}: {{.Type}}
    {{- end}}
    {{- range .Operations }}
    {{.Name}}({{range $index, $param := .Params}}{{if $index}}, {{end}}{{.Type}} {{.Name}}{{end}}): {{ .Return.Type }}
    {{- end}}
    {{- range .Signals }}
    signal {{.Name}}({{range $index, $param := .Params}}{{if $index}}, {{end}}{{.Type}} {{.Name}}{{end}})
    {{- end}}
}
{{- end -}}