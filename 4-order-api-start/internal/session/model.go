package session

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	SID       string    `gorm:"size:64;uniqueIndex" validate:"required" json:"sid"`
	UserID    uint      `gorm:"index;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	Confirmed bool      `gorm:"default:false"`
	Code      string    `gorm:"size:255;not null" validate:"required" json:"code"`
}
