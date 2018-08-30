package main

import (
	"fmt"

	"github.com/tylerflint/survey"
)

// the questions to ask
var simpleQs = []*survey.Question{
	{
		Name: "name",
		Prompt: &survey.Input{
			Message: "What is your name",
		},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "nickname",
		Prompt: &survey.Input{
			Message: "What is your nickname",
		},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "color",
		Prompt: &survey.Select{
			Message: "How would you like to build the runtime",
			Options: []string{
				"Automatically with Heroku buildpacks",
				"Manually with a custom Dockerfile",
			},
		},
		Validate: survey.Required,
	},
	// {
	// 	Name: "password",
	// 	Prompt: &survey.Password{
	// 		Message: "Password",
	// 	},
	// },
}

func main() {
	answers := struct {
		Name  		string
		Color 		string
		Password 	string
		Nickname	string
	}{}

	// ask the question
	err := survey.Ask(simpleQs, &answers)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
