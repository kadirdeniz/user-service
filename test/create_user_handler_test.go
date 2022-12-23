package test

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
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
		Data: map[string]string{
			"token": "token",
		},
	}

	request := dto.CreateUserRequestExample
	json.NewEncoder(&buf).Encode(request)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockIRepository(ctrl)
	mockRepository.
		EXPECT().
		IsEmailExists(request.Email).
		Return(false).
		Times(1)

	mockRepository.
		EXPECT().
		IsNicknameExists(request.Nickname).
		Return(false).
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

	asd, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(asd, &responseObj)
	if err != nil {
		t.Errorf("Cannot unmarshal response: %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, mockResponse.Status, responseObj.Status)
	assert.Equal(t, mockResponse.Message, responseObj.Message)
}
