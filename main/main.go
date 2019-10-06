package main

import (
	"qna/main/config"
	"qna/main/server"
)

func main() {
	config.SetupDatabase()
	config.SetupNewRelic()
	config.SetupRedis()

	srv := server.New(config.InitializeApp)
	srv.ListenAndServe()
}
