template: {{.Meta.Layer.Template}}
name: {{.Module.Name}}
version: {{.Module.Version}}

interfaces:
{{- range .Module.Interfaces }}
    - name: {{.Name}}
      id: {{.Id}}
      properties:
{{- range .Properties }}
        - name: {{.Name}}
          id: {{.Id}}
{{- end }}
      operations:
{{- range .Operations }}
        - name: {{.Name}}
          id: {{.Id}}
{{- end }}
      signals:
{{- range .Signals }}
        - name: {{.Name}}
          id: {{.Id}}
{{- end }}
{{- end }}

