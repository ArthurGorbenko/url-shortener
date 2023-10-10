package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port             string `mapstructure:"PORT"`
	ConnectionString string `mapstructure:"DATABASE_URL"`
}

var AppConfig *Config

func LoadAppConfig() {
	log.Println("Loading Server Configurations...")

	godotenv.Load(".env")

	AppConfig = &Config{
		ConnectionString: os.Getenv("DATABASE_URL"),
		Port:             os.Getenv("PORT"),
	}

	// Additional validation if needed
	if AppConfig.Port == "" || AppConfig.ConnectionString == "" {
		log.Fatal("Configuration error: PORT and DATABASE_URL must be set")
	}

	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
