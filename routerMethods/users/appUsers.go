package users

import (
	"Plug-Ins/databases/mysql"
	"Plug-Ins/databases/redisServer"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gomodule/redigo/redis"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var jwtKey []byte = []byte("secret")

type customClaims struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"IsAdmin"`
	jwt.RegisteredClaims
}

func GetToken(num string) string {
	claims := customClaims{
		Username: num,
		IsAdmin:  num == "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(1 * time.Hour)},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}
	return tokenString
}

// LoginJwt 登入
func LoginJwt(ctx *gin.Context) {
	var req LoginRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		return
	}
	login(req.Phone, req.Password)
	token := GetToken(strconv.Itoa(req.Phone))

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功", "data": token})
}

func login(phone int, pwd string) {

	sqlStr := fmt.Sprintf(`select userinfo_password from userinfos where userinfo_phone=%d`, phone)
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
	userinfo := LoginRequest{}

	if err := ctx.ShouldBind(&userinfo); err != nil {
		panic(map[string]interface{}{
			"code": "400",
			"msg":  "格式错误",
		})

	}

	regQuery(userinfo.Phone, userinfo.Password)

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功"})
}

func regQuery(phone int, pwd string) {

	mysqlSelect := mysql.SelectMysql(fmt.Sprintf(`select userinfo_phone from userinfos where userinfo_phone=%d`, phone))
	if len(mysqlSelect) == 1 {
		panic("用户名已存在！")
	}
	mysql.InsUpdDelMysql(fmt.Sprintf(`insert into userinfos(userinfo_id, userinfo_phone, userinfo_password) values("%s", %d, "%s")`, RandCreator(64), phone, pwd))

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

// QueryByPhone 发送验证码接口，查询用户是否存在
func QueryByPhone(ctx *gin.Context) {
	var phoneNum = LoginRequest{}
	err := ctx.BindJSON(&phoneNum)
	if err != nil {
		panic("Json错误")
	}

	// 从redis 连接池中拿出连接
	get := redisServer.RedisDb.Get()
	defer get.Close()

	// 验证手机号是否发送过验证码
	redis := redisServer.ExistsRedis(strconv.Itoa(phoneNum.Phone), get)
	if !redis {
		panic("验证码已发送请等待")
	}

	// 验证手机号是否存在
	selectMysql := mysql.SelectMysql(fmt.Sprintf(`select  userinfo_phone from userinfos where userinfo_phone = %d`, phoneNum.Phone))
	fmt.Println(len(selectMysql))
	if len(selectMysql) == 0 {
		panic("手机号错误")
	}

	//生成随机数
	verifyCode := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	client := &http.Client{}
	requestBody := fmt.Sprintf("phoneNumber=%s&smsSignId=%s&verifyCode=%s", "18355320102", "0000", verifyCode)
	var jsonStr = []byte(requestBody)
	requst, err1 := http.NewRequest("POST",
		"https://miitangs09.market.alicloudapi.com/v1/tools/sms/code/sender",
		bytes.NewBuffer(jsonStr))

	if err1 != nil {
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
	response, err2 := client.Do(requst)

	if err2 != nil {
		panic(map[string]interface{}{
			"code": 1,
			"msg":  "验证码发送失败，请联系管理员",
		})
	}
	defer response.Body.Close()
	body, err3 := io.ReadAll(response.Body)
	if err3 != nil {
		panic(map[string]interface{}{
			"code": 1,
			"msg":  "验证码发送失败，请联系管理员",
		})
	}
	var tempMap map[string]interface{}
	err4 := json.Unmarshal(body, &tempMap)
	if err4 != nil {
		panic("Json错误")
	}

	// 保存手机号对应对验证码
	redisServer.SetRedis(strconv.Itoa(phoneNum.Phone), verifyCode, 300, get)

	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  tempMap,
	})

}

func LoginByCode(ctx *gin.Context) {
	var phoneNum = LoginByCodeRequest{}
	err := ctx.BindJSON(&phoneNum)
	if err != nil {
		panic("Json解析错误")
	}

	get := redisServer.RedisDb.Get()
	defer func(get redis.Conn) {
		err := get.Close()
		if err != nil {
			panic(err)
		}
	}(get)

	// redis 拿出手机号 验证码
	getRedis := redisServer.GetRedis(phoneNum.Phone, get)
	if getRedis != phoneNum.Code {
		panic("验证码错误")
	}

	token := GetToken(phoneNum.Phone)
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功", "data": token})
}
