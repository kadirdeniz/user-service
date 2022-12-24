package pkg

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type MongoDBConfig struct {
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Database   string `yaml:"database"`
	Collection string `yaml:"collection"`
}

type Configs struct {
	MongoDB MongoDBConfig
}

func NewConfigs() *Configs {
	return &Configs{
		MongoDB: MongoDBConfig{},
	}
}

func (c *Configs) ReadYamlFile(path string, object interface{}) (interface{}, error) {

	var mongoDBConfig MongoDBConfig

	yamlFile, err := ioutil.ReadFile("configs/mongodb.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, mongoDBConfig)
	if err != nil {
		return nil, err
	}
	fmt.Println(mongoDBConfig)
	return object, nil
}

func (c *Configs) ReadConfigFiles() (*Configs, error) {
	mongoDBConfig, err := c.ReadYamlFile("configs/mongodb.yaml", c.MongoDB)
	if err != nil {
		return nil, err
	}

	c.MongoDB = mongoDBConfig.(MongoDBConfig)

	return c, nil
}
