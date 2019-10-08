package services

import (
	"github.com/jinzhu/gorm"
	"qna/main/dtos"
)

type AnswerService struct {
	db *gorm.DB
}

func NewAnswerService(db *gorm.DB) AnswerService {
	return AnswerService{db: db}
}

func (service *AnswerService) GetAnswers(id uint) []dtos.Answer {
	var answers []dtos.Answer
	service.db.Find(&answers)
	return answers
}
