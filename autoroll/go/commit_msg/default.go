package commit_msg

import "text/template"

const (
	// TmplNameDefault is the name of the default commit message template which
	// is suitable for most rollers.
	TmplNameDefault = "default"
)

var (
	// tmplNameCommitMsg is the name of the overall commit message template.
	tmplNameCommitMsg = "commitMsg"

	// tmplCommitMsg utilizes the skeleton with the default block definitions.
	// It is the primary entry point which is executed for every commit message.
	// Custom commit messages may modify it by overriding any of the blocks
	// defined above, including the commitMsg itself for a completely custom
	// message.
	tmplCommitMsg = template.Must(parseCommitMsgTemplate(tmplFooterDefault, tmplNameCommitMsg, `
{{- define "subject" }}{{ template "defaultSubject" . }}{{ end -}}
{{- define "manual" }}{{ template "defaultManualMessage" . }}{{ end -}}
{{- define "revisions" }}{{ template "defaultRevisions" . }}{{ end -}}
{{- define "boilerplate" }}{{ template "defaultBoilerplate" . }}{{ end -}}
{{- define "footer" }}{{ template "defaultFooter" . }}{{ end -}}
{{ template "skeleton" . }}`))

	tmplNameSkeleton = "skeleton"
	// tmplSkeleton defines the basic structure for all commit messages.
	tmplSkeleton = template.Must(parseCommitMsgTemplate(nil, tmplNameSkeleton,
		`{{ template "subject" . }}

{{ template "manual" . }}

{{ template "revisions" . }}

{{ template "boilerplate" . }}

{{ template "footer" . }}
`))

	tmplNameSubjectDefault = "defaultSubject"
	tmplSubjectDefault     = template.Must(parseCommitMsgTemplate(tmplSkeleton, tmplNameSubjectDefault,
		`{{ if .ManualRollRequester }}Manual r{{ else }}R{{ end }}oll {{ .ChildName }} from {{ .RollingFrom }} to {{ .RollingTo }}{{ if .IncludeRevisionCount}} ({{ len .Revisions }} revision{{ if gt (len .Revisions) 1 }}s{{ end }}){{ end }}`))

	tmplNameManualDefault = "defaultManualMessage"
	tmplManualDefault     = template.Must(parseCommitMsgTemplate(tmplSubjectDefault, tmplNameManualDefault,
		`{{ if .ManualRollRequester }}Manual roll requested by {{.ManualRollRequester}}{{ end }}`))

	tmplNameRevisionsDefault = "defaultRevisions"
	tmplRevisionsDefault     = template.Must(parseCommitMsgTemplate(tmplManualDefault, tmplNameRevisionsDefault,
		`{{ if .ChildLogURL }}{{ .ChildLogURL }}

{{ end -}}
{{- if .IncludeLog -}}
{{ range .Revisions }}{{ .Timestamp.Format "2006-01-02" }} {{ .Author }} {{ .Description }}
{{ end }}
{{ end -}}
{{ if len .TransitiveDeps -}}
Also rolling transitive DEPS:
{{ range .TransitiveDeps }}  {{ .String }}
{{ end }}
{{- end }}`))

	tmplNameBoilerplateDefault = "defaultBoilerplate"
	tmplBoilerplateDefault     = template.Must(parseCommitMsgTemplate(tmplRevisionsDefault, tmplNameBoilerplateDefault,
		`If this roll has caused a breakage, revert this CL and stop the roller
using the controls here:
{{.ServerURL}}
Please CC {{stringsJoin (mergeNoDuplicates .Reviewers .Contacts) ","}} on the revert to ensure that a human
is aware of the problem.

{{ if .ChildBugLink -}}To file a bug in {{ .ChildName }}: {{ .ChildBugLink }}{{ end }}
{{ if .ParentBugLink -}}To file a bug in {{ .ParentName }}: {{ .ParentBugLink }}{{ end }}

To report a problem with the AutoRoller itself, please file a bug:
https://issues.skia.org/issues/new?component=1389291&template=1850622

Documentation for the AutoRoller is here:
https://skia.googlesource.com/buildbot/+doc/main/autoroll/README.md`))

	tmplNameFooterDefault = "defaultFooter"
	tmplFooterDefault     = template.Must(parseCommitMsgTemplate(tmplBoilerplateDefault, tmplNameFooterDefault,
		`{{ if .CqExtraTrybots -}}
Cq-Include-Trybots: {{ stringsJoin .CqExtraTrybots ";" }}
{{ end -}}
{{ if .CqDoNotCancelTrybots -}}
Cq-Do-Not-Cancel-Tryjobs: true
{{ end -}}
{{ if .BugProject -}}
Bug: {{ if .Bugs }}{{ stringsJoin .Bugs "," }}{{ else }}None{{ end }}
{{ end -}}
{{ if .IncludeTbrLine -}}
Tbr: {{ stringsJoin .Reviewers "," }}
{{ end -}}
{{ if .IncludeTests -}}
{{ range .Tests }}Test: {{.}}
{{ end -}}
{{- end -}}
{{ if .ExtraFooters -}}
{{ range .ExtraFooters }}{{.}}
{{ end -}}
{{- end -}}
`))
)
