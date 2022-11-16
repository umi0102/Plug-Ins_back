package mysql

import (
	"fmt"
)

func RegQuery(name string, pwd string) bool {
	sqlStr := "select username from userinfo where username=? limit 1"
	row, _ := db.Query(sqlStr, name)
	ok := row.Next()
	defer row.Close()
	if ok {
		return false
	}
	defer row.Close()
	sqlStr1 := "insert into userinfo(username, password) values(?, ?)"
	res, err := db.Exec(sqlStr1, name, pwd)
	if err != nil {
		return false
	}
	l, err1 := res.RowsAffected()
	if err1 != nil {
		return false
	}
	if l == 1 {
		return true
	}
	return false

}

func Login(username string, pwd string) bool {
	var password string
	sqlStr := "select password from userinfo where username=(?)"
	err := db.QueryRow(sqlStr, username).Scan(&password)
	if err != nil {
		return false
	}
	if pwd == password {
		return true
	}
	return false

	//sqlStr := "select  developer from developer where projectname = (?)"
	//res, err := db.Query(sqlStr)
	//if err != nil {
	//	return DevelopBack{
	//		State: false,
	//	}
	//
	//}
	//userdata := ProjectName{}
	//userArr := make([]ProjectName, 0)
	//for res.Next() {
	//	res.Scan(&userdata.Name)
	//	userArr = append(userArr, userdata)
	//}
	//fmt.Println(userArr)
	//return DevelopBack{
	//	State: true,
	//	Data:  userArr,
	//}
}

// AddProject 添加项目名称
func AddProject(name string, creator string, identityType string) bool {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("开启事务失败")
		return false
	}

	sqlStr := "insert into projectlist(name, creator) values(?, ?)"
	sqlStr1 := "insert into developer(projectname, developer,identityType) values(?, ?, ?)"

	res, err := tx.Exec(sqlStr, name, creator)
	if err != nil {
		tx.Rollback()
		return false
	}
	row, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return false
	}
	if row != 1 {
		tx.Rollback()
		return false
	}
	res1, err := tx.Exec(sqlStr1, name, creator, identityType)
	if err != nil {
		tx.Rollback()
		return false
	}
	row1, err := res1.RowsAffected()
	if err != nil {
		tx.Rollback()
		return false
	}
	if row1 != 1 {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true

}

func GetProjectList() projectBack {
	sqlStr := "select nameid, name, creator, finished from projectlist"
	res, err := db.Query(sqlStr)
	if err != nil {
		return projectBack{
			State: false,
		}

	}
	userdata := projectlist{}
	userArr := make([]projectlist, 0)
	for res.Next() {
		res.Scan(&userdata.Nameid, &userdata.Name, &userdata.Creator, &userdata.Finished)
		userArr = append(userArr, userdata)
	}
	fmt.Println(userArr)
	return projectBack{
		State: true,
		Data:  userArr,
	}
}

// GetDeveloperList 获取项目人员
func GetDeveloperList() DevelopBack {
	sqlStr := "select  developer from developer where projectname = (?)"
	res, err := db.Query(sqlStr)
	if err != nil {
		return DevelopBack{
			State: false,
		}

	}
	userdata := ProjectName{}
	userArr := make([]ProjectName, 0)
	for res.Next() {
		res.Scan(&userdata.Name)
		userArr = append(userArr, userdata)
	}
	fmt.Println(userArr)
	return DevelopBack{
		State: true,
		Data:  userArr,
	}
}
func JoinPro(projectname string, developer string, identityType string) bool {
	sqlStr := "insert into developer(projectname,developer,identityType) values(?, ?,?)"
	res, err := db.Exec(sqlStr, projectname, developer, identityType)
	if err != nil {
		return false
	}
	row, err := res.RowsAffected()
	if err != nil {
		return false
	}
	if row == 1 {
		return true
	}
	return false
}

func DeleteDeveloper(projectname string, developer string) bool {
	//DELETE FROM table_name
	//WHERE some_column=some_value;
	sqlStr := "delete from developer where (projectname,developer) = (?,?)"
	res, err := db.Exec(sqlStr, projectname, developer)
	if err != nil {
		return false
	}
	row, err := res.RowsAffected()
	if err != nil {
		return false
	}
	if row >= 1 {
		return true
	}
	return false
}
