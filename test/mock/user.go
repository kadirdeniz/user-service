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

var MockUsers = []*user.User{
	{
		ID:        primitive.NewObjectID(),
		FirstName: "John",
		LastName:  "Doe",
		Nickname:  "johndoe",
		Email:     "johndoe@mail.com",
		Password:  "password",
	},
	{
		ID:        primitive.NewObjectID(),
		FirstName: "Jane",
		LastName:  "Doe",
		Nickname:  "janedoe",
		Email:     "janedoe@mail.com",
		Password:  "password",
	},
	{
		ID:        primitive.NewObjectID(),
		FirstName: "John",
		LastName:  "Smith",
		Nickname:  "johnsmith",
		Email:     "johnsmith@mail.com",
		Password:  "password",
	},
}
