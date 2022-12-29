package fiber

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"user-service/internal"
	"user-service/internal/user"
	"user-service/pkg"
	"user-service/pkg/dto"
)

type Handler struct {
	repository internal.IRepository
	service    user.IService
}

type IHandler interface {
	CreateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
}

func NewHandler(repository internal.IRepository, service user.IService) IHandler {
	return &Handler{
		repository: repository,
		service:    service,
	}
}

func (h Handler) CreateUser(c *fiber.Ctx) error {

	var CreateUserRequest dto.CreateUserRequest

	if err := c.BodyParser(&CreateUserRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.SomethingWentWrongResponse)
	}

	isEmailExists, err := h.repository.IsEmailExists(CreateUserRequest.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.SomethingWentWrongResponse)
	}

	if isEmailExists {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.EmailAlreadyExistsResponse)
	}

	isNicknameExists, err := h.repository.IsNicknameExists(CreateUserRequest.Nickname)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.SomethingWentWrongResponse)
	}

	if isNicknameExists {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.NicknameAlreadyExistsResponse)
	}

	userObj := user.New(
		CreateUserRequest.FirstName,
		CreateUserRequest.LastName,
		CreateUserRequest.Nickname,
		CreateUserRequest.Email,
		CreateUserRequest.Password,
	)

	if err := h.repository.Upsert(userObj); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.SomethingWentWrongResponse)
	}

	return c.Status(fiber.StatusOK).JSON(pkg.UserCreatedSuccessResponse)
}

func (h Handler) DeleteUser(c *fiber.Ctx) error {

	userId, _ := primitive.ObjectIDFromHex(c.Params("id"))

	userObj, err := h.repository.GetUserByID(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.SomethingWentWrongResponse)
	}

	if userObj.ID == primitive.NilObjectID {
		return c.Status(fiber.StatusNotFound).JSON(pkg.UserNotFoundResponse)
	}

	if err := h.repository.DeleteUserByID(userId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.SomethingWentWrongResponse)
	}

	return c.Status(fiber.StatusOK).JSON(pkg.UserDeletedSuccessResponse)
}

func (h Handler) UpdateUser(c *fiber.Ctx) error {

	userId, _ := primitive.ObjectIDFromHex(c.Params("id"))

	userObj, err := h.repository.GetUserByID(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.SomethingWentWrongResponse)
	}

	if userObj.ID == primitive.NilObjectID {
		return c.Status(fiber.StatusNotFound).JSON(pkg.UserNotFoundResponse)
	}

	var UpdateUserRequest dto.UpdateUserRequest

	if err := pkg.JSONEncoder(c.Body(), &UpdateUserRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.SomethingWentWrongResponse)
	}

	isEmailExists, err := h.repository.IsEmailExists(UpdateUserRequest.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.SomethingWentWrongResponse)
	}

	if isEmailExists {
		return c.Status(fiber.StatusConflict).JSON(pkg.EmailAlreadyExistsResponse)
	}

	isNicknameExists, err := h.repository.IsNicknameExists(UpdateUserRequest.Nickname)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.SomethingWentWrongResponse)
	}

	if isNicknameExists {
		return c.Status(fiber.StatusConflict).JSON(pkg.NicknameAlreadyExistsResponse)
	}

	if UpdateUserRequest.FirstName != "" {
		userObj.FirstName = UpdateUserRequest.FirstName
	}

	if UpdateUserRequest.LastName != "" {
		userObj.LastName = UpdateUserRequest.LastName
	}

	if UpdateUserRequest.Nickname != "" {
		userObj.Nickname = UpdateUserRequest.Nickname
	}

	if UpdateUserRequest.Email != "" {
		userObj.Email = UpdateUserRequest.Email
	}

	if err := h.repository.Upsert(userObj); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.SomethingWentWrongResponse)
	}

	return c.Status(fiber.StatusOK).JSON(pkg.UserUpdatedSuccessResponse)
}

func (h Handler) GetUser(c *fiber.Ctx) error {
	userId, _ := primitive.ObjectIDFromHex(c.Params("id"))

	userObj, err := h.repository.GetUserByID(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.SomethingWentWrongResponse)
	}

	if userObj.ID == primitive.NilObjectID {
		return c.Status(fiber.StatusNotFound).JSON(pkg.UserNotFoundResponse)
	}

	return c.Status(fiber.StatusOK).JSON(pkg.NewResponse(true, "User found", userObj))
}

func (h Handler) GetUsers(c *fiber.Ctx) error {
	users, err := h.repository.GetUsers()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.SomethingWentWrongResponse)
	}

	return c.Status(fiber.StatusOK).JSON(pkg.NewResponse(true, "Users found", users))
}
