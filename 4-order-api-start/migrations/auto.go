package main

import (
	"order-api-start/configs"
	"order-api-start/internal/product"
	"order-api-start/internal/session"
	"order-api-start/internal/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	conf := configs.LoadConfig()

	db, err := gorm.Open(postgres.Open(conf.DSN), &gorm.Config{})
	if err != nil {
		panic("No connection to database")
	}

	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&session.Session{})
}
