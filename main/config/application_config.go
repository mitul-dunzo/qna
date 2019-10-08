package config

import (
	"github.com/gorilla/mux"
	"qna/main/clients"
	"qna/main/orchestrator"
	"qna/main/services"
)

func InitializeApp() func(mux *mux.Router) {
	db := GetDB()
	redis := GetRedis()

	smsClient := clients.NewSmsClient()

	jwtService := services.NewJwtService()
	otpService := services.NewOtpService(redis, &smsClient)
	userService := services.NewUserService(db, &jwtService)
	questionService := services.NewQuestionService(db)

	loginOrchestrator := orchestrator.NewLoginOrchestrator(&otpService, &userService)
	authMiddleware := orchestrator.NewAuthenticationMiddleware(jwtService)
	questionOrch := orchestrator.NewQuestionOrchestrator(&questionService)

	return func(mux *mux.Router) {
		mux.Use(authMiddleware.Check)

		loginRouter := mux.PathPrefix("/auth").Subrouter()
		loginOrchestrator.Handle(loginRouter)

		questionRouter := mux.PathPrefix("/questions").Subrouter()
		questionOrch.Handle(questionRouter)
	}
}
