name: {{.Module.Name}}
version: {{.Module.Version}}
interfaces:
{{- range .Module.Interfaces }}
    - name: {{.Name}}
{{- end }}
