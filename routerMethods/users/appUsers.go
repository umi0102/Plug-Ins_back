package users

import (
	"Plug-Ins/databases/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var jwtKey []byte = []byte("secret")

type customClaims struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"IsAdmin"`
	jwt.RegisteredClaims
}

// LoginJwt 登入
func LoginJwt(ctx *gin.Context) {
	var req LoginRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		return
	}

	claims := customClaims{
		Username: req.Username,
		IsAdmin:  req.Username == "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(1 * time.Hour)},
		},
	}

	login(req.Username, req.Password)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err1 := token.SignedString(jwtKey)
	if err1 != nil {
		log.Println(err1.Error())

		panic(map[string]interface{}{
			"code": "400",
			"msg":  "Token过期",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功", "data": tokenString})
}

func login(username string, pwd string) {

	sqlStr := fmt.Sprintf(`select userinfo_password from userinfos where userinfo_phone="%s"`, username)
	mysqlSelect := mysql.SelectMysql(sqlStr)
	if mysqlSelect["userinfo_password"] != pwd {

		panic(map[string]interface{}{
			"code": "400",
			"msg":  "密码错误",
		})
	}
}

// Regist 注册
func Regist(ctx *gin.Context) {
	userinfo := Userinfo{}

	if err := ctx.ShouldBind(&userinfo); err != nil {
		panic(map[string]interface{}{
			"code": "400",
			"msg":  "格式错误",
		})

	}

	if len(userinfo.Name) == 0 {
		panic(map[string]interface{}{
			"code": "400",
			"msg":  "格式错误",
		})

	}
	regQuery(userinfo.Name, userinfo.Password)

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功"})
}

func regQuery(name string, pwd string) {

	mysqlSelect := mysql.SelectMysql(fmt.Sprintf(`select userinfo_phone from userinfos where userinfo_phone="%s"`, name))
	if len(mysqlSelect) == 1 {
		panic("用户名已存在！")
	}
	mysql.InsUpdDelMysql(fmt.Sprintf(`insert into userinfos(userinfo_id, userinfo_phone, userinfo_password) values("%s", "%s", "%s")`, RandCreator(64), name, pwd))

}
func RandCreator(l int) string {
	str := "0123456789abcdefghigklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+"
	strList := []byte(str)

	var result []byte
	i := 0

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i < l {
		newStr := strList[r.Intn(len(strList))]
		result = append(result, newStr)
		i = i + 1
	}
	return string(result)
}
