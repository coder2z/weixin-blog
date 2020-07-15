package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Redis struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	PoolSize int    `yaml:"poolSize"`
}

type WX struct {
	AppID     string `yaml:"appID"`
	AppSecret string `yaml:"appSecret"`
}

type Config struct {
	Redis Redis `yaml:"redis"`
	WX    WX    `yaml:"wx"`
}

func (c *Config) GetConf() *Config {
	yamlFile, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		log.Panic(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Panic(err.Error())
	}
	return c
}
