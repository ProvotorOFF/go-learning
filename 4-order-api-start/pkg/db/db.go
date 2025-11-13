package db

import (
	"errors"
	"order-api-start/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) (*Db, error) {
	db, err := gorm.Open(postgres.Open(conf.DSN), &gorm.Config{})
	if err != nil {
		return nil, errors.New("no connection to database")
	}
	return &Db{db}, nil
}
