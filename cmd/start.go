package cmd

import (
	"github.com/urfave/cli"
	"schedule-management-api/router"
)

var Start = cli.Command{
	Name: "start",
	Usage: "Start server",
	Action: func(appContext *cli.Context) (err error) {
		router.InitRouter()
		return
	},
}
