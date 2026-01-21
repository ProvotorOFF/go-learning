package user

import (
	"order-api-start/internal/order"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Phone  string `json:"phone" validate:"required,e164" gorm:"uniqueIndex"`
	Orders []order.Order
}
