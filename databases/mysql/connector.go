package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var MysqlDb *gorm.DB

func init() {

	dbStr := "root:******@tcp(*********:3306)/my_db_01?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	MysqlDb, err = gorm.Open(mysql.Open(dbStr), &gorm.Config{})
	if err != nil {

		log.Println(err)
		panic(err)
	}
	s, _ := MysqlDb.DB()

	// 设置连接池，空闲连接
	s.SetMaxIdleConns(50)
	// 打开链接
	s.SetMaxOpenConns(100)
}
