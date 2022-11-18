package users

import (
	"Plug-Ins/databases/mysql"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"io"
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

// QueryByPhone 验证码接口，查询用户是否存在

var logger = log.Default()

type QueryByPhoneCallBack struct {
	Code        int
	SendCodeMsg string
}

// Createcode 生成随机四位数字
func Createcode() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000)) //这里面前面的04v是和后面的1000相对应的
}

type phones struct {
	Phone int `json:"phone"`
}

func QueryByPhone(ctx *gin.Context) {
	var phoneNum = phones{}
	err := ctx.BindJSON(&phoneNum)
	if err != nil {
		panic("Json错误")
	}

	selectMysql := mysql.SelectMysql(fmt.Sprintf(`select  userinfo_phone from userinfos where userinfo_phone = %d`, phoneNum.Phone))
	fmt.Println(len(selectMysql))
	if len(selectMysql) == 0 {
		panic("手机号错误")
	}

	//生成随机数
	verifyCode := Createcode()
	client := &http.Client{}
	requestBody := fmt.Sprintf("phoneNumber=%s&smsSignId=%s&verifyCode=%s", "18355320102", "0000", verifyCode)
	var jsonStr = []byte(requestBody)
	requst, err := http.NewRequest("POST",
		"https://miitangs09.market.alicloudapi.com/v1/tools/sms/code/sender",
		bytes.NewBuffer(jsonStr))

	if err != nil {
		//验证码请求发送失败联系管理员
		panic(map[string]interface{}{
			"code": 1,
			"msg":  "验证码发送失败，请联系管理员",
		})
	}

	requst.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	requst.Header.Add("Authorization", "APPCODE 23043a6b4ea44c56b5ebe9b65800d3fb")
	//随机字符串
	requst.Header.Add("X-Ca-Nonce", verifyCode)
	response, err := client.Do(requst)
	if err != nil {
		panic(map[string]interface{}{
			"code": 1,
			"msg":  "验证码发送失败，请联系管理员",
		})
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(map[string]interface{}{
			"code": 1,
			"msg":  "验证码发送失败，请联系管理员",
		})
	}

	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  string(body),
	})

}
