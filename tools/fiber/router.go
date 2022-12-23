package fiber

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"user-service/internal/user"
)

func Router() {
	err := StartServer(8000)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("User service started")
}

func StartServer(port int) error {

	app := fiber.New()

	handler := NewHandler(
		user.NewRepository(),
		user.NewService(),
	)

	app.Post("/user", handler.CreateUser)
	app.Delete("/user/:id", handler.DeleteUser)
	app.Get("/user/:id", handler.GetUser)
	app.Get("/users", handler.GetUsers)
	app.Put("/user/:id", handler.UpdateUser)

	return app.Listen(fmt.Sprintf(":%d", port))
}
