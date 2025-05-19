package commands

import (
	"github.com/Metudu/ged/internal/db"
	"github.com/urfave/cli/v2"
)

var List *cli.Command = &cli.Command{
	Name: "list",
	Usage: "list the desktop files",
	Action: ListAction,
}

func ListAction(ctx *cli.Context) error {
	db.Refresh()
	db.List()
	return nil
}