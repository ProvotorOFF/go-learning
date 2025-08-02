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
	key, ok := os.LookupEnv("KEY")
	if !ok {
		fmt.Println("Нет KEY - дальнейшее выполнение может завершиться с ошибками!")
	}
	return &Config{key}
}
