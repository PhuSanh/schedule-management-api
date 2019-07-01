package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
	"schedule-management-api/cmd"
	_ "schedule-management-api/config"
)

func main() {

	app := cli.NewApp()

	app.Name = "schedule management"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		cmd.Migrate,
		cmd.Start,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
