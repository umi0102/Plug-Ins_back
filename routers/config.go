package routers

import (
	"Plug-Ins/middlewares"
	"github.com/gin-gonic/gin"
)

func RouterService() {
	router := gin.Default()
	router.Use(middlewares.Cors())
	//权限路由
	api := router.Group("/api")
	api.Use(middlewares.AuthRequired())
	api.POST("/addProject", AddProject)      //新增项目
	api.POST("/developer", GetDeveloperList) //获取项目人员
	api.POST("/joinProject", JoinProject)    //加入项目
	api.POST("/leaveProject", LeaveProject)  //退出项目/删除成员
	//非权限路由
	router.POST("/user/regist", Regist)              //注册
	router.POST("/user/login", middlewares.LoginJwt) //登录
	api.GET("/user/projects", GetProjectList)        //项目列表

	router.Run(":8080")
}
