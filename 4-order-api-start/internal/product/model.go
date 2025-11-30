package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name" validate:"required"`
	Description string         `json:"description" validate:"required"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]" validate:"required,dive,url"`
}
