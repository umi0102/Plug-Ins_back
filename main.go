package main

import (
	"Plug-Ins/app"
	_ "Plug-Ins/databases/mysql"
	_ "Plug-Ins/databases/redisServer"
	"Plug-Ins/ws"
)

func main() {
	app.Create()
	ws.Wsconnect()
	app.RouterService()

}
