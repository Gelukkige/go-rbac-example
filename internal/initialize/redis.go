package initialize

import (
	"fmt"
	"go-rbac-example/internal/global"
	"log"

	"github.com/go-redis/redis/v8"
)

func RedisInit() {
	redisConfig := global.Config.Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	ctx := redisClient.Context()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	global.RedisClient = redisClient
	log.Println("Redis连接成功!")
}
