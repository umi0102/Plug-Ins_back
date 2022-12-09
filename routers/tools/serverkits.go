package tools

import (
	"Plug-Ins/databases/mysql"
	"Plug-Ins/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserinfo 获取个人信息
func GetUserinfo(ctx *gin.Context) {
	get, _ := ctx.Get("phone")
	res := mysql.SelectMysql(fmt.Sprintf(`select userinfo_name,userinfo_usericon,userinfo_phone,userinfo_name from userinfos where userinfo_phone = "%s"`, get))
	iconPath := res["userinfo_usericon"].(string)

	image, err := routers.GetImage(iconPath)
	if err != nil {
		fmt.Println("2121")
		return
	}
	res["userinfo_usericon"] = image
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
func GetServeList(ctx *gin.Context) {
	res := mysql.SelectAllMysql(`SELECT * FROM kits_server`)
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})

}
