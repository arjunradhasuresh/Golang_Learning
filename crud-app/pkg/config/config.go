package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *Config

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
}

type ServerConfig struct {
	Port string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func InitConfig() {
	viper.SetConfigFile("conf/config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed to read config file: " + err.Error())
	}
	if viper.Unmarshal(&config); err != nil {
		log.Fatal("Error in unmarshall in config")
	}
	log.Println("Successfully loaded configuraion")
}

func GetConfig() *Config {
	return config
}
