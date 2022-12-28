package fiber

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"user-service/internal"
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

	// Read config files
	if err := pkg.NewConfigs().ReadConfigFiles(); err != nil {
		log.Fatal(err)
	}

	// Connect to MongoDB
	db, err := mongodb.NewMongoDB(pkg.AppConfigs.MongoDB).Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to Redis
	redis, err := redis.NewRedis(pkg.AppConfigs.Redis).Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create repository
	repository := internal.NewRepository(db, redis)

	// Create service
	service := user.NewService()

	// Create handler
	handler := NewHandler(repository, service)

	app := fiber.New()

	user := app.Group("/user")

	user.Post("/", handler.CreateUser)
	user.Put("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUser)
	user.Get("/:id", handler.GetUser)
	user.Get("/", handler.GetUsers)

	return app.Listen(fmt.Sprintf(":%d", port))
}
