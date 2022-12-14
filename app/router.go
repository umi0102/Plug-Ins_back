package app

import (
	"Plug-Ins/middlewares"
	"Plug-Ins/routers"
	"Plug-Ins/routers/projects"
	"Plug-Ins/routers/tools"
	"Plug-Ins/routers/users"
	"Plug-Ins/ws"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"reflect"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

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

	router.Use(middlewares.Cors())

	//权限路由
	api := router.Group("/api")
	{
		api.Use(middlewares.AuthRequired())
		api.POST("/addProject", middlewares.InterceptRequests(100), projects.AddProject)       //新增项目
		api.POST("/developer", middlewares.InterceptRequests(100), projects.GetDeveloperList)  //获取项目人员
		api.POST("/joinProject", middlewares.InterceptRequests(100), projects.JoinProject)     //加入项目
		api.POST("/leaveProject", middlewares.InterceptRequests(100), projects.LeaveProject)   //退出项目/删除成员
		api.GET("/user/projects", middlewares.InterceptRequests(100), projects.GetProjectList) //项目列表
		api.POST("/checkToken", middlewares.InterceptRequests(100), projects.CheckToken)       //验证token是否可用
		api.POST("/userinfo", middlewares.InterceptRequests(100), users.GetUserinfo)           //获取个人信息
		api.POST("/imgupload", middlewares.InterceptRequests(100), routers.UploadImage)        //上传图片
		api.POST("/tools/servelist", middlewares.InterceptRequests(100), tools.GetServeList)   //获取服务器列表
	}

	//非权限路由
	group := router.Group("/user")
	{
		group.POST("/regist", middlewares.InterceptRequests(30), users.Regist)           //注册
		group.POST("/login", middlewares.InterceptRequests(30), users.LoginJwt)          //登录
		group.POST("/sendcode", middlewares.InterceptRequests(10), users.QueryByPhone)   //发送验证码
		group.POST("/loginbycode", middlewares.InterceptRequests(30), users.LoginByCode) //验证码登陆

		//group.POST("/aa", middlewares.InterceptRequests(30), users.Sss) //验证码登陆

	}

	router.GET("/socket", func(context *gin.Context) {
		ws.Handler(context.Writer, context.Request)

	})

	router.Run(":8080")
}
