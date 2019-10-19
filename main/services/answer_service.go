package services

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"qna/main/constants"
	"qna/main/dtos"
)

type AnswerService struct {
	db *gorm.DB
}

func NewAnswerService(db *gorm.DB) AnswerService {
	return AnswerService{db: db}
}

func (service *AnswerService) GetAnswersForQuestion(id uint) []dtos.Answer {
	var answers []dtos.Answer
	service.db.Where(constants.QuestionIdQuery, id).Find(&answers)
	return answers
}

func (service *AnswerService) NewAnswer(answer *dtos.Answer, userId uint) (*dtos.Answer, error) {
	answer.UserId = userId
	err := service.db.Create(answer).Error
	if err != nil {
		logrus.Error("Couldn't save answer: ", err.Error())
		return nil, err
	}
	return answer, nil
}

func (service *AnswerService) GetAnswer(id uint) (*dtos.Answer, error) {
	ans := dtos.Answer{
		ID: id,
	}
	err := service.db.First(&ans, ans).Error
	if err != nil {
		logrus.Error("Couldn't find answer: ", err.Error())
		return nil, err
	}

	return &ans, nil
}
