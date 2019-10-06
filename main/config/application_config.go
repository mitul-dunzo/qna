package config

import (
	"github.com/gorilla/mux"
	"qna/main/clients"
	"qna/main/orchestrator"
	"qna/main/services"
)

func InitializeApp(mux *mux.Router) {
	db := GetDB()
	redis := GetRedis()

	smsClient := clients.NewSmsClient()

	jwtService := services.NewJwtService()
	otpService := services.NewOtpService(redis, &smsClient)
	userService := services.NewUserService(db, &jwtService)

	loginOrchestrator := orchestrator.NewLoginOrchestrator(&otpService, &userService)

	router := mux.PathPrefix("/auth").Subrouter()
	if router != nil {
		loginOrchestrator.Handle(router)
		return
	}

}
