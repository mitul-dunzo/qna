package server

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type Server struct {
	*mux.Router
	Address string
}

func New(route func(router *mux.Router)) *Server {
	r := mux.NewRouter()
	addr := os.Getenv("ServerAddress")
	s := Server{
		Router:  r,
		Address: addr,
	}
	r.Use()
	route(r)
	return &s
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
