package dockertest

import (
	"github.com/ory/dockertest/v3"
	"log"
	"user-service/tools/mongodb"
)

func NewDockerTest() {
	// Setup
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	environmentVariables := []string{
		"MONGO_INITDB_ROOT_USERNAME=root",
		"MONGO_INITDB_ROOT_PASSWORD=password",
	}

	resource, err := pool.Run("mongo", "5.0", environmentVariables)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err = pool.Retry(func() error {
		_, err := mongodb.NewMongoDB("root", "password", "localhost", resource.GetPort("27017/tcp")).Connect()
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}
