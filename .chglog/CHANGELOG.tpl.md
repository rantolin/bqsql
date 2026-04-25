{{ if .Versions -}}
# CHANGELOG

All notable changes to this project will be documented in this file.

{{ range .Versions }}
<a name="{{ .Tag.Name }}"></a>
## {{ .Tag.Name }} ({{ .Tag.Date.Format "2006-01-02" }})

{{ range .CommitGroups -}}
### {{ .Title }}

{{ range .Commits -}}
- {{ if .Scope }}**{{ .Scope }}:** {{ end }}{{ .Subject }}{{ if .Hash }} ([{{ .Hash.Short }}]({{ $.Info.RepositoryURL }}/commit/{{ .Hash.Long }})){{ end }}
{{ end }}
{{ end -}}

{{- if .NoteGroups -}}
{{ range .NoteGroups -}}
### {{ .Title }}

{{ range .Notes }}
{{ .Body }}
{{ end }}
{{ end -}}
{{ end -}}
{{ end -}}
{{ end -}}
