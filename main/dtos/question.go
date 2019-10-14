package dtos

type Question struct {
	ID     uint `gorm:"PRIMARY_KEY"`
	UserId uint `json:"user_id"`
	Text   string
}

func (Question) TableName() string {
	return "questions"
}

type UserQuestion struct {
	ID      uint          `json:"id"`
	Text    string        `json:"text"`
	Answers *[]UserAnswer `json:"answers"`
}

func (q UserQuestion) Question() Question {
	return Question{
		ID:   q.ID,
		Text: q.Text,
	}
}

func (q Question) UserQuestion(answers *[]UserAnswer) UserQuestion {
	return UserQuestion{
		ID:      q.ID,
		Text:    q.Text,
		Answers: answers,
	}
}

func UserQuestions(questions *[]Question) []UserQuestion {
	var ques []UserQuestion
	qs := *questions
	for i := 0; i < len(qs); i++ {
		ques = append(ques, qs[i].UserQuestion(nil))
	}
	return ques
}
