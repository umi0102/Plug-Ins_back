package mysql

import (
	"log"
)

// SelectMysql 查询
func SelectMysql(sqlStr string) map[string]interface{} {
	data := make(map[string]interface{}, 0)
	scan := MysqlDb.Raw(sqlStr).Scan(&data)
	if scan.Error != nil {
		log.Println(scan.Error.Error())
		panic(scan.Error.Error())
	}
	return data

}

// InsUpdDelMysql  插入，修改，删除
func InsUpdDelMysql(sqlStr string) {
	tx := MysqlDb.Exec(sqlStr)
	if tx.Error != nil {
		panic(tx.Error.Error())
	}
	if tx.RowsAffected == 0 {
		panic("插入,修改,删除操作错误")
	}

}

// CreateTableMysql 创建表
func CreateTableMysql(tables ...interface{}) {

	migrator := MysqlDb.Migrator()
	for _, v := range tables {
		err := migrator.AutoMigrate(&v)
		if err != nil {
			panic(err)
		}
	}
}
