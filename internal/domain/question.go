package domain

// Question struct
type Question struct {
	ID        uint   `gorm:"primarykey"`
	Statement string `gorm:"not null" json:"statement"`
	Answer    string `json:"answer,omitempty"`
	UserID    uint   `gorm:"not null" json:"user_id,omitempty"`
	User      User   `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT" json:"-"`
}

// IsEmpty checks if all the fields are filled
func (question Question) IsEmpty(checkID bool) bool {
	if (question.ID == 0 && checkID) ||
		question.Statement == "" ||
		question.UserID == 0 {

		return true
	}

	return false
}
