package dtos

import "github.com/jinzhu/gorm"

type Answer struct {
	gorm.Model
	QuestionId uint
	UserId     uint
	Text       string
}

type UserAnswer struct {
	ID   uint
	Text string
}

func (a Answer) UserAnswer() UserAnswer {
	return UserAnswer{
		ID:   a.ID,
		Text: a.Text,
	}
}

func UserAnswers(answers []Answer) []UserAnswer {
	var ans []UserAnswer
	for i := 0; i < len(answers); i++ {
		ans = append(ans, answers[i].UserAnswer())
	}
	return ans
}
