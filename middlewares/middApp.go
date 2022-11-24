package middlewares

import (
	"Plug-Ins/databases/redisServer"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gomodule/redigo/redis"
)

// Cors 跨域
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		context.Header("Access-Control-Max-Age", "1728000")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}

// 用户验证
type CustomClaims struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"IsAdmin"`
	jwt.RegisteredClaims
}

var jwtKey []byte = []byte("secret")

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })
		if err != nil {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{})
			return
		}
		//验证过期
		claims, ok := token.Claims.(*CustomClaims)
		if !ok && !token.Valid || !claims.VerifyExpiresAt(time.Now(), false) {
			//token解析失败
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "登陆失效",
			})
			return
		}
		// 发送token信息
		username := claims.Username
		ctx.Set("phone", username)
		ctx.Next()
	}
}

// InterceptRequests IP限制
func InterceptRequests(num int) gin.HandlerFunc {
	return func(context *gin.Context) {

		get := redisServer.RedisDb.Get()
		defer func(get redis.Conn) {
			err := get.Close()
			if err != nil {
				context.Abort()
				panic(err)
			}
		}(get)

		ip := context.ClientIP()
		if len(ip) == 0 {
			context.Abort()
			panic("IP错误")
		}

		keyRedis := fmt.Sprintf("%s-%s", context.Request.URL, ip)
		existsRedis := redisServer.ExistsRedis(keyRedis, get)
		if !existsRedis {
			redisServer.SetRedis(keyRedis, 1, 60, get)
		}
		getRedis := redisServer.GetRedis(keyRedis, get)
		redisServer.ExpireRedis(keyRedis, 60, get)
		redisServer.IncrRedis(keyRedis, get)

		res, err := strconv.Atoi(getRedis)
		if err != nil {
			panic(err)
		}

		if res >= num {
			context.Abort()
			panic("拒绝请求")

		}

		context.Next()
	}
}
