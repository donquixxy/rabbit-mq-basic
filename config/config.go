package config

import (
	"log"
	"micro-company/database/mysql"
	"micro-company/database/postgre"

	"github.com/spf13/viper"
)

type Application struct {
	Name   string `yaml:"name"`
	Server string `yaml:"server"`
}

type AppConfiguration struct {
	Application
	Database *postgre.PostgreConfig
	Mysql    *mysql.MysqlConfig
}

func InitConfig() *AppConfiguration {
	return readConfig()
}

func readConfig() *AppConfiguration {
	viper.AddConfigPath("./")

	viper.SetConfigName("config")

	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("failed to read config %v", err)
	}

	var appConfig *AppConfiguration

	err = viper.Unmarshal(&appConfig)

	if err != nil {
		log.Fatalf("failed to unmarshal config %v", err)
	}

	log.Println("Success reading config")
	return appConfig
}
