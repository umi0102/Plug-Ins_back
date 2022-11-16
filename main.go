package main

import (
	"Plug-Ins/databases/mysql"
	"Plug-Ins/databases/redis"
	"Plug-Ins/routers"
)

func main() {
	redis.Setup()
	mysql.OpenSql()

	routers.RouterService()
}
