package dockertest

import (
	"errors"
	"github.com/ory/dockertest/v3"
	"log"
	"user-service/pkg"
	"user-service/tools/mongodb"
)

type DockerTest struct{}

var MongoDBEnvirontmentVariables = []string{
	"MONGO_INITDB_ROOT_USERNAME=" + pkg.AppConfigs.MongoDB.Username,
	"MONGO_INITDB_ROOT_PASSWORD=" + pkg.AppConfigs.MongoDB.Password,
	"MONGO_INITDB_DATABASE=" + pkg.AppConfigs.MongoDB.Database,
}

const MongoDBImage = "mongo"
const MongoDBTag = "5.0"

type Dockertest struct {
	Pool     *dockertest.Pool
	Resource *dockertest.Resource
}

func NewDockertest(endpoint string) *Dockertest {
	pool, err := dockertest.NewPool(endpoint)
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return &Dockertest{
		Pool: pool,
	}
}

func (d *Dockertest) RunMongoDB() error {

	var err error

	d.Resource, err = d.Pool.Run(MongoDBImage, MongoDBTag, MongoDBEnvirontmentVariables)
	if err != nil {
		return errors.New("Could not start resource")
	}

	if err = d.Pool.Retry(func() error {
		_, err := mongodb.NewMongoDB(pkg.AppConfigs.MongoDB).Connect()
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.New("Could not connect to docker")
	}

	return nil
}

func (d *Dockertest) Purge() error {
	if err := d.Pool.Purge(d.Resource); err != nil {
		return errors.New("Could not purge resource")
	}
	return nil
}
