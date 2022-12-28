package test

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
	"user-service/internal/user"
	"user-service/pkg"
	"user-service/tools/dockertest"
	"user-service/tools/redis"
)

var userMock = &user.User{
	ID:        primitive.NewObjectID(),
	FirstName: "John",
	LastName:  "Doe",
	Nickname:  "johndoe",
	Email:     "johndoe@mail.com",
	Password:  "password",
}

var redisConfig = pkg.RedisConfig{
	Host: "localhost",
	Port: "6379",
	//Password: "",
	DB: 0,
}

func TestRedis(t *testing.T) {

	var dockerContainer = dockertest.NewDockertest("")

	if err := dockerContainer.RunRedis(redisConfig); err != nil {
		t.Fatal(err)
	}

	newRedis(t)
	set(t)
	get(t)

	dockerContainer.Purge()
}

func newRedis(t *testing.T) {

	redis := redis.NewRedis(redisConfig)
	assert.NotNil(t, redis)
}

func set(t *testing.T) {

	redis := redis.NewRedis(redisConfig)
	redis.Connect()

	err := redis.SetUser(userMock, time.Minute*20)
	assert.Nil(t, err)
}

func get(t *testing.T) {

	redis := redis.NewRedis(redisConfig)
	redis.Connect()

	user, err := redis.GetUserByID(userMock.ID)
	assert.Nil(t, err)
	assert.Equal(t, userMock.ID, user.ID)
	assert.Equal(t, userMock.FirstName, user.FirstName)
	assert.Equal(t, userMock.LastName, user.LastName)
	assert.Equal(t, userMock.Nickname, user.Nickname)
	assert.Equal(t, userMock.Email, user.Email)
}
