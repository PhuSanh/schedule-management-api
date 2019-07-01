package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

type MysqlConn struct {
	*gorm.DB
}

var instance *MysqlConn
var once sync.Once

func NewConn(username, password, database string, maxIdle, maxOpen int) (*MysqlConn, error) {
	var err error = nil
	once.Do(func() {
		params := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, database)
		db, err := gorm.Open("mysql", params)
		if err != nil {
			return
		}

		db.SingularTable(true)
		db.DB().SetMaxIdleConns(maxIdle)
		db.DB().SetMaxOpenConns(maxOpen)
		instance = &MysqlConn{db}
	})
	return instance, err
}

func (conn *MysqlConn) Stop() error {
	if conn.DB != nil {
		err := conn.DB.Close()
		if err != nil {
			return err
		} else {
			conn.DB = nil
			return nil
		}
	} else {
		return nil
	}
}