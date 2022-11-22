package main

import (
	"Plug-Ins/app"
	_ "Plug-Ins/databases/mysql"
	_ "Plug-Ins/databases/redisServer"
)

func main() {
	app.Create()
	app.RouterService()

}
