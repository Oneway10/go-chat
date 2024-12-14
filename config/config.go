package config

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/viper"
)

var Mysql *mysqlConfig

type mysqlConfig struct {
	ServerHost string
	ServerPort string
	Username   string
	Password   string
	Database   string
}

func Init() {
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath("../../conf")
	err := viper.ReadInConfig()
	if err != nil {
		hlog.Fatal("error reading config file, %s", err)
	}

	Mysql = loadMysqlConfig()
}

func loadMysqlConfig() *mysqlConfig {
	s := "Mysql."
	return &mysqlConfig{
		ServerHost: viper.GetString(s + "ServerHost"),
		ServerPort: viper.GetString(s + "ServerPort"),
		Username:   viper.GetString(s + "Username"),
		Password:   viper.GetString(s + "Password"),
		Database:   viper.GetString(s + "Database"),
	}
}
