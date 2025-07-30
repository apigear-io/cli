module {{.Module.Name}} {{.Module.Version}}
{{- range .Module.Enums }}

enum {{Camel .Name}} {
    {{- range .Members }}
    {{Camel .Name}} = {{.Value}}
    {{- end}}
}
{{- end }}

{{- range .Module.Structs }}

struct {{Camel .Name}} {
    {{- range .Fields -}}
    {{camel .Name}}: {{.Type}}
    {{- end}}
}
{{- end }}

{{- range .Module.Interfaces }}

interface {{Camel .Name}} {
    {{- range .Properties }}
    {{camel .Name}}: {{.Type}}
    {{- end}}
    {{- range .Operations }}
    {{camel .Name}}({{range $index, $param := .Params}}{{if $index}}, {{end}}{{.Type}} {{camel .Name}}{{end}}): {{.Return}}
    {{- end}}
    {{- range .Signals }}
    signal {{camel .Name}}({{range $index, $param := .Params}}{{if $index}}, {{end}}{{.Type}} {{camel .Name}}{{end}})
    {{- end}}
}
{{- end -}}