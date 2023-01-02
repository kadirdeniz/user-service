package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"user-service/internal/user"
	"user-service/pkg"
)

var MongoClient *mongo.Database

const UserCollection = "users"

var CTX = context.TODO()

type MongoDB struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

type MongoDBInterface interface {
	GetMongoDBURI() string
	Connect() (*MongoDB, error)
	GetUserCollection() *mongo.Collection
	GetDatabase() *mongo.Database
	Upsert(user *user.User) error
	IsEmailExists(email string) (bool, error)
	IsNicknameExists(nickname string) (bool, error)
	DeleteUserByID(id primitive.ObjectID) error
	GetUserByID(id primitive.ObjectID) (*user.User, error)
	GetUsers() ([]*user.User, error)
	CreateUsers(users []interface{}) error
	FlushUsers() error
}

func NewMongoDB(config pkg.MongoDBConfig) MongoDBInterface {
	return &MongoDB{
		Username: config.Username,
		Password: config.Password,
		Host:     config.Host,
		Port:     config.Port,
		DBName:   config.Database,
	}
}

func (m *MongoDB) GetMongoDBURI() string {
	return "mongodb://" + m.Username + ":" + m.Password + "@" + m.Host + ":" + m.Port
}

func (m *MongoDB) GetUserCollection() *mongo.Collection {
	return MongoClient.Collection(UserCollection)
}

func (m *MongoDB) GetDatabase() *mongo.Database {
	return MongoClient
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

	MongoClient = client.Database(m.DBName)

	return m, nil
}
