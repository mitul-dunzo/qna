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
)

type LoginOrchestrator struct {
	otpService  services.IOtpService
	userService services.IUserService
}

func NewLoginOrchestrator(otpService services.IOtpService, userService services.IUserService) LoginOrchestrator {
	return LoginOrchestrator{
		otpService:  otpService,
		userService: userService,
	}
}

func (orch *LoginOrchestrator) Handle(r *mux.Router) {
	utils.Instrument(r, constants.LoginEp, orch.login).Methods(http.MethodPost)
	utils.Instrument(r, constants.VerifyOtpEp, orch.verifyOtp).Methods(http.MethodPost)
}

func (orch *LoginOrchestrator) login(w http.ResponseWriter, r *http.Request) {
	var userDetails dtos.UserDetails
	err := utils.ReadHTTPBody(r, &userDetails)
	if err != nil {
		utils.SendBadRequestError(w)
		return
	}

	// We should add checks on email and phone number here.
	err = orch.otpService.SendOtp(&userDetails)
	if err != nil {
		logrus.Error("Couldn't write to response: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (orch *LoginOrchestrator) verifyOtp(w http.ResponseWriter, r *http.Request) {
	var otpResult dtos.OtpData
	err := utils.ReadHTTPBody(r, &otpResult)
	if err != nil {
		utils.SendBadRequestError(w)
		return
	}

	userDetails, err := orch.otpService.ValidateOtp(otpResult.PhoneNumber, otpResult.Otp)
	if err != nil {
		if err == constants.InvalidOtpError {
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
