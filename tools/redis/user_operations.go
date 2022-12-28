package redis

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"user-service/internal/user"
	"user-service/pkg"
)

const userPrefix = "user:"

var TTL = 20 * time.Minute

func (r *Redis) GetUserByID(userId primitive.ObjectID) (*user.User, error) {
	var userObj *user.User

	userStr, err := r.GetRedisClient().Get(CTX, userPrefix+userId.Hex()).Result()
	if err != nil {
		return nil, err
	}

	err = pkg.JSONEncoder([]byte(userStr), &userObj)
	if err != nil {
		return new(user.User), err
	}

	return userObj, nil
}

func (r *Redis) SetUser(user *user.User, ttl time.Duration) error {

	decodedUser, err := pkg.JSONDecoder(user)
	if err != nil {
		return err
	}

	err = r.GetRedisClient().Set(CTX, userPrefix+user.ID.Hex(), decodedUser, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}
