package redisServer

import (
	"github.com/gomodule/redigo/redis"
)

var (
	RedisDb *redis.Pool
)

func init() {
	RedisDb = &redis.Pool{
		MaxIdle:     10,  //最大空闲链接数量
		MaxActive:   0,   //表示和数据库最大链接数，0表示，并发不限制数量
		IdleTimeout: 100, //最大空闲时间，用完链接后100秒后就回收到链接池
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(0), redis.DialPassword("Fjw0102*"))
		},
	}
}
