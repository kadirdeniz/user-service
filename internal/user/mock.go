package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userId, _ = primitive.ObjectIDFromHex("63add115ed3628749870398e")

var MockUser = &User{
	ID:        userId,
	FirstName: "John",
	LastName:  "Doe",
	Nickname:  "johndoe",
	Email:     "johndoe@mail.com",
	Password:  "123456",
}

var MockUsers = []*User{
	MockUser,
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
