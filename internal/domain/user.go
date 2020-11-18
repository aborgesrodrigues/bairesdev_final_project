package domain

import (
	"time"
)

// User struct
type User struct {
	ID       uint      `gorm:"primarykey"`
	Username string    `gorm:"not null" json:"username"`
	Name     string    `gorm:"not null" json:"name"`
	Birthday time.Time `gorm:"not null" json:"birthday"`
}
