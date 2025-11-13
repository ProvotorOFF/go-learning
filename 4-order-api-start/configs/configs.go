package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Print("Cannot load .env")
	}
	return &Config{
		DSN: os.Getenv("DSN"),
	}
}
