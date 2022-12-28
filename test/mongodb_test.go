package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"user-service/pkg"
	"user-service/tools/dockertest"
	"user-service/tools/mongodb"
)

var mongoConfig = pkg.MongoDBConfig{
	Username: "admin",
	Password: "admin",
	Host:     "localhost",
	Port:     "27017",
	Database: "test",
}

func TestMongoDB(t *testing.T) {

	var dockerContainer = dockertest.NewDockertest("")

	if err := dockerContainer.RunMongoDB(mongoConfig); err != nil {
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

	mongo := mongodb.NewMongoDB(mongoConfig)
	mongodbURI := mongo.GetMongoDBURI()

	assert.Equal(t, "mongodb://"+mongoConfig.Username+":"+mongoConfig.Password+"@"+mongoConfig.Host+":"+mongoConfig.Port+"", mongodbURI)
}

func newMongoDB(t *testing.T) {

	mongo := mongodb.NewMongoDB(mongoConfig)
	db, err := mongo.Connect()

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, db)
}

func getUserCollection(t *testing.T) {

	db, err := mongodb.NewMongoDB(mongoConfig).Connect()

	userCollection := db.GetUserCollection()

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, userCollection)
}

func getDatabase(t *testing.T) {

	db, err := mongodb.NewMongoDB(mongoConfig).Connect()

	database := db.GetDatabase()

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, database)
}

func upsert(t *testing.T) {

	db, err := mongodb.NewMongoDB(mongoConfig).Connect()

	err = db.Upsert(userMock)

	assert.Equal(t, nil, err)
}

func isEmailExists(t *testing.T) {

	db, err := mongodb.NewMongoDB(mongoConfig).Connect()

	emailExists, err := db.IsEmailExists(userMock.Email)

	assert.Equal(t, nil, err)
	assert.Equal(t, true, emailExists)
}

func isNicknameExists(t *testing.T) {

	db, err := mongodb.NewMongoDB(mongoConfig).Connect()

	nicknameExists, err := db.IsNicknameExists(userMock.Nickname)

	assert.Equal(t, nil, err)
	assert.Equal(t, true, nicknameExists)
}

func getUserByID(t *testing.T) {

	db, err := mongodb.NewMongoDB(mongoConfig).Connect()

	user, err := db.GetUserByID(userMock.ID)

	assert.Equal(t, nil, err)
	assert.Equal(t, userMock, user)
}

func getUsers(t *testing.T) {

	db, err := mongodb.NewMongoDB(mongoConfig).Connect()

	users, err := db.GetUsers()

	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(users))
}

func deleteUserByID(t *testing.T) {

	db, err := mongodb.NewMongoDB(mongoConfig).Connect()

	err = db.DeleteUserByID(userMock.ID)

	assert.Equal(t, nil, err)
}
