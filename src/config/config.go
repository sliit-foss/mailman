package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

type Config struct {
	Port                  string `mapstructure:"PORT"`
	MongoConnectionString string `mapstructure:"MONGO_CONNECTION_STRING"`
	JWTSecret             string `mapstructure:"JWT_SECRET"`
}

var Env *Config

func Load() {
	Env = load()
}

func setDefaults() {
	viper.SetDefault("PORT", "8080")
}

func load() (config *Config) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
