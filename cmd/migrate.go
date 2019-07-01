package cmd

import (
	"github.com/urfave/cli"
	"schedule-management-api/database"
	"schedule-management-api/model"
	"schedule-management-api/setting"
)

func migrateAction(appContext *cli.Context) {
	setting.InitMysql()
	database.MysqlConn.AutoMigrate(&model.User{})
	defer database.MysqlConn.Stop()
}

var Migrate = cli.Command{
	Name: "migrate",
	Usage: "migrate db",
	Action: func(appContext *cli.Context) error {
		migrateAction(appContext)
		return nil
	},
}