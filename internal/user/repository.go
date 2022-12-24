package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockgen -source=repository.go -destination=../../test/mock/mock_repository.go -package=mock
type IRepository interface {
	Upsert(user *User) error
	IsEmailExists(email string) (bool, error)
	IsNicknameExists(nickname string) (bool, error)
	GetUserByID(id primitive.ObjectID) (*User, error)
	GetUsers() ([]*User, error)
}

type Repository struct {
	Collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) IRepository {
	return &Repository{
		Collection: collection,
	}
}

func (r Repository) Upsert(user *User) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) IsEmailExists(email string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) IsNicknameExists(nickname string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetUserByID(id primitive.ObjectID) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetUsers() ([]*User, error) {
	//TODO implement me
	panic("implement me")
}
