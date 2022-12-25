package redis

import (
	"github.com/go-redis/redis/v8"
	"user-service/pkg"
)

func NewRedis(configs pkg.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     configs.Host + ":" + configs.Port,
		Password: configs.Password,
		DB:       configs.DB,
	})
}
