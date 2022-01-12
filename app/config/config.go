package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

var Config = struct {
	DB struct {
		Host     string `yml:"host"`
		Port     int    `yml:"port"`
		User     string `yml:"user"`
		Password string `yml:"password"`
		DBName   string `yml:"dbname"`
	}
}{}

func GetConfiguration(confPath string) string {
	configor.Load(&Config, confPath)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Config.DB.Host, Config.DB.Port, Config.DB.User, Config.DB.Password, Config.DB.DBName)

	return psqlInfo
}
