package test

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
	"user-service/internal/user"
	"user-service/pkg/dto"
	"user-service/test/mock"
	tool_fiber "user-service/tools/fiber"
)

var buf bytes.Buffer

func Test_CreateUserHandler(t *testing.T) {

	var responseObj dto.Response
	mockResponse := dto.Response{
		Status:  true,
		Message: "User created",
		Data:    nil,
	}

	request := dto.CreateUserRequestExample
	json.NewEncoder(&buf).Encode(request)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockIRepository(ctrl)
	mockRepository.
		EXPECT().
		IsEmailExists(request.Email).
		Return(false, nil).
		Times(1)

	mockRepository.
		EXPECT().
		IsNicknameExists(request.Nickname).
		Return(false, nil).
		Times(1)

	mockRepository.
		EXPECT().
		Upsert(gomock.Any()).
		Return(nil).
		Times(1)

	mockService := mock.NewMockIService(ctrl)

	handler := tool_fiber.NewHandler(mockRepository, mockService)

	app := fiber.New()
	app.Post("/user", handler.CreateUser)

	req := httptest.NewRequest("POST", "/user", &buf)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	responseBody, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(responseBody, &responseObj)
	if err != nil {
		t.Errorf("Cannot unmarshal response: %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, mockResponse.Status, responseObj.Status)
	assert.Equal(t, mockResponse.Message, responseObj.Message)
}

func Test_CreateUserHandlerEmailAlreadyExists(t *testing.T) {

	var responseObj dto.Response
	mockResponse := dto.Response{
		Status:  false,
		Message: "Email already exists",
		Data:    nil,
	}

	request := dto.CreateUserRequestExample
	json.NewEncoder(&buf).Encode(request)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockIRepository(ctrl)
	mockRepository.
		EXPECT().
		IsEmailExists(request.Email).
		Return(true, nil).
		Times(1)

	mockService := mock.NewMockIService(ctrl)

	handler := tool_fiber.NewHandler(mockRepository, mockService)

	app := fiber.New()
	app.Post("/user", handler.CreateUser)

	req := httptest.NewRequest("POST", "/user", &buf)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	responseBody, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(responseBody, &responseObj)
	if err != nil {
		t.Errorf("Cannot unmarshal response: %v", err)
	}

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, mockResponse.Status, responseObj.Status)
	assert.Equal(t, mockResponse.Message, responseObj.Message)
}

func Test_CreateUserHandlerNicknameAlreadyExists(t *testing.T) {

	var responseObj dto.Response
	mockResponse := dto.Response{
		Status:  false,
		Message: "Nickname already exists",
		Data:    nil,
	}

	request := dto.CreateUserRequestExample
	json.NewEncoder(&buf).Encode(request)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockIRepository(ctrl)
	mockRepository.
		EXPECT().
		IsEmailExists(request.Email).
		Return(false, nil).
		Times(1)

	mockRepository.
		EXPECT().
		IsNicknameExists(request.Nickname).
		Return(true, nil).
		Times(1)

	mockService := mock.NewMockIService(ctrl)

	handler := tool_fiber.NewHandler(mockRepository, mockService)

	app := fiber.New()
	app.Post("/user", handler.CreateUser)

	req := httptest.NewRequest("POST", "/user", &buf)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	responseBody, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(responseBody, &responseObj)
	if err != nil {
		t.Errorf("Cannot unmarshal response: %v", err)
	}

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, mockResponse.Status, responseObj.Status)
	assert.Equal(t, mockResponse.Message, responseObj.Message)
}

func Test_GetUserById(t *testing.T) {
	var responseObj dto.Response
	mockResponse := dto.Response{
		Status:  true,
		Message: "User found",
		Data:    nil,
	}

	userId := primitive.NewObjectID()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockIRepository(ctrl)
	mockRepository.
		EXPECT().
		GetUserByID(userId).
		Return(&user.User{
			ID:        userId,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "test@mail.com",
			Nickname:  "test",
			Password:  "password",
		}, nil).
		Times(1)

	mockService := mock.NewMockIService(ctrl)

	handler := tool_fiber.NewHandler(mockRepository, mockService)

	app := fiber.New()
	app.Get("/user/:id", handler.GetUser)

	req := httptest.NewRequest("GET", "/user/"+userId.Hex(), nil)

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	responseBody, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(responseBody, &responseObj)
	if err != nil {
		t.Errorf("Cannot unmarshal response: %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, mockResponse.Status, responseObj.Status)
	assert.Equal(t, mockResponse.Message, responseObj.Message)
}

func Test_GetUserByIdUserNotFound(t *testing.T) {
	var responseObj dto.Response
	mockResponse := dto.Response{
		Status:  false,
		Message: "User not found",
		Data:    nil,
	}

	userId := primitive.NewObjectID()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockIRepository(ctrl)
	mockRepository.
		EXPECT().
		GetUserByID(userId).
		Return(&user.User{}, nil).
		Times(1)

	mockService := mock.NewMockIService(ctrl)

	handler := tool_fiber.NewHandler(mockRepository, mockService)

	app := fiber.New()
	app.Get("/user/:id", handler.GetUser)

	req := httptest.NewRequest("GET", "/user/"+userId.Hex(), nil)

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	responseBody, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(responseBody, &responseObj)
	if err != nil {
		t.Errorf("Cannot unmarshal response: %v", err)
	}

	assert.Equal(t, 404, resp.StatusCode)
	assert.Equal(t, mockResponse.Status, responseObj.Status)
	assert.Equal(t, mockResponse.Message, responseObj.Message)
}

func Test_GetUsers(t *testing.T) {
	var responseObj dto.Response
	mockResponse := dto.Response{
		Status:  true,
		Message: "Users found",
		Data:    nil,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockIRepository(ctrl)
	mockRepository.
		EXPECT().
		GetUsers().
		Return([]*user.User{
			{
				ID:        primitive.NewObjectID(),
				FirstName: "John",
				LastName:  "Doe",
				Email:     "test@mail.com",
				Nickname:  "test",
				Password:  "password",
			},
			{
				ID:        primitive.NewObjectID(),
				FirstName: "John",
				LastName:  "Doe",
				Email:     "test2@mail.com",
				Nickname:  "test2",
				Password:  "password",
			},
		}, nil).
		Times(1)

	mockService := mock.NewMockIService(ctrl)

	handler := tool_fiber.NewHandler(mockRepository, mockService)

	app := fiber.New()
	app.Get("/users", handler.GetUsers)

	req := httptest.NewRequest("GET", "/users", nil)

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	responseBody, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(responseBody, &responseObj)
	if err != nil {
		t.Errorf("Cannot unmarshal response: %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, mockResponse.Status, responseObj.Status)
	assert.Equal(t, mockResponse.Message, responseObj.Message)
}

func Test_DeleteUser(t *testing.T) {
	var responseObj dto.Response
	mockResponse := dto.Response{
		Status:  true,
		Message: "User deleted",
		Data:    nil,
	}

	userId := primitive.NewObjectID()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockIRepository(ctrl)
	mockRepository.
		EXPECT().
		GetUserByID(userId).
		Return(&user.User{
			ID:        userId,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "test@mail.com",
			Nickname:  "test",
			Password:  "password",
		}, nil).
		Times(1)

	mockRepository.
		EXPECT().
		DeleteUserByID(userId).
		Return(nil).
		Times(1)

	mockService := mock.NewMockIService(ctrl)

	handler := tool_fiber.NewHandler(mockRepository, mockService)

	app := fiber.New()
	app.Delete("/user/:id", handler.DeleteUser)

	req := httptest.NewRequest("DELETE", "/user/"+userId.Hex(), nil)

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	responseBody, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(responseBody, &responseObj)
	if err != nil {
		t.Errorf("Cannot unmarshal response: %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, mockResponse.Status, responseObj.Status)
	assert.Equal(t, mockResponse.Message, responseObj.Message)
}

func Test_DeleteUserUserNotFound(t *testing.T) {
	var responseObj dto.Response
	mockResponse := dto.Response{
		Status:  false,
		Message: "User not found",
		Data:    nil,
	}

	userId := primitive.NewObjectID()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockIRepository(ctrl)
	mockRepository.
		EXPECT().
		GetUserByID(userId).
		Return(&user.User{}, nil).
		Times(1)

	mockService := mock.NewMockIService(ctrl)

	handler := tool_fiber.NewHandler(mockRepository, mockService)

	app := fiber.New()
	app.Delete("/user/:id", handler.DeleteUser)

	req := httptest.NewRequest("DELETE", "/user/"+userId.Hex(), nil)

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	responseBody, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(responseBody, &responseObj)
	if err != nil {
		t.Errorf("Cannot unmarshal response: %v", err)
	}

	assert.Equal(t, 404, resp.StatusCode)
	assert.Equal(t, mockResponse.Status, responseObj.Status)
	assert.Equal(t, mockResponse.Message, responseObj.Message)
}

func Test_UpdateUser(t *testing.T) {
	var responseObj dto.Response
	mockResponse := dto.Response{
		Status:  true,
		Message: "User updated",
		Data:    nil,
	}

	var buf bytes.Buffer

	request := dto.UpdateUserRequestExample
	json.NewEncoder(&buf).Encode(request)

	userId := primitive.NewObjectID()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockIRepository(ctrl)
	mockRepository.
		EXPECT().
		Upsert(&user.User{
			ID:        userId,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe@mail.com",
			Nickname:  "johndoe",
		}).
		Return(nil).
		Times(1)

	mockRepository.
		EXPECT().
		GetUserByID(userId).
		Return(&user.User{
			ID:        userId,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe@mail.com",
			Nickname:  "johndoe",
		}, nil).
		Times(1)

	mockRepository.
		EXPECT().
		IsEmailExists(request.Email).
		Return(false, nil).
		Times(1)

	mockRepository.
		EXPECT().
		IsNicknameExists(request.Nickname).
		Return(false, nil).
		Times(1)

	mockService := mock.NewMockIService(ctrl)

	handler := tool_fiber.NewHandler(mockRepository, mockService)

	app := fiber.New()

	app.Put("/user/:id", handler.UpdateUser)

	req := httptest.NewRequest("PUT", "/user/"+userId.Hex(), &buf)

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	responseBody, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(responseBody, &responseObj)
	if err != nil {
		t.Errorf("Cannot unmarshal response: %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, mockResponse.Status, responseObj.Status)
	assert.Equal(t, mockResponse.Message, responseObj.Message)
}

func Test_UpdateUserUserNotFound(t *testing.T) {
	var responseObj dto.Response
	mockResponse := dto.Response{
		Status:  false,
		Message: "User not found",
		Data:    nil,
	}

	userId := primitive.NewObjectID()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockIRepository(ctrl)
	mockRepository.
		EXPECT().
		GetUserByID(userId).
		Return(&user.User{}, nil).
		Times(1)

	mockService := mock.NewMockIService(ctrl)

	handler := tool_fiber.NewHandler(mockRepository, mockService)

	app := fiber.New()

	app.Put("/user/:id", handler.UpdateUser)

	req := httptest.NewRequest("PUT", "/user/"+userId.Hex(), strings.NewReader(`{
		"firstName": "John",
		"lastName": "Doe",
		"email": "test@mail.com",
		"nickname": "test",
		"password": "password"
	}`))

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	responseBody, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(responseBody, &responseObj)
	if err != nil {
		t.Errorf("Cannot unmarshal response: %v", err)
	}

	assert.Equal(t, 404, resp.StatusCode)
	assert.Equal(t, mockResponse.Status, responseObj.Status)
	assert.Equal(t, mockResponse.Message, responseObj.Message)
}

func Test_UpdateUserEmailAlreadyExists(t *testing.T) {
	var responseObj dto.Response
	mockResponse := dto.Response{
		Status:  false,
		Message: "Email already exists",
		Data:    nil,
	}

	var buf bytes.Buffer

	request := dto.UpdateUserRequestExample
	json.NewEncoder(&buf).Encode(request)

	userId := primitive.NewObjectID()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockIRepository(ctrl)
	mockRepository.
		EXPECT().
		GetUserByID(userId).
		Return(&user.User{
			ID:        userId,
			FirstName: "John",
			LastName:  "Doe",
			Nickname:  "johndoe",
			Email:     "johndoe@mail.com",
		}, nil).
		Times(1)

	mockRepository.
		EXPECT().
		IsEmailExists(request.Email).
		Return(true, nil).
		Times(1)

	mockService := mock.NewMockIService(ctrl)

	handler := tool_fiber.NewHandler(mockRepository, mockService)

	app := fiber.New()

	app.Put("/user/:id", handler.UpdateUser)

	req := httptest.NewRequest("PUT", "/user/"+userId.Hex(), &buf)

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	responseBody, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(responseBody, &responseObj)
	if err != nil {
		t.Errorf("Cannot unmarshal response: %v", err)
	}

	assert.Equal(t, 409, resp.StatusCode)
	assert.Equal(t, mockResponse.Status, responseObj.Status)
	assert.Equal(t, mockResponse.Message, responseObj.Message)
}

func Test_UpdateUserNicknameAlreadyExists(t *testing.T) {
	var responseObj dto.Response
	mockResponse := dto.Response{
		Status:  false,
		Message: "Nickname already exists",
		Data:    nil,
	}

	userId := primitive.NewObjectID()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var buf bytes.Buffer

	request := dto.UpdateUserRequestExample
	json.NewEncoder(&buf).Encode(request)

	mockRepository := mock.NewMockIRepository(ctrl)
	mockRepository.
		EXPECT().
		GetUserByID(userId).
		Return(&user.User{
			ID:        userId,
			FirstName: "John",
			LastName:  "Doe",
			Nickname:  "johndoe",
			Email:     "johndoe@mail.com",
		}, nil).
		Times(1)

	mockRepository.
		EXPECT().
		IsEmailExists(request.Email).
		Return(false, nil).
		Times(1)

	mockRepository.
		EXPECT().
		IsNicknameExists(request.Nickname).
		Return(true, nil).
		Times(1)

	mockService := mock.NewMockIService(ctrl)

	handler := tool_fiber.NewHandler(mockRepository, mockService)

	app := fiber.New()

	app.Put("/user/:id", handler.UpdateUser)

	req := httptest.NewRequest("PUT", "/user/"+userId.Hex(), &buf)

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	responseBody, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(responseBody, &responseObj)
	if err != nil {
		t.Errorf("Cannot unmarshal response: %v", err)
	}

	assert.Equal(t, 409, resp.StatusCode)
	assert.Equal(t, mockResponse.Status, responseObj.Status)
	assert.Equal(t, mockResponse.Message, responseObj.Message)
}
