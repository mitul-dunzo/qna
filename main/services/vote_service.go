package services

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"qna/main/dtos"
)

type VoteService struct {
	db *gorm.DB
}

func NewVoteService(db *gorm.DB) VoteService {
	return VoteService{
		db: db,
	}
}

func (s *VoteService) validateVote(vote int) (string, error) {
	if vote == 1 {
		return dtos.Upvote, nil
	} else if vote == -1 {
		return dtos.Downvote, nil
	}
	return "", errors.New("invalid vote")
}

func (s *VoteService) VoteQuestion(q *dtos.Question, v int, userId uint) error {
	if userId == q.UserId {
		logrus.Error("Voting on your question")
		return errors.New("can't vote on your question")
	}

	voteId, err := s.validateVote(v)
	if err != nil {
		logrus.Error("Invalid vote: ", err.Error())
		return err
	}

	logrus.Debug("q user id: ", q.UserId)
	vote := dtos.Vote{
		UpvotedUserId: q.UserId,
		Type:          "ques",
		EntityId:      q.ID,
		Vote:          voteId,
	}
	err = s.db.Create(&vote).Error
	if err != nil {
		logrus.Error("Couldn't save vote: ", err.Error())
		return err
	}

	return nil
}

func (s *VoteService) VoteAnswer(a *dtos.Answer, v int, userId uint) error {
	if userId == a.UserId {
		logrus.Error("Voting on your answer")
		return errors.New("can't vote on your answer")
	}

	voteId, err := s.validateVote(v)
	if err != nil {
		logrus.Error("Invalid vote: ", err.Error())
		return err
	}

	vote := dtos.Vote{
		UpvotedUserId: a.UserId,
		Type:          "ans",
		EntityId:      a.ID,
		Vote:          voteId,
	}
	err = s.db.Create(&vote).Error
	if err != nil {
		logrus.Error("Couldn't save vote: ", err.Error())
		return err
	}

	return nil
}
