package fiber

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"user-service/internal/user"
	"user-service/pkg"
	"user-service/tools/mongodb"
	"user-service/tools/redis"
)

func Router() {
	err := StartServer(8000)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("User service started")
}

func StartServer(port int) error {

	// Connect to MongoDB
	db, err := mongodb.NewMongoDB(pkg.AppConfigs.MongoDB).Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to Redis
	redis := redis.NewRedis(pkg.AppConfigs.Redis)

	// Create repository
	repository := user.NewRepository(db.GetDatabase(), redis)

	// Create service
	service := user.NewService()

	// Create handler
	handler := NewHandler(repository, service)

	app := fiber.New()

	app.Post("/user", handler.CreateUser)
	app.Put("/user/:id", handler.UpdateUser)
	app.Delete("/user/:id", handler.DeleteUser)
	app.Get("/user/:id", handler.GetUser)
	app.Get("/users", handler.GetUsers)

	return app.Listen(fmt.Sprintf(":%d", port))
}
