package domain

// Question struct
type Question struct {
	ID        uint   `gorm:"primarykey"`
	Statement string `gorm:"not null" json:"statement"`
	Answer    string `json:"answer,omitempty"`
	UserID    uint   `gorm:"not null" json:"user_id,omitempty"`
	User      User   `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT" json:"-"`
}

/*
func (srv QuestionService) checkStruct(question domain.Question, checkID bool) bool {
	// Entry log
	srv.logger.Info("Called checkStruct",
		zap.String("question", fmt.Sprintf("%#v", question)),
		zap.String("checkID", strconv.FormatBool(checkID)),
	)

	if (question.ID == 0 && checkID) ||
		question.Statement == "" ||
		question.UserID == 0 {

		return false
	}

	return true
}
*/
