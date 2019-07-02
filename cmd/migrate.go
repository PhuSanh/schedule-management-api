package cmd

import (
	"github.com/urfave/cli"
	"schedule-management-api/database"
	"schedule-management-api/model"
	"schedule-management-api/setting"
)

func migrateAction(appContext *cli.Context) {
	input := appContext.String("table")
	_ = setting.InitMysql()
	switch input {
	case "user":
		database.MysqlConn.AutoMigrate(&model.User{})
	case "user-category":
		database.MysqlConn.AutoMigrate(&model.UserCategory{})
	}
	defer database.MysqlConn.Stop()
}

var Migrate = cli.Command{
	Name: "migrate",
	Usage: "migrate db",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "table, tb",
			Usage: "List the name of table should be run",
		},
	},
	Action: func(appContext *cli.Context) error {
		migrateAction(appContext)
		return nil
	},
}