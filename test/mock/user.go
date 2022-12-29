package mock

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"user-service/internal/user"
)

var userId, _ = primitive.ObjectIDFromHex("63add115ed3628749870398e")

var MockUser = &user.User{
	ID:        userId,
	FirstName: "John",
	LastName:  "Doe",
	Nickname:  "johndoe",
	Email:     "johndoe@mail.com",
	Password:  "123456",
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
