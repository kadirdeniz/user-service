package fiber_test

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io"
	"net/http/httptest"
	"testing"
	"user-service/internal/user"
	"user-service/pkg"
	"user-service/test/mock"
	fibertools "user-service/tools/fiber"
)

var testingObj *testing.T

func TestHandler(t *testing.T) {
	testingObj = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler Suite")
}

var _ = Describe("Handler", Ordered, func() {

	ctrl := gomock.NewController(testingObj)
	defer ctrl.Finish()

	var mockRepository *mock.MockIRepository
	var mockService *mock.MockIService

	BeforeEach(func() {
		mockRepository = mock.NewMockIRepository(ctrl)
		mockService = mock.NewMockIService(ctrl)
	})

	Context("GetUserByID", func() {
		When("user is found", func() {
			It("should return user", func() {
				mock.MockRepositoryGetUser(*mockRepository)
				getUserByIDHandler(mockRepository, mockService, pkg.UserFoundResponse)
			})
		})
		When("user is not found", func() {
			It("should return error", func() {
				mock.MockRepositoryGetUserNotFound(*mockRepository)
				getUserByIDHandler(mockRepository, mockService, pkg.UserNotFoundResponse)
			})
		})
	})

	Context("DeleteUserByID", func() {
		When("user is found", func() {
			It("should return user", func() {
				mock.MockRepositoryGetUser(*mockRepository)
				mock.MockRepositoryDeleteUser(*mockRepository)
				deleteUserByIDHandler(mockRepository, mockService, pkg.UserDeletedSuccessResponse)
			})
		})
		When("user is not found", func() {
			It("should return error", func() {
				mock.MockRepositoryGetUserNotFound(*mockRepository)
				mock.MockRepositoryDeleteUserNotFound(*mockRepository)
				deleteUserByIDHandler(mockRepository, mockService, pkg.UserNotFoundResponse)
			})
		})
	})

	Context("CreateUser", func() {
		When("email and nickname not exists", func() {
			It("should return user", func() {
				mock.MockRepositoryIsEmailExistsFalse(*mockRepository)
				mock.MockRepositoryIsNicknameExistsFalse(*mockRepository)
				mock.MockRepositoryUpsert(*mockRepository)
				createUserHandler(mockRepository, mockService, pkg.UserCreatedSuccessResponse)
			})
		})

		When("email exists", func() {
			It("should return error", func() {
				mock.MockRepositoryIsEmailExistsTrue(*mockRepository)
				mock.MockRepositoryIsNicknameExistsFalse(*mockRepository)
				createUserHandler(mockRepository, mockService, pkg.EmailAlreadyExistsResponse)
			})
		})

		When("nickname exists", func() {
			It("should return error", func() {
				mock.MockRepositoryIsNicknameExistsTrue(*mockRepository)
				mock.MockRepositoryIsEmailExistsFalse(*mockRepository)
				mock.MockRepositoryUpsert(*mockRepository)
				createUserHandler(mockRepository, mockService, pkg.NicknameAlreadyExistsResponse)
			})
		})
	})

	Context("UpdateUser", func() {
		When("email and nickname not exists", func() {
			It("should return user", func() {
				mock.MockRepositoryIsEmailExistsFalse(*mockRepository)
				mock.MockRepositoryIsNicknameExistsFalse(*mockRepository)
				mock.MockRepositoryGetUser(*mockRepository)
				mock.MockRepositoryUpsert(*mockRepository)
				updateUserHandler(mockRepository, mockService, pkg.UserUpdatedSuccessResponse)
			})
		})

		When("email exists", func() {
			It("should return error", func() {
				mock.MockRepositoryIsEmailExistsTrue(*mockRepository)
				mock.MockRepositoryIsNicknameExistsFalse(*mockRepository)
				mock.MockRepositoryGetUser(*mockRepository)
				mock.MockRepositoryUpsert(*mockRepository)
				updateUserHandler(mockRepository, mockService, pkg.EmailAlreadyExistsResponse)
			})
		})

		When("nickname exists", func() {
			It("should return error", func() {
				mock.MockRepositoryIsNicknameExistsTrue(*mockRepository)
				mock.MockRepositoryIsEmailExistsFalse(*mockRepository)
				mock.MockRepositoryGetUser(*mockRepository)
				mock.MockRepositoryUpsert(*mockRepository)
				updateUserHandler(mockRepository, mockService, pkg.NicknameAlreadyExistsResponse)
			})
		})

		When("user not found", func() {
			It("should return error", func() {
				mock.MockRepositoryIsNicknameExistsTrue(*mockRepository)
				mock.MockRepositoryIsEmailExistsFalse(*mockRepository)
				mock.MockRepositoryGetUserNotFound(*mockRepository)
				mock.MockRepositoryUpsert(*mockRepository)
				updateUserHandler(mockRepository, mockService, pkg.UserNotFoundResponse)
			})
		})
	})

	Context("GetUsers", func() {
		When("users are found", func() {
			It("should return users", func() {
				mock.MockRepositoryGetUsers(*mockRepository)
				getUsersHandler(mockRepository, mockService, pkg.UsersFoundResponse)
			})
		})
	})
})

func createUserHandler(mockRepository *mock.MockIRepository, mockService *mock.MockIService, mockResponse pkg.Response) {
	var responseObj pkg.Response

	byteRequest, _ := pkg.JSONDecoder(mock.CreateUserRequestExample)

	handler := fibertools.NewHandler(mockRepository, mockService)

	app := fiber.New()
	app.Post("/user", handler.CreateUser)

	req := httptest.NewRequest("POST", "/user", bytes.NewReader(byteRequest))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	Expect(err).To(BeNil())

	responseBody, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(responseBody, &responseObj)
	Expect(err).To(BeNil())

	Expect(responseObj.Status).To(Equal(mockResponse.Status))
	Expect(responseObj.Message).To(Equal(mockResponse.Message))
}

func updateUserHandler(mockRepository *mock.MockIRepository, mockService *mock.MockIService, mockResponse pkg.Response) {
	var responseObj pkg.Response

	byteRequest, _ := pkg.JSONDecoder(mock.UpdateUserRequestExample)

	handler := fibertools.NewHandler(mockRepository, mockService)

	app := fiber.New()
	app.Put("/user/:id", handler.UpdateUser)

	req := httptest.NewRequest("PUT", "/user/"+user.MockUser.ID.Hex(), bytes.NewReader(byteRequest))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	Expect(err).To(BeNil())

	responseBody, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(responseBody, &responseObj)
	Expect(err).To(BeNil())

	Expect(responseObj.Status).To(Equal(mockResponse.Status))
	Expect(responseObj.Message).To(Equal(mockResponse.Message))
}

func getUserByIDHandler(mockRepository *mock.MockIRepository, mockService *mock.MockIService, mockResponse pkg.Response) {
	var responseObj pkg.Response

	handler := fibertools.NewHandler(mockRepository, mockService)

	app := fiber.New()
	app.Get("/user/:id", handler.GetUser)

	req := httptest.NewRequest("GET", "/user/"+user.MockUser.ID.Hex(), nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	Expect(err).To(BeNil())

	responseBody, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(responseBody, &responseObj)
	Expect(err).To(BeNil())

	Expect(responseObj.Status).To(Equal(mockResponse.Status))
	Expect(responseObj.Message).To(Equal(mockResponse.Message))
}

func deleteUserByIDHandler(mockRepository *mock.MockIRepository, mockService *mock.MockIService, mockResponse pkg.Response) {
	var responseObj pkg.Response

	handler := fibertools.NewHandler(mockRepository, mockService)

	app := fiber.New()
	app.Delete("/user/:id", handler.DeleteUser)

	req := httptest.NewRequest("DELETE", "/user/"+user.MockUser.ID.Hex(), nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	Expect(err).To(BeNil())

	responseBody, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(responseBody, &responseObj)
	Expect(err).To(BeNil())

	Expect(responseObj.Status).To(Equal(mockResponse.Status))
	Expect(responseObj.Message).To(Equal(mockResponse.Message))
}

func getUsersHandler(mockRepository *mock.MockIRepository, mockService *mock.MockIService, mockResponse pkg.Response) {
	var responseObj pkg.Response

	handler := fibertools.NewHandler(mockRepository, mockService)

	app := fiber.New()
	app.Get("/users", handler.GetUsers)

	req := httptest.NewRequest("GET", "/users", nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	Expect(err).To(BeNil())

	responseBody, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(responseBody, &responseObj)
	Expect(err).To(BeNil())

	Expect(responseObj.Status).To(Equal(mockResponse.Status))
	Expect(responseObj.Message).To(Equal(mockResponse.Message))
}
