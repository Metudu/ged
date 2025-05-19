package commands

import (
	"log"

	"github.com/Metudu/ged/internal/db"
	"github.com/urfave/cli/v2"
)

var Refresh *cli.Command = &cli.Command{
	Name: "refresh",
	Usage: "refresh the ged database",
	Action: RefreshAction,
}

func RefreshAction(ctx *cli.Context) error {
	db.Refresh()
	log.Println("Refreshed!")
	return nil
}