package main

import (
	"log"

	"github.com/Metudu/ged/cmd"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	cmd.Execute()
}