package commands

import (
	"errors"
	"fmt"
	"log"

	"github.com/Metudu/ged/internal/db"
	"github.com/urfave/cli/v2"
)

var Hide *cli.Command = &cli.Command{
	Name: "show",
	Usage: "make a launcher visible",
	Action: ShowAction,
	BashComplete: HideComplete,
}

func HideComplete(ctx *cli.Context) {
	if ctx.NArg() > 0 {
		return
	}

	for _, file := range db.GetNames() {
		fmt.Println(file)
	}
}

func HideAction(ctx *cli.Context) error {
	if ctx.Args().Len() == 0 {
		return errors.New("at least one name must be passed as argument")
	}

	for _, name := range ctx.Args().Slice() {
		db.Hide(name)
		log.Printf("Visibility of %s is set to true!\n", name)
	}

	db.List()
	return nil
}