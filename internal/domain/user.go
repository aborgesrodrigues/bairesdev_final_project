package domain

import (
	"time"

	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	Username string    `gorm:"not null" json:"username"`
	Name     string    `gorm:"not null" json:"name"`
	Birthday time.Time `gorm:"not null" json:"birthday"`
}
