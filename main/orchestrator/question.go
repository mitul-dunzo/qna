package orchestrator

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"qna/main/dtos"
	"qna/main/services"
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
	r.HandleFunc("/", orch.getQuestions).Methods(http.MethodGet)
	r.HandleFunc("/new/", orch.addQuestion).Methods(http.MethodPost)
}

func (orch *QuestionOrchestrator) getQuestions(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	pageString := queryParams.Get("page")

	logrus.Debug(pageString)

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
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logrus.Error("Couldn't read from body request: ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var question dtos.UserQuestion
	err = json.Unmarshal(b, &question)
	if err != nil {
		logrus.Error("Incorrect format of question: ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ques := question.Question()
	q, err := orch.q.AddQuestion(&ques)
	if err != nil {
		logrus.Error("Couldn't save question: ", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
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
