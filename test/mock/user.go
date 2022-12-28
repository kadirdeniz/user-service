package mock

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"user-service/internal/user"
)

var MockUser = &user.User{
	ID:        primitive.NewObjectID(),
	FirstName: "John",
	LastName:  "Doe",
	Nickname:  "johndoe",
	Email:     "johndoe@mail.com",
	Password:  "password",
}
