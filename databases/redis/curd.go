package redis

import (
	"fmt"
	"time"
)

func SetRedis(key string, value string, t int64) bool {
	expire := time.Duration(t) * time.Second
	if err := MyRedis.Set(ctx, key, value, expire).Err(); err != nil {
		return false
	}
	return true
}

func GetRedis(key string) string {
	result, err := MyRedis.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return result
}

func DelRedis(key string) bool {
	_, err := MyRedis.Del(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func ExpireRedis(key string, t int64) bool {
	// 延长过期时间
	expire := time.Duration(t) * time.Second
	if err := MyRedis.Expire(ctx, key, expire).Err(); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
