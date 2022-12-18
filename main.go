package main

import (
	"Plug-Ins/app"
	_ "Plug-Ins/databases/mysql"
	_ "Plug-Ins/databases/redisServer"
)

func main() {
	//runtime.GOMAXPROCS(8)
	app.Create()
	//启动socket协程

	//ws.Wsconnect() //启动websocket

	app.RouterService()

}
