package dto

type CreateUserRequest struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Nickname  string `json:"nickname" bson:"nickname"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
}
