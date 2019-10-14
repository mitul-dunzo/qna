package orchestrator

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"qna/main/dtos"
	"qna/main/services"
	"qna/main/utils"
)

type LoginOrchestrator struct {
	otpService  *services.OtpService
	userService *services.UserService
}

func NewLoginOrchestrator(otpService *services.OtpService, userService *services.UserService) LoginOrchestrator {
	return LoginOrchestrator{
		otpService:  otpService,
		userService: userService,
	}
}

func (orch *LoginOrchestrator) Handle(r *mux.Router) {
	utils.Instrument(r, "/login", orch.login).Methods(http.MethodPost)
	utils.Instrument(r, "/verify-otp", orch.verifyOtp).Methods(http.MethodPost)
}

func (orch *LoginOrchestrator) login(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logrus.Error("Couldn't read from body request: ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var userDetails dtos.UserDetails
	err = json.Unmarshal(b, &userDetails)
	if err != nil {
		logrus.Error("Couldn't unmarshal from body request: ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// We should add a checks on email and phone number here.
	err = orch.otpService.SendOtp(&userDetails)
	if err != nil {
		logrus.Error("Couldn't write to response: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (orch *LoginOrchestrator) verifyOtp(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logrus.Error("Couldn't read from body request: ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var otpResult dtos.OtpData
	err = json.Unmarshal(b, &otpResult)
	if err != nil {
		logrus.Error("Couldn't unmarshal from body request: ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userDetails, err := orch.otpService.ValidateOtp(otpResult.PhoneNumber, otpResult.Otp)
	if err != nil {
		if err.Error() == services.InvalidOtp {
			http.Error(w, "Invalid OTP", http.StatusUnauthorized)
			return
		}
		logrus.Error("Couldn't check otp: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jwt, err := orch.userService.CreateUser(userDetails)
	if err != nil {
		logrus.Error("Couldn't create jwt: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := dtos.OtpValidationSuccess{Jwt: jwt}
	output, err := json.Marshal(result)
	if err != nil {
		logrus.Error("Couldn't write to response: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	_, _ = w.Write(output)
}
