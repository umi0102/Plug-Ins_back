package app

import (
	"Plug-Ins/middlewares"
	"Plug-Ins/routerMethods/projects"
	"Plug-Ins/routerMethods/users"
	"github.com/gin-gonic/gin"
	"reflect"
)

func RouterService() {
	router := gin.Default()

	router.Use(func(context *gin.Context) {
		defer func() {
			if a := recover(); a != nil {
				j := make(map[string]interface{}, 0)
				if reflect.TypeOf(a) != reflect.TypeOf(j) {
					j["code"] = "-1"
					j["msg"] = a
					context.JSON(500, j)
					return
				}
				context.JSON(500, a)
			}
		}()
		context.Next()
	})
	router.Use(middlewares.GetRequestIP())
	router.Use(middlewares.Cors())
	//router.Use(middlewares.InterceptRequests())
	//权限路由
	api := router.Group("/api")
	{
		api.Use(middlewares.AuthRequired())
		api.POST("/addProject", projects.AddProject)       //新增项目
		api.POST("/developer", projects.GetDeveloperList)  //获取项目人员
		api.POST("/joinProject", projects.JoinProject)     //加入项目
		api.POST("/leaveProject", projects.LeaveProject)   //退出项目/删除成员
		api.GET("/user/projects", projects.GetProjectList) //项目列表
		api.POST("/checkToken", projects.CheckToken)       //验证token是否可用
		api.POST("/userinfo", users.GetUserinfo)
	}

	//非权限路由
	group := router.Group("/user")
	{
		group.POST("/regist", users.Regist)           //注册
		group.POST("/login", users.LoginJwt)          //登录
		group.POST("/sendcode", users.QueryByPhone)   //发送验证码
		group.POST("/loginbycode", users.LoginByCode) //验证码登陆

	}

	router.Run(":8080")
}
