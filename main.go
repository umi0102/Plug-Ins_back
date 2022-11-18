package main

import (
	"Plug-Ins/App"
	_ "Plug-Ins/databases/mysql"
	_ "Plug-Ins/databases/redisServer"
)

func main() {
	App.Create()
	App.RouterService()

}
