package app_template

const AIWorkGenerationEmailHTML = `
{{.Title}}<br><br>
{{.Message}}<br><br>
{{.ButtonText}}: <a href="{{.Link}}">{{.Link}}</a>
`

const AIWorkGenerationEmailPlain = `
{{.Title}}

{{.Message}}

{{.ButtonText}}: {{.Link}}
`
