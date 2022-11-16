package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var (
	MyRedis *redis.Client
	ctx     = context.Background()
)

func Setup() {
	MyRedis = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		//Password: "",
		DB: 0, // use default DB
	})
	_, err := MyRedis.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Redis connect ping failed, err:", err)
		return
	}
	fmt.Println("Redis connect succeeded")

	return
}
