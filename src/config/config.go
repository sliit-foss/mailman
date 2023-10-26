package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

type Config struct {
	Port                  string `mapstructure:"PORT"`
	MongoConnectionString string `mapstructure:"MONGO_CONNECTION_STRING"`
}

var Env *Config

func Load() {
	Env = loadEnvVariables()
}

func loadEnvVariables() (config *Config) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
