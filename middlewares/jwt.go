package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"Plug-Ins/databases/mysql"
	"Plug-Ins/databases/redis"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey []byte = []byte("secret")

type customClaims struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"IsAdmin"`
	jwt.RegisteredClaims
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthRequired gin jwt 认证中间件
func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
		token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": -1, "msg": fmt.Sprintf("access token parse error: %v", err)})
			return
		}
		if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
			if !claims.VerifyExpiresAt(time.Now(), false) {
				ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": -1, "msg": "access token expired"})
				return
			}
			ctx.Set("claims", claims)
		} else {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": -1, "msg": fmt.Sprintf("Claims parse error: %v", err)})
			return
		}
		ctx.Next()
	}
}

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

	if mysql.Login(req.Username, req.Password) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		if tokenString, err := token.SignedString(jwtKey); err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "token生成失败: " + err.Error()})
		} else {
			if redis.SetRedis(req.Username, tokenString, 9999999999) {
				ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功", "data": tokenString})
			}
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "密码不正确"})
	}
}
