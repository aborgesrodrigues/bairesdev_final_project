package domain

// Question struct
type Question struct {
	ID        uint   `gorm:"primarykey"`
	Statement string `gorm:"not null" json:"statement"`
	Answer    string `json:"answer"`
	UserID    uint   `gorm:"not null" json:"user_id,omitempty"`
	User      User   `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT" json:"user"`
}
