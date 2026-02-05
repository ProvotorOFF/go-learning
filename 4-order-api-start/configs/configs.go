package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN    string
	Secret string
}

func LoadConfig() *Config {
	env := os.Getenv("APP_ENV")

	envFile := ".env"
	if env == "testing" {
		envFile = ".env.testing"
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("Cannot load %s", envFile)
	}

	return &Config{
		DSN:    os.Getenv("DSN"),
		Secret: os.Getenv("SECRET"),
	}
}
