package mock

import "user-service/pkg"

var MongoConfig = pkg.MongoDBConfig{
	Username: "admin",
	Password: "admin",
	Host:     "localhost",
	Port:     "27017",
	Database: "test",
}

var RedisConfig = pkg.RedisConfig{
	Host: "localhost",
	Port: "6379",
	//Password: "",
	DB: 0,
}
