package mock

import "user-service/pkg/dto"

var UpdateUserRequestExample = dto.UpdateUserRequest{
	FirstName: "John",
	LastName:  "Doe",
	Nickname:  "johndoe",
	Email:     "johndoe@mail.com",
}

var CreateUserRequestExample = dto.CreateUserRequest{
	FirstName: "John",
	LastName:  "Doe",
	Nickname:  "johndoe",
	Email:     "johndoe@mail.com",
	Password:  "123456",
}
