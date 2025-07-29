package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env не загружен. Будут использованы переменные окружения")
	}
	key := os.Getenv("KEY")
	if key == "" {
		panic("Приложение невозможно запустить без ключа")
	}
	return &Config{key}
}
