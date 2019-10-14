package services

import (
	"errors"
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

	if !service.canUserAsk(userId) {
		return nil, errors.New("user doesn't have enough votes")
	}

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

func (service *QuestionService) canUserAsk(id uint) bool {
	db := service.db.Where("upvoted_user_id = ?", id)
	var upvoteCount int
	err := db.Where("vote = ?", dtos.Upvote).Find(&dtos.Vote{}).Count(&upvoteCount).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Error("Couldn't fetch upvotes: ", err.Error())
		return false
	}

	var downvoteCount int
	err = db.Where("vote = ?", dtos.Downvote).Find(&dtos.Vote{}).Count(&downvoteCount).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Error("Couldn't fetch downvotes: ", err.Error())
		return false
	}

	return (2*upvoteCount)-downvoteCount >= 20
}
