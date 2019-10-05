package services

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type login struct {
	PhoneNumber string `json:"phone_number"`
}

type GenericSuccess struct {
	Status string `json:"status"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logrus.Error("Couldn't read from body request: ", err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	var loginResult login
	err = json.Unmarshal(b, &loginResult)
	if err != nil {
		logrus.Error("Couldn't unmarshal from body request: ", err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	err = SendOtp(loginResult.PhoneNumber)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	result := GenericSuccess{Status: "Ok"}
	output, err := json.Marshal(result)
	if err != nil {
		logrus.Error("Couldn't write to response: ", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
