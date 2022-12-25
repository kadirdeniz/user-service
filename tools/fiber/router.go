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

var Configs = pkg.NewConfigs()

func Router() {

	// Read config files
	if err := Configs.ReadConfigFiles(); err != nil {
		log.Fatal(err)
	}

	err := StartServer(8000)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("User service started")
}

func StartServer(port int) error {

	// database connection
	mongoDBConnection, err := mongodb.NewMongoDB(Configs.MongoDB).Connect()
	if err != nil {
		return err
	}

	// redis connection
	redisConnection := redis.NewRedis(Configs.Redis)

	handler := NewHandler(
		user.NewRepository(mongoDBConnection.GetUserCollection(), redisConnection),
		user.NewService(),
	)

	app := fiber.New()

	app.Post("/user", handler.CreateUser)
	app.Put("/user/:id", handler.UpdateUser)
	app.Delete("/user/:id", handler.DeleteUser)
	app.Get("/user/:id", handler.GetUser)
	app.Get("/users", handler.GetUsers)

	return app.Listen(fmt.Sprintf(":%d", port))
}
