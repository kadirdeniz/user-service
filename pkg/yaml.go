package pkg

import (
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

type RedisConfig struct {
	URL      string `yaml:"url"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
	Prefix   string `yaml:"prefix"`
}

type Configs struct {
	MongoDB MongoDBConfig `yaml:"mongodb"`
	Redis   RedisConfig   `yaml:"redis"`
}

func NewConfigs() *Configs {
	return &Configs{
		MongoDB: MongoDBConfig{},
		Redis:   RedisConfig{},
	}
}

func (c *Configs) ReadYamlFile(path string) error {

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Configs) ReadConfigFiles() error {
	err := c.ReadYamlFile("configs/mongodb.yaml")
	if err != nil {
		return err
	}

	err = c.ReadYamlFile("configs/redis.yaml")
	if err != nil {
		return err
	}

	return nil
}
