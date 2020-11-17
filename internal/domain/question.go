package domain

import "gorm.io/gorm"

// Question struct
type Question struct {
	gorm.Model
	Statement string `gorm:"not null" json:"statement"`
	Answer    string `json:"answer"`
	UserID    uint   `gorm:"not null" json:"user_id,omitempty"`
	User      User   `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT" json:"user"`
}
