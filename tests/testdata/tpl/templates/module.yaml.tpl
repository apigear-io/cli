template: {{.Meta.Layer.Template}}
name: {{.Module.Name}}
version: {{.Module.Version}}
interfaces:
{{- range .Module.Interfaces }}
    - name: {{.Name}}
{{- end }}


{{- range .System.Modules }}
{{- $module := . }}
{{- range .Interfaces }}
{{- $interface := . }}
{{- range .Operations }}
{{- $operation := . }}
mod: {{ $module.Id }} / int: {{ $interface.Id }} / op: {{ $operation.Id }}
{{- end }}
{{- range .Properties }}
{{- $property := . }}
mod: {{ $module.Id }} / int: {{ $interface.Id }} / prop: {{ $property.Id }}
{{- end }}
{{- range .Signals }}
{{- $signal := . }}
mod: {{ $module.Id }} / int: {{ $interface.Id }} / sig: {{ $signal.Id }}
{{- end }}
{{- end }}
{{- range .Structs }}
{{- $structs := . }}
{{- end }}
{{- end }}