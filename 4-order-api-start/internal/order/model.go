package order

import (
	"order-api-start/internal/product"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID   uint              `gorm:"index;not null" json:"user_id" validate:"required,gt=0"`
	Products []product.Product `gorm:"many2many:order_products" json:"product_ids" validate:"required,min=1,dive,gt=0"`
}
