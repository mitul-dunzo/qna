package services

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"qna/main/constants"
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
		return constants.Upvote, nil
	} else if vote == -1 {
		return constants.Downvote, nil
	}
	return "", constants.InvalidVoteError
}

func (s *VoteService) VoteQuestion(q *dtos.Question, v int, userId uint) error {
	if userId == q.UserId {
		logrus.Error("Voting on your question")
		return constants.VotingOnOwnQuestionError
	}

	voteId, err := s.validateVote(v)
	if err != nil {
		logrus.Error("Invalid vote: ", err.Error())
		return err
	}

	vote := dtos.Vote{
		UpvotedUserId: q.UserId,
		Type:          constants.QuesVote,
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
		return constants.VotingOnOwnAnswerError
	}

	voteId, err := s.validateVote(v)
	if err != nil {
		logrus.Error("Invalid vote: ", err.Error())
		return err
	}

	vote := dtos.Vote{
		UpvotedUserId: a.UserId,
		Type:          constants.AnsVote,
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
