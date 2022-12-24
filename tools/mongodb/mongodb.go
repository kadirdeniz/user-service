package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Username string
	Password string
	Host     string
	Port     string
	Database *mongo.Database
}

type MongoDBInterface interface {
	GetMongoDBURI() string
	Connect() (*MongoDB, error)
	GetUserCollection() *mongo.Collection
	GetDatabase() *mongo.Database
}

func NewMongoDB(username, password, host, port string) MongoDBInterface {
	return &MongoDB{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
	}
}

func (m *MongoDB) GetMongoDBURI() string {
	return "mongodb://" + m.Username + ":" + m.Password + "@" + m.Host + ":" + m.Port
}

func (m *MongoDB) GetUserCollection() *mongo.Collection {
	return m.Database.Collection("user")
}

func (m *MongoDB) GetDatabase() *mongo.Database {
	return m.Database
}

func (m *MongoDB) Connect() (*MongoDB, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.GetMongoDBURI()))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	m.Database = client.Database("user")

	return m, nil
}
