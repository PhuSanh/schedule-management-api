package database

import (
	"schedule-management-api/database/mysql"
	"schedule-management-api/database/redis"
)

var (
	MysqlConn *mysql.MysqlConn
	RedisConn *redis.RedisConn
)