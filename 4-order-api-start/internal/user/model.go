package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone string `json:"phone" validate:"required,e164" gorm:"uniqueIndex"`
}
