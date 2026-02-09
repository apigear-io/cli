{{- range $key, $value := .Features }}
{{$key}}: {{$value}}
{{- end }}