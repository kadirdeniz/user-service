package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userId, _ = primitive.ObjectIDFromHex("63add115ed3628749870398e")
var userId2, _ = primitive.ObjectIDFromHex("63add115ed36287498703999")
var userId3, _ = primitive.ObjectIDFromHex("63add115ed36287498703988")

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
		ID:        userId2,
		FirstName: "Jane",
		LastName:  "Doe",
		Nickname:  "janedoe",
		Email:     "janedoe@mail.com",
		Password:  "password",
	},
	{
		ID:        userId3,
		FirstName: "John",
		LastName:  "Smith",
		Nickname:  "johnsmith",
		Email:     "johnsmith@mail.com",
		Password:  "password",
	},
}

var MockUsers2 = []interface{}{
	&User{
		ID:        userId2,
		FirstName: "Jane",
		LastName:  "Doe",
		Nickname:  "janedoe",
		Email:     "janedoe@mail.com",
		Password:  "password",
	},
	&User{
		ID:        userId3,
		FirstName: "John",
		LastName:  "Smith",
		Nickname:  "johnsmith",
		Email:     "johnsmith@mail.com",
		Password:  "password",
	},
}
