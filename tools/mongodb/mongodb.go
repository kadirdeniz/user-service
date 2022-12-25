package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"user-service/pkg"
)

type MongoDB struct {
	Username   string
	Password   string
	Host       string
	Port       string
	DBName     string
	Collection string
	Database   *mongo.Database
}

type MongoDBInterface interface {
	GetMongoDBURI() string
	Connect() (*MongoDB, error)
	GetUserCollection() *mongo.Collection
	GetDatabase() *mongo.Database
}

func NewMongoDB(config pkg.MongoDBConfig) MongoDBInterface {
	return &MongoDB{
		Username:   config.Username,
		Password:   config.Password,
		Host:       config.Host,
		Port:       config.Port,
		DBName:     config.Database,
		Collection: config.Collection,
	}
}

func (m *MongoDB) GetMongoDBURI() string {
	return "mongodb://" + m.Username + ":" + m.Password + "@" + m.Host + ":" + m.Port
}

func (m *MongoDB) GetUserCollection() *mongo.Collection {
	return m.Database.Collection(m.Collection)
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

	m.Database = client.Database(m.DBName)

	return m, nil
}
