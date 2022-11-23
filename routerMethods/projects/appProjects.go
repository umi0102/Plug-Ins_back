package projects

import (
	"Plug-Ins/databases/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddProject 添加项目
func AddProject(ctx *gin.Context) {
	projectList := ProjectList{}
	err := ctx.BindJSON(&projectList)
	if err != nil {
		panic(map[string]interface{}{
			"code": "400",
			"msg":  "Json格式错误",
		})
	}
	addProject(projectList.Name, projectList.Creator, projectList.IdentityType)

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "项目Ok"})
}

// addProject 添加项目名称
func addProject(name string, creator string, identityType string) {
	tx := mysql.MysqlDb.Begin()

	sqlStr := fmt.Sprintf(`insert into projectlist(name, creator) values("%s", "%s")`, name, creator)
	sqlStr1 := fmt.Sprintf(`insert into developer(projectname, developer,identityType) values("%s", "%s", "%s")`, name, creator, identityType)

	res := tx.Exec(sqlStr)
	if res.Error != nil || res.RowsAffected == 0 {
		tx.Rollback()
		panic(map[string]interface{}{
			"code": "400",
			"msg":  res.Error.Error(),
		})
	}

	res1 := tx.Exec(sqlStr1)
	if res1.Error != nil || res1.RowsAffected == 0 {
		tx.Rollback()
		panic(map[string]interface{}{
			"code": "400",
			"msg":  res.Error.Error(),
		})
	}

	tx.Commit()
}

// GetDeveloperList 获取项目人员
func GetDeveloperList(ctx *gin.Context) {
	Name := ProjectName{}
	err := ctx.BindJSON(&Name)
	if err != nil {
		panic(map[string]interface{}{
			"code": "400",
			"msg":  "Json格式错误",
		})
	}

	mysqlSelect := mysql.SelectMysql(fmt.Sprintf(`select developer from developer where projectname = ("%s")`, Name.Name))
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": mysqlSelect,
	})
}

// JoinProject 用户加入项目
func JoinProject(ctx *gin.Context) {
	Type := JoinDeveloper{}
	err := ctx.BindJSON(&Type)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "json格式不正确",
		})
	}
	mysql.InsUpdDelMysql(fmt.Sprintf(`insert into developer(projectname,developer,identityType) values("%s", "%s","%s")`, Type.Projectname, Type.Developer, Type.IdentityType))
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": 200,
		"msg":  "加入项目成功",
	})
}

// LeaveProject 退出项目/删除成员
func LeaveProject(ctx *gin.Context) {
	Leave := LeaveDeveloper{}
	err := ctx.BindJSON(&Leave)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "json格式不正确",
		})
	}

	mysql.InsUpdDelMysql(fmt.Sprintf(`delete from developer where (projectname,developer) = ("%s","%s")`, Leave.Projectname, Leave.Developer))

	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": 200,
		"msg":  "用户删除成功",
	})

}

// GetProjectList 获取项目列表
func GetProjectList(ctx *gin.Context) {

	res := mysql.SelectMysql(`select nameid, name, creator, finished from projectlist`)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": res,
	})
}

// CheckToken 验证token
func CheckToken(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "success",
	})
}

//根据token解析用户名

// GetUserinfo 获取个人信息
func GetUserinfo(ctx *gin.Context) {
	info := Userinfo{}
	fmt.Println(info.Name)
	res := mysql.SelectMysql(fmt.Sprintf(`select userinfo_name,userinfo_usericon,userinfo_phone,userinfo_name from userinfos where userinfo_name = ("%s")`, info.Name))
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": res,
	})
}
