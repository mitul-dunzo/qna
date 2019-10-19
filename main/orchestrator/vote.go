package orchestrator

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"qna/main/constants"
	"qna/main/dtos"
	"qna/main/services"
	"qna/main/utils"
)

type VoteOrchestrator struct {
	q *services.QuestionService
	a *services.AnswerService
	v *services.VoteService
}

func NewVoteOrchestrator(q *services.QuestionService, a *services.AnswerService, v *services.VoteService) VoteOrchestrator {
	return VoteOrchestrator{
		q: q,
		a: a,
		v: v,
	}
}

func (orch *VoteOrchestrator) Handle(r *mux.Router) {
	utils.Instrument(r, constants.VoteEp, orch.vote).Methods(http.MethodPost)
}

func (orch *VoteOrchestrator) vote(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserId(r)
	if err != nil {
		utils.SendBadRequestError(w)
		return
	}

	var vote dtos.UserVote
	err = utils.ReadHTTPBody(r, &vote)
	if err != nil {
		utils.SendBadRequestError(w)
		return
	}

	if vote.QuestionId > 0 && vote.AnswerId > 0 {
		logrus.Error("Trying to vote on both ques and answer")
		http.Error(w, "Send either question or answer", http.StatusBadRequest)
		return
	}

	if vote.QuestionId > 0 {
		var q *dtos.Question
		q, err = orch.q.GetQuestion(vote.QuestionId)
		if err != nil {
			logrus.Error("Invalid question: ", err.Error())
			http.Error(w, "Invalid question id", http.StatusBadRequest)
			return
		}
		err = orch.v.VoteQuestion(q, vote.Vote, userId)
	} else {
		var a *dtos.Answer
		a, err = orch.a.GetAnswer(vote.AnswerId)
		if err != nil {
			logrus.Error("Invalid answer: ", err.Error())
			http.Error(w, "Invalid answer id", http.StatusBadRequest)
			return
		}
		err = orch.v.VoteAnswer(a, vote.Vote, userId)
	}

	if err != nil {
		logrus.Error("Couldn't vote: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
