package orchestrator

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"qna/main/constants"
	"qna/main/dtos"
	"qna/main/services"
	"qna/main/utils"
	"strconv"
)

type QuestionOrchestrator struct {
	q *services.QuestionService
	a *services.AnswerService
}

func NewQuestionOrchestrator(q *services.QuestionService, a *services.AnswerService) QuestionOrchestrator {
	return QuestionOrchestrator{
		q: q,
		a: a,
	}
}

func (orch *QuestionOrchestrator) Handle(r *mux.Router) {
	utils.Instrument(r, constants.GetQuestionsEp, orch.getQuestions).Methods(http.MethodGet)
	utils.Instrument(r, constants.GetQuestionEp, orch.getQuestion).Methods(http.MethodGet)
	utils.Instrument(r, constants.AddQuestionEp, orch.addQuestion).Methods(http.MethodPost)
	utils.Instrument(r, constants.AddAnswerEp, orch.newAnswer).Methods(http.MethodPost)
}

func (orch *QuestionOrchestrator) getQuestions(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	pageString := queryParams.Get("page")
	pageNo, err := strconv.Atoi(pageString)
	if err != nil {
		logrus.Error("Couldn't get page number: ", err.Error())
		http.Error(w, "Add page number", http.StatusBadRequest)
		return
	}

	questions := orch.q.GetQuestions(pageNo)

	respObj := map[string]interface{}{
		"questions": dtos.UserQuestions(questions),
		"page_no":   pageNo,
	}
	resp, err := json.Marshal(respObj)
	if err != nil {
		logrus.Error("Couldn't parse question: ", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(resp)
}

func (orch *QuestionOrchestrator) addQuestion(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserId(r)
	if err != nil {
		utils.SendBadRequestError(w)
		return
	}

	var question dtos.UserQuestion
	err = utils.ReadHTTPBody(r, &question)
	if err != nil {
		utils.SendBadRequestError(w)
		return
	}

	ques := question.Question()
	q, err := orch.q.AddQuestion(&ques, userId)
	if err != nil {
		logrus.Error("Couldn't save question: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	question = q.UserQuestion(nil)
	resp, err := json.Marshal(question)
	if err != nil {
		logrus.Error("Can't convert body into json: ", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}

	_, _ = w.Write(resp)
}

func (orch *QuestionOrchestrator) getQuestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	ques, err := orch.q.GetQuestion(uint(id))
	if err != nil {
		logrus.Error("Couldn't find question")
		http.Error(w, "Question doesn't exist", http.StatusBadRequest)
		return
	}

	ans := orch.a.GetAnswersForQuestion(uint(id))
	uans := dtos.UserAnswers(ans)
	q := ques.UserQuestion(&uans)

	resp, err := json.Marshal(&q)
	if err != nil {
		logrus.Error("Could't convert question to json: ", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(resp)
}

func (orch *QuestionOrchestrator) newAnswer(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserId(r)
	if err != nil {
		utils.SendBadRequestError(w)
		return
	}

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var answer dtos.UserAnswer
	err = utils.ReadHTTPBody(r, &answer)
	if err != nil {
		utils.SendBadRequestError(w)
		return
	}

	ans := answer.Answer(uint(id))
	a, err := orch.a.NewAnswer(&ans, userId)
	if err != nil {
		logrus.Error("Failed to save answer: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	answer = a.UserAnswer()
	resp, err := json.Marshal(answer)
	if err != nil {
		logrus.Error("Can't convert body into json: ", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}

	_, _ = w.Write(resp)
}
