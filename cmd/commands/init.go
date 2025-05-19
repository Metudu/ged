package commands

import (
	"log"

	"github.com/Metudu/ged/internal/db"
	"github.com/urfave/cli/v2"
)

var Init *cli.Command = &cli.Command{
	Name: "init",
	Usage: "get ged ready to use",
	Action: InitAction,
}

func InitAction(ctx *cli.Context) error {
	db.Initialize()
	log.Println("Initialized!")
	return nil
}