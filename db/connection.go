package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mysqlDB *gorm.DB

func GetMySQLDB() *gorm.DB {
	if mysqlDB == nil {
		var err error
		mysqlDB, err = gorm.Open("mysql", "root@/orderAppApi?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			panic("failed to connect to MySQL database")
		}
	}
	return mysqlDB
}
