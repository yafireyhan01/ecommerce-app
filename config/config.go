package config

import (
	"github.com/spf13/viper"
	"log"
)

var AppConfig *Config

type Config struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	APIPort    string
	JWTSecret  string
}

func LoadConfig() {
	viper.SetConfigFile("app.env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	viper.AutomaticEnv() // read in environment variables that match

	AppConfig = &Config{
		DBDriver:   viper.GetString("DB_DRIVER"),
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBName:     viper.GetString("DB_NAME"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		APIPort:    viper.GetString("API_PORT"),
		JWTSecret:  viper.GetString("JWT_SECRET"),
	}
}
