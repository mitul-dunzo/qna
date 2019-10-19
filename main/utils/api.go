package utils

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"qna/main/constants"
	"reflect"
)

func ReadHTTPBody(req *http.Request, body interface{}) error {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logrus.Error("Couldn't read from body: ", err.Error())
		return err
	}
	if reflect.ValueOf(body).Kind() == reflect.Ptr {
		err := json.Unmarshal(reqBody, body)
		return err
	}
	return nil
}

func SendBadRequestError(w http.ResponseWriter) {
	http.Error(w, "bad request", http.StatusBadRequest)
}

func GetUserId(r *http.Request) (uint, error) {
	userId, ok := r.Context().Value("user_id").(uint)
	if !ok {
		logrus.Error("No user id present")
		return 0, constants.NoUserIdPresent
	}
	return userId, nil
}
