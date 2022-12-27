package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"user-service/tools/mongodb"
	"user-service/tools/redis"
)

//go:generate mockgen -source=repository.go -destination=../../test/mock/mock_repository.go -package=mock
type IRepository interface {
	Upsert(user *User) error
	IsEmailExists(email string) (bool, error)
	IsNicknameExists(nickname string) (bool, error)
	DeleteUserByID(id primitive.ObjectID) error
	GetUserByID(id primitive.ObjectID) (*User, error)
	GetUsers() ([]*User, error)
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

func (r Repository) Upsert(user *User) error {
	if err := r.MongoDBInterface.Upsert(user); err != nil {
		return err
	}

	if err := r.RedisClient.SetUser(user, 20); err != nil {
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

func (r Repository) GetUserByID(id primitive.ObjectID) (*User, error) {

	user, err := r.RedisClient.GetUserByID(id)
	if err != nil {
		log.Printf("Error while getting user from redis:", err)
	}

	if user != nil {
		return user, nil
	}

	user, err = r.MongoDBInterface.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r Repository) GetUsers() ([]*User, error) {
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

	user := &User{
		ID: id,
	}

	if err := r.RedisClient.SetUser(user, 0); err != nil {
		log.Printf("Error while deleting user from redis:", err)
	}

	return nil
}
