package constants

import "errors"

var (
	NoUserIdPresent          = errors.New("no user id")
	WrongTokenError          = errors.New("wrong token")
	NotEnoughVotesError      = errors.New("user doesn't have enough votes")
	InvalidVoteError         = errors.New("invalid vote")
	VotingOnOwnQuestionError = errors.New("can't vote on your own question")
	VotingOnOwnAnswerError   = errors.New("can't vote on your answer")
	FailedToSendSmsError     = errors.New("failed to send SMS")
	InvalidOtpError          = errors.New("invalid otp")
)
