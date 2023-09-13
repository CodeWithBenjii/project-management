package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL      string
	DatabasePassword string
	DatabaseUsername string
	DatabaseName     string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed Loading .env file: %v", err)
	}

	viper.AutomaticEnv()
	viper.SetConfigType("env")
	viper.SetDefault("DATABASE_URL", "localhost")
	viper.SetDefault("DATABASE_USERNAME", "postgres")
	viper.SetDefault("DATABASE_PASSWORD", "postgres")
	viper.SetDefault("DATABASE_NAME", "postgres")

	config := &Config{
		DatabaseURL:      viper.GetString("DATABASE_URL"),
		DatabasePassword: viper.GetString("DATABASE_PASSWORD"),
		DatabaseUsername: viper.GetString("DATABASE_USERNAME"),
		DatabaseName:     viper.GetString("DATABASE_NAME"),
	}

	return config, nil
}
