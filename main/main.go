package main

import (
	"qna/main/config"
	"qna/main/server"
	"qna/main/utils"
)

func main() {
	config.SetupDatabase()
	config.SetupRedis()
	utils.SetupNewRelic()
	router := config.InitializeApp()

	srv := server.New(router)
	srv.ListenAndServe()
}
