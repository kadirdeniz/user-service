package pkg

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

var AppConfigs Configs

type MongoDBConfig struct {
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Database   string `yaml:"database"`
	Collection string `yaml:"collection"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Configs struct {
	MongoDB MongoDBConfig `yaml:"mongodb"`
	Redis   RedisConfig   `yaml:"redis"`
}

func NewConfigs() *Configs {
	return &AppConfigs
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

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Join(filepath.Dir(b), "./..")

	err := c.ReadYamlFile(filepath.Join(basepath, "./configs/mongodb.yaml"))
	if err != nil {
		return err
	}

	err = c.ReadYamlFile(filepath.Join(basepath, "./configs/redis.yaml"))
	if err != nil {
		return err
	}

	return nil
}
