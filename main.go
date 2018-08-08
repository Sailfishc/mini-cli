package main

import (
	"fmt"
	"os"

	"./api"
	"github.com/mkideal/cli"
)

const SailfishVersion = "1.0"

type App struct {
}

// root Command
type rootT struct {
	cli.Helper
	Version bool `cli:"v, version" usage:"display mini-cli version"`
}

var root = &cli.Command{
	Desc: "root of mini-cli",
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		if len(ctx.FormValues()) == 0 {
			ctx.WriteUsage()
			return nil
		}
		argv := ctx.Argv().(*rootT)
		if argv.Version {
			ctx.String("\n" + SailfishVersion + "\n")
		}
		return nil
	},
}

func main() {
	if err := cli.Root(root,
		cli.Tree(cli.HelpCommand("display help information")),
		cli.Tree(api.New()),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
