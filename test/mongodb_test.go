package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"user-service/pkg"
	"user-service/tools/dockertest"
	"user-service/tools/mongodb"
)

var redisConfig = pkg.MongoDBConfig{
	Username: "admin",
	Password: "admin",
	Host:     "localhost",
	Port:     "27017",
	Database: "test",
}

func TestMongoDB(t *testing.T) {

	var dockerContainer = dockertest.NewDockertest("")

	if err := dockerContainer.RunMongoDB(redisConfig); err != nil {
		t.Fatal(err)
	}

	getMongoDBURI(t)
	newMongoDB(t)
	getUserCollection(t)
	getDatabase(t)
	upsert(t)
	isEmailExists(t)
	isNicknameExists(t)
	getUserByID(t)
	getUsers(t)
	deleteUserByID(t)

	dockerContainer.Purge()
}

func getMongoDBURI(t *testing.T) {

	mongo := mongodb.NewMongoDB(redisConfig)
	mongodbURI := mongo.GetMongoDBURI()

	assert.Equal(t, "mongodb://admin:admin@localhost:27017", mongodbURI)
}

func newMongoDB(t *testing.T) {

	mongo := mongodb.NewMongoDB(redisConfig)
	db, err := mongo.Connect()

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, db)
}

func getUserCollection(t *testing.T) {

	db, err := mongodb.NewMongoDB(redisConfig).Connect()

	userCollection := db.GetUserCollection()

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, userCollection)
}

func getDatabase(t *testing.T) {

	db, err := mongodb.NewMongoDB(redisConfig).Connect()

	database := db.GetDatabase()

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, database)
}

func upsert(t *testing.T) {

	db, err := mongodb.NewMongoDB(redisConfig).Connect()

	err = db.Upsert(&userMock)

	assert.Equal(t, nil, err)
}

func isEmailExists(t *testing.T) {

	db, err := mongodb.NewMongoDB(redisConfig).Connect()

	emailExists, err := db.IsEmailExists(userMock.Email)

	assert.Equal(t, nil, err)
	assert.Equal(t, true, emailExists)
}

func isNicknameExists(t *testing.T) {

	db, err := mongodb.NewMongoDB(redisConfig).Connect()

	nicknameExists, err := db.IsNicknameExists(userMock.Nickname)

	assert.Equal(t, nil, err)
	assert.Equal(t, true, nicknameExists)
}

func getUserByID(t *testing.T) {

	db, err := mongodb.NewMongoDB(redisConfig).Connect()

	user, err := db.GetUserByID(userMock.ID)

	assert.Equal(t, nil, err)
	assert.Equal(t, userMock, *user)
}

func getUsers(t *testing.T) {

	db, err := mongodb.NewMongoDB(redisConfig).Connect()

	users, err := db.GetUsers()

	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(users))
}

func deleteUserByID(t *testing.T) {

	db, err := mongodb.NewMongoDB(redisConfig).Connect()

	err = db.DeleteUserByID(userMock.ID)

	assert.Equal(t, nil, err)
}
