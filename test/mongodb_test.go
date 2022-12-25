package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"user-service/pkg"
	"user-service/tools/mongodb"
)

var config = pkg.MongoDBConfig{
	Username:   "admin",
	Password:   "admin",
	Host:       "localhost",
	Port:       "27017",
	Database:   "user",
	Collection: "users",
}

func Test_GetMongoDBURI(t *testing.T) {

	mongo := mongodb.NewMongoDB(config)

	mongodbURI := mongo.GetMongoDBURI()

	assert.Equal(t, "mongodb://admin:admin@localhost:27017", mongodbURI)
}

func Test_NewMongoDB(t *testing.T) {

	mongo := mongodb.NewMongoDB(config)

	db, err := mongo.Connect()

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, db)
}

func Test_GetUserCollection(t *testing.T) {

	db, err := mongodb.NewMongoDB(config).Connect()

	userCollection := db.GetUserCollection()

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, userCollection)
}

func Test_GetDatabase(t *testing.T) {

	db, err := mongodb.NewMongoDB(config).Connect()

	database := db.GetDatabase()

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, database)
}
