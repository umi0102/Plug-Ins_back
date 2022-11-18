package App

import (
	"Plug-Ins/middlewares"
	"Plug-Ins/routerMethods/projects"
	"Plug-Ins/routerMethods/users"
	"github.com/gin-gonic/gin"
)

func RouterService() {
	router := gin.Default()
	router.Use(func(context *gin.Context) {
		defer func() {
			if a := recover(); a != nil {
				context.JSON(500, a)
			}
		}()
		context.Next()
	})
	router.Use(middlewares.Cors())

	//权限路由
	api := router.Group("/api")
	{
		api.Use(middlewares.AuthRequired())
		api.POST("/addProject", projects.AddProject)       //新增项目
		api.POST("/developer", projects.GetDeveloperList)  //获取项目人员
		api.POST("/joinProject", projects.JoinProject)     //加入项目
		api.POST("/leaveProject", projects.LeaveProject)   //退出项目/删除成员
		api.GET("/user/projects", projects.GetProjectList) //项目列表
	}

	//非权限路由

	router.POST("/user/regist", users.Regist)  //注册
	router.POST("/user/login", users.LoginJwt) //登录

	router.Run(":8080")
}
