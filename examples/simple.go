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
	{
		Name: "band",
		Prompt: &survey.Select{
			Message: "What is your favorite band",
			Options: []string{
				"Beatles",
				"Nsync",
				"The Monkeys",
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
	{
		Name: "secret",
		Prompt: &survey.Password{
			Message: "What's your super secret API key",
		},
	},
	{
		Name: "crazy",
		Prompt: &survey.Confirm{
			Message: "You're a bit crazy, right",
			Default: true,
		},
	},
}

func main() {
	answers := struct {
		Name  		string
		Color 		string
		Band			string
		Password 	string
		Nickname	string
		Crazy			bool
		Secret		string
	}{}

	// ask the question
	err := survey.Ask(simpleQs, &answers)
	
	fmt.Printf("\n\n\n")
	fmt.Println(answers)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
