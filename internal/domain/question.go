package domain

import "gorm.io/gorm"

// Question struct
type Question struct {
	gorm.Model
	Statement string `gorm:"not null" json:"statement"`
	Answer    string `json:"answer"`
	User      User   `gorm:"not null" json:"user"`
}
