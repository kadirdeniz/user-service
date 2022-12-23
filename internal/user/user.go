package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	FirsName string             `json:"first_name" bson:"first_name"`
	LastName string             `json:"last_name" bson:"last_name"`
	Nickname string             `json:"nickname" bson:"nickname"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"-" bson:"password"`
}
