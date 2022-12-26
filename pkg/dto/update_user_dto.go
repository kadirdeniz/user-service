package dto

type UpdateUserRequest struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Nickname  string `json:"nickname" bson:"nickname"`
	Email     string `json:"email" bson:"email"`
}

var UpdateUserRequestExample = UpdateUserRequest{
	FirstName: "John",
	LastName:  "Doe",
	Nickname:  "johndoe",
	Email:     "johndoe@mail.com",
}
