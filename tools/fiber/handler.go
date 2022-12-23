package fiber

import (
	"github.com/gofiber/fiber/v2"
	"user-service/internal/user"
)

type Handler struct {
	repository user.IRepository
	service    user.IService
}

type IHandler interface {
	CreateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
}

func NewHandler(repository user.IRepository, service user.IService) IHandler {
	return &Handler{
		repository: repository,
		service:    service,
	}
}

func (h Handler) CreateUser(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h Handler) DeleteUser(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h Handler) UpdateUser(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetUser(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetUsers(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
