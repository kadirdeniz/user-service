package redis

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"user-service/internal/user"
)

const userPrefix = "user:"

func (r *Redis) GetUserByID(userId primitive.ObjectID) (*user.User, error) {
	var user *user.User

	err := r.GetRedisClient().Get(CTX, userPrefix+userId.Hex()).Scan(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Redis) SetUser(user *user.User, ttl time.Duration) error {

	err := r.GetRedisClient().Set(CTX, userPrefix+user.ID.Hex(), user, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}
