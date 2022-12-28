package internal

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"user-service/internal/user"
	"user-service/tools/mongodb"
	"user-service/tools/redis"
)

//go:generate mockgen -source=repository.go -destination=../../test/mock/mock_repository.go -package=mock
type IRepository interface {
	Upsert(user *user.User) error
	IsEmailExists(email string) (bool, error)
	IsNicknameExists(nickname string) (bool, error)
	DeleteUserByID(id primitive.ObjectID) error
	GetUserByID(id primitive.ObjectID) (*user.User, error)
	GetUsers() ([]*user.User, error)
}

type Repository struct {
	MongoDBInterface mongodb.MongoDBInterface
	RedisClient      redis.RedisInterface
}

func NewRepository(mongoDBInterface mongodb.MongoDBInterface, redisClient redis.RedisInterface) IRepository {
	return &Repository{
		MongoDBInterface: mongoDBInterface,
		RedisClient:      redisClient,
	}
}

func (r Repository) Upsert(user *user.User) error {
	if err := r.MongoDBInterface.Upsert(user); err != nil {
		return err
	}

	if err := r.RedisClient.SetUser(user, redis.TTL); err != nil {
		log.Printf("Error while setting user to redis:", err)
		return nil
	}

	return nil
}

func (r Repository) IsEmailExists(email string) (bool, error) {
	isEmailExits, err := r.MongoDBInterface.IsEmailExists(email)
	if err != nil {
		return false, err
	}

	return isEmailExits, nil
}

func (r Repository) IsNicknameExists(nickname string) (bool, error) {
	isNicknameExists, err := r.MongoDBInterface.IsNicknameExists(nickname)
	if err != nil {
		return false, err
	}

	return isNicknameExists, nil
}

func (r Repository) GetUserByID(id primitive.ObjectID) (*user.User, error) {

	userObj, err := r.RedisClient.GetUserByID(id)
	if err != nil {
		log.Printf("Error while getting user from redis:", err)
	} else if userObj != nil {
		return userObj, nil
	}

	userObj, err = r.MongoDBInterface.GetUserByID(id)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return new(user.User), nil
		}
		return nil, err
	}

	return userObj, nil
}

func (r Repository) GetUsers() ([]*user.User, error) {
	users, err := r.MongoDBInterface.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r Repository) DeleteUserByID(id primitive.ObjectID) error {
	if err := r.MongoDBInterface.DeleteUserByID(id); err != nil {
		return err
	}

	user := &user.User{
		ID: id,
	}

	if err := r.RedisClient.SetUser(user, 0); err != nil {
		log.Printf("Error while deleting user from redis:", err)
	}

	return nil
}
