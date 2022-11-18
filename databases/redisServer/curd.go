package redisServer

import (
	"github.com/gomodule/redigo/redis"
)

func SetRedis(key string, value string, t int64, get redis.Conn) {

	_, err := get.Do("SETEX", key, t, value)
	if err != nil {
		panic(err)
	}

}

func GetRedis(key string, get redis.Conn) string {

	result, err := redis.String(get.Do("GET", key))
	if err != nil {
		panic(err)
	}
	return result
}

func DelRedis(key string, get redis.Conn) {

	_, err := get.Do("DEL", key)

	if err != nil {
		panic(err)
	}
}

func ExpireRedis(key string, t int64, get redis.Conn) {
	// 延长过期时间

	_, err := get.Do("EXPIRE", key, t)
	if err != nil {
		panic(err)
	}
}

func ExistsRedis(key string, get redis.Conn) bool {
	do, err := redis.Bool(get.Do("EXISTS", key))
	if err != nil {
		panic("redis错误")
	}
	return do
}
