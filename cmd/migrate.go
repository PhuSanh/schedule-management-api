package cmd

import (
	"fmt"
	"github.com/urfave/cli"
	"schedule-management-api/database"
	"schedule-management-api/model"
	"schedule-management-api/setting"
)

var tables = map[string]interface{} {
	"user": 			&model.User{},
	"user-category":	&model.UserCategory{},
	"schedule": 		&model.Schedule{},
	"group": 			&model.Group{},
	"user-group": 		&model.UserGroup{},
}

func migrateAction(appContext *cli.Context) {
	input := appContext.String("table")
	if _, ok := tables[input]; !ok {
		fmt.Println("Table " + input + " not found")
		return
	}
	_ = setting.InitMysql()
	database.MysqlConn.AutoMigrate(tables[input])
	fmt.Println("Migrate table " + input + " success")
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