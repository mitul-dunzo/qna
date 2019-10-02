package main

import (
	"qna/main/config"
	"qna/main/server"
)

func main() {
	config.SetupDatabase()
	config.SetupNewRelic()

	srv := server.New()
	srv.ListenAndServe()
}
