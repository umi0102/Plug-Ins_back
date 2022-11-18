package main

import (
	"Plug-Ins/App"
	_ "Plug-Ins/databases/mysql"
	_ "Plug-Ins/databases/redis"
)

func main() {
	App.Create()
	App.RouterService()

}
