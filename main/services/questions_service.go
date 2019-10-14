package services

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"qna/main/dtos"
)

type QuestionService struct {
	db *gorm.DB
}

func NewQuestionService(db *gorm.DB) QuestionService {
	return QuestionService{db: db}
}

func (service *QuestionService) GetQuestions(page int) *[]dtos.Question {
	var questions []dtos.Question

	var limit = 10
	service.db.Offset((page - 1) * limit).Limit(limit).Find(&questions)
	return &questions
}

func (service *QuestionService) AddQuestion(q *dtos.Question, userId uint) (*dtos.Question, error) {
	q.UserId = userId
	err := service.db.Create(q).Error
	if err != nil {
		logrus.Error("Couldn't save question: ", err.Error())
		return nil, err
	}
	return q, nil
}

func (service *QuestionService) GetQuestion(id uint) (*dtos.Question, error) {
	var question dtos.Question
	err := service.db.Where(&dtos.Question{
		ID: id,
	}).First(&question).Error
	if err != nil {
		logrus.Error("Couldn't find question", err.Error())
		return nil, err
	}
	return &question, nil
}
