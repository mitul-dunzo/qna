package main

import (
	"qna/main/config"
	"qna/main/server"
)

func main() {
	config.SetupDatabase()
	config.SetupNewRelic()
	config.SetupRedis()
	router := config.InitializeApp()

	srv := server.New(router)
	srv.ListenAndServe()
}
