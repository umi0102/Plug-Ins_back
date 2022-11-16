package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func OpenSql() {
	db, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/my_db_01")
	pingErr := db.Ping()
	if pingErr != nil {
		panic("pingErr")
	}
	fmt.Println("sql：ok")
}
