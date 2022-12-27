package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"user-service/internal/user"
	"user-service/pkg"
)

var CTX = context.Background()

var RedisClient *redis.Client

type Redis struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type RedisInterface interface {
	GetRedisClient() *redis.Client
	Connect() (*Redis, error)
	GetUserByID(userId primitive.ObjectID) (*user.User, error)
	SetUser(user *user.User, ttl time.Duration) error
}

func NewRedis(configs pkg.RedisConfig) RedisInterface {
	return &Redis{
		Host:     configs.Host,
		Port:     configs.Port,
		Password: configs.Password,
		DB:       configs.DB,
	}
}

func (r *Redis) GetRedisClient() *redis.Client {
	return RedisClient
}

func (r *Redis) Connect() (*Redis, error) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     r.Host + ":" + r.Port,
		Password: r.Password,
		DB:       r.DB,
	})

	return r, nil
}
