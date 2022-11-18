package redis

import (
	"github.com/gomodule/redigo/redis"
)

func SetRedis(key string, value string, t int64, get redis.Conn) {
	defer func(get redis.Conn) {
		err := get.Close()
		if err != nil {
			panic(err)
		}
	}(get)
	_, err := get.Do("SETEX", key, t, value)
	if err != nil {
		panic(err)
	}

}

func GetRedis(key string, get redis.Conn) string {

	defer func(get redis.Conn) {
		err := get.Close()
		if err != nil {
			panic(err)
		}
	}(get)
	result, err := redis.String(get.Do("GET", key))
	if err != nil {
		panic(err)
	}
	return result
}

func DelRedis(key string, get redis.Conn) {

	defer func(get redis.Conn) {
		err := get.Close()
		if err != nil {
			panic(err)
		}
	}(get)
	_, err := get.Do("DEL", key)
	if err != nil {
		panic(err)
	}
}

func ExpireRedis(key string, t int64, get redis.Conn) {
	// 延长过期时间
	defer func(get redis.Conn) {
		err := get.Close()
		if err != nil {
			panic(err)
		}
	}(get)
	_, err := get.Do("EXPIRE", key, t)
	if err != nil {
		panic(err)
	}
}
