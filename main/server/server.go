package server

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"qna/main/handler"
	"time"
)

type Server struct {
	*mux.Router
	Address string
}

func New() *Server {
	r := mux.NewRouter()
	addr := "0.0.0.0:8000"
	s := Server{
		Router:  r,
		Address: addr,
	}
	s.setupComponents()
	return &s
}

func (s Server) setupComponents() {
	s.HandleFunc("/", handler.Welcome).Methods(http.MethodGet)
}

func (s Server) ListenAndServe() {
	loggedRouter := handlers.LoggingHandler(os.Stdout, s.Router)
	srv := &http.Server{
		Addr:         s.Address,
		Handler:      loggedRouter,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	logrus.Info("Server starting at addr: ", s.Address)
	logrus.Fatal(srv.ListenAndServe())
}
