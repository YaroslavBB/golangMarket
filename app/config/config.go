package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

type Config struct {
	DB struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"db"`
}

func NewConfig(confPath string) *Config {
	var conf Config

	configor.Load(&conf, confPath)

	return &conf
}

func (c *Config) GetConfiguration() string {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.DB.Host, c.DB.Port, c.DB.User, c.DB.Password, c.DB.DBName)

	return psqlInfo
}
