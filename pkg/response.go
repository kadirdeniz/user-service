package pkg

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(status bool, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

var UserCreatedSuccessResponse = NewResponse(true, "User created", nil)
var UserUpdatedSuccessResponse = NewResponse(true, "User updated", nil)
var UserDeletedSuccessResponse = NewResponse(true, "User deleted", nil)
var UserNotFoundResponse = NewResponse(false, "User not found", nil)
var UserFoundResponse = NewResponse(true, "User found", nil)
var UsersFoundResponse = NewResponse(true, "Users found", nil)
var EmailAlreadyExistsResponse = NewResponse(false, "Email already exists", nil)
var NicknameAlreadyExistsResponse = NewResponse(false, "Nickname already exists", nil)
