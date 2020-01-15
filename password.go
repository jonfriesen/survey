package survey

import (
	"fmt"

	"github.com/tylerflint/survey/core"
	"github.com/tylerflint/survey/terminal"
)

/*
Password is like a normal Input but the text shows up as *'s and there is no default. Response
type is a string.

	password := ""
	prompt := &survey.Password{ Message: "Please type your password" }
	survey.AskOne(prompt, &password, nil)
*/
type Password struct {
	core.Renderer
	Message string
	Help    string
}

type PasswordTemplateData struct {
	Password
	ShowHelp bool
	Done     bool
}

// Templates with Color formatting. See Documentation: https://github.com/mgutz/ansi#style-format
var PasswordQuestionTemplate = `
{{- if .ShowHelp }}{{- color "cyan"}}{{ HelpIcon }} {{ .Help }}{{color "reset"}}{{"\n"}}{{end}}
{{- if .Done}}
	{{- color "green+h"}}{{ DoneIcon }} {{color "reset"}}
{{- else }}
	{{- color "green+hb"}}{{ InputIcon }} {{color "reset"}}
{{- end}}
{{- if not .Done}}{{- color "default+hb"}}{{ .Message }}{{color "reset"}}
{{- else}}{{- color "default+h"}}{{ .Message }}{{color "reset"}}{{end}}
{{- if and .Help (not .ShowHelp)}}{{color "cyan"}}[{{ HelpInputRune }} for help]{{color "reset"}} {{end}}
{{- if not .Done}}: {{color "cyan"}}
{{- else}}{{color "reset"}}{{end}}`

func (p *Password) Prompt() (line interface{}, err error) {
	// render the question template
	out, err := core.RunTemplate(
		PasswordQuestionTemplate,
		PasswordTemplateData{Password: *p},
	)
	fmt.Fprint(terminal.NewAnsiStdout(p.Stdio().Out), out)
	if err != nil {
		return "", err
	}

	rr := p.NewRuneReader()
	rr.SetTermMode()
	defer rr.RestoreTermMode()

	// no help msg?  Just return any response
	if p.Help == "" {
		line, err := rr.ReadLine('*')
		return string(line), err
	}

	cursor := p.NewCursor()

	// process answers looking for help prompt answer
	for {
		line, err := rr.ReadLine('*')
		if err != nil {
			return string(line), err
		}

		if string(line) == string(core.HelpInputRune) {
			// terminal will echo the \n so we need to jump back up one row
			cursor.PreviousLine(1)

			err = p.Render(
				PasswordQuestionTemplate,
				PasswordTemplateData{Password: *p, ShowHelp: true},
			)
			if err != nil {
				return "", err
			}
			continue
		}
		return string(line), err
	}
}

// Cleanup hides the string with a fixed number of characters.
func (p *Password) Cleanup(val interface{}) error {
	cursor := p.NewCursor()
	cursor.PreviousLine(1)

	return p.Render(
		PasswordQuestionTemplate,
		PasswordTemplateData{Password: *p, Done: true},
	)
}
