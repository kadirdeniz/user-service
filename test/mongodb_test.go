package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"user-service/tools/mongodb"
)

func Test_GetMongoDBURI(t *testing.T) {

	mongo := mongodb.NewMongoDB("admin", "admin", "localhost", "27017")

	mongodbURI := mongo.GetMongoDBURI()

	assert.Equal(t, "mongodb://admin:admin@localhost:27017", mongodbURI)
}

func Test_NewMongoDB(t *testing.T) {

	mongo := mongodb.NewMongoDB("admin", "admin", "localhost", "27017")

	db, err := mongo.Connect()

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, db)
}

func Test_GetUserCollection(t *testing.T) {

	db, err := mongodb.NewMongoDB("admin", "admin", "localhost", "27017").Connect()

	userCollection := db.GetUserCollection()

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, userCollection)
}
