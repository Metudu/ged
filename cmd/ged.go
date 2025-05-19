package cmd

import (
	"log"
	"os"

	"github.com/Metudu/ged/cmd/commands"
	"github.com/urfave/cli/v2"
)

var ged *cli.App = &cli.App{
	EnableBashCompletion: true,
	Commands: []*cli.Command{
		commands.Init,
		commands.Refresh,
		commands.List,
		commands.Show,
		commands.Hide,
	},
}

func Execute() {
	if err := ged.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}