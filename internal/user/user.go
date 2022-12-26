package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Nickname  string             `json:"nickname" bson:"nickname"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"-" bson:"password"`
}

func New(firstName, lastName, nickname, email, password string) *User {
	return &User{
		ID:        primitive.NewObjectID(),
		FirstName: firstName,
		LastName:  lastName,
		Nickname:  nickname,
		Email:     email,
		Password:  password,
	}
}
