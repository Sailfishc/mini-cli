package api

import (
	"./templates"
	"github.com/AlecAivazis/survey"
	"github.com/mkideal/cli"
)

// New new mini-cli application
func New() *cli.Command {
	return &cli.Command{
		Name:        "new",
		Desc:        "create maven application by template",
		Text:        `    sailfish new <name>`,
		Argv:        func() interface{} { return new(newT) },
		CanSubRoute: true,
		Fn: func(ctx *cli.Context) error {
			argv := ctx.Argv().(*newT)
			argv.Version = "1.0"

			prompt := &survey.Input{
				Message: "please input package name(e.g: com.sailfish.example):",
			}
			survey.AskOne(prompt, &argv.PackageName, nil)

			buildPrompt := &survey.Select{
				Message: "select the build tool:",
				Options: []string{"Maven", "Gradle"},
			}
			survey.AskOne(buildPrompt, &argv.BuildTool, nil)
			return templates.New(ctx, argv.BaseConfig)
		},
	}
}

type newT struct {
	cli.Helper
	templates.BaseConfig
}
