package redisServer

import (
	"github.com/gomodule/redigo/redis"
)

func SetRedis(key string, value interface{}, t int64, get redis.Conn) {

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

// ExpireRedis 延长过期时间
func ExpireRedis(key string, t int64, get redis.Conn) {

	_, err := get.Do("EXPIRE", key, t)
	if err != nil {
		panic(err)
	}
}

// ExistsRedis 检查给定 key 是否存在。
func ExistsRedis(key string, get redis.Conn) bool {
	do, err := redis.Bool(get.Do("EXISTS", key))
	if err != nil {
		panic("redis错误")
	}
	return do
}

// IncrRedis 将 key 中储存的数字值增一。
func IncrRedis(key string, get redis.Conn) {
	_, err := get.Do("INCR", key)
	if err != nil {
		panic(err)
	}
}

// HmsetRedis 设置到哈希表 key 中
func HmsetRedis(key string, field string, value string, get redis.Conn) {
	_, err := get.Do("HMSET", key, field, value)
	if err != nil {
		panic(err)
	}

}

// HincrbyRedis 为哈希表 key 中的指定字段的整数值加上增量 increment 。
func HincrbyRedis(key string, field string, increment int, get redis.Conn) {
	_, err := get.Do("HINCRBY", key, field, increment)
	if err != nil {
		panic(err)
	}
}

// HgetRedis 获取存储在哈希表中指定字段的值。
func HgetRedis(key string, field string, get redis.Conn) int {
	do, err := redis.Int(get.Do("HGET", key, field))
	if err != nil {
		panic(err)
	}
	return do
}
