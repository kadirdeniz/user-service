package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"user-service/internal"
	"user-service/tools/dockertest"
	"user-service/tools/mongodb"
	"user-service/tools/redis"
)

var mongoContainer = dockertest.NewDockertest("")
var redisContainer = dockertest.NewDockertest("")

func Test_UserRepository(t *testing.T) {

	if err := mongoContainer.RunMongoDB(mongoConfig); err != nil {
		t.Fatal(err)
	}

	if err := redisContainer.RunRedis(); err != nil {
		t.Fatal(err)
	}

	// Run tests here

	userRepositoryUpsert(t)
	userRepositoryGetUserByID(t)
	userRepositoryGetUsers(t)
	userRepositoryIsEmailExists(t)
	userRepositoryIsNicknameExists(t)
	userRepositoryDeleteUserByID(t)

	mongoContainer.Purge()
	redisContainer.Purge()
}

func userRepositoryUpsert(t *testing.T) {

	var mongoDBInterface = mongodb.NewMongoDB(mongoConfig)
	var redisClient = redis.NewRedis(redisConfig)

	repo := internal.NewRepository(mongoDBInterface, redisClient)

	err := repo.Upsert(userMock)

	assert.Nil(t, err)
}

func userRepositoryDeleteUserByID(t *testing.T) {

	var mongoDBInterface = mongodb.NewMongoDB(mongoConfig)
	var redisClient = redis.NewRedis(redisConfig)

	repo := internal.NewRepository(mongoDBInterface, redisClient)

	err := repo.DeleteUserByID(userMock.ID)

	assert.NotNil(t, err)
}

func userRepositoryGetUserByID(t *testing.T) {

	var mongoDBInterface = mongodb.NewMongoDB(mongoConfig)
	var redisClient = redis.NewRedis(redisConfig)

	repo := internal.NewRepository(mongoDBInterface, redisClient)

	user, err := repo.GetUserByID(userMock.ID)

	assert.Nil(t, err)
	assert.Equal(t, userMock.ID, user.ID)
	assert.Equal(t, userMock.Email, user.Email)
	assert.Equal(t, userMock.Nickname, user.Nickname)
	assert.Equal(t, userMock.FirstName, user.FirstName)
	assert.Equal(t, userMock.LastName, user.LastName)
}

func userRepositoryGetUsers(t *testing.T) {

	var mongoDBInterface = mongodb.NewMongoDB(mongoConfig)
	var redisClient = redis.NewRedis(redisConfig)

	repo := internal.NewRepository(mongoDBInterface, redisClient)

	users, err := repo.GetUsers()

	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, 1, len(users))
}

func userRepositoryIsEmailExists(t *testing.T) {

	var mongoDBInterface = mongodb.NewMongoDB(mongoConfig)
	var redisClient = redis.NewRedis(redisConfig)

	repo := internal.NewRepository(mongoDBInterface, redisClient)

	exists, err := repo.IsEmailExists(userMock.Email)

	assert.Nil(t, err)
	assert.True(t, exists)
}

func userRepositoryIsNicknameExists(t *testing.T) {

	var mongoDBInterface = mongodb.NewMongoDB(mongoConfig)
	var redisClient = redis.NewRedis(redisConfig)

	repo := internal.NewRepository(mongoDBInterface, redisClient)

	exists, err := repo.IsNicknameExists(userMock.Nickname)

	assert.Nil(t, err)
	assert.True(t, exists)
}
