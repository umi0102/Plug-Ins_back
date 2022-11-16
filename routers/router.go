package routers

import (
	"Plug-Ins/databases/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Regist(ctx *gin.Context) {
	userinfo := Userinfo{}
	fmt.Println(userinfo)

	if err := ctx.ShouldBind(&userinfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(userinfo)
	if len(userinfo.Name) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"code": 500, "msg": ""})
		return
	}
	if mysql.RegQuery(userinfo.Name, userinfo.Password) {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功"})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "用户已注册"})
}

// GetProjectList 获取项目列表
func GetProjectList(ctx *gin.Context) {
	res := mysql.GetProjectList()
	if res.State {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": res.Data,
		})
	}
}

// AddProject 添加项目
func AddProject(ctx *gin.Context) {
	projectList := ProjectList{}
	err := ctx.BindJSON(&projectList)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "json格式不正确",
		})
	}
	if mysql.AddProject(projectList.Name, projectList.Creator, projectList.IdentityType) {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "新建项目成功"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "项目新建失败"})

}

func GetDeveloperList(ctx *gin.Context) {
	Name := ProjectName{}
	err := ctx.BindJSON(&Name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "json格式不正确",
		})
	}
}

func JoinProject(ctx *gin.Context) {
	Type := JoinDeveloper{}
	err := ctx.BindJSON(&Type)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "json格式不正确",
		})
	}
	if mysql.JoinPro(Type.Projectname, Type.Developer, Type.IdentityType) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "加入项目成功",
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": 400,
		"msg":  "加入项目失败",
	})
}

func LeaveProject(ctx *gin.Context) {
	Leave := LeaveDeveloper{}
	err := ctx.BindJSON(&Leave)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "json格式不正确",
		})
	}
	if mysql.DeleteDeveloper(Leave.Projectname, Leave.Developer) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "人员删除成功",
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": 400,
		"msg":  "用户删除失败",
	})

}
