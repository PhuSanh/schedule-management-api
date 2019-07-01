package setting

import (
	"log"
	"schedule-management-api/config"
	"schedule-management-api/database"
	"schedule-management-api/database/mysql"
)

var cfg = config.GetConfig()

func InitMysql() (err error) {
	log.Println(cfg.Name)
	database.MysqlConn, err = mysql.NewConn(cfg.Mysql.Username, cfg.Mysql.Password, cfg.Mysql.DatabaseName, cfg.Mysql.MaxIdleConn, cfg.Mysql.MaxOpenConn)
	if err != nil {
		log.Fatal("Connect database fail")
	} else {
		log.Println("Connect database success")
	}
	return
}

//func InitRedis() (err error) {
//	database.RedisConn, err = redis.NewConn(cfg.Redis.IP, cfg.Redis.Port, cfg.Redis.Password)
//	if err != nil {
//		logger.Error("Connect redis fail")
//	} else {
//		logger.Info("Connect redis success", logger.LogFields{})
//	}
//	return
//}
//
//func InitLogger() (err error) {
//	logger.NewLogger()
//	return
//}